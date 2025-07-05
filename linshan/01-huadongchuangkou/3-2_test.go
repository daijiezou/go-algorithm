package _1_huadongchuangkou

import "testing"

func Test_findLengthOfShortestSubarray_2(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{arr: []int{1, 2, 3, 10, 4, 2, 3, 5}},
			want: 3,
		},
		{
			name: "1",
			args: args{arr: []int{10, 13, 17, 21, 15, 15, 9, 17, 22, 22, 13}},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLengthOfShortestSubarray_2(tt.args.arr); got != tt.want {
				t.Errorf("findLengthOfShortestSubarray_2() = %v, want %v", got, tt.want)
			}
		})
	}
}
