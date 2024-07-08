package _1_listnode

type ListNodeHeap []*ListNode

func (l ListNodeHeap) Len() int {
	return len(l)
}

func (l ListNodeHeap) Less(i, j int) bool {
	return l[i].Val < l[j].Val
}

func (l ListNodeHeap) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l ListNodeHeap) Push(x any) {
	l = append(l, x.(*ListNode))
}

func (l ListNodeHeap) Pop() any {
	x := l[len(l)-1]
	l = l[0 : len(l)-1]
	return x
}
