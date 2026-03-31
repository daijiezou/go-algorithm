package _4_datastruct

import "testing"

func Test_pickGifts(t *testing.T) {
	type args struct {
		gifts []int
		k     int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "1",
			args: args{
				gifts: []int{25, 64, 9, 4, 100},
				k:     4,
			},
			want: 29,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pickGifts(tt.args.gifts, tt.args.k); got != tt.want {
				t.Errorf("pickGifts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_halveArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				nums: []int{5, 19, 8, 1},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := halveArray(tt.args.nums); got != tt.want {
				t.Errorf("halveArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maximumProduct_2233(t *testing.T) {
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
			name: "1",
			args: args{nums: []int{6, 3, 3, 2}, k: 2},
			want: 216,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumProduct_2233(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("maximumProduct_2233() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_smallestChair(t *testing.T) {
	type args struct {
		times        [][]int
		targetFriend int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "!",
			args: args{
				times: [][]int{
					{3, 10},
					{1, 5},
					{2, 6},
				},
				targetFriend: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestChair(tt.args.times, tt.args.targetFriend); got != tt.want {
				t.Errorf("smallestChair() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getNumberOfBacklogOrders(t *testing.T) {
	type args struct {
		orders [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{orders: [][]int{
				{10, 5, 0},
				{15, 2, 1},
				{25, 1, 1},
				{30, 4, 0},
			}},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getNumberOfBacklogOrders(tt.args.orders); got != tt.want {
				t.Errorf("getNumberOfBacklogOrders() = %v, want %v", got, tt.want)
			}
		})
	}
}
