package _1_listNodeAndArray

import (
	"math/rand"
	"testing"
)

func Test_subarraySum(t *testing.T) {
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
				nums: []int{1, 1, 1},
				k:    2,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subarraySum(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("subarraySum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxSubArrayLen(t *testing.T) {
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
				nums: []int{1, -1, 5, -2, 3},
				k:    3,
			},
			want: 4,
		},
		{
			name: "2",
			args: args{
				nums: []int{1, -1, 5, -2, 3, -1, 1, -1, 1},
				k:    3,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxSubArrayLen(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("MaxSubArrayLen() = %v, want %v", got, tt.want)
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
			name: "1",
			args: args{hours: []int{6, 6, 9}},
			want: 1,
		},
		{
			name: "2",
			args: args{hours: []int{9, 9, 6, 0, 6, 6, 9}},
			want: 3,
		},
		{
			name: "2",
			args: args{hours: []int{6, 9, 9}},
			want: 3,
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

func TestConstructor1(t *testing.T) {
	N := 2000
	Max := 1000000
	nums := make([]int, 0)
	for i := 0; i < N; i++ {
		nums = append(nums, rand.Intn(Max)-Max/2)
	}
}
