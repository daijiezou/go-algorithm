package leetcode

import "testing"

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
