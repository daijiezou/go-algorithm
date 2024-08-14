package leetcode

import (
	"reflect"
	"testing"
)

func Test_numberOfRightTriangles(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
		{
			name: "case 1",
			args: args{
				grid: [][]int{
					{0, 1, 0},
					{0, 1, 1},
					{0, 1, 0},
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfRightTriangles(tt.args.grid); got != tt.want {
				t.Errorf("numberOfRightTriangles() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxPointsInsideSquare(t *testing.T) {
	type args struct {
		points [][]int
		s      string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				points: [][]int{
					{2, 2}, {-1, -2}, {-4, 4}, {-3, 1}, {3, -3},
				},
				s: "abdca",
			},
			want: 2,
		},
		{
			name: "case 1",
			args: args{
				points: [][]int{
					{-35, -3}, {17, 28}, {28, -28}, {25, -1}, {25, -16}, {1, -21},
				},
				s: "ffcbea",
			},
			want: 5,
		},
		{
			name: "case 3",
			args: args{
				points: [][]int{
					{-1, 0},
				},
				s: "a",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPointsInsideSquare2(tt.args.points, tt.args.s); got != tt.want {
				t.Errorf("maxPointsInsideSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findIntegers(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				n: 5,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findIntegers(tt.args.n); got != tt.want {
				t.Errorf("findIntegers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minimumAddedInteger(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				nums1: []int{4, 20, 16, 12, 8},
				nums2: []int{14, 18, 10},
			},
			want: -2,
		},
		{
			name: "case 2",
			args: args{
				nums1: []int{4, 6, 3, 1, 4, 2, 10, 9, 5},
				nums2: []int{5, 10, 3, 2, 6, 1, 9},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumAddedInteger2(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("minimumAddedInteger() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxUncrossedLines(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				nums1: []int{2, 5, 1, 2, 5},
				nums2: []int{10, 5, 2, 1, 5, 2},
			},
			want: 3,
		},
		{
			name: "case 2",
			args: args{
				nums1: []int{1, 3, 7, 1, 7, 5},
				nums2: []int{1, 9, 2, 5, 1},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxUncrossedLines2(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("maxUncrossedLines() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMagicDict(t *testing.T) {
	type args struct {
		req1 []string
		req2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case 1",
			args: args{
				req1: []string{"hello", "leetcode"},
				req2: "hhllo",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MagicDict(tt.args.req1, tt.args.req2); got != tt.want {
				t.Errorf("MagicDict() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isArraySpecial2(t *testing.T) {
	type args struct {
		nums    []int
		queries [][]int
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		{
			name: "case 1",
			args: args{
				nums:    []int{1, 1},
				queries: [][]int{{0, 1}},
			},
			want: []bool{false},
		},
		{
			name: "case 2",
			args: args{
				nums:    []int{2, 8, 10, 9},
				queries: [][]int{{1, 3}},
			},
			want: []bool{false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isArraySpecial2(tt.args.nums, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("isArraySpecial2() = %v, want %v", got, tt.want)
			}
		})
	}
}
