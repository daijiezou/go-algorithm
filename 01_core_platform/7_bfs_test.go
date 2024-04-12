package _1_core_platform

import "testing"

func TestBFS(t *testing.T) {
	type args struct {
		target string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "1",
			args: args{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			BFS(tt.args.target)
		})
	}
}
