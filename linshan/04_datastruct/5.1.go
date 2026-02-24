package _4_datastruct

import "container/heap"

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
