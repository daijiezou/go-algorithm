package _1_listNodeAndArray

import (
	"container/heap"
)

type PriorityQueue [][]int

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i][0]+pq[i][1] < pq[j][0]+pq[j][1]
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.([]int))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func kSmallestPairs2(nums1 []int, nums2 []int, k int) [][]int {
	pq := make(PriorityQueue, 0)
	for i := 0; i < len(nums1); i++ {
		pq.Push([]int{nums1[i], nums2[0], 0})

	}
	res := make([][]int, 0)
	for pq.Len() > 0 && k > 0 {
		cur := heap.Pop(&pq).([]int)
		k--
		next_index := cur[2] + 1
		if next_index < len(nums2) {
			heap.Push(&pq, []int{cur[0], nums2[next_index], next_index})
		}
		pair := []int{cur[0], cur[1]}
		res = append(res, pair)
	}
	return res
}
