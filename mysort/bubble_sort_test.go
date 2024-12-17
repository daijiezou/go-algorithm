package mysort

import "testing"

func TestBubbleSort2(t *testing.T) {
	type args struct {
		arrayList []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "1",
			args: args{
				arrayList: []int{3, 2, 1, 3, 1, 32, 10, 100},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Insert(tt.args.arrayList)
		})
	}
}

func TestBubbleSort21(t *testing.T) {
	type args struct {
		arrayList []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "1",
			args: args{
				arrayList: []int{3, 2, 1, 3, 1, 32, 10, 100},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			BubbleSort2(tt.args.arrayList)
		})
	}
}
