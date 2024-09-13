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
