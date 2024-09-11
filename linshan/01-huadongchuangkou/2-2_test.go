package _1_huadongchuangkou

import (
	"reflect"
	"testing"
)

func Test_smallestRange(t *testing.T) {
	type args struct {
		nums [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case 1",
			args: args{
				nums: [][]int{{1, 3, 5, 7, 9, 10}, {2, 4, 6, 8, 10}},
			},
			want: []int{10, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestRange(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("smallestRange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkInclusion(t *testing.T) {
	type args struct {
		t string
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{
				t: "ABC",
				s: "ADOBECODEBANC",
			},
			want: "BANC",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minWindow(tt.args.t, tt.args.s); got != tt.want {
				t.Errorf("minWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}
