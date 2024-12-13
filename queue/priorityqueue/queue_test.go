package priorityqueue

import (
	"fmt"
	"testing"
)

func TestNewMinPQ(t *testing.T) {
	pq := NewMinPQ(0)
	pq.insert(5)
	pq.insert(1)
	pq.insert(2)
	pq.insert(3)
	fmt.Println(pq.pq)
	fmt.Println(pq.pop())
	fmt.Println(pq.pq)
}

type ListNode struct {
	Val  int
	Next *ListNode
}
