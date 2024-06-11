package main

import (
	"reflect"
	"testing"
)

func TestGet(t *testing.T) {
	type args struct {
		nums []int
		avg  float64
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "case1",
			args: args{
				nums: []int{0, 1, 2, 3, 4},
				avg:  1,
			},
			want: []string{"0-2"},
		},
		{
			name: "case2",
			args: args{
				nums: []int{0, 0, 100, 2, 2, 99, 0, 2},
				avg:  2,
			},
			want: []string{"0-1", "3-4", "6-7"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Get(tt.args.nums, tt.args.avg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_evaluateExpression(t *testing.T) {
	type args struct {
		expr string
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := evaluateExpression(tt.args.expr)
			if (err != nil) != tt.wantErr {
				t.Errorf("evaluateExpression() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("evaluateExpression() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isOperator(t *testing.T) {
	type args struct {
		c rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isOperator(tt.args.c); got != tt.want {
				t.Errorf("isOperator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestValidExpression(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestValidExpression(tt.args.s); got != tt.want {
				t.Errorf("longestValidExpression() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromXiaoqu(t *testing.T) {
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
				nums: []int{2, 2, 3},
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromXiaoqu(tt.args.nums); got != tt.want {
				t.Errorf("FromXiaoqu() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nonRecursiveDFS(t *testing.T) {
	type args struct {
		lengths []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				lengths: []int{7, 3, 4, 5, 6, 5, 12, 13},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countRightTriangles(tt.args.lengths); got != tt.want {
				t.Errorf("nonRecursiveDFS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandleExcel(t *testing.T) {
	type args struct {
		cells []string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "case1",
			args: args{
				cells: []string{"1", "2<A>00"},
			},
		},
		{
			name: "case2",
			args: args{
				cells: []string{"<F>", "2<A>00", "3<B>00", "4<A>00", "5<B>00", "3"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			HandleExcel(tt.args.cells)
		})
	}
}

func Test_maxPizzaSum(t *testing.T) {
	type args struct {
		pizzaSizes []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				pizzaSizes: []int{8, 2, 10, 5, 7},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPizzaSum(tt.args.pizzaSizes); got != tt.want {
				t.Errorf("maxPizzaSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
