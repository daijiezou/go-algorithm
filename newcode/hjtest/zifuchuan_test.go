package main

import "testing"

func Test_countPeaks(t *testing.T) {
	type args struct {
		heights []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				heights: []int{0, 1, 2, 3, 2, 4},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPeaks(tt.args.heights); got != tt.want {
				t.Errorf("countPeaks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countValidStrings(t *testing.T) {
	type args struct {
		characters []byte
		N          int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				characters: []byte{'a', 'a', 'b'},
				N:          2,
			},
			want: 2,
		},
		{
			name: "2",
			args: args{
				characters: []byte{'a', 'b', 'c'},
				N:          2,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countValidStrings(tt.args.characters, tt.args.N); got != tt.want {
				t.Errorf("countValidStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countPaths(t *testing.T) {
	type args struct {
		grid [][]int
		rows int
		cols int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				grid: [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}},
				rows: 3,
				cols: 3,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPaths(tt.args.grid, tt.args.rows, tt.args.cols); got != tt.want {
				t.Errorf("countPaths() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxTasks(t *testing.T) {
	type args struct {
		tasks []Task
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				tasks: []Task{
					{
						start: 1,
						end:   1,
					},
					{
						start: 1,
						end:   2,
					},
					{
						start: 1,
						end:   3,
					},
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxTasks(tt.args.tasks); got != tt.want {
				t.Errorf("maxTasks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxBananas(t *testing.T) {
	type args struct {
		numbers []int
		N       int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				numbers: []int{1, 2, 2, 7, 3, 6, 1},
				N:       3,
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxBananas(tt.args.numbers, tt.args.N); got != tt.want {
				t.Errorf("maxBananas() = %v, want %v", got, tt.want)
			}
		})
	}
}
