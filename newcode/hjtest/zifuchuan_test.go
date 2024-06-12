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
