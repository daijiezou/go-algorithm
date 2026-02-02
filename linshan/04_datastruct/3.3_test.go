package _4_datastruct

import "testing"

func Test_removeDuplicates1209(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				s: "deeedbbcccbdaa",
				k: 3,
			},
			want: "aa",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates1209(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("removeDuplicates1209() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeSubstring(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				s: "(())",
				k: 1,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeSubstring(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("removeSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
