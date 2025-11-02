package _4_datastruct

import (
	"reflect"
	"testing"
)

func Test_reverseNum(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{num: 110},
			want: 11,
		},
		{
			name: "1",
			args: args{num: 213},
			want: 312,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseNum(tt.args.num); got != tt.want {
				t.Errorf("reverseNum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findIndices(t *testing.T) {
	type args struct {
		nums            []int
		indexDifference int
		valueDifference int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{valueDifference: 5, indexDifference: 3,
				nums: []int{0, 9, 0, 5, 9}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findIndices(tt.args.nums, tt.args.indexDifference, tt.args.valueDifference); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findIndices() = %v, want %v", got, tt.want)
			}
		})
	}
}
