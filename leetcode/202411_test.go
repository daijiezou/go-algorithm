package leetcode

import (
	"fmt"
	"testing"
)

func Test_minChanges(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				n: 13,
				k: 4,
			},
			want: 2,
		},
		{
			name: "case2",
			args: args{
				n: 14,
				k: 13,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minChanges(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("minChanges() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_shoppingOffers(t *testing.T) {
	type args struct {
		price   []int
		special [][]int
		needs   []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				price:   []int{2, 5},
				special: [][]int{{3, 0, 5}, {1, 2, 10}},
				needs:   []int{3, 2},
			},
			want: 14,
		},
		{
			name: "case1",
			args: args{
				price:   []int{2, 3, 4},
				special: [][]int{{1, 1, 0, 4}, {2, 2, 1, 9}},
				needs:   []int{1, 2, 1},
			},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shoppingOffers(tt.args.price, tt.args.special, tt.args.needs); got != tt.want {
				t.Errorf("shoppingOffers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConstructor2(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want NeighborSum
	}{
		{
			name: "case1",
			args: args{
				grid: [][]int{{0, 1, 2}, {3, 4, 5}, {6, 7, 8}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ns := Constructor2(tt.args.grid)
			fmt.Println(ns)
			fmt.Println(ns.AdjacentSum(1))
		})
	}
}

func Test_minCost2(t *testing.T) {
	type args struct {
		n    int
		cuts []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				n:    7,
				cuts: []int{1, 3, 4, 5},
			},
			want: 16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCost2(tt.args.n, tt.args.cuts); got != tt.want {
				t.Errorf("minCost2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countKConstraintSubstrings(t *testing.T) {
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
			name: "case1",
			args: args{
				s: "1010101",
				k: 2,
			},
			want: 25,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countKConstraintSubstrings(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("countKConstraintSubstrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
