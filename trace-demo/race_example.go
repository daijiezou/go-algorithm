package main

import (
	"sync"
	"sync/atomic"
)

// 演示数据竞争和解决方案

// 错误示例：存在数据竞争
type UnsafeCounter struct {
	count int64
}

func (c *UnsafeCounter) Increment() {
	c.count++ // 数据竞争！
}

func (c *UnsafeCounter) Get() int64 {
	return c.count // 数据竞争！
}

// 解决方案 1：使用 Mutex
type MutexCounter struct {
	mu    sync.Mutex
	count int64
}

func (c *MutexCounter) Increment() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

func (c *MutexCounter) Get() int64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

// 解决方案 2：使用 Atomic（更高效）
type AtomicCounter struct {
	count atomic.Int64
}

func (c *AtomicCounter) Increment() {
	c.count.Add(1)
}

func (c *AtomicCounter) Get() int64 {
	return c.count.Load()
}

// False Sharing 示例
// 错误：多个 goroutine 修改相邻的内存位置
type FalseSharingCounters struct {
	counter1 int64 // 可能在同一个 cache line
	counter2 int64 // 可能在同一个 cache line
}

// 正确：使用 padding 避免 false sharing
type PaddedCounters struct {
	counter1 int64
	_        [7]int64 // padding: 64 bytes (cache line size)
	counter2 int64
	_        [7]int64
}

// 演示数据竞争的函数
func DemoDataRace() int64 {
	counter := &UnsafeCounter{}
	var wg sync.WaitGroup
	
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter.Increment()
			}
		}()
	}
	
	wg.Wait()
	return counter.Get() // 结果不确定！
}

// 使用 Atomic 的正确版本
func DemoAtomicCounter() int64 {
	counter := &AtomicCounter{}
	var wg sync.WaitGroup
	
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter.Increment()
			}
		}()
	}
	
	wg.Wait()
	return counter.Get() // 结果确定：1000000
}

// 演示 False Sharing 的影响
func DemoFalseSharing() {
	counters := &FalseSharingCounters{}
	var wg sync.WaitGroup
	
	// 两个 goroutine 修改相邻的计数器
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 10000000; i++ {
			atomic.AddInt64(&counters.counter1, 1)
		}
	}()
	
	go func() {
		defer wg.Done()
		for i := 0; i < 10000000; i++ {
			atomic.AddInt64(&counters.counter2, 1)
		}
	}()
	
	wg.Wait()
}

// 使用 Padding 避免 False Sharing
func DemoPaddedCounters() {
	counters := &PaddedCounters{}
	var wg sync.WaitGroup
	
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 10000000; i++ {
			atomic.AddInt64(&counters.counter1, 1)
		}
	}()
	
	go func() {
		defer wg.Done()
		for i := 0; i < 10000000; i++ {
			atomic.AddInt64(&counters.counter2, 1)
		}
	}()
	
	wg.Wait()
}
