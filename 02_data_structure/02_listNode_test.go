package _2_data_structure

import (
	"fmt"
	"testing"
)

func Test_deleteDuplicatesUnsorted(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "1",
			args: args{head: genListNode([]int{1, 2, 2, 3, 3, 3, 10, 9, 8, 7})},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := deleteDuplicatesUnsorted(tt.args.head)
			for res != nil {
				fmt.Println(res.Val)
				res = res.Next
			}
		})
	}
}

func genListNode(in []int) *ListNode {
	dummy := &ListNode{Val: -1}
	p := dummy
	for i := 0; i < len(in); i++ {
		p.Next = &ListNode{Val: in[i]}
		p = p.Next
	}
	return dummy.Next
}

func Test_nthUglyNumber(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{n: 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nthUglyNumber(tt.args.n); got != tt.want {
				t.Errorf("nthUglyNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_kthSmallest(t *testing.T) {
	type args struct {
		matrix [][]int
		k      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{

				matrix: [][]int{
					{1, 5, 9},
					{10, 11, 13},
					{12, 13, 15},
				},
				k: 8,
			},
			want: 13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthSmallest(tt.args.matrix, tt.args.k); got != tt.want {
				t.Errorf("kthSmallest() = %v, want %v", got, tt.want)
			}
		})
	}
}
