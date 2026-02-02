package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/debug"
	"runtime/trace"
	"sync"
	"time"
)

// 模拟一个需要处理的任务
type Task struct {
	ID   int
	Data []byte
}

// 模拟 CPU 密集型工作
func processTask(task Task) Result {
	// 模拟一些计算工作
	sum := 0
	for i := 0; i < 10000; i++ {
		sum += i
	}
	
	// 模拟更多内存分配（触发 GC）
	// 每个任务分配约 100KB
	temp := make([]byte, 100*1024)
	for i := range temp {
		temp[i] = byte(sum % 256)
	}
	
	// 额外分配一些临时对象
	_ = make([]int, 1000)
	_ = make(map[int]string)
	
	return Result{
		TaskID: task.ID,
		Value:  sum,
	}
}

type Result struct {
	TaskID int
	Value  int
}

// 方法 1: 单线程处理
func singleThreaded(tasks []Task) []Result {
	results := make([]Result, 0, len(tasks))
	for _, task := range tasks {
		result := processTask(task)
		results = append(results, result)
	}
	return results
}

// 方法 2: Fan-out 模式（为每个任务创建 goroutine）
func fanOut(tasks []Task) []Result {
	results := make([]Result, len(tasks))
	var wg sync.WaitGroup
	
	for i, task := range tasks {
		wg.Add(1)
		go func(idx int, t Task) {
			defer wg.Done()
			results[idx] = processTask(t)
		}(i, task)
	}
	
	wg.Wait()
	return results
}

// 方法 3: Worker Pool 模式
func workerPool(tasks []Task, numWorkers int) []Result {
	results := make([]Result, len(tasks))
	taskChan := make(chan struct {
		idx  int
		task Task
	}, len(tasks))
	
	var wg sync.WaitGroup
	
	// 启动 worker
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for job := range taskChan {
				results[job.idx] = processTask(job.task)
			}
		}()
	}
	
	// 发送任务
	for i, task := range tasks {
		taskChan <- struct {
			idx  int
			task Task
		}{i, task}
	}
	close(taskChan)
	
	wg.Wait()
	return results
}

// 性能统计
type Stats struct {
	Duration   time.Duration
	MemAlloc   uint64
	TotalAlloc uint64
	GCTime     time.Duration
	NumGC      uint32
}

func measurePerformance(name string, fn func()) Stats {
	// 强制 GC 以获得干净的起点
	runtime.GC()
	
	var memBefore, memAfter runtime.MemStats
	runtime.ReadMemStats(&memBefore)
	
	start := time.Now()
	fn()
	duration := time.Since(start)
	
	runtime.ReadMemStats(&memAfter)
	
	return Stats{
		Duration:   duration,
		MemAlloc:   memAfter.Alloc,
		TotalAlloc: memAfter.TotalAlloc - memBefore.TotalAlloc,
		GCTime:     time.Duration(memAfter.PauseTotalNs - memBefore.PauseTotalNs),
		NumGC:      memAfter.NumGC - memBefore.NumGC,
	}
}

func printStats(name string, stats Stats) {
	gcPercent := float64(stats.GCTime) / float64(stats.Duration) * 100
	fmt.Printf("\n=== %s ===\n", name)
	fmt.Printf("执行时间: %v\n", stats.Duration)
	fmt.Printf("内存分配: %.2f MB\n", float64(stats.TotalAlloc)/1024/1024)
	fmt.Printf("当前内存: %.2f MB\n", float64(stats.MemAlloc)/1024/1024)
	fmt.Printf("GC 次数:  %d\n", stats.NumGC)
	fmt.Printf("GC 时间:  %v (%.1f%%)\n", stats.GCTime, gcPercent)
}

func main() {
	mode := flag.String("mode", "all", "运行模式: single, fanout, pool, pool-tuned, all")
	numTasks := flag.Int("tasks", 50000, "任务数量（默认 50000，确保触发 GC）")
	traceFile := flag.String("trace", "", "trace 输出文件")
	gogc := flag.Int("gogc", 100, "GOGC 值")
	flag.Parse()
	
	// 设置 GOGC
	if *gogc != 100 {
		old := debug.SetGCPercent(*gogc)
		fmt.Printf("GOGC 从 %d 设置为 %d\n", old, *gogc)
	}
	
	// 启动 trace
	if *traceFile != "" {
		f, err := os.Create(*traceFile)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		
		if err := trace.Start(f); err != nil {
			log.Fatal(err)
		}
		defer trace.Stop()
		fmt.Printf("Trace 输出到: %s\n", *traceFile)
	}
	
	// 生成任务
	tasks := make([]Task, *numTasks)
	for i := range tasks {
		tasks[i] = Task{
			ID:   i,
			Data: make([]byte, 100),
		}
	}
	
	fmt.Printf("任务数量: %d\n", *numTasks)
	fmt.Printf("CPU 核心数: %d\n", runtime.NumCPU())
	
	// 运行不同模式
	if *mode == "all" || *mode == "single" {
		stats := measurePerformance("单线程", func() {
			singleThreaded(tasks)
		})
		printStats("单线程处理", stats)
	}
	
	if *mode == "all" || *mode == "fanout" {
		stats := measurePerformance("Fan-out", func() {
			fanOut(tasks)
		})
		printStats("Fan-out 模式", stats)
	}
	
	if *mode == "all" || *mode == "pool" {
		numWorkers := runtime.NumCPU()
		stats := measurePerformance("Worker Pool", func() {
			workerPool(tasks, numWorkers)
		})
		printStats(fmt.Sprintf("Worker Pool (%d workers)", numWorkers), stats)
	}
	
	if *mode == "pool-tuned" {
		numWorkers := runtime.NumCPU()
		stats := measurePerformance("Worker Pool (Tuned)", func() {
			workerPool(tasks, numWorkers)
		})
		printStats(fmt.Sprintf("Worker Pool Tuned (%d workers, GOGC=%d)", numWorkers, *gogc), stats)
	}
}
