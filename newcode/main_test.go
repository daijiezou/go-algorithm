package newcode

import "testing"

func Test_checkIsBor(t *testing.T) {
	type args struct {
		x string
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{
				x: "dacbb",
				s: "adcbb",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkIsBor(tt.args.x, tt.args.s); got != tt.want {
				t.Errorf("checkIsBor() = %v, want %v", got, tt.want)
			}
		})
	}
}
