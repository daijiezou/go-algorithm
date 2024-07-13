package leetcode

import "testing"

func Test_incremovableSubarrayCount(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				nums: []int{6, 5, 7, 8},
			},
			want: 7,
		},
		{
			name: "case 2",
			args: args{
				nums: []int{8, 7, 6, 6},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := incremovableSubarrayCount(tt.args.nums); got != tt.want {
				t.Errorf("incremovableSubarrayCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_canSortArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case 1",
			args: args{
				nums: []int{8, 4, 2, 30, 15},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canSortArray(tt.args.nums); got != tt.want {
				t.Errorf("canSortArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
