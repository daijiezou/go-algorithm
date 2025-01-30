package meeting150

import (
	"container/heap"
)

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}
func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// 寻找第K大的元素
func findKthLargest(nums []int, k int) int {
	h := IntHeap{}
	heap.Init(&h)
	for _, e := range nums {
		// 每个元素都要过一遍二叉堆
		heap.Push(&h, e)

	}
	for i := len(nums); i > k; i-- {
		heap.Pop(&h)
	}
	return heap.Pop(&h).(int)
}
