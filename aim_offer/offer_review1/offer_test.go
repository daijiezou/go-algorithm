package offer_review1

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_hasPath(t *testing.T) {
	type args struct {
		matrix [][]byte
		word   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				matrix: [][]byte{
					{'a', 'b', 'c', 'e'},
					{'s', 'f', 'c', 's'},
					{'a', 'd', 'e', 'e'},
				},
				word: "abcced",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasPath(tt.args.matrix, tt.args.word); got != tt.want {
				t.Errorf("hasPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_movingCount(t *testing.T) {
	type args struct {
		threshold int
		rows      int
		cols      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				threshold: 5,
				rows:      10,
				cols:      10,
			},
			want: 21,
		},
		{
			name: "1",
			args: args{
				threshold: 10,
				rows:      1,
				cols:      100,
			},
			want: 29,
		},
		{
			name: "1",
			args: args{
				threshold: 0,
				rows:      1,
				cols:      3,
			},
			want: 1,
		},
		{
			name: "1",
			args: args{
				threshold: 1,
				rows:      2,
				cols:      3,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := movingCount(tt.args.threshold, tt.args.rows, tt.args.cols); got != tt.want {
				t.Errorf("movingCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pow(t *testing.T) {
	type args struct {
		x float64
		n int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "1",
			args: args{
				x: 3,
				n: 4,
			},
			want: 81.0,
		},
		{
			name: "1",
			args: args{
				x: 2,
				n: 3,
			},
			want: 8.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pow(tt.args.x, tt.args.n); got != tt.want {
				t.Errorf("pow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_printMatrix(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{matrix: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
				{13, 14, 15, 16},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := printMatrix2(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("printMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsPopOrder(t *testing.T) {
	type args struct {
		pushV []int
		popV  []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				pushV: []int{2, 1, 0},
				popV:  []int{1, 2, 0},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPopOrder(tt.args.pushV, tt.args.popV); got != tt.want {
				t.Errorf("IsPopOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVerifySquenceOfBST(t *testing.T) {
	type args struct {
		sequence []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{sequence: []int{4, 6, 7, 5}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := VerifySquenceOfBST(tt.args.sequence); got != tt.want {
				t.Errorf("VerifySquenceOfBST() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindPath(t *testing.T) {
	root := &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val:   5,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 7},
		},
		Right: &TreeNode{Val: 12},
	}

	got := FindPath(root, 22)

	// 说明：FindPath 返回的路径顺序和内部顺序，可能与期望不一致。
	// 如果你严格比较，可以根据返回顺序调整期望；下面给一个常见的期望写法：
	want1 := [][]int{{10, 12}, {10, 5, 7}}
	want2 := [][]int{{10, 5, 7}, {10, 12}}

	if !reflect.DeepEqual(got, want1) && !reflect.DeepEqual(got, want2) {
		t.Fatalf("FindPath() = %#v, want %#v or %#v", got, want1, want2)
	}
}

func TestPermutation(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "1",
			args: args{str: "aab"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Permutation(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Permutation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindGreatestSumOfSubArray(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{array: []int{1, -2, 3, 10, -4, 7, 2, -5}},
			want: 18,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindGreatestSumOfSubArray(tt.args.array); got != tt.want {
				t.Errorf("FindGreatestSumOfSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findNthDigit(t *testing.T) {
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
			want: 1,
		},
		{
			name: "1",
			args: args{n: 13},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findNthDigit(tt.args.n); got != tt.want {
				t.Errorf("findNthDigit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNumberOf1Between1AndN_Solution(t *testing.T) {
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
			args: args{n: 233},
			want: 154,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumberOf1Between1AndN_Solution(tt.args.n); got != tt.want {
				t.Errorf("NumberOf1Between1AndN_Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solve(t *testing.T) {
	type args struct {
		nums string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{nums: "12"},
			want: 2,
		},
		{
			name: "1",
			args: args{nums: "31717126241541717"},
			want: 192,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.nums); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxValue(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{grid: [][]int{
				{1, 3, 1},
				{1, 5, 1},
				{4, 2, 1},
			}},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxValue(tt.args.grid); got != tt.want {
				t.Errorf("maxValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{s: "abcabcbb"},
			want: 3,
		},
		{
			name: "1",
			args: args{s: ""},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetUglyNumber_Solution(t *testing.T) {
	type args struct {
		index int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{index: 7},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetUglyNumber_Solution(tt.args.index); got != tt.want {
				t.Errorf("GetUglyNumber_Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetNumberOfK(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				nums: []int{1, 2, 3, 3, 3, 3, 4, 5},
				k:    3,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetNumberOfK(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("GetNumberOfK() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLeftRotateString(t *testing.T) {
	type args struct {
		str string
		n   int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				str: "abcXYZdef",
				n:   3,
			},
			want: "XYZdefabc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LeftRotateString(tt.args.str, tt.args.n); got != tt.want {
				t.Errorf("LeftRotateString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverse(t *testing.T) {
	type args struct {
		strs  []byte
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "1",
			args: args{
				strs:  []byte{'a', 'b', 'c'},
				left:  0,
				right: 2,
			},
		},
		{
			name: "1",
			args: args{
				strs: []byte{'a', 'b', 'c', 'x', 'y', 'z'},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reverse(tt.args.strs)
			fmt.Println(string(tt.args.strs))
		})
	}
}

func Test_maxInWindows(t *testing.T) {
	type args struct {
		num  []int
		size int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				num:  []int{2, 3, 4, 2, 6, 2, 5, 1},
				size: 3,
			},
			want: []int{4, 4, 6, 6, 6, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxInWindows(tt.args.num, tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxInWindows() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLastRemaining_Solution(t *testing.T) {
	type args struct {
		n int
		m int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				n: 4,
				m: 3,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LastRemaining_Solution(tt.args.n, tt.args.m); got != tt.want {
				t.Errorf("LastRemaining_Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLastRemaining_Debug(t *testing.T) {
	t.Run("n=5, m=3", func(t *testing.T) {
		fmt.Println("--- Running test for n=5, m=3 ---")
		LastRemaining_Debug(5, 3)
		fmt.Println("---------------------------------")
	})

	t.Run("n=6, m=4", func(t *testing.T) {
		fmt.Println("\n--- Running test for n=6, m=4 ---")
		LastRemaining_Debug(6, 4)
		fmt.Println("---------------------------------")
	})
}

func Test_maxProfit(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{prices: []int{8, 9, 2, 5, 4, 7, 1}},
			want: 5,
		},
		{
			name: "2",
			args: args{prices: []int{2, 4, 1}},
			want: 2,
		},
		{
			name: "3",
			args: args{prices: []int{3, 2, 1}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStrToInt(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "basic positive",
			args: args{str: "42"},
			want: 42,
		},
		{
			name: "leading whitespace and sign",
			args: args{str: "   -42"},
			want: -42,
		},
		{
			name: "with trailing words",
			args: args{str: "4193 with words"},
			want: 4193,
		},
		{
			name: "starts with words",
			args: args{str: "words and 987"},
			want: 0,
		},
		{
			name: "positive overflow",
			args: args{str: "2147483648"},
			want: 2147483647, // math.MaxInt32
		},
		{
			name: "negative overflow",
			args: args{str: "-2147483649"},
			want: -2147483648, // math.MinInt32
		},
		{
			name: "empty string",
			args: args{str: ""},
			want: 0,
		},
		{
			name: "only whitespace",
			args: args{str: "   "},
			want: 0,
		},
		{
			name: "plus sign only",
			args: args{str: "+"},
			want: 0,
		},
		{
			name: "just a number",
			args: args{str: "-91283472332"},
			want: -2147483648,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StrToInt(tt.args.str); got != tt.want {
				t.Errorf("StrToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseSentence(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{str: "nowcoder. a am I"},
			want: "I am a nowcoder.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseSentence(tt.args.str); got != tt.want {
				t.Errorf("ReverseSentence() = %v, want %v", got, tt.want)
			}
		})
	}
}
