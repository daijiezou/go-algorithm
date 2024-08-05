package leetcode

import "testing"

func Test_numberOfRightTriangles(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
		{
			name: "case 1",
			args: args{
				grid: [][]int{
					{0, 1, 0},
					{0, 1, 1},
					{0, 1, 0},
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfRightTriangles(tt.args.grid); got != tt.want {
				t.Errorf("numberOfRightTriangles() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxPointsInsideSquare(t *testing.T) {
	type args struct {
		points [][]int
		s      string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				points: [][]int{
					{2, 2}, {-1, -2}, {-4, 4}, {-3, 1}, {3, -3},
				},
				s: "abdca",
			},
			want: 2,
		},
		{
			name: "case 1",
			args: args{
				points: [][]int{
					{-35, -3}, {17, 28}, {28, -28}, {25, -1}, {25, -16}, {1, -21},
				},
				s: "ffcbea",
			},
			want: 5,
		},
		{
			name: "case 3",
			args: args{
				points: [][]int{
					{-1, 0},
				},
				s: "a",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPointsInsideSquare2(tt.args.points, tt.args.s); got != tt.want {
				t.Errorf("maxPointsInsideSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findIntegers(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				n: 5,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findIntegers(tt.args.n); got != tt.want {
				t.Errorf("findIntegers() = %v, want %v", got, tt.want)
			}
		})
	}
}
