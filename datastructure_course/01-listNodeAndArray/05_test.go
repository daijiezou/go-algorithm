package _1_listNodeAndArray

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewDifference(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want *Difference
	}{
		{
			name: "1",
			args: args{nums: []int{8, 2, 6, 3, 1}},
			want: &Difference{
				diff: []int{8, -6, 4, -3, -2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDifference(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDifference() = %v, want %v", got, tt.want)
			} else {
				fmt.Println(got.Result())
			}

		})

	}
}
