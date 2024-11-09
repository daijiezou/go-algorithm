package _3_monotone_stack

import "testing"

func Test_removeDuplicateLetters(t *testing.T) {
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
			args: args{
				s: "bcabc",
			},
			want: "abc",
		},
		{
			name: "2",
			args: args{
				s: "cbacdcbc",
			},
			want: "acdb",
		},
		{
			name: "3",
			args: args{
				s: "cdadabcc",
			},
			want: "adbc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicateLetters(tt.args.s); got != tt.want {
				t.Errorf("removeDuplicateLetters() = %v, want %v", got, tt.want)
			}
		})
	}
}
