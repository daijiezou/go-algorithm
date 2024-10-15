package _2_binary_search

import (
	"fmt"
	"math"
	"testing"
)

func Test_findRadius(t *testing.T) {
	type args struct {
		houses  []int
		heaters []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				houses:  []int{1, 2, 3},
				heaters: []int{2},
			},
			want: 1,
		},
		{
			name: "case 2",
			args: args{
				houses:  []int{1, 5},
				heaters: []int{10},
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRadius(tt.args.houses, tt.args.heaters); got != tt.want {
				t.Errorf("findRadius() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repairCars(t *testing.T) {

	fmt.Println(math.Log2(16))
	type args struct {
		ranks []int
		cars  int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "case 1",
			args: args{
				ranks: []int{4, 2, 3, 1},
				cars:  10,
			},
			want: 16,
		},
		{
			name: "case 2",
			args: args{
				ranks: []int{1, 1, 3, 3},
				cars:  74,
			},
			want: 576,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := repairCars(tt.args.ranks, tt.args.cars); got != tt.want {
				t.Errorf("repairCars() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getRepairCars(t *testing.T) {
	type args struct {
		ranks []int
		k     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				ranks: []int{4, 2, 3, 1},
				k:     64,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getRepairCars(tt.args.ranks, tt.args.k); got != tt.want {
				t.Errorf("getRepairCars() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findRadius1(t *testing.T) {
	type args struct {
		houses  []int
		heaters []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				houses:  []int{1, 5},
				heaters: []int{10},
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRadius(tt.args.houses, tt.args.heaters); got != tt.want {
				t.Errorf("findRadius() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_check(t *testing.T) {
	type args struct {
		houses  []int
		heaters []int
		r       int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case 1",
			args: args{
				houses:  []int{1, 5},
				heaters: []int{10},
				r:       5,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := check(tt.args.houses, tt.args.heaters, tt.args.r); got != tt.want {
				t.Errorf("check() = %v, want %v", got, tt.want)
			}
		})
	}
}
