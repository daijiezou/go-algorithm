package leetcode

import (
	"fmt"
	"testing"
)

func Test_mincostTickets(t *testing.T) {
	type args struct {
		days  []int
		costs []int
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
				days:  []int{1, 4, 6, 7, 8, 20},
				costs: []int{2, 7, 15},
			},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mincostTickets(tt.args.days, tt.args.costs); got != tt.want {
				t.Errorf("mincostTickets() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minSpeedOnTime(t *testing.T) {
	fmt.Println(1 / 3)
	type args struct {
		dist []int
		hour float64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				dist: []int{1, 3, 2},
				hour: 2.7,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSpeedOnTime(tt.args.dist, tt.args.hour); got != tt.want {
				t.Errorf("minSpeedOnTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minCost(t *testing.T) {
	type args struct {
		maxTime     int
		edges       [][]int
		passingFees []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				maxTime: 30,
				edges: [][]int{
					{0, 1, 10},
					{1, 2, 10},
					{2, 5, 10},
					{0, 3, 1},
					{3, 4, 10},
					{4, 5, 15},
				},
				passingFees: []int{5, 1, 2, 20, 20, 3},
			},
			want: 11,
		},
		{
			name: "case 2",
			args: args{
				maxTime: 29,
				edges: [][]int{
					{0, 1, 10},
					{1, 2, 10},
					{2, 5, 10},
					{0, 3, 1},
					{3, 4, 10},
					{4, 5, 15},
				},
				passingFees: []int{5, 1, 2, 20, 20, 3},
			},
			want: 48,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCost(tt.args.maxTime, tt.args.edges, tt.args.passingFees); got != tt.want {
				t.Errorf("minCost() = %v, want %v", got, tt.want)
			}
		})
	}
}
