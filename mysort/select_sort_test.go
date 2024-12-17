package mysort

import (
	"reflect"
	"testing"
)

func Test_selectSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				nums: []int{5, 4, 3, 2, 1},
			},
			want: []int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := selectSort(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("selectSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
