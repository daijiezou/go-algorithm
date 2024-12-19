package main

import (
	"reflect"
	"testing"
)

func Test_nSumTarget(t *testing.T) {
	type args struct {
		nums   []int
		n      int
		start  int
		target int64
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "1",
			args: args{
				nums:   []int{1, 0, -1, 0, -2, 2},
				n:      4,
				start:  0,
				target: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nSumTarget(tt.args.nums, tt.args.n, tt.args.start, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nSumTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}
