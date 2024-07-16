package __graph

import (
	"reflect"
	"testing"
)

func Test_findOrder1(t *testing.T) {
	type args struct {
		numCourses    int
		prerequisites [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				numCourses:    2,
				prerequisites: [][]int{{1, 0}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findOrder1(tt.args.numCourses, tt.args.prerequisites); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findOrder1() = %v, want %v", got, tt.want)
			}
		})
	}
}
