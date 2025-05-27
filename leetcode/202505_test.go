package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_pushDominoes(t *testing.T) {
	type args struct {
		dominoes string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				dominoes: ".L.R...LR..L..",
			},
			want: "LL.RR.LLRRLL..",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pushDominoes(tt.args.dominoes); got != tt.want {
				t.Errorf("pushDominoes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minDominoRotations(t *testing.T) {
	type args struct {
		tops    []int
		bottoms []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "!",
			args: args{
				tops:    []int{2, 1, 2, 4, 2, 2},
				bottoms: []int{5, 2, 6, 2, 3, 2},
			},
			want: 2,
		},
		{
			name: "!",
			args: args{
				tops:    []int{3, 5, 1, 2, 3},
				bottoms: []int{3, 6, 3, 3, 4},
			},
			want: -1,
		},
		{
			name: "!",
			args: args{
				tops:    []int{1, 2, 1, 1, 1, 2, 2, 2},
				bottoms: []int{2, 1, 2, 2, 2, 2, 2, 2},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDominoRotations1(tt.args.tops, tt.args.bottoms); got != tt.want {
				t.Errorf("minDominoRotations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minTimeToReach(t *testing.T) {
	type args struct {
		moveTime [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{moveTime: [][]int{
				{0, 4},
				{4, 4},
			}},
			want: 6,
		},
		{
			name: "2",
			args: args{moveTime: [][]int{
				{94, 79, 62, 27, 69, 84},
				{6, 32, 11, 82, 42, 30},
			}},
			want: 74,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minTimeToReach(tt.args.moveTime); got != tt.want {
				t.Errorf("minTimeToReach() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minSum(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			args: args{
				nums1: []int{3, 2, 0, 1, 0},
				nums2: []int{6, 5, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSum(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("minSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findEvenNumbers(t *testing.T) {
	type args struct {
		digits []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{digits: []int{2, 2, 8, 8, 2}},
			want: []int{222, 228, 282, 288, 822, 828, 882},
		},
		{
			name: "1",
			args: args{digits: []int{2, 1, 3, 0}},
			want: []int{102, 120, 130, 132, 210, 230, 302, 310, 312, 320},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findEvenNumbers2(tt.args.digits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findEvenNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkIsPalindrome(b *testing.B) {

	for i := 0; i < b.N; i++ {
		findEvenNumbers([]int{2, 2, 8, 8, 2})
	}
}

func Test_sortColors(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "1",
			args: args{
				nums: []int{2, 0, 2, 1, 1, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortColors(tt.args.nums)
		})
		fmt.Println(tt.args.nums)
	}
}

func Test_isZeroArray(t *testing.T) {
	type args struct {
		nums    []int
		queries [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				nums: []int{1, 0, 1},
				queries: [][]int{
					{0, 2},
				},
			},
			want: true,
		},
		{
			name: "2",
			args: args{
				nums: []int{4, 3, 2, 1},
				queries: [][]int{
					{1, 3},
					{0, 2},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isZeroArray(tt.args.nums, tt.args.queries); got != tt.want {
				t.Errorf("isZeroArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minZeroArray(t *testing.T) {
	type args struct {
		nums    []int
		queries [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				nums: []int{2, 0, 2},
				queries: [][]int{
					{0, 2, 1},
					{0, 2, 1},
					{1, 1, 3}},
			},
			want: 2,
		},
		{
			name: "1",
			args: args{
				nums: []int{4, 3, 2, 1},
				queries: [][]int{
					{1, 3, 2},
					{0, 2, 1}},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minZeroArray(tt.args.nums, tt.args.queries); got != tt.want {
				t.Errorf("minZeroArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				words: []string{"lc", "cl", "gg"},
			},
			want: 6,
		},
		{
			name: "1",
			args: args{
				words: []string{"ab", "ty", "yt", "lc", "cl", "ab"},
			},
			want: 8,
		},
		{
			name: "1",
			args: args{
				words: []string{"cc", "ll", "xx"},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome2(tt.args.words); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
