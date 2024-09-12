package leetcode

import "testing"

func Test_maxNumOfMarkedIndices(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		wantRes int
	}{
		{
			name: "case 1",
			args: args{
				nums: []int{42, 83, 48, 10, 24, 55, 9, 100, 10, 17, 17, 99, 51, 32, 16, 98, 99, 31, 28, 68, 71, 14, 64, 29, 15, 40},
			},
			wantRes: 26,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := maxNumOfMarkedIndices(tt.args.nums); gotRes != tt.wantRes {
				t.Errorf("maxNumOfMarkedIndices() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
