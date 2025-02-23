package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_validPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{s: "abca"},
			want: true,
		},
		{
			name: "1",
			args: args{s: "abc"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validPalindrome(tt.args.s); got != tt.want {
				t.Errorf("validPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_subsetsWithDup(t *testing.T) {
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
			args: args{nums: []int{1, 1, 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := permuteUnique(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("subsetsWithDup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConstructor0218(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want RangeFreqQuery
	}{
		{
			name: "1",
			args: args{arr: []int{12, 33, 4, 56, 22, 2, 34, 33, 22, 12, 34, 56}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ins := Constructor0218(tt.args.arr)
			fmt.Println(ins.Query(1, 2, 4))
			fmt.Println(ins.Query(0, 11, 33))
		})
	}
}

func Test_evenOddBit(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{n: 50},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := evenOddBit(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("evenOddBit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_similarPairs(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{words: []string{"aba", "aabb", "abcd", "bac", "aabc"}},
			want: 2,
		},
		{
			name: "2",
			args: args{words: []string{"aabb", "ab", "ba"}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := similarPairs(tt.args.words); got != tt.want {
				t.Errorf("similarPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
