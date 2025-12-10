package _4_datastruct

import "testing"

func Test_subarraySum2(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				nums: []int{1, 2, 3},
				k:    3,
			},
			want: 2,
		},
		{
			name: "1",
			args: args{
				nums: []int{1, 1, 1},
				k:    2,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subarraySum2(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("subarraySum2() = %v, want %v", got, tt.want)
			}
		})
	}
}
