package aim_offer

import "testing"

func Test_dup(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				nums: []int{2, 3, 1, 0, 2, 5, 3},
			},
			want: []int{2, 3},
		},
		{
			name: "2",
			args: args{
				nums: []int{1, 0, 2, 3, 4, 5},
			},
			want: []int{-1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := dup(tt.args.nums)
			flag := false
			for _, w := range tt.want {
				if w == got {
					flag = true
					return
				}
			}
			if !flag {
				t.Error(got)
			}
		})
	}
}

func Test_dup2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				nums: []int{2, 3, 1, 0, 2, 5, 3},
			},
			want: []int{2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := dup(tt.args.nums)
			flag := false
			for _, w := range tt.want {
				if w == got {
					flag = true
					return
				}
			}
			if !flag {
				t.Error(got)
			}
		})
	}
}

func Test_findTargetIn2DPlants(t *testing.T) {
	type args struct {
		plants [][]int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "exist",
			args: args{
				plants: [][]int{
					{2, 3, 6, 8},
					{4, 5, 8, 9},
					{5, 9, 10, 12},
				},
				target: 8,
			},
			want: true,
		},
		{
			name: "no-exist",
			args: args{
				plants: [][]int{
					{1, 3, 5},
					{2, 5, 7},
				},
				target: 4,
			},
			want: false,
		},
		{
			name: "no-exist",
			args: args{
				plants: [][]int{
					{-5},
				},
				target: -2,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTargetIn2DPlants(tt.args.plants, tt.args.target); got != tt.want {
				t.Errorf("findTargetIn2DPlants() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestSubsequence(t *testing.T) {
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
			name: "1",
			args: args{
				s: "1001010",
				k: 5,
			},
			want: 5,
		},
		{
			name: "1",
			args: args{
				s: "00101001",
				k: 1,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestSubsequence(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("longestSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
