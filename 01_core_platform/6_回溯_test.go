package _1_core_platform

import (
	"reflect"
	"testing"
)

func Test_subsets2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "1",
			args: args{nums: []int{1, 2, 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subsets2(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("subsets2() = %v, want %v", got, tt.want)
			}
		})
	}
}
