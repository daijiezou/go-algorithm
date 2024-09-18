package _1_huadongchuangkou

import "testing"

func Test_countSubarrays(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{

			name: "case 1",
			args: args{
				nums: []int{1, 3, 2, 3, 3},
				k:    2,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSubarrays(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("countSubarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
