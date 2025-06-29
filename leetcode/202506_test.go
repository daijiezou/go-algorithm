package leetcode

import (
	"reflect"
	"testing"
)

func Test_distributeCandies(t *testing.T) {
	type args struct {
		n     int
		limit int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "1",
			args: args{
				n:     5,
				limit: 2,
			},
			want: 3,
		},
		{
			name: "3",
			args: args{
				n:     3,
				limit: 3,
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := distributeCandies(tt.args.n, tt.args.limit); got != tt.want {
				t.Errorf("distributeCandies() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxCandies(t *testing.T) {
	type args struct {
		status         []int
		candies        []int
		keys           [][]int
		containedBoxes [][]int
		initialBoxes   []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				status:  []int{1, 0, 1, 0},
				candies: []int{7, 5, 4, 100},
				keys: [][]int{
					{},
					{},
					{1},
					{},
				},
				containedBoxes: [][]int{
					{1, 2},
					{3},
					{},
					{},
				},
				initialBoxes: []int{0},
			},
			want: 16,
		},
		{
			args: args{
				status:  []int{1, 0, 1, 0},
				candies: []int{7, 5, 4, 100},
				keys: [][]int{
					{},
					{},
					{1},
					{3},
				},
				containedBoxes: [][]int{
					{1, 2},
					{3},
					{},
					{},
				},
				initialBoxes: []int{0},
			},
			want: 16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxCandies(tt.args.status, tt.args.candies, tt.args.keys, tt.args.containedBoxes, tt.args.initialBoxes); got != tt.want {
				t.Errorf("maxCandies() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_answerString(t *testing.T) {
	type args struct {
		word       string
		numFriends int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				word:       "dbca",
				numFriends: 2,
			},
			want: "dbc",
		},
		{
			name: "2",
			args: args{
				word:       "gggg",
				numFriends: 4,
			},
			want: "g",
		},
		{
			name: "3",
			args: args{
				word:       "aann",
				numFriends: 2,
			},
			want: "nn",
		},
		{
			name: "3",
			args: args{
				word:       "gh",
				numFriends: 1,
			},
			want: "gh",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := answerString(tt.args.word, tt.args.numFriends); got != tt.want {
				t.Errorf("answerString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_smallestEquivalentString(t *testing.T) {
	type args struct {
		s1      string
		s2      string
		baseStr string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				s1:      "parker",
				s2:      "morris",
				baseStr: "parser",
			},
			want: "makkek",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestEquivalentString(tt.args.s1, tt.args.s2, tt.args.baseStr); got != tt.want {
				t.Errorf("smallestEquivalentString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_clearStars(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{s: "aaba*"},
			want: "aab",
		},
		{
			name: "1",
			args: args{s: "ed"},
			want: "ed",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := clearStars(tt.args.s); got != tt.want {
				t.Errorf("clearStars() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minMaxDifference(t *testing.T) {
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
			args: args{num: 11891},
			want: 99009,
		},
		{
			name: "1",
			args: args{num: 99999},
			want: 99999,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minMaxDifference(tt.args.num); got != tt.want {
				t.Errorf("minMaxDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxDiff(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "123456",
			args: args{num: 123456},
			want: 820000,
		},
		{
			name: "111",
			args: args{num: 111},
			want: 888,
		},
		{
			name: "1101057",
			args: args{num: 1101057},
			want: 8808050,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDiff(tt.args.num); got != tt.want {
				t.Errorf("maxDiff() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maximumDifference(t *testing.T) {
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
			args: args{nums: []int{7, 1, 5, 4}},
			want: 4,
		},
		{
			name: "1",
			args: args{nums: []int{9, 4, 3, 2}},
			want: -1,
		},
		{
			name: "1",
			args: args{nums: []int{1, 5, 2, 10}},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumDifference(tt.args.nums); got != tt.want {
				t.Errorf("maximumDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxDistance2(t *testing.T) {
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
			name: "1",
			args: args{
				s: "NSWWEW",
				k: 3,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDistance2(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("maxDistance2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minimumDeletions(t *testing.T) {
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
				word: "dabdcbdcdcd",
				k:    2,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumDeletions(tt.args.word, tt.args.k); got != tt.want {
				t.Errorf("minimumDeletions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findKDistantIndices(t *testing.T) {
	type args struct {
		nums []int
		key  int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				nums: []int{3, 4, 9, 1, 3, 9, 5},
				key:  9,
				k:    1,
			},
			want: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name: "1",
			args: args{
				nums: []int{2, 2, 2, 2, 2},
				key:  2,
				k:    2,
			},
			want: []int{0, 1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKDistantIndices(tt.args.nums, tt.args.key, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findKDistantIndices() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxSubsequence(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				nums: []int{-1, -2, 3, 4},
				k:    3,
			},
			want: []int{-1, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubsequence(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
