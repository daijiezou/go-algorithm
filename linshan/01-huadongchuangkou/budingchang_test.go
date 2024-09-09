package _1_huadongchuangkou

import "testing"

func Test_equalSubstring(t *testing.T) {
	type args struct {
		s       string
		t       string
		maxCost int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				s:       "abcd",
				t:       "bcdf",
				maxCost: 3,
			},
			want: 3,
		},
		{
			name: "case 2",
			args: args{
				s:       "pxezla",
				t:       "loewbi",
				maxCost: 25,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := equalSubstring(tt.args.s, tt.args.t, tt.args.maxCost); got != tt.want {
				t.Errorf("equalSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
