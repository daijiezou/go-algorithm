package _1_listNodeAndArray

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewDifference(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want *Difference
	}{
		{
			name: "1",
			args: args{nums: []int{8, 2, 6, 3, 1}},
			want: &Difference{
				diff: []int{8, -6, 4, -3, -2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDifference(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDifference() = %v, want %v", got, tt.want)
			} else {
				fmt.Println(got.Result())
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
			name: "bookings",
			args: args{
				bookings: [][]int{{1, 2, 10}, {2, 3, 20}, {2, 5, 25}},
				n:        5,
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

func Test_carPooling(t *testing.T) {
	type args struct {
		trips    [][]int
		capacity int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				trips:    [][]int{{2, 1, 5}, {3, 5, 7}},
				capacity: 3,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := carPooling(tt.args.trips, tt.args.capacity); got != tt.want {
				t.Errorf("carPooling() = %v, want %v", got, tt.want)
			}
		})
	}
}
