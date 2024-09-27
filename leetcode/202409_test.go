package leetcode

import "testing"

func Test_maxNumOfMarkedIndices(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		wantRes int
	}{
		{
			name: "case 1",
			args: args{
				nums: []int{42, 83, 48, 10, 24, 55, 9, 100, 10, 17, 17, 99, 51, 32, 16, 98, 99, 31, 28, 68, 71, 14, 64, 29, 15, 40},
			},
			wantRes: 26,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := maxNumOfMarkedIndices(tt.args.nums); gotRes != tt.wantRes {
				t.Errorf("maxNumOfMarkedIndices() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func Test_maximumRobots(t *testing.T) {
	type args struct {
		chargeTimes  []int
		runningCosts []int
		budget       int64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				chargeTimes:  []int{3, 6, 1, 3, 4},
				runningCosts: []int{2, 1, 3, 4, 5},
				budget:       25,
			},
			want: 3,
		},
		{
			name: "case 2",
			args: args{
				chargeTimes:  []int{11, 12, 19},
				runningCosts: []int{10, 8, 7},
				budget:       19,
			},
			want: 0,
		},
		{
			name: "case 3",
			args: args{
				chargeTimes:  []int{11, 12, 74, 67, 37, 87, 42, 34, 18, 90, 36, 28, 34, 20},
				runningCosts: []int{18, 98, 2, 84, 7, 57, 54, 65, 59, 91, 7, 23, 94, 20},
				budget:       937,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumRobots(tt.args.chargeTimes, tt.args.runningCosts, tt.args.budget); got != tt.want {
				t.Errorf("maximumRobots() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_latestTimeCatchTheBus(t *testing.T) {
	type args struct {
		buses      []int
		passengers []int
		capacity   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				buses:      []int{10, 20},
				passengers: []int{2, 17, 18, 19},
				capacity:   2,
			},
			want: 16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := latestTimeCatchTheBus(tt.args.buses, tt.args.passengers, tt.args.capacity); got != tt.want {
				t.Errorf("latestTimeCatchTheBus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxScoreSightseeingPair(t *testing.T) {
	type args struct {
		values []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				values: []int{8, 1, 5, 2, 6},
			},
			want: 11,
		},
		{
			name: "case 2",
			args: args{
				values: []int{1, 2},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxScoreSightseeingPair(tt.args.values); got != tt.want {
				t.Errorf("maxScoreSightseeingPair() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_takeCharacters(t *testing.T) {
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
			name: "case 1",
			args: args{
				s: "aabaaaacaabc",
				k: 2,
			},
			want: 8,
		},
		{
			name: "case 2",
			args: args{
				s: "a",
				k: 1,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := takeCharacters2(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("takeCharacters() = %v, want %v", got, tt.want)
			}
		})
	}
}
