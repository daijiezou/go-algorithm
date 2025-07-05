package _1_huadongchuangkou

import "testing"

func Test_minimumLength(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{s: "aabccabba"},
			want: 3,
		},
		{
			name: "1",
			args: args{s: "cabaabac"},
			want: 0,
		},
		{
			name: "1",
			args: args{s: "aabaaa"},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumLength(tt.args.s); got != tt.want {
				t.Errorf("minimumLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_purchasePlans(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				nums:   []int{2, 2, 1, 9},
				target: 10,
			},
			want: 4,
		},
		{
			name: "1",
			args: args{
				nums:   []int{2, 5, 3, 5},
				target: 6,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := purchasePlans(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("purchasePlans() = %v, want %v", got, tt.want)
			}
		})
	}
}
