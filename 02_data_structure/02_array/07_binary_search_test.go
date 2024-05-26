package _2_array

import "testing"

func Test_fDays(t *testing.T) {
	type args struct {
		weights  []int
		capacity int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				weights:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				capacity: 15,
			},
			want: 5,
		},
		{
			name: "case2",
			args: args{
				weights:  []int{1, 2, 3, 1, 1},
				capacity: 4,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fDays(tt.args.weights, tt.args.capacity); got != tt.want {
				t.Errorf("fDays() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_shipWithinDays(t *testing.T) {
	type args struct {
		weights []int
		days    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				weights: []int{1, 2, 3, 1, 1},
				days:    4,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shipWithinDays(tt.args.weights, tt.args.days); got != tt.want {
				t.Errorf("shipWithinDays() = %v, want %v", got, tt.want)
			}
		})
	}
}
