package _1_listNodeAndArray

import (
	"reflect"
	"testing"
)

func Test_myMin(t *testing.T) {
	type args struct {
		a int
		b int
		c int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				a: 10,
				b: 12,
				c: 10,
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myMin(tt.args.a, tt.args.b, tt.args.c); got != tt.want {
				t.Errorf("myMin() = %v, want %v", got, tt.want)
			}
		})
	}
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
			want: 12,
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

func Test_merge2SmallestPairs(t *testing.T) {
	type args struct {
		shudui1 [][]int
		shudui2 [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "1",
			args: args{
				shudui1: [][]int{{1, 2}, {1, 3}},
				shudui2: [][]int{{1, 1}, {1, 4}},
			},
			want: [][]int{{1, 1}, {1, 2}, {1, 3}, {3, 4}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge2SmallestPairs(tt.args.shudui1, tt.args.shudui2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge2SmallestPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mergeNum(t *testing.T) {
	type args struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "1",
			args: args{
				nums1: []int{1, 2, 3, 0, 0, 0},
				m:     3,
				nums2: []int{2, 5, 6},
				n:     3,
			},
		},
		{
			name: "2",
			args: args{
				nums1: []int{0},
				m:     0,
				nums2: []int{1},
				n:     1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mergeNum(tt.args.nums1, tt.args.m, tt.args.nums2, tt.args.n)
		})
	}
}

func Test_sortTransformedArray(t *testing.T) {
	type args struct {
		nums []int
		a    int
		b    int
		c    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				nums: []int{-4, -2, 2, 4},
				a:    1,
				b:    3,
				c:    5,
			},
			want: []int{3, 9, 15, 33},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortTransformedArray(tt.args.nums, tt.args.a, tt.args.b, tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortTransformedArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseWords(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{s: " a good   example "},
			want: "example good a",
		},
		{
			name: "2",
			args: args{s: "  hello world  "},
			want: "world hello",
		},
		{
			name: "3",
			args: args{s: "  Bob    Loves  Alice   "},
			want: "Alice Loves Bob",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseWords(tt.args.s); got != tt.want {
				t.Errorf("reverseWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
