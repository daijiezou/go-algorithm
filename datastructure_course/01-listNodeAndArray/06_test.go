package _1_listNodeAndArray

import (
	"reflect"
	"testing"
)

func Test_leftBound(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				nums:   []int{1, 2, 3, 3, 3, 5, 7},
				target: 3,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := leftBound(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("leftBound() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rightBound(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				nums:   []int{1, 2, 3, 3, 3, 5, 7},
				target: 3,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rightBound(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("rightBound() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findClosestElements(t *testing.T) {
	type args struct {
		arr []int
		k   int
		x   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				arr: []int{0, 0, 1, 2, 3, 3, 4, 7, 7, 8},
				k:   3,
				x:   5,
			},
			want: []int{3, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findClosestElements(tt.args.arr, tt.args.k, tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findClosestElements() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findPeakElement(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				nums: []int{1, 2, 1, 3, 5, 6, 4},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findPeakElement(tt.args.nums)
			t.Log(got)
		})
	}
}
