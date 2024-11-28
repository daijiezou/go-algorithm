package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_minChanges(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				n: 13,
				k: 4,
			},
			want: 2,
		},
		{
			name: "case2",
			args: args{
				n: 14,
				k: 13,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minChanges(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("minChanges() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_shoppingOffers(t *testing.T) {
	type args struct {
		price   []int
		special [][]int
		needs   []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				price:   []int{2, 5},
				special: [][]int{{3, 0, 5}, {1, 2, 10}},
				needs:   []int{3, 2},
			},
			want: 14,
		},
		{
			name: "case1",
			args: args{
				price:   []int{2, 3, 4},
				special: [][]int{{1, 1, 0, 4}, {2, 2, 1, 9}},
				needs:   []int{1, 2, 1},
			},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shoppingOffers(tt.args.price, tt.args.special, tt.args.needs); got != tt.want {
				t.Errorf("shoppingOffers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConstructor2(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want NeighborSum
	}{
		{
			name: "case1",
			args: args{
				grid: [][]int{{0, 1, 2}, {3, 4, 5}, {6, 7, 8}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ns := Constructor2(tt.args.grid)
			fmt.Println(ns)
			fmt.Println(ns.AdjacentSum(1))
		})
	}
}

func Test_minCost2(t *testing.T) {
	type args struct {
		n    int
		cuts []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				n:    7,
				cuts: []int{1, 3, 4, 5},
			},
			want: 16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCost2(tt.args.n, tt.args.cuts); got != tt.want {
				t.Errorf("minCost2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countKConstraintSubstrings(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				s: "1010101",
				k: 2,
			},
			want: 25,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countKConstraintSubstrings(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("countKConstraintSubstrings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countKConstraintSubstrings2(t *testing.T) {
	type args struct {
		s       string
		k       int
		queries [][]int
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{
			name: "case1",
			args: args{
				s:       "1010101",
				k:       2,
				queries: [][]int{{0, 6}},
			},
			want: []int64{26},
		},
		{
			name: "case2",
			args: args{
				s:       "1010101",
				k:       1,
				queries: [][]int{{0, 5}, {1, 4}, {2, 3}},
			},
			want: []int64{15, 9, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countKConstraintSubstrings2(tt.args.s, tt.args.k, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countKConstraintSubstrings2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minFlips(t *testing.T) {
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
			args: args{grid: [][]int{{1, 0, 0}, {0, 0, 0}, {0, 0, 1}}},
			want: 2,
		},
		{
			name: "1",
			args: args{grid: [][]int{{0, 1}, {0, 1}, {0, 0}}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minFlips(tt.args.grid); got != tt.want {
				t.Errorf("minFlips() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_numFriendRequests(t *testing.T) {
	type args struct {
		ages []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{ages: []int{16, 17, 18}},
			want: 2,
		},
		{
			name: "2",
			args: args{ages: []int{20, 30, 100, 110, 120}},
			want: 3,
		},
		{
			name: "3",
			args: args{ages: []int{16, 16}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numFriendRequests(tt.args.ages); got != tt.want {
				t.Errorf("numFriendRequests() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_imageSmoother(t *testing.T) {
	type args struct {
		img [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "1",
			args: args{img: [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}},
			want: [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := imageSmoother(tt.args.img); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("imageSmoother() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_shortestDistanceAfterQueries(t *testing.T) {
	type args struct {
		n       int
		queries [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case1",
			args: args{
				n:       5,
				queries: [][]int{{2, 4}, {0, 2}, {0, 4}},
			},
			want: []int{3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortestDistanceAfterQueries(tt.args.n, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("shortestDistanceAfterQueries() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nonSpecialCount(t *testing.T) {
	type args struct {
		l int
		r int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				l: 4,
				r: 16,
			},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nonSpecialCount(tt.args.l, tt.args.r); got != tt.want {
				t.Errorf("nonSpecialCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_winningPlayerCount(t *testing.T) {
	type args struct {
		n    int
		pick [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				n:    5,
				pick: [][]int{{0, 0}, {1, 0}, {1, 0}, {2, 1}, {2, 1}, {2, 0}},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := winningPlayerCount(tt.args.n, tt.args.pick); got != tt.want {
				t.Errorf("winningPlayerCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_networkDelayTime(t *testing.T) {
	type args struct {
		times [][]int
		n     int
		k     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				times: [][]int{{1, 2, 1}, {2, 3, 2}, {1, 3, 4}},
				n:     3,
				k:     1,
			},
			want: 3,
		},
		{
			name: "case1",
			args: args{
				times: [][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}},
				n:     4,
				k:     2,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := networkDelayTime3(tt.args.times, tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("networkDelayTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_numberOfAlternatingGroups2(t *testing.T) {
	type args struct {
		colors []int
		k      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				colors: []int{0, 1, 0, 1, 0},
				k:      3,
			},
			want: 3,
		},
		{
			name: "1",
			args: args{
				colors: []int{0, 1, 0, 0, 1, 0, 1},
				k:      6,
			},
			want: 2,
		},
		{
			name: "1",
			args: args{
				colors: []int{0, 0, 1, 1},
				k:      3,
			},
			want: 0,
		},
		{
			name: "1",
			args: args{
				colors: []int{0, 1, 0, 1},
				k:      3,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfAlternatingGroups2(tt.args.colors, tt.args.k); got != tt.want {
				t.Errorf("numberOfAlternatingGroups2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countOfPairs(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				nums: []int{2, 3, 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countOfPairs(tt.args.nums); got != tt.want {
				t.Errorf("countOfPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countOfPairs1(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	fmt.Println(int(1e9+7) == 1000000007)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countOfPairs(tt.args.nums); got != tt.want {
				t.Errorf("countOfPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
