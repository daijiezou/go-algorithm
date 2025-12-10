package _4_datastruct

import (
	"reflect"
	"testing"
)

func Test_subarraySum(t *testing.T) {
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
				nums: []int{2, 3, 1},
			},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subarraySum(tt.args.nums); got != tt.want {
				t.Errorf("subarraySum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isArraySpecial(t *testing.T) {
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
			name: "1",
			args: args{
				nums: []int{4, 3, 1, 6},
				queries: [][]int{
					[]int{0, 2},
					{2, 3},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isArraySpecial(tt.args.nums, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("isArraySpecial() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxAbsoluteSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "common",
			args: args{nums: []int{2, -5, 1, -4, 3, -2}},
			want: 8,
		},
		{
			name: "common",
			args: args{nums: []int{1, -3, 2, 3, -4}},
			want: 5,
		},
		{
			name: "common3",
			args: args{nums: []int{-1}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxAbsoluteSum(tt.args.nums); got != tt.want {
				t.Errorf("maxAbsoluteSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
