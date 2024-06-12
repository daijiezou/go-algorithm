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
			if got := maxTask2(tt.args.tasks); got != tt.want {
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

func Test_summarizeString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case1",
			args: args{
				s: "aabbcc",
			},
			want: "a2b2c2",
		},
		{
			name: "case2",
			args: args{
				s: "bAaAcBb",
			},
			want: "a3b2b2c0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := summarize(tt.args.s); got != tt.want {
				t.Errorf("summarizeString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countLuckyNumbers(t *testing.T) {
	type args struct {
		k int
		n int
		m int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				k: 10,
				n: 2,
				m: 4,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countLuckyNumbers(tt.args.k, tt.args.n, tt.args.m); got != tt.want {
				t.Errorf("countLuckyNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getMaxValue(t *testing.T) {
	type args struct {
		tasks []Task2
		T     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				tasks: []Task2{
					{
						SLA:   1,
						Value: 2,
					},
					{
						SLA:   1,
						Value: 3,
					},
					{
						SLA:   1,
						Value: 4,
					},
					{
						SLA:   1,
						Value: 5,
					},
				},
				T: 3,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMaxValue(tt.args.tasks, tt.args.T); got != tt.want {
				t.Errorf("getMaxValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
