package _3_array

import (
	"reflect"
	"testing"
)

func Test_twoSumTarget(t *testing.T) {
	type args struct {
		nums   []int
		target int
		start  int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "case 1",
			args: args{
				nums:   []int{1, 1, 1, 2, 2, 3, 3},
				target: 4,
				start:  0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSumTarget(tt.args.nums, tt.args.start, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSumTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_threeSumTarget(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "case 1",
			args: args{
				nums:   []int{1, 2, 3},
				target: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSumTarget(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("threeSumTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}
