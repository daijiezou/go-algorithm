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

func Test_numSubarraysWithSum(t *testing.T) {
	type args struct {
		nums []int
		goal int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				nums: []int{1, 0, 1, 0, 1},
				goal: 2,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numSubarraysWithSum(tt.args.nums, tt.args.goal); got != tt.want {
				t.Errorf("numSubarraysWithSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findMaxLength(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{nums: []int{0, 1}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxLength(tt.args.nums); got != tt.want {
				t.Errorf("findMaxLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maximumSubarraySum(t *testing.T) {
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
			args: args{
				nums: []int{-636, -784, -356, -832, -797, -978, -651, -667, -907, -900, -168, -697, -879, -998, -126, -900, -542, -553, -268, -374, -710, -768, -727, -975, -106, -756, -462, -815, -276, -163, -301, -822, -367, -685, -581, -488, -763, -612, -847, -730, -479, -874, -221, -912, -229, -543, -876, -845, -424, -215, -819, -164, -840, -525, -987, -291, -658, -168, -382, -781, -951, -396, -228, -394, -445, -863, -290, -675, -289, -950, -885, -228, -624, -236, -437, -246, -302, -741, -243, -419, -851, -980, -667, -661, -140, -893, -328, -354, -359, -845, -396, -309, -450, -941, -310, -119, -614, -854, -599, -605},
				k:    8,
			},
			want: -1088,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumSubarraySum(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("maximumSubarraySum() = %v, want %v", got, tt.want)
			}
		})
	}
}
