package _1_core_platform

import "testing"

func Test_getlongestPalindrome(t *testing.T) {
	type args struct {
		s     string
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "babad",
			args: args{
				s:     "babad",
				left:  2,
				right: 2,
			},
			want: "aba",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getlongestPalindrome(tt.args.s, tt.args.left, tt.args.right); got != tt.want {
				t.Errorf("getlongestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
