package _1_jianzhioffer

import "testing"

func Test_duplicate2(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{numbers: []int{2, 3, 1, 0, 2, 5}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := duplicate2(tt.args.numbers); got != tt.want {
				t.Errorf("duplicate2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getDup(t *testing.T) {
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
			args: args{nums: []int{3, 2, 1, 4, 4, 5, 6, 7}},
			want: 4,
		},
		{
			name: "1",
			args: args{nums: []int{1, 7, 3, 4, 5, 6, 8, 2, 8}},
			want: 8,
		},
		{
			name: "1",
			args: args{nums: []int{1, 1}},
			want: 1,
		},
		{
			name: "1",
			args: args{nums: []int{3, 2, 1, 3, 4, 5, 6, 7}},
			want: 3,
		},
		{
			name: "1",
			args: args{nums: []int{1, 2, 3, 4, 5, 6, 7}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDup(tt.args.nums); got != tt.want {
				t.Errorf("getDup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReplaceSpace(t *testing.T) {
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
			args: args{
				s: "A B",
			},
			want: "A%20B",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReplaceSpace(tt.args.s); got != tt.want {
				t.Errorf("ReplaceSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fob(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		n    int
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := fob(tt.n)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("fob() = %v, want %v", got, tt.want)
			}
		})
	}
}
