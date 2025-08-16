package aim_offer

import (
	"reflect"
	"testing"
)

func TestPrintMinNumber(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{numbers: []int{11, 3}},
			want: "113",
		},
		{
			name: "1",
			args: args{numbers: []int{3, 32, 321}},
			want: "321323",
		},
		{
			name: "1",
			args: args{numbers: []int{1, 3, 22}},
			want: "1223",
		},
		{
			name: "1",
			args: args{numbers: []int{39, 397, 6}},
			want: "393976",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PrintMinNumber(tt.args.numbers); got != tt.want {
				t.Errorf("PrintMinNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solve(t *testing.T) {
	type args struct {
		nums string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{nums: "31717126241541717"},
			want: 192,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.nums); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxValue(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				grid: [][]int{
					{1, 3, 1},
					{1, 5, 1},
					{4, 2, 1},
				},
			},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxValue(tt.args.grid); got != tt.want {
				t.Errorf("maxValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{s: "bbbbbb"},
			want: 1,
		},
		{
			name: "1",
			args: args{s: "abcabcbb"},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetUglyNumber_Solution(t *testing.T) {
	type args struct {
		index int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				index: 7,
			},
			want: 8,
		},
		{
			name: "1",
			args: args{
				index: 11,
			},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetUglyNumber_Solution(tt.args.index); got != tt.want {
				t.Errorf("GetUglyNumber_Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetMissNumber(t *testing.T) {
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
			args: args{nums: []int{0, 1, 2, 3, 4, 5, 7, 8, 9}},
			want: 6,
		},
		{
			name: "1",
			args: args{nums: []int{0, 1, 2, 3}},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMissNumber(tt.args.nums); got != tt.want {
				t.Errorf("GetMissNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetNumSameAsIndex(t *testing.T) {
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
			args: args{nums: []int{-3, -1, 1, 3, 5}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetNumSameAsIndex(tt.args.nums); got != tt.want {
				t.Errorf("GetNumSameAsIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindNumsAppearOnce(t *testing.T) {
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
			args: args{nums: []int{1, 2, 3, 3, 2, 9}},
			want: []int{1, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindNumsAppearOnce(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindNumsAppearOnce() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindContinuousSeq(t *testing.T) {
	type args struct {
		sum int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "1",
			args: args{sum: 9},
			want: [][]int{
				{2, 3, 4},
				{4, 5},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindContinuousSeq(tt.args.sum); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindContinuousSeq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLeftRotateString(t *testing.T) {
	type args struct {
		str string
		n   int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				str: "abcXYZdef",
				n:   3,
			},
			want: "XYZdefabc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LeftRotateString2(tt.args.str, tt.args.n); got != tt.want {
				t.Errorf("LeftRotateString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxInWindows(t *testing.T) {
	type args struct {
		num  []int
		size int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				num:  []int{2, 3, 4, 2, 6, 2, 5, 1},
				size: 3,
			},
			want: []int{4, 4, 6, 6, 6, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxInWindows(tt.args.num, tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxInWindows() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLastRemaining_Solution(t *testing.T) {
	type args struct {
		n int
		m int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				n: 5,
				m: 3,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LastRemaining_Solution(tt.args.n, tt.args.m); got != tt.want {
				t.Errorf("LastRemaining_Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
