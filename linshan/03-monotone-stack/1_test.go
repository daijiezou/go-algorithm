package _3_monotone_stack

import (
	"reflect"
	"testing"
)

func Test_dailyTemperatures(t *testing.T) {
	type args struct {
		temperatures []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case1",
			args: args{
				temperatures: []int{73, 74, 75, 71, 69, 72, 76, 73},
			},
			want: []int{1, 1, 4, 2, 1, 1, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dailyTemperatures(tt.args.temperatures); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dailyTemperatures() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_finalPrices(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case1",
			args: args{
				prices: []int{8, 4, 6, 2, 3},
			},
			want: []int{4, 2, 4, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := finalPrices2(tt.args.prices); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("finalPrices() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nextGreaterElement(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case1",
			args: args{
				nums1: []int{4, 1, 2},
				nums2: []int{1, 3, 4, 2},
			},
			want: []int{-1, 3, -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextGreaterElement(tt.args.nums1, tt.args.nums2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nextGreaterElement() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nextGreaterElements2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case1",
			args: args{
				nums: []int{1, 2, 1},
			},
			want: []int{2, -1, 2},
		},
		{
			name: "case2",
			args: args{
				nums: []int{1, 2, 3, 4, 3},
			},
			want: []int{2, 3, 4, -1, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextGreaterElements2(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nextGreaterElements2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxWidthRamp(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				nums: []int{6, 0, 8, 2, 1, 5},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxWidthRamp(tt.args.nums); got != tt.want {
				t.Errorf("maxWidthRamp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maximumScore(t *testing.T) {
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
			name: "case1",
			args: args{
				nums: []int{1, 4, 3, 7, 4, 5},
				k:    3,
			},
			want: 15,
		},
		{
			name: "case2",
			args: args{
				nums: []int{6569, 9667, 3148, 7698, 1622, 2194, 793, 9041, 1670, 1872},
				k:    5,
			},
			want: 9732,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumScore3(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("maximumScore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_largestRectangleArea(t *testing.T) {
	type args struct {
		heights []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				heights: []int{2, 1, 5, 6, 2, 3},
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestRectangleArea(tt.args.heights); got != tt.want {
				t.Errorf("largestRectangleArea() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_trap(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trap3(tt.args.height); got != tt.want {
				t.Errorf("trap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_subArrayRanges(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "case1",
			args: args{
				nums: []int{1, 2, 3},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subArrayRanges(tt.args.nums); got != tt.want {
				t.Errorf("subArrayRanges() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxSumMinProduct(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				nums: []int{1, 2, 3, 2},
			},
			want: 14,
		},
		{
			name: "case2",
			args: args{
				nums: []int{3, 1, 5, 6, 4, 2},
			},
			want: 60,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSumMinProduct(tt.args.nums); got != tt.want {
				t.Errorf("maxSumMinProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeKdigits(t *testing.T) {
	type args struct {
		num string
		k   int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case1",
			args: args{
				num: "1432219",
				k:   3,
			},
			want: "1219",
		},
		{
			name: "case2",
			args: args{
				num: "10200",
				k:   1,
			},
			want: "200",
		},
		{
			name: "10",
			args: args{
				num: "10",
				k:   2,
			},
			want: "0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeKdigits(tt.args.num, tt.args.k); got != tt.want {
				t.Errorf("removeKdigits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestWPI(t *testing.T) {
	type args struct {
		hours []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				hours: []int{6, 9, 6},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestWPI(tt.args.hours); got != tt.want {
				t.Errorf("longestWPI() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_find132pattern(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case1",
			args: args{
				nums: []int{1, 2, 3},
			},
			want: false,
		},
		{
			name: "case2",
			args: args{
				nums: []int{3, 1, 4, 2},
			},
			want: true,
		},
		{
			name: "case3",
			args: args{
				nums: []int{-1, 3, 2, 0},
			},
			want: true,
		},
		{
			name: "case4",
			args: args{
				nums: []int{-2, 1, 1},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := find132pattern(tt.args.nums); got != tt.want {
				t.Errorf("find132pattern() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mostCompetitive(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case1",
			args: args{
				nums: []int{2, 4, 3, 3, 5, 4, 9, 6},
				k:    4,
			},
			want: []int{2, 3, 3, 4},
		},
		{
			name: "case2",
			args: args{
				nums: []int{3, 5, 2, 6},
				k:    2,
			},
			want: []int{2, 6},
		},
		{
			name: "case3",
			args: args{
				nums: []int{71, 18, 52, 29, 55, 73, 24, 42, 66, 8, 80, 2},
				k:    3,
			},
			want: []int{8, 80, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mostCompetitive(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mostCompetitive() = %v, want %v", got, tt.want)
			}
		})
	}
}
