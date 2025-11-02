package _5_grid_chart

import (
	"reflect"
	"testing"
)

func Test_pondSizes(t *testing.T) {
	type args struct {
		land [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "name1",
			args: args{land: [][]int{
				{0, 2, 1, 0},
				{0, 1, 0, 1},
				{1, 1, 0, 1},
				{0, 1, 0, 1},
			}},
			want: []int{1, 2, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pondSizes(tt.args.land); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pondSizes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_largestArea(t *testing.T) {
	type args struct {
		grid []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				grid: []string{
					"110",
					"231",
					"221",
				},
			},
			want: 1,
		},
		{
			name: "case 2",
			args: args{
				grid: []string{
					"111",
					"111",
					"111",
				},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestArea(tt.args.grid); got != tt.want {
				t.Errorf("largestArea() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_islandPerimeter(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				grid: [][]int{
					{0, 1, 0, 0},
					{1, 1, 1, 0},
					{0, 1, 0, 0},
					{1, 1, 0, 0},
				},
			},
			want: 16,
		},
		{
			name: "case 2",
			args: args{
				grid: [][]int{
					{1, 1},
				},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := islandPerimeter(tt.args.grid); got != tt.want {
				t.Errorf("islandPerimeter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findMaxFish(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				grid: [][]int{{0, 2, 1, 0}, {4, 0, 0, 3}, {1, 0, 0, 4}, {0, 3, 2, 0}},
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxFish(tt.args.grid); got != tt.want {
				t.Errorf("findMaxFish() = %v, want %v", got, tt.want)
			}
		})
	}
}
