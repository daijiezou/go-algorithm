package _1_huadongchuangkou

import "testing"

func Test_maxVowels(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				s: "abciiidef",
				k: 3,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxVowels(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("maxVowels() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minimumDifference(t *testing.T) {
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
			name: "test1",
			args: args{
				nums: []int{
					9, 4, 1, 7,
				},
				k: 2,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumDifference(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("minimumDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findMaxAverage(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "test1",
			args: args{
				nums: []int{1, 12, -5, -6, 50, 3},
				k:    4,
			},
			want: 12.75,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxAverage2(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("findMaxAverage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxSatisfied(t *testing.T) {
	type args struct {
		customers []int
		grumpy    []int
		minutes   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				customers: []int{1, 0, 1, 2, 1, 1, 7, 5},
				grumpy:    []int{0, 1, 0, 1, 0, 1, 0, 1},
				minutes:   3,
			},
			want: 16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSatisfied(tt.args.customers, tt.args.grumpy, tt.args.minutes); got != tt.want {
				t.Errorf("maxSatisfied() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxSum(t *testing.T) {
	type args struct {
		nums []int
		m    int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "test1",
			args: args{
				nums: []int{1, 2, 2},
				m:    2,
				k:    2,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSum(tt.args.nums, tt.args.m, tt.args.k); got != tt.want {
				t.Errorf("maxSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
