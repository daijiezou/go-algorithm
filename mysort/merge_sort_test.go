package mysort

import (
	"reflect"
	"testing"
)

func Test_mergeSort(t *testing.T) {
	type args struct {
		sumList []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{
				sumList: []int{5, 2, 1, 4, 10, 11, 13, 0, 0, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeSort(tt.args.sumList); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sort(t *testing.T) {
	type args struct {
		nums []int
		lo   int
		hi   int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test1",
			args: args{
				nums: []int{5, 2, 1, 4, 10, 11, 13, 0, 0, 0},
				lo:   0,
				hi:   9,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sort(tt.args.nums, tt.args.lo, tt.args.hi)
		})
	}
}

func Test_mergeSort3(t *testing.T) {
	type args struct {
		sumList []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{
				sumList: []int{5, 2, 1, 4, 10, 11, 13, 0, 0, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeSort3(tt.args.sumList); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeSort3() = %v, want %v", got, tt.want)
			}
		})
	}
}
