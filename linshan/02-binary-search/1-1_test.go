package _2_binary_search

import "testing"

func TestFirstMax(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "case 1",
			args: args{
				nums:   []int{1, 3, 5, 5, 6, 6, 6, 7, 8, 10},
				target: 5,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FirstMax(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("FirstMax() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maximumCount(t *testing.T) {
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
				nums: []int{-2, -1, -1, 1, 2, 3},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumCount(tt.args.nums); got != tt.want {
				t.Errorf("maximumCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
