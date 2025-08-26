package offer_review1

import (
	"reflect"
	"testing"
)

func Test_hasPath(t *testing.T) {
	type args struct {
		matrix [][]byte
		word   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				matrix: [][]byte{
					{'a', 'b', 'c', 'e'},
					{'s', 'f', 'c', 's'},
					{'a', 'd', 'e', 'e'},
				},
				word: "abcced",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasPath(tt.args.matrix, tt.args.word); got != tt.want {
				t.Errorf("hasPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_movingCount(t *testing.T) {
	type args struct {
		threshold int
		rows      int
		cols      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				threshold: 5,
				rows:      10,
				cols:      10,
			},
			want: 21,
		},
		{
			name: "1",
			args: args{
				threshold: 10,
				rows:      1,
				cols:      100,
			},
			want: 29,
		},
		{
			name: "1",
			args: args{
				threshold: 0,
				rows:      1,
				cols:      3,
			},
			want: 1,
		},
		{
			name: "1",
			args: args{
				threshold: 1,
				rows:      2,
				cols:      3,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := movingCount(tt.args.threshold, tt.args.rows, tt.args.cols); got != tt.want {
				t.Errorf("movingCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pow(t *testing.T) {
	type args struct {
		x float64
		n int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "1",
			args: args{
				x: 3,
				n: 4,
			},
			want: 81.0,
		},
		{
			name: "1",
			args: args{
				x: 2,
				n: 3,
			},
			want: 8.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pow(tt.args.x, tt.args.n); got != tt.want {
				t.Errorf("pow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_printMatrix(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{matrix: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
				{13, 14, 15, 16},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := printMatrix2(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("printMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsPopOrder(t *testing.T) {
	type args struct {
		pushV []int
		popV  []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				pushV: []int{2, 1, 0},
				popV:  []int{1, 2, 0},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPopOrder(tt.args.pushV, tt.args.popV); got != tt.want {
				t.Errorf("IsPopOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVerifySquenceOfBST(t *testing.T) {
	type args struct {
		sequence []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{sequence: []int{4, 6, 7, 5}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := VerifySquenceOfBST(tt.args.sequence); got != tt.want {
				t.Errorf("VerifySquenceOfBST() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindPath(t *testing.T) {
	root := &TreeNode{
        Val: 10,
        Left: &TreeNode{
            Val: 5,
            Left:  &TreeNode{Val: 4},
            Right: &TreeNode{Val: 7},
        },
        Right: &TreeNode{Val: 12},
    }

    got := FindPath(root, 22)

    // 说明：FindPath 返回的路径顺序和内部顺序，可能与期望不一致。
    // 如果你严格比较，可以根据返回顺序调整期望；下面给一个常见的期望写法：
    want1 := [][]int{{10, 12}, {10, 5, 7}}
    want2 := [][]int{{10, 5, 7}, {10, 12}}

    if !reflect.DeepEqual(got, want1) && !reflect.DeepEqual(got, want2) {
        t.Fatalf("FindPath() = %#v, want %#v or %#v", got, want1, want2)
    }
}
