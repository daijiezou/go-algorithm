package _4_datastruct

import (
	"container/heap"
	"math"
	"sort"
)

type myInts []int

func (m myInts) Len() int {
	return len(m)
}

func (m myInts) Less(i, j int) bool {
	return m[i] > m[j]
}

func (m myInts) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *myInts) Push(x any) {
	*m = append(*m, x.(int))
}

// pop应该弹出最后一个，而不是第一个
func (m *myInts) Pop() any {
	n := len(*m)
	x := (*m)[n-1]
	*m = (*m)[:n-1]
	return x
}

func lastStoneWeight(stones []int) int {
	newStones := myInts(stones)
	q := &newStones
	heap.Init(q)
	for q.Len() > 1 {
		x, y := heap.Pop(q).(int), heap.Pop(q).(int)
		if x > y {
			heap.Push(q, x-y)
		}
	}
	if q.Len() > 0 {
		return (*q)[0]
	}
	return 0
}

/*
给你一个整数数组 gifts ，表示各堆礼物的数量。每一秒，你需要执行以下操作：

选择礼物数量最多的那一堆。
如果不止一堆都符合礼物数量最多，从中选择任一堆即可。
将堆中的礼物数量减少到堆中原来礼物数量的平方根，向下取整。
返回在 k 秒后剩下的礼物数量。
https://leetcode.cn/problems/take-gifts-from-the-richest-pile/description/
*/

func pickGifts(gifts []int, k int) int64 {
	h := &hp{gifts}
	heap.Init(h)
	for i := 0; i < k && gifts[0] > 1; i++ {
		gifts[0] = int(math.Sqrt(float64(gifts[0]))) // 直接修改堆顶
		heap.Fix(h, 0)
	}
	sum := 0
	for i := 0; i < len(gifts); i++ {
		sum += gifts[i]
	}
	return int64(sum)
}

type hp struct{ sort.IntSlice }

func (h hp) Less(i, j int) bool {
	return h.IntSlice[i] > h.IntSlice[j]
} // 最大堆

func (h *hp) Push(x any) {
	h.IntSlice = append(h.IntSlice, x.(int))
}

func (h *hp) Pop() any {
	n := len(h.IntSlice)
	x := h.IntSlice[n-1]
	h.IntSlice = h.IntSlice[:n-1]
	return x
}
