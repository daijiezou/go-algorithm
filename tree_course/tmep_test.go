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

func TestParseTreeAndRightView(t *testing.T) {
	testCases := []struct {
		name      string
		input     string
		wantLevel [][]int
		wantRight []int
	}{
		{
			name:      "测试1: 1(2(7,8),3)",
			input:     "1(2(7,8),3)",
			wantLevel: [][]int{{1}, {2, 3}, {7, 8}},
			wantRight: []int{1, 3, 8},
		},
		{
			name:      "测试2: 1(2(7,8),4(5))",
			input:     "1(2(7,8),4(5))",
			wantLevel: [][]int{{1}, {2, 4}, {7, 8, 5}},
			wantRight: []int{1, 4, 5},
		},
		{
			name:      "测试3: 1(2,3(4,5))",
			input:     "1(2,3(4,5))",
			wantLevel: [][]int{{1}, {2, 3}, {4, 5}},
			wantRight: []int{1, 3, 5},
		},
		{
			name:      "测试4: 单节点",
			input:     "1",
			wantLevel: [][]int{{1}},
			wantRight: []int{1},
		},
		{
			name:      "测试5: 1(2,3)",
			input:     "1(2,3)",
			wantLevel: [][]int{{1}, {2, 3}},
			wantRight: []int{1, 3},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fmt.Printf("\n输入: %s\n", tc.input)

			// 解析树
			root := ParseTree(tc.input)

			// 层序遍历
			levels := LevelOrder(root)
			fmt.Printf("层序遍历: %v\n", levels)

			// 右视图
			rightView := RightSideView(root)
			fmt.Printf("右视图: %v\n", rightView)

			// 验证结果
			if !equalIntSliceSlice(levels, tc.wantLevel) {
				t.Errorf("层序遍历错误, got %v, want %v", levels, tc.wantLevel)
			}
			if !equalIntSlice(rightView, tc.wantRight) {
				t.Errorf("右视图错误, got %v, want %v", rightView, tc.wantRight)
			}
		})
	}
}

// 辅助函数：比较两个一维整数切片
func equalIntSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// 辅助函数：比较两个二维整数切片
func equalIntSliceSlice(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if !equalIntSlice(a[i], b[i]) {
			return false
		}
	}
	return true
}

func TestRightSideViewDirect(t *testing.T) {
	testCases := []struct {
		name      string
		input     string
		wantRight []int
	}{
		{
			name:      "测试1: (1(2(7,8),3))",
			input:     "(1(2(7,8),3))",
			wantRight: []int{1, 3, 8},
		},
		{
			name:      "测试2: (1(2(7,8),4(5)))",
			input:     "(1(2(7,8),4(5)))",
			wantRight: []int{1, 4, 5},
		},
		{
			name:      "测试3: (1(2,3(4,5)))",
			input:     "(1(2,3(4,5)))",
			wantRight: []int{1, 3, 5},
		},
		{
			name:      "测试4: 单节点 (1)",
			input:     "(1)",
			wantRight: []int{1},
		},
		{
			name:      "测试5: (1(2,3))",
			input:     "(1(2,3))",
			wantRight: []int{1, 3},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fmt.Printf("\n输入: %s\n", tc.input)

			// 直接解析右视图（不构建树）
			rightView := RightSideViewDirect(tc.input)
			fmt.Printf("右视图（直接解析）: %v\n", rightView)

			// 验证结果
			if !equalIntSlice(rightView, tc.wantRight) {
				t.Errorf("右视图错误, got %v, want %v", rightView, tc.wantRight)
			}
		})
	}
}
