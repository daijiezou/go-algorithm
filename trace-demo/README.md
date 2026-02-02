# Go Trace 性能分析演示

这个项目演示了如何使用 Go 的 trace 和 profiling 工具来分析和优化并发程序的性能。

## 核心概念

### Profiling vs Trace

- **Profiling**：显示"正在发生什么"
  - CPU 在哪里花费时间
  - 内存在哪里分配
  - 哪些函数最耗时

- **Trace**：显示"什么没有发生"
  - Goroutine 为什么在等待
  - 调度器的行为
  - GC 的影响
  - 系统调用阻塞

### 并发的本质

**Out of Order Execution（乱序执行）**：现代 CPU 和编译器会重排指令以提高性能，这在并发环境下可能导致数据竞争。

## 项目结构

```
trace-demo/
├── main.go              # 主程序，实现四种处理模式
├── race_example.go      # 数据竞争示例和解决方案
├── benchmark_test.go    # 性能测试
├── run.sh              # 自动化运行脚本
└── README.md           # 本文件
```

## 四种处理模式

### 1. 单线程（Single-threaded）
```go
// 顺序处理所有任务
for _, task := range tasks {
    result := processTask(task)
}
```
- ✅ 简单、可预测
- ❌ 无法利用多核
- 📊 预期：~1000ms，6MB 内存，7% GC 时间

### 2. Fan-out 模式
```go
// 为每个任务创建一个 goroutine
for _, task := range tasks {
    go func(t Task) {
        processTask(t)
    }(task)
}
```
- ✅ 最大化并发
- ❌ 大量 goroutine 创建开销
- ❌ 内存分配激增，GC 压力大
- 📊 预期：~264ms，55MB 内存，65% GC 时间

### 3. Worker Pool 模式
```go
// 固定数量的 worker goroutines
for i := 0; i < numWorkers; i++ {
    go worker(taskChan)
}
```
- ✅ 控制并发数量
- ✅ 减少 goroutine 创建开销
- ⚠️  默认 GOGC 下 GC 频繁
- 📊 预期：~388ms，13MB 内存，74% GC 时间

### 4. Worker Pool + GOGC 调优
```bash
GOGC=1000 go run .
```
- ✅ 减少 GC 频率
- ✅ 最佳性能
- ⚠️  使用更多内存
- 📊 预期：~136ms，70MB 内存，13% GC 时间

## 数据竞争解决方案

### 问题：数据竞争
```go
type UnsafeCounter struct {
    count int64
}

func (c *UnsafeCounter) Increment() {
    c.count++ // 数据竞争！
}
```

### 解决方案 1：Mutex
```go
type MutexCounter struct {
    mu    sync.Mutex
    count int64
}

func (c *MutexCounter) Increment() {
    c.mu.Lock()
    c.count++
    c.mu.Unlock()
}
```

### 解决方案 2：Atomic（推荐）
```go
type AtomicCounter struct {
    count atomic.Int64
}

func (c *AtomicCounter) Increment() {
    c.count.Add(1)
}
```

## False Sharing 问题

### 什么是 False Sharing？

当多个 CPU 核心修改同一个 cache line 中的不同变量时，会导致 cache line 频繁失效，严重影响性能。

### 错误示例
```go
type FalseSharingCounters struct {
    counter1 int64 // 可能在同一个 cache line
    counter2 int64 // 可能在同一个 cache line
}
```

### 正确做法：添加 Padding
```go
type PaddedCounters struct {
    counter1 int64
    _        [7]int64 // padding: 64 bytes (cache line size)
    counter2 int64
    _        [7]int64
}
```

## 快速开始

### 1. 运行所有模式
```bash
go run . -mode=all -tasks=10000
```

### 2. 生成 trace 文件
```bash
go run . -mode=fanout -tasks=10000 -trace=fanout.trace
go tool trace fanout.trace
```

### 3. 对比不同 GOGC 值
```bash
# 默认 GOGC=100
go run . -mode=pool -tasks=10000 -trace=pool.trace

# 优化 GOGC=1000
go run . -mode=pool-tuned -tasks=10000 -trace=pool-tuned.trace -gogc=1000
```

