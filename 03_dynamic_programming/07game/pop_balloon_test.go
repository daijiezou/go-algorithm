package _7game

import "testing"

func Test_maxCoins(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				nums: []int{3, 1, 5, 8},
			},
			want: 167,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxCoins(tt.args.nums); got != tt.want {
				t.Errorf("maxCoins() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxCoinsBackTrack(t *testing.T) {
	type args struct {
		nums  []int
		score int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxCoinsBackTrack(tt.args.nums, tt.args.score); got != tt.want {
				t.Errorf("maxCoinsBackTrack() = %v, want %v", got, tt.want)
			}
		})
	}
}
