package aim_offer

import "testing"

func Test_inventoryManagement(t *testing.T) {
	type args struct {
		stock []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				stock: []int{4, 5, 6, 7, 0, 1, 2},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inventoryManagement(tt.args.stock); got != tt.want {
				t.Errorf("inventoryManagement() = %v, want %v", got, tt.want)
			}
		})
	}
}
