# 使用指南

## 快速开始

### 1. 基础运行

```bash
# 运行所有模式
go run . -mode=all -tasks=10000

# 只运行特定模式
go run . -mode=single -tasks=10000
go run . -mode=fanout -tasks=10000
go run . -mode=pool -tasks=10000
go run . -mode=pool-tuned -tasks=10000 -gogc=1000
```

### 2. 生成 Trace 文件

```bash
# 创建 traces 目录
mkdir -p traces

# 生成不同模式的 trace
go run . -mode=fanout -tasks=10000 -trace=traces/fanout.trace
go run . -mode=pool -tasks=10000 -trace=traces/pool.trace
go run . -mode=pool-tuned -tasks=10000 -trace=traces/pool-tuned.trace -gogc=1000
```

### 3. 查看 Trace

```bash
# 打开 trace 可视化界面
go tool trace traces/fanout.trace

# 浏览器会自动打开，显示以下视图：
# - View trace: 时间线视图
# - Goroutine analysis: Goroutine 分析
# - Network blocking profile: 网络阻塞分析
# - Synchronization blocking profile: 同步阻塞分析
# - Syscall blocking profile: 系统调用阻塞分析
# - Scheduler latency profile: 调度延迟分析
```

## Trace 分析实战

### 场景 1: 对比 Fan-out 和 Worker Pool

```bash
# 生成两个 trace
go run . -mode=fanout -tasks=10000 -trace=traces/fanout.trace
go run . -mode=pool -tasks=10000 -trace=traces/pool.trace

# 分别查看
go tool trace traces/fanout.trace
go tool trace traces/pool.trace
```

**观察要点**：
- Fan-out: 大量 goroutine 同时创建和销毁
- Worker Pool: 固定数量的 goroutine，复用性好

### 场景 2: GC 影响分析

```bash
# 默认 GOGC=100
go run . -mode=pool -tasks=10000 -trace=traces/pool-gc100.trace

# 优化 GOGC=1000
go run . -mode=pool-tuned -tasks=10000 -trace=traces/pool-gc1000.trace -gogc=1000

# 对比查看
go tool trace traces/pool-gc100.trace
go tool trace traces/pool-gc1000.trace
```

**观察要点**：
- GC 事件频率
- GC 暂停时间
- 程序执行时间占比

### 场景 3: 调度器行为

在 trace 界面中：
1. 点击 "View trace"
2. 观察 "PROCS" 行（每个 P 的执行情况）
3. 查看 goroutine 的调度切换
4. 识别是否有 P 空闲（说明没有充分利用 CPU）

## Benchmark 测试

### 基础 Benchmark

```bash
# 运行所有 benchmark
go test -bench=. -benchmem

# 运行特定 benchmark
go test -bench=BenchmarkFanOut -benchmem
go test -bench=BenchmarkWorkerPool -benchmem

# 增加运行时间以获得更准确的结果
go test -bench=. -benchmem -benchtime=5s
```

### 对比不同 GOGC 值

```bash
# GOGC=100（默认）
GOGC=100 go test -bench=BenchmarkWorkerPool -benchmem -benchtime=3s

# GOGC=1000（优化）
GOGC=1000 go test -bench=BenchmarkWorkerPool -benchmem -benchtime=3s

# 对比输出
```

### False Sharing 对比

```bash
# 对比有无 padding 的性能差异
go test -bench='BenchmarkFalseSharing|BenchmarkPaddedCounters' -benchmem -benchtime=3s
```

**预期结果**：PaddedCounters 应该明显更快

## 数据竞争检测

### 运行 Race Detector

```bash
# 检测主程序
go run -race . -mode=single -tasks=100

# 检测测试
go test -race

# 检测 benchmark
go test -race -bench=BenchmarkDataRace
```

**注意**：
- Race detector 会显著降低性能（5-10x）
- 只在开发和测试时使用
- 生产环境不要启用

### 修复数据竞争

如果检测到数据竞争，输出类似：
```
WARNING: DATA RACE
Write at 0x... by goroutine 7:
  main.(*UnsafeCounter).Increment()
      /path/to/file.go:15

Previous read at 0x... by goroutine 6:
  main.(*UnsafeCounter).Get()
      /path/to/file.go:19
```

解决方案：
1. 使用 `sync.Mutex`
2. 使用 `atomic` 操作
3. 使用 channel 通信

## CPU Profiling

### 生成 CPU Profile

```bash
# 运行 benchmark 并生成 profile
go test -bench=BenchmarkWorkerPool -cpuprofile=cpu.prof -benchtime=5s

# 查看 profile
go tool pprof cpu.prof

# 在 pprof 交互界面中：
(pprof) top10        # 显示 top 10 函数
(pprof) list main    # 显示 main 包的详细信息
(pprof) web          # 生成调用图（需要 graphviz）
```

### 常用 pprof 命令

```bash
# 文本模式
go tool pprof -text cpu.prof

# 生成 PDF 调用图
go tool pprof -pdf cpu.prof > cpu.pdf

# Web 界面
go tool pprof -http=:8080 cpu.prof
```

