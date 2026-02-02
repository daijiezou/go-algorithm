package main

import (
	"runtime"
	"testing"
)

// 生成测试任务
func generateTasks(n int) []Task {
	tasks := make([]Task, n)
	for i := range tasks {
		tasks[i] = Task{
			ID:   i,
			Data: make([]byte, 100),
		}
	}
	return tasks
}

// Benchmark: 单线程
func BenchmarkSingleThreaded(b *testing.B) {
	tasks := generateTasks(1000)
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		singleThreaded(tasks)
	}
}

// Benchmark: Fan-out
func BenchmarkFanOut(b *testing.B) {
	tasks := generateTasks(1000)
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		fanOut(tasks)
	}
}

// Benchmark: Worker Pool
func BenchmarkWorkerPool(b *testing.B) {
	tasks := generateTasks(1000)
	numWorkers := runtime.NumCPU()
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		workerPool(tasks, numWorkers)
	}
}

// Benchmark: 数据竞争示例
func BenchmarkDataRace(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DemoDataRace()
	}
}

// Benchmark: Atomic Counter
func BenchmarkAtomicCounter(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DemoAtomicCounter()
	}
}

// Benchmark: False Sharing
func BenchmarkFalseSharing(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DemoFalseSharing()
	}
}

// Benchmark: Padded Counters
func BenchmarkPaddedCounters(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DemoPaddedCounters()
	}
}

// 不同 GOGC 值的对比
func BenchmarkWorkerPoolGOGC100(b *testing.B) {
	tasks := generateTasks(1000)
	numWorkers := runtime.NumCPU()
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		workerPool(tasks, numWorkers)
	}
}

func BenchmarkWorkerPoolGOGC1000(b *testing.B) {
	tasks := generateTasks(1000)
	numWorkers := runtime.NumCPU()
	// 注意：实际使用时通过环境变量设置 GOGC=1000
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		workerPool(tasks, numWorkers)
	}
}
