package leetcode

import (
	"reflect"
	"testing"
)

func Test_partition(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "1",
			args: args{
				s: "aab",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partition(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minCut(t *testing.T) {
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
			args: args{
				s: "aab",
			},
			want: 1,
		},
		{
			name: "1",
			args: args{
				s: "ab",
			},
			want: 1,
		},
		{
			name: "1",
			args: args{
				s: "b",
			},
			want: 0,
		},
		{
			name: "1",
			args: args{
				s: "cdd",
			},
			want: 1,
		},
		{
			name: "1",
			args: args{
				s: "ababababababababababababcbabababababababababababa",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCut(tt.args.s); got != tt.want {
				t.Errorf("minCut() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maximumBeauty(t *testing.T) {
	type args struct {
		items   [][]int
		queries []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				items: [][]int{
					{1, 2},
					{3, 2},
					{2, 4},
					{5, 6},
					{3, 5},
				},
				queries: []int{1, 2, 3, 4, 5, 6},
			},
			want: []int{2, 4, 5, 5, 6, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumBeauty(tt.args.items, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maximumBeauty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sumOfBeauties(t *testing.T) {
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
			args: args{
				nums: []int{1, 2, 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumOfBeauties(tt.args.nums); got != tt.want {
				t.Errorf("sumOfBeauties() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countOfSubstrings(t *testing.T) {
	type args struct {
		word string
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
				word: "ieaouqqieaouqq",
				k:    1,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countOfSubstrings(tt.args.word, tt.args.k); got != tt.want {
				t.Errorf("countOfSubstrings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minSwaps(t *testing.T) {
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
			args: args{s: "]]][[["},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSwaps(tt.args.s); got != tt.want {
				t.Errorf("minSwaps() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findMatrix(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "1",
			args: args{nums: []int{1, 3, 4, 1, 2, 3, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMatrix(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maximumOr(t *testing.T) {
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
			name: "1",
			args: args{
				nums: []int{12, 9},
				k:    1,
			},
			want: 30,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumOr(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("maximumOr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_canBeValid(t *testing.T) {
	type args struct {
		s      string
		locked string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				s:      "))()))",
				locked: "010100",
			},
			want: true,
		},
		{
			name: "2",
			args: args{
				s:      "()((",
				locked: "1111",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canBeValid(tt.args.s, tt.args.locked); got != tt.want {
				t.Errorf("canBeValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countPrefixes(t *testing.T) {
	type args struct {
		words []string
		s     string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				s:     "w",
				words: []string{"feh", "w", "w", "lwd", "c", "s", "vk", "zwlv", "n", "w", "sw", "qrd", "w", "w", "mqe", "w", "w", "w", "gb", "w", "qy", "xs", "br", "w", "rypg", "wh", "g", "w", "w", "fh", "w", "w", "sccy"},
			},
			want: 14,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPrefixes(tt.args.words, tt.args.s); got != tt.want {
				t.Errorf("countPrefixes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minimumSum(t *testing.T) {
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
			name: "1",
			args: args{
				n: 5,
				k: 4,
			},
			want: 18,
		},
		{
			name: "1",
			args: args{
				n: 2,
				k: 3,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumSum(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("minimumSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestCycle(t *testing.T) {
	type args struct {
		edges []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{edges: []int{3, 3, 4, 2, 3}},
			want: 3,
		},
		{
			name: "1",
			args: args{edges: []int{2, -1, 3, 1}},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCycle(tt.args.edges); got != tt.want {
				t.Errorf("longestCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_addSpaces(t *testing.T) {
	type args struct {
		s      string
		spaces []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "!",
			args: args{
				s:      "LeetcodeHelpsMeLearn",
				spaces: []int{8, 13, 15},
			},
			want: "Leetcode Helps Me Learn",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addSpaces2(tt.args.s, tt.args.spaces); got != tt.want {
				t.Errorf("addSpaces() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mostPoints(t *testing.T) {
	type args struct {
		questions [][]int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "1",
			args: args{questions: [][]int{
				{3, 2},
				{4, 3},
				{4, 4},
				{2, 5},
			}},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got :=
				mostPoints2(tt.args.questions); got != tt.want {
				t.Errorf("mostPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
