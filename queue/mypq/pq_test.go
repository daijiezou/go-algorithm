package mypq

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	pq := NewPQ()
	pq.insert(5)
	pq.insert(1)
	fmt.Println(pq.nums)
	pq.insert(2)
	fmt.Println(pq.nums)
	pq.insert(3)
	fmt.Println(pq.nums)
	fmt.Println(pq.Pop())
	fmt.Println(pq.nums)
}
