package my_string

import "testing"

func TestMakeSmallestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{s: "seven"},
			want: "neven",
		},
		{
			name: "test2",
			args: args{s: "abcd"},
			want: "abba",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MakeSmallestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("MakeSmallestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
