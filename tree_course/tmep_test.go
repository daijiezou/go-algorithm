package tree_course

import (
	"fmt"
	"testing"
)

func TestSizeListNode(t *testing.T) {
	type args struct {
		node *ListNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{node: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val:  3,
						Next: nil,
					},
				},
			}},
			want: 3,
		},
		{
			name: "test2",
			args: args{node: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
			}},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SizeListNode(tt.args.node); got != tt.want {
				t.Errorf("SizeListNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderPrintArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			args: args{nums: []int{1, 2, 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			OrderPrintArray(tt.args.nums)
		})
	}
}

func TestOrderPrintNode(t *testing.T) {
	type args struct {
		node *ListNode
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test2",
			args: args{node: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			OrderPrintNode(tt.args.node)
		})
	}
}

func TestTempTest(t *testing.T) {
	a := 1
	b := 3
	c := 3
	d := 4

	fmt.Println(a ^ b ^ c ^ d)
}
