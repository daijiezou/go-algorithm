package _4_datastruct

import "testing"

func Test_maximumTripletValue(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "1",
			args: args{nums: []int{12, 6, 1, 2, 7}},
			want: 77,
		},
		{
			name: "2",
			args: args{nums: []int{1, 10, 3, 4, 19}},
			want: 133,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumTripletValue(tt.args.nums); got != tt.want {
				t.Errorf("maximumTripletValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maximizeWin(t *testing.T) {
	type args struct {
		prizePositions []int
		k              int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				prizePositions: []int{1, 1, 2, 2, 3, 3, 5},
				k:              2,
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximizeWin(tt.args.prizePositions, tt.args.k); got != tt.want {
				t.Errorf("maximizeWin() = %v, want %v", got, tt.want)
			}
		})
	}
}
