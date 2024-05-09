package mysort

import "testing"

func TestQuickSort(t *testing.T) {
	type args struct {
		arr   []int
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "1",
			args: args{
				arr:   []int{1, 20, 1, 0, 22, 33, 4, 10, -1, -8},
				left:  0,
				right: 9,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			QuickSort(tt.args.arr, tt.args.left, tt.args.right)
			t.Log(tt.args.arr)
		})

	}
}
