package aim_offer

import (
	"reflect"
	"testing"
)

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

func Test_replaceSpace(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{s: "We Are Happy"},
			want: "We%20Are%20Happy",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := replaceSpace(tt.args.s); got != tt.want {
				t.Errorf("replaceSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_check(t *testing.T) {
	type args struct {
		threshold int
		row       int
		col       int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				threshold: 100,
				row:       22,
				col:       33,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := check(tt.args.threshold, tt.args.row, tt.args.col); got != tt.want {
				t.Errorf("check() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_movingCount(t *testing.T) {
	type args struct {
		threshold int
		rows      int
		cols      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				threshold: 10,
				rows:      1,
				cols:      100,
			},
			want: 29,
		},
		{
			name: "1",
			args: args{
				threshold: 15,
				rows:      20,
				cols:      20,
			},
			want: 29,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := movingCount(tt.args.threshold, tt.args.rows, tt.args.cols); got != tt.want {
				t.Errorf("movingCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cutRope(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				n: 8,
			},
			want: 18,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cutRope(tt.args.n); got != tt.want {
				t.Errorf("cutRope() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNumberOf1(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{n: 10},
			want: 2,
		},
		{
			name: "1",
			args: args{n: -11},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumberOf1(tt.args.n); got != tt.want {
				t.Errorf("NumberOf1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_printNumbers(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "1",
			args: args{n: 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := printNumbers(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("printNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
