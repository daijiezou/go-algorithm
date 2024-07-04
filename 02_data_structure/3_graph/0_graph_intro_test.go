package __graph

import (
	"reflect"
	"testing"
)

func Test_allPathsSourceTarget(t *testing.T) {
	type args struct {
		graph [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "case1",
			args: args{
				graph: [][]int{{1, 2}, {3}, {3}, {}},
			},
			want: [][]int{{0, 1, 3}, {0, 2, 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := allPathsSourceTarget(tt.args.graph); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("allPathsSourceTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_canFinish(t *testing.T) {
	type args struct {
		numCourses    int
		prerequisites [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case1",
			args: args{
				numCourses:    2,
				prerequisites: [][]int{{1, 0}, {0, 1}},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canFinish(tt.args.numCourses, tt.args.prerequisites); got != tt.want {
				t.Errorf("canFinish() = %v, want %v", got, tt.want)
			}
		})
	}
}