### 4. 运行 Benchmark
```bash
go test -bench=. -benchmem
```

### 5. 检测数据竞争
```bash
go run -race . -mode=single -tasks=100
```

### 6. 使用自动化脚本
```bash
chmod +x run.sh
./run.sh
```

## Trace 分析技巧

### 打开 trace 文件
```bash
go tool trace traces/fanout.trace
```

### 关键指标

1. **Goroutine Analysis**
   - 查看 goroutine 的创建和销毁
   - 识别 goroutine 泄漏

2. **Network/Syscall Blocking**
   - 查看系统调用阻塞时间
   - 识别 I/O 瓶颈

3. **Synchronization Blocking**
   - 查看锁竞争
   - 识别 channel 阻塞

4. **Scheduler Latency**
   - 查看调度延迟
   - 识别 CPU 饥饿

5. **GC Events**
   - 查看 GC 暂停时间
   - 分析 GC 频率

## 性能优化建议

### 1. 选择合适的并发模式
- **CPU 密集型**：Worker Pool（worker 数 = CPU 核心数）
- **I/O 密集型**：Fan-out（可以创建更多 goroutine）
- **混合型**：根据实际情况调整

### 2. GC 调优
```bash
# 减少 GC 频率（使用更多内存）
GOGC=1000 go run .

# 限制最大内存使用
GOMEMLIMIT=1GiB go run .
```

### 3. 避免数据竞争
- 使用 `go run -race` 检测
- 优先使用 `atomic` 而非 `mutex`
- 使用 channel 通信而非共享内存

### 4. 避免 False Sharing
- 为频繁修改的变量添加 padding
- 使用局部变量累积，最后再写入共享变量

### 5. 减少内存分配
- 复用对象（sync.Pool）
- 预分配切片容量
- 避免不必要的字符串拼接

## 关键洞察

1. **"Throwing goroutines at the problem" 是好的第一步**
   - Go 的调度器和 GC 为 web 服务场景优化
   - 先实现并发，再根据 trace 优化

2. **GOMEMLIMIT 比 GOGC 语义更清晰**
   - GOGC：基于堆增长百分比触发 GC
   - GOMEMLIMIT：设置内存上限，更直观

3. **Trace 比 Profiling 更适合诊断并发问题**
   - Profiling：找到热点函数
   - Trace：找到等待和阻塞的原因

## 常见问题

### Q: 为什么 Worker Pool 比 Fan-out 慢？
A: 默认 GOGC=100 导致 GC 过于频繁。调整 GOGC=1000 后，Worker Pool 性能最佳。

### Q: 如何选择 Worker 数量？
A: 
- CPU 密集型：`runtime.NumCPU()`
- I/O 密集型：可以更多，根据实际测试调整

### Q: GOGC 设置多少合适？
A: 
- 默认 100（堆增长 100% 触发 GC）
- 内存充足时可设置 200-1000
- 使用 GOMEMLIMIT 更安全

### Q: 如何判断是否有数据竞争？
A: 
```bash
go run -race .
go test -race
```

## 参考资源

- [Go Execution Tracer](https://golang.org/pkg/runtime/trace/)
- [Go Memory Model](https://golang.org/ref/mem)
- [Profiling Go Programs](https://blog.golang.org/profiling-go-programs)
- [GC Guide](https://tip.golang.org/doc/gc-guide)

## 实验建议

1. 修改任务数量，观察性能变化
2. 调整 Worker 数量，找到最佳值
3. 对比不同 GOGC 值的影响
4. 使用 `-race` 检测数据竞争
5. 对比 False Sharing 前后的性能差异

## 总结

这个项目展示了：
- ✅ 如何使用 trace 分析并发程序
- ✅ 不同并发模式的性能特征
- ✅ GC 调优的重要性
- ✅ 数据竞争的检测和解决
- ✅ False Sharing 的影响和避免方法

通过实际运行和分析 trace，你可以深入理解 Go 的并发模型和性能优化技巧。
