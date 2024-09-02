package leetcode

import (
	"reflect"
	"testing"
)

func Test_numberOfRightTriangles(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
		{
			name: "case 1",
			args: args{
				grid: [][]int{
					{0, 1, 0},
					{0, 1, 1},
					{0, 1, 0},
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfRightTriangles(tt.args.grid); got != tt.want {
				t.Errorf("numberOfRightTriangles() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxPointsInsideSquare(t *testing.T) {
	type args struct {
		points [][]int
		s      string
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
					{2, 2}, {-1, -2}, {-4, 4}, {-3, 1}, {3, -3},
				},
				s: "abdca",
			},
			want: 2,
		},
		{
			name: "case 1",
			args: args{
				points: [][]int{
					{-35, -3}, {17, 28}, {28, -28}, {25, -1}, {25, -16}, {1, -21},
				},
				s: "ffcbea",
			},
			want: 5,
		},
		{
			name: "case 3",
			args: args{
				points: [][]int{
					{-1, 0},
				},
				s: "a",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPointsInsideSquare2(tt.args.points, tt.args.s); got != tt.want {
				t.Errorf("maxPointsInsideSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findIntegers(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				n: 5,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findIntegers(tt.args.n); got != tt.want {
				t.Errorf("findIntegers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minimumAddedInteger(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				nums1: []int{4, 20, 16, 12, 8},
				nums2: []int{14, 18, 10},
			},
			want: -2,
		},
		{
			name: "case 2",
			args: args{
				nums1: []int{4, 6, 3, 1, 4, 2, 10, 9, 5},
				nums2: []int{5, 10, 3, 2, 6, 1, 9},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumAddedInteger2(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("minimumAddedInteger() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxUncrossedLines(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				nums1: []int{2, 5, 1, 2, 5},
				nums2: []int{10, 5, 2, 1, 5, 2},
			},
			want: 3,
		},
		{
			name: "case 2",
			args: args{
				nums1: []int{1, 3, 7, 1, 7, 5},
				nums2: []int{1, 9, 2, 5, 1},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxUncrossedLines2(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("maxUncrossedLines() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMagicDict(t *testing.T) {
	type args struct {
		req1 []string
		req2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case 1",
			args: args{
				req1: []string{"hello", "leetcode"},
				req2: "hhllo",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MagicDict(tt.args.req1, tt.args.req2); got != tt.want {
				t.Errorf("MagicDict() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isArraySpecial2(t *testing.T) {
	type args struct {
		nums    []int
		queries [][]int
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		{
			name: "case 1",
			args: args{
				nums:    []int{1, 1},
				queries: [][]int{{0, 1}},
			},
			want: []bool{false},
		},
		{
			name: "case 2",
			args: args{
				nums:    []int{2, 8, 10, 9},
				queries: [][]int{{1, 3}},
			},
			want: []bool{false},
		},
		{
			name: "case 3",
			args: args{
				nums:    []int{4, 3, 1, 6},
				queries: [][]int{{0, 2}, {2, 3}},
			},
			want: []bool{false, true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isArraySpecial2(tt.args.nums, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("isArraySpecial2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxScore(t *testing.T) {
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
					{9, 5, 7, 3},
					{8, 9, 6, 1},
					{6, 7, 14, 3},
					{2, 5, 3, 1},
				},
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxScore(tt.args.grid); got != tt.want {
				t.Errorf("maxScore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minimumOperationsToMakeKPeriodic(t *testing.T) {
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
			name: "case 1",
			args: args{
				word: "leetcoleet",
				k:    2,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumOperationsToMakeKPeriodic(tt.args.word, tt.args.k); got != tt.want {
				t.Errorf("minimumOperationsToMakeKPeriodic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkRecord(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				n: 10101,
			},
			want: 183236316,
		},
		{
			name: "case 2",
			args: args{
				n: 3,
			},
			want: 19,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkRecord(tt.args.n); got != tt.want {
				t.Errorf("checkRecord() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_jump(t *testing.T) {
	type args struct {
		j int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				j: 2,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := jump(tt.args.j); got != tt.want {
				t.Errorf("jump() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_waysToReachStair(t *testing.T) {
	type args struct {
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				k: 0,
			},
			want: 2,
		},
		{
			name: "case 2",
			args: args{
				k: 1,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := waysToReachStair(tt.args.k); got != tt.want {
				t.Errorf("waysToReachStair() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_canPartitionKSubsets(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case 1",
			args: args{
				nums: []int{4, 3, 2, 3, 5, 2, 1},
				k:    4,
			},
			want: true,
		},
		{
			name: "case 2",
			args: args{
				nums: []int{1, 1, 1, 1, 2, 2, 2, 2},
				k:    4,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canPartitionKSubsets(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("canPartitionKSubsets() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getImportance(t *testing.T) {
	type args struct {
		employees []*Employee
		id        int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				employees: []*Employee{
					{
						Id:           1,
						Importance:   5,
						Subordinates: []int{2, 3},
					},
					{
						Id:           2,
						Importance:   3,
						Subordinates: []int{},
					},
					{
						Id:           3,
						Importance:   3,
						Subordinates: []int{},
					},
				},
				id: 1,
			},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getImportance(tt.args.employees, tt.args.id); got != tt.want {
				t.Errorf("getImportance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkBalance(t *testing.T) {
	type args struct {
		sByte []byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case 1",
			args: args{
				sByte: []byte{'a', 'b', 'c', 'd', 'e', 'e'},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkBalance(tt.args.sByte); got != tt.want {
				t.Errorf("checkBalance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minimumSubstringsInPartition(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				s: "abababaccddb",
			},
			want: 2,
		},
		{
			name: "case 2",
			args: args{
				s: "fabccddg",
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumSubstringsInPartition(tt.args.s); got != tt.want {
				t.Errorf("minimumSubstringsInPartition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sumDigitDifferences(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "case 1",
			args: args{
				nums: []int{13, 23, 12},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumDigitDifferences(tt.args.nums); got != tt.want {
				t.Errorf("sumDigitDifferences() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_canMakeSquare(t *testing.T) {
	type args struct {
		grid [][]byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case 1",
			args: args{
				grid: [][]byte{
					{'B', 'B', 'B'},
					{'B', 'W', 'W'},
					{'B', 'B', 'B'},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canMakeSquare(tt.args.grid); got != tt.want {
				t.Errorf("canMakeSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}
