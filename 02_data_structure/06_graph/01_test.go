package _6_graph

import "testing"

func Test_findCircleNum(t *testing.T) {
	type args struct {
		isConnected [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				isConnected: [][]int{
					{1, 1, 0},
					{1, 1, 0},
					{0, 0, 1},
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findCircleNum(tt.args.isConnected); got != tt.want {
				t.Errorf("findCircleNum() = %v, want %v", got, tt.want)
			}
		})
	}
}
