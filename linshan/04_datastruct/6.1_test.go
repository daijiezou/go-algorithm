package _4_datastruct

import (
	"reflect"
	"testing"
)

func Test_longestWord(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{words: []string{"a", "banana", "app", "appl", "ap", "apply", "apple"}},
			want: "apple",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestWord(tt.args.words); got != tt.want {
				t.Errorf("longestWord() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sumPrefixScores(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{words: []string{"abc", "ab", "bc", "b"}},
			want: []int{5, 4, 3, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumPrefixScores(tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sumPrefixScores() = %v, want %v", got, tt.want)
			}
		})
	}
}
