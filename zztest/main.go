package main

import (
	"fmt"
	"runtime"
	"unsafe"
)

func main() {
	// 测试栈上变量
	var stackVar int = 42
	stackAddr := uintptr(unsafe.Pointer(&stackVar))

	// 测试堆上变量（通过逃逸分析）
	heapVar := makeHeapVar()
	heapAddr := uintptr(unsafe.Pointer(heapVar))

	// 大数组测试（可能触发栈扩容）
	var largeArray [10000]int
	largeArrayAddr := uintptr(unsafe.Pointer(&largeArray))

	// 获取内存统计
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	fmt.Println("=== 内存地址分析 ===")
	fmt.Printf("栈变量地址:     0x%016x\n", stackAddr)
	fmt.Printf("堆变量地址:     0x%016x\n", heapAddr)
	fmt.Printf("大数组地址:     0x%016x\n", largeArrayAddr)

	fmt.Printf("\n=== 内存统计信息 ===\n")
	fmt.Printf("HeapAlloc:      0x%x (%d bytes) - 当前堆分配大小\n", m.HeapAlloc, m.HeapAlloc)
	fmt.Printf("HeapSys:        0x%x (%d bytes) - 堆从OS获取的总内存\n", m.HeapSys, m.HeapSys)
	fmt.Printf("StackInuse:     0x%x (%d bytes) - 栈正在使用的内存\n", m.StackInuse, m.StackInuse)
	fmt.Printf("StackSys:       0x%x (%d bytes) - 栈从OS获取的总内存\n", m.StackSys, m.StackSys)

	fmt.Printf("\n=== 地址特征分析 ===\n")
	fmt.Printf("栈地址前缀: 0x%x (通常是 0x14... 或 0x7ff...)\n", stackAddr>>40)
	fmt.Printf("堆地址前缀: 0x%x (通常是 0xc0...)\n", heapAddr>>40)

	fmt.Printf("\n=== 判断结果 ===\n")
	if (stackAddr>>40) == 0x14 || (stackAddr>>40) == 0x7ff {
		fmt.Println("✓ stackVar 很可能在栈上")
	}
	if (heapAddr >> 40) == 0xc0 {
		fmt.Println("✓ heapVar 很可能在堆上")
	}

	fmt.Printf("\n=== 提示 ===\n")
	fmt.Println("运行以下命令查看逃逸分析:")
	fmt.Println("  go build -gcflags='-m -l' main.go")
}

// 强制变量逃逸到堆上
func makeHeapVar() *int {
	x := 100
	return &x // 返回指针，x 会逃逸到堆
}

type DListNode struct {
	Key, Val int
	Pre      *DListNode
	Next     *DListNode
}

type LRUCache struct {
	cap   int
	Cache map[int]*DListNode
	Head  *DListNode
	Tail  *DListNode
}

func New(cap int) *LRUCache {
	if cap <= 0 {
		panic("cap value is invalid")
	}
	head := &DListNode{}
	tail := &DListNode{}
	head.Next = tail
	tail.Pre = head
	return &LRUCache{
		cap:   cap,
		Cache: make(map[int]*DListNode),
		Head:  head,
		Tail:  tail,
	}
}

func (l *LRUCache) DeleteNode(node *DListNode) {
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre
}

func (l *LRUCache) AddToHead(node *DListNode) {
	node.Next = l.Head.Next
	node.Pre = l.Head
	l.Head.Next.Pre = node
	l.Head.Next = node
}

func (l *LRUCache) RemoveTail() {
	node := l.Tail.Pre
	delete(l.Cache, node.Key)
	l.Tail.Pre.Pre.Next = l.Tail
	l.Tail.Pre = l.Tail.Pre.Pre
}

func (l *LRUCache) MovetoHead(node *DListNode) {
	l.DeleteNode(node)
	l.AddToHead(node)
}

func (l *LRUCache) Get(key int) int {
	if node, ok := l.Cache[key]; ok {
		l.MovetoHead(node)
		return node.Val

	}
	return -1
}

func (l *LRUCache) Put(key int, val int) {
	if node, ok := l.Cache[key]; ok {
		node.Val = val
		l.MovetoHead(node)
		return
	}
	node := &DListNode{
		Key: key,
		Val: val,
	}
	l.Cache[key] = node
	l.AddToHead(node)
	if len(l.Cache) > l.cap {
		l.RemoveTail()
	}
}