## Memory Profiling

### 生成 Memory Profile

```bash
# 运行 benchmark 并生成 profile
go test -bench=BenchmarkWorkerPool -memprofile=mem.prof -benchtime=5s

# 查看 profile
go tool pprof mem.prof

# 在 pprof 交互界面中：
(pprof) top10              # 显示分配最多的函数
(pprof) list processTask   # 查看具体函数的分配
(pprof) alloc_space        # 按总分配量排序
(pprof) alloc_objects      # 按分配对象数排序
```

### 分析内存泄漏

```bash
# 查看当前存活的对象
go tool pprof -inuse_space mem.prof

# 查看总分配量
go tool pprof -alloc_space mem.prof
```

## 实验建议

### 实验 1: 任务数量影响

```bash
# 测试不同任务数量
for n in 100 1000 5000 10000 50000; do
    echo "Tasks: $n"
    go run . -mode=all -tasks=$n
    echo ""
done
```

### 实验 2: Worker 数量优化

修改 `main.go` 中的 worker 数量：
```go
// 测试不同 worker 数量
for _, numWorkers := range []int{1, 2, 4, 8, 16, 32} {
    stats := measurePerformance(fmt.Sprintf("Pool-%d", numWorkers), func() {
        workerPool(tasks, numWorkers)
    })
    printStats(fmt.Sprintf("Worker Pool (%d workers)", numWorkers), stats)
}
```

### 实验 3: GOGC 调优

```bash
# 测试不同 GOGC 值
for gogc in 50 100 200 500 1000 2000; do
    echo "GOGC=$gogc"
    GOGC=$gogc go run . -mode=pool -tasks=10000
    echo ""
done
```

### 实验 4: GOMEMLIMIT 测试

```bash
# 限制内存使用
GOMEMLIMIT=100MiB go run . -mode=pool -tasks=10000
GOMEMLIMIT=500MiB go run . -mode=pool -tasks=10000
GOMEMLIMIT=1GiB go run . -mode=pool -tasks=10000
```

## 常见问题排查

### 问题 1: 程序很慢

**排查步骤**：
1. 生成 trace: `go run . -trace=slow.trace`
2. 查看 trace: `go tool trace slow.trace`
3. 检查：
   - GC 是否频繁？→ 调整 GOGC
   - Goroutine 是否阻塞？→ 检查锁和 channel
   - CPU 是否充分利用？→ 调整并发数

### 问题 2: 内存占用高

**排查步骤**：
1. 生成 memory profile: `go test -bench=. -memprofile=mem.prof`
2. 查看分配: `go tool pprof -alloc_space mem.prof`
3. 优化：
   - 复用对象（sync.Pool）
   - 预分配切片
   - 减少不必要的分配

### 问题 3: GC 暂停时间长

**排查步骤**：
1. 查看 trace 中的 GC 事件
2. 调整 GOGC 或 GOMEMLIMIT
3. 减少堆上分配
4. 考虑使用 `runtime.GC()` 手动触发

### 问题 4: 数据竞争

**排查步骤**：
1. 运行: `go run -race .`
2. 查看报告，定位竞争位置
3. 使用 atomic 或 mutex 修复

## 性能优化检查清单

- [ ] 使用 trace 识别瓶颈
- [ ] 使用 pprof 找到热点函数
- [ ] 检测并修复数据竞争
- [ ] 优化 GOGC 或使用 GOMEMLIMIT
- [ ] 选择合适的并发模式
- [ ] 避免 False Sharing
- [ ] 减少内存分配
- [ ] 复用对象（sync.Pool）
- [ ] 预分配切片容量
- [ ] 使用 benchmark 验证优化效果

## 自动化脚本

使用提供的 `run.sh` 脚本一键运行所有测试：

```bash
./run.sh
```

脚本会：
1. ✅ 运行所有模式并生成 trace
2. ✅ 运行 benchmark 测试
3. ✅ 检测数据竞争
4. ✅ 生成 CPU 和 Memory profile
5. ✅ 提供查看命令

## 学习路径

1. **第一天**：运行基础示例，理解四种模式
2. **第二天**：学习 trace 工具，分析 goroutine 行为
3. **第三天**：学习 pprof，优化 CPU 和内存
4. **第四天**：实验 GOGC 调优，理解 GC 影响
5. **第五天**：学习数据竞争检测和修复
6. **第六天**：深入 False Sharing 和 cache line
7. **第七天**：综合优化，达到最佳性能

## 推荐阅读

- [Go Execution Tracer Design Doc](https://docs.google.com/document/d/1FP5apqzBgr7ahCCgFO-yoVhk4YZrNIDNf9RybngBc14/pub)
- [Profiling Go Programs](https://blog.golang.org/profiling-go-programs)
- [Go GC Guide](https://tip.golang.org/doc/gc-guide)
- [The Go Memory Model](https://golang.org/ref/mem)
- [Effective Go - Concurrency](https://golang.org/doc/effective_go#concurrency)

祝你学习愉快！🚀
