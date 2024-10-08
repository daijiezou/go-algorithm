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

func Test_getTotalTrips(t *testing.T) {
	type args struct {
		time []int
		t    int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "case 1",
			args: args{
				time: []int{1, 2, 3},
				t:    3,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getTotalTrips(tt.args.time, tt.args.t); got != tt.want {
				t.Errorf("getTotalTrips() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minRefuelStops(t *testing.T) {
	type args struct {
		target    int
		startFuel int
		stations  [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{

		{
			name: "1",
			args: args{
				target:    100,
				startFuel: 10,
				stations: [][]int{
					{10, 60},
					{20, 30},
					{30, 30},
					{60, 40},
				},
			},
			want: 2,
		},
		{
			name: "2",
			args: args{
				target:    100,
				startFuel: 25,
				stations: [][]int{
					{25, 25},
					{50, 25},
					{75, 25},
				},
			},
			want: 3,
		},
		{
			name: "3",
			args: args{
				target:    1000000000,
				startFuel: 26136899,
				stations:  [][]int{{17654548, 460787121}, {67802923, 34444712}, {243977947, 259740557}, {438730568, 343225863}, {574211102, 423090989}, {577337718, 237883992}, {780977723, 314461540}, {848603056, 144394709}, {881449326, 364937682}, {918179140, 289252804}},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minRefuelStops(tt.args.target, tt.args.startFuel, tt.args.stations); got != tt.want {
				t.Errorf("minRefuelStops() = %v, want %v", got, tt.want)
			}
		})
	}
}
