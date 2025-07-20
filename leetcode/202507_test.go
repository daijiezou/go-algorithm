package leetcode

import "testing"

func Test_countDays(t *testing.T) {
	type args struct {
		days     int
		meetings [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				days: 10,
				meetings: [][]int{
					{5, 7},
					{1, 3},
					{9, 10},
				},
			},
			want: 2,
		},
		{
			name: "1",
			args: args{
				days: 57,
				meetings: [][]int{
					{3, 49},
					{23, 44},
					{21, 56},
					{26, 55},
					{23, 52},
					{2, 9},
					{1, 48},
					{3, 31},
				},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countDays(tt.args.days, tt.args.meetings); got != tt.want {
				t.Errorf("countDays() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_matchPlayersAndTrainers(t *testing.T) {
	type args struct {
		players  []int
		trainers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				players:  []int{4, 7, 9},
				trainers: []int{8, 2, 5, 8},
			},
			want: 2,
		},
		{
			name: "1",
			args: args{
				players:  []int{1, 1, 1},
				trainers: []int{10},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := matchPlayersAndTrainers(tt.args.players, tt.args.trainers); got != tt.want {
				t.Errorf("matchPlayersAndTrainers() = %v, want %v", got, tt.want)
			}
		})
	}
}
