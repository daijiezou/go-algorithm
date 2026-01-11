package _4_datastruct

import (
	"reflect"
	"testing"
)

func Test_numberOfPoints(t *testing.T) {
	type args struct {
		nums [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{nums: [][]int{
				{3, 6},
				{1, 5},
				{4, 7},
			}},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfPoints(tt.args.nums); got != tt.want {
				t.Errorf("numberOfPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_corpFlightBookings(t *testing.T) {
	type args struct {
		bookings [][]int
		n        int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				bookings: [][]int{
					{1, 2, 10},
					{2, 3, 20},
					{2, 5, 25},
				},
				n: 5,
			},
			want: []int{10, 55, 45, 25, 25},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := corpFlightBookings(tt.args.bookings, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("corpFlightBookings() = %v, want %v", got, tt.want)
			}
		})
	}
}
