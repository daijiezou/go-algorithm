package newcode

import "testing"

func TestYuhuashi(t *testing.T) {
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
				nums: []int{1, 1, 2, 2, 3, 3},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Yuhuashi(tt.args.nums); got != tt.want {
				t.Errorf("Yuhuashi() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLongestExpression(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "case1",
			args: args{
				s: "a+1--23*45-+67*89b",
			},
			want: 23 * 45,
		},
		{
			name: "case2",
			args: args{
				s: "1-2abcd",
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestExpression(tt.args.s); got != tt.want {
				t.Errorf("LongestExpression() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCrimeTime(t *testing.T) {
	type args struct {
		s1 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case1",
			args: args{
				s1: "20:12",
			},
			want: "20:20",
		},
		{
			name: "case2",
			args: args{
				s1: "23:59",
			},
			want: "22:22",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CrimeTime(tt.args.s1); got != tt.want {
				t.Errorf("CrimeTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWangluofuwuqi(t *testing.T) {
	type args struct {
		nums [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				nums: [][]int{{1, 1, 1}, {1, 0, 1}, {0, 1, 0}},
			},
			want: 5,
		},
		{
			name: "case2",
			args: args{
				nums: [][]int{{1, 0}, {1, 1}},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestServerNetwork(tt.args.nums); got != tt.want {
				t.Errorf("Wangluofuwuqi() = %v, want %v", got, tt.want)
			}
		})
	}
}
