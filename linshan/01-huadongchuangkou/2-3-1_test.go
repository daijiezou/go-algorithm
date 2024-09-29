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

func Test_countGood(t *testing.T) {
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
				nums: []int{3, 1, 4, 3, 2, 2, 4},
				k:    2,
			},
			want: 4,
		},
		{

			name: "case 2",
			args: args{
				nums: []int{1, 1, 1, 1, 1},
				k:    10,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countGood(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("countGood() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countSubarrays1(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSubarrays(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("countSubarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_numSubarrayProductLessThanK(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numSubarrayProductLessThanK(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("numSubarrayProductLessThanK() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_numberOfSubstrings(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfSubstrings(tt.args.s); got != tt.want {
				t.Errorf("numberOfSubstrings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countCompleteSubarrays(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{

			name: "case 1",
			args: args{
				nums: []int{1, 3, 1, 2, 2},
			},
			want: 4,
		},
		{

			name: "case 2",
			args: args{
				nums: []int{1, 1, 1, 1},
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countCompleteSubarrays(tt.args.nums); got != tt.want {
				t.Errorf("countCompleteSubarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
