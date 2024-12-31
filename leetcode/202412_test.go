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

func Test_getWeight(t *testing.T) {
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
			args: args{num: 12},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getWeight(tt.args.num); got != tt.want {
				t.Errorf("getWeight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getKth2(t *testing.T) {
	type args struct {
		lo int
		hi int
		k  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				lo: 12,
				hi: 15,
				k:  2,
			},
			want: 13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getKth2(tt.args.lo, tt.args.hi, tt.args.k); got != tt.want {
				t.Errorf("getKth2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_eatenApples(t *testing.T) {
	type args struct {
		apples []int
		days   []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				apples: []int{1, 2, 3, 5, 2},
				days:   []int{3, 2, 1, 4, 2},
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := eatenApples(tt.args.apples, tt.args.days); got != tt.want {
				t.Errorf("eatenApples() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rankTeams(t *testing.T) {
	type args struct {
		votes []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{votes: []string{"ABC", "ACB", "ABC", "ACB", "ACB"}},
			want: "ACB",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rankTeams(tt.args.votes); got != tt.want {
				t.Errorf("rankTeams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findMid(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "1",
			args: args{nums: []int{7, 2, 1, 10, 5, 8, 3}},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMedian(tt.args.nums); got != tt.want {
				t.Errorf("findMid() = %v, want %v", got, tt.want)
			}
		})
	}
}
