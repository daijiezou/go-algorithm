package _024

import "testing"

func Test_maximumOddBinaryNumber(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{s: "0101"},
			want: "1001",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumOddBinaryNumber(tt.args.s); got != tt.want {
				t.Errorf("maximumOddBinaryNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maximumBinaryString(t *testing.T) {
	type args struct {
		binary string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				binary: "00011",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumBinaryString(tt.args.binary); got != tt.want {
				t.Errorf("maximumBinaryString() = %v, want %v", got, tt.want)
			}
		})
	}
}
