package leetcode_one_question_one_day

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
