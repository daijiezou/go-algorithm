package labuladong_homework

import "testing"

func Test_generateBinaryNumber(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test1",
			args: args{n: 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			generateBinaryNumber(tt.args.n)
		})
	}
}
