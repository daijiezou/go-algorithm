package leetcode

import (
	"reflect"
	"testing"
)

func Test_minMovesToCaptureTheQueen(t *testing.T) {
	type args struct {
		a int
		b int
		c int
		d int
		e int
		f int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				a: 5,
				b: 3,
				c: 3,
				d: 4,
				e: 5,
				f: 2,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minMovesToCaptureTheQueen(tt.args.a, tt.args.b, tt.args.c, tt.args.d, tt.args.e, tt.args.f); got != tt.want {
				t.Errorf("minMovesToCaptureTheQueen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_knightDialer(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				n: 2,
			},
			want: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := knightDialer(tt.args.n); got != tt.want {
				t.Errorf("knightDialer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getFinalState(t *testing.T) {
	type args struct {
		nums       []int
		k          int
		multiplier int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				nums:       []int{2, 1, 3, 5, 6},
				k:          5,
				multiplier: 2,
			},
			want: []int{8, 4, 6, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getFinalState(tt.args.nums, tt.args.k, tt.args.multiplier); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getFinalState() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minSetSize(t *testing.T) {
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
			args: args{
				arr: []int{3, 3, 3, 3, 5, 5, 5, 2, 2, 7},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSetSize(tt.args.arr); got != tt.want {
				t.Errorf("minSetSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_closestRoom(t *testing.T) {
	type args struct {
		rooms   [][]int
		queries [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				rooms:   [][]int{{23, 22}, {6, 20}, {15, 6}, {22, 19}, {2, 10}, {21, 4}, {10, 18}, {16, 1}, {12, 7}, {5, 22}},
				queries: [][]int{{12, 5}, {15, 15}, {21, 6}, {15, 1}, {23, 4}, {15, 11}, {1, 24}, {3, 19}, {25, 8}, {18, 6}},
			},
			want: []int{12, 10, 22, 15, 23, 10, -1, 5, 23, 15},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := closestRoom(tt.args.rooms, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("closestRoom() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minValidStrings(t *testing.T) {
	type args struct {
		words  []string
		target string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				words:  []string{"abc", "aaaaa", "bcdef"},
				target: "aabcdabc",
			},
			want: 3,
		},
		{
			name: "2",
			args: args{
				words:  []string{"b", "ccacc", "a"},
				target: "cccaaaacba",
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minValidStrings(tt.args.words, tt.args.target); got != tt.want {
				t.Errorf("minValidStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_groupAnagrams(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "1",
			args: args{strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := groupAnagrams(tt.args.strs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("groupAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestConsecutive(t *testing.T) {
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
			args: args{nums: []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestConsecutive(tt.args.nums); got != tt.want {
				t.Errorf("longestConsecutive() = %v, want %v", got, tt.want)
			}
		})
	}
}
