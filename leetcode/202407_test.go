package leetcode

import (
	"reflect"
	"testing"
)

func Test_incremovableSubarrayCount(t *testing.T) {
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
				nums: []int{6, 5, 7, 8},
			},
			want: 7,
		},
		{
			name: "case 2",
			args: args{
				nums: []int{8, 7, 6, 6},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := incremovableSubarrayCount(tt.args.nums); got != tt.want {
				t.Errorf("incremovableSubarrayCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_canSortArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case 1",
			args: args{
				nums: []int{8, 4, 2, 30, 15},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canSortArray(tt.args.nums); got != tt.want {
				t.Errorf("canSortArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxIncreaseKeepingSkyline(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				grid: [][]int{
					{3, 0, 8, 4},
					{2, 4, 5, 7},
					{9, 2, 6, 3},
					{0, 3, 1, 0},
				},
			},
			want: 35,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxIncreaseKeepingSkyline(tt.args.grid); got != tt.want {
				t.Errorf("maxIncreaseKeepingSkyline() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_accountsMerge(t *testing.T) {
	type args struct {
		accounts [][]string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "case 1",
			args: args{
				accounts: [][]string{
					{"John", "johnsmith@mail.com", "john_newyork@mail.com"},
					{"John", "johnsmith@mail.com", "john00@mail.com"},
					{"Mary", "mary@mail.com"},
					{"John", "johnnybravo@mail.com"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := accountsMerge(tt.args.accounts); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("accountsMerge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_permutation(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "case 1",
			args: args{
				n: 3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			permutation(tt.args.n)
		})
	}
}

func Test_minimumMoves(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				grid: [][]int{
					{0, 2, 0},
					{3, 0, 0},
					{3, 1, 0},
				},
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumMoves(tt.args.grid); got != tt.want {
				t.Errorf("minimumMoves() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maximumDetonation(t *testing.T) {
	type args struct {
		bombs [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				bombs: [][]int{
					{4, 4, 3},
					{4, 4, 3},
				},
			},
			want: 2,
		},
		{
			name: "case 2",
			args: args{
				bombs: [][]int{
					{1, 2, 3}, {2, 3, 1}, {3, 4, 2}, {4, 5, 3}, {5, 6, 4},
				},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumDetonation(tt.args.bombs); got != tt.want {
				t.Errorf("maximumDetonation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sumOfPowers(t *testing.T) {
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
			name: "case 1",
			args: args{
				nums: []int{1, 2, 3, 4},
				k:    3,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumOfPowers(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("sumOfPowers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_relocateMarbles(t *testing.T) {
	type args struct {
		nums     []int
		moveFrom []int
		moveTo   []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := relocateMarbles(tt.args.nums, tt.args.moveFrom, tt.args.moveTo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("relocateMarbles() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minimumOperations(t *testing.T) {
	type args struct {
		num string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				num: "2908305",
			},
			want: 3,
		},
		{
			name: "case 2",
			args: args{
				num: "10",
			},
			want: 1,
		},
		{
			name: "case 3",
			args: args{
				num: "50986431",
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumOperations2(tt.args.num); got != tt.want {
				t.Errorf("minimumOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calPoints(t *testing.T) {
	type args struct {
		operations []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "case 1",
			args: args{
				operations: []string{
					"5", "2", "C", "D", "+",
				},
			},
			want: 30,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calPoints(tt.args.operations); got != tt.want {
				t.Errorf("calPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getGoodIndices(t *testing.T) {
	type args struct {
		variables [][]int
		target    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "case 1",
			args: args{
				variables: [][]int{
					{528, 818, 733, 438},
				},
				target: 256,
			},
			want: []int{0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getGoodIndices(tt.args.variables, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getGoodIndices() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pow_mod(t *testing.T) {
	type args struct {
		x   int
		y   int
		mod int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				x:   2,
				y:   11,
				mod: 10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pow_mod(tt.args.x, tt.args.y, tt.args.mod); got != tt.want {
				t.Errorf("pow_mod() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minRectanglesToCoverPoints(t *testing.T) {
	type args struct {
		points [][]int
		w      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				points: [][]int{
					{2, 1}, {1, 0}, {1, 4}, {1, 8}, {3, 5}, {4, 6},
				},
				w: 1,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minRectanglesToCoverPoints(tt.args.points, tt.args.w); got != tt.want {
				t.Errorf("minRectanglesToCoverPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxmiumScore(t *testing.T) {
	type args struct {
		cards []int
		cnt   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				cards: []int{1, 2, 8, 9},
				cnt:   3,
			},
			want: 18,
		},
		{
			name: "case 4",
			args: args{
				cards: []int{7, 6, 4, 6},
				cnt:   1,
			},
			want: 6,
		},
		{
			name: "case 2",
			args: args{
				cards: []int{13, 12, 10, 19, 19, 4, 16, 10, 2, 9, 2, 13, 13, 15, 5, 19, 3, 13, 17, 4, 18, 19, 8, 1, 19, 18, 17, 14, 6, 9, 6, 11, 4},
				cnt:   4,
			},
			want: 76,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxmiumScore(tt.args.cards, tt.args.cnt); got != tt.want {
				t.Errorf("maxmiumScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
