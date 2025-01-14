package leetcode

import (
	"fmt"
	"testing"
)

func Test_convertToBinary(t *testing.T) {
	type args struct {
		num string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{num: "2080"},
			want: "100000100000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertToBinary(tt.args.num); got != tt.want {
				t.Errorf("convertToBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertDateToBinary(t *testing.T) {
	type args struct {
		date string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{date: "1900-01-01"},
			want: "11101101100-1-1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertDateToBinary(tt.args.date); got != tt.want {
				t.Errorf("convertDateToBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConstructorCalendar(t *testing.T) {
	cal := ConstructorCalendar()
	fmt.Println(cal.Book(47, 50))
	fmt.Println(cal.Book(33, 41))
	fmt.Println(cal.Book(39, 45))
	fmt.Println(cal.Book(33, 42))
	fmt.Println(cal.Book(25, 32))
	fmt.Println(cal.Book(26, 35))
	fmt.Println(cal.Book(19, 25))
	fmt.Println(cal.Book(3, 8))
	fmt.Println(cal.Book(8, 13))
	fmt.Println(cal.Book(18, 27))
	fmt.Println(cal.Calendar)
}

func TestConstructorATM(t *testing.T) {
	atm := ConstructorATM()
	atm.Deposit([]int{0, 0, 1, 2, 1})
	fmt.Println(atm.Withdraw(600))
	atm.Deposit([]int{0, 1, 0, 1, 1})
	fmt.Println(atm.Withdraw(600))

	fmt.Println(atm.Withdraw(550))
	//fmt.Println(atm.BanknotesCount)
}

func Test_maxConsecutive(t *testing.T) {
	type args struct {
		bottom  int
		top     int
		special []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				bottom:  2,
				top:     9,
				special: []int{4, 6},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxConsecutive(tt.args.bottom, tt.args.top, tt.args.special); got != tt.want {
				t.Errorf("maxConsecutive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_largestGoodInteger(t *testing.T) {
	type args struct {
		num string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{num: "6777133339"},
			want: "777",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestGoodInteger(tt.args.num); got != tt.want {
				t.Errorf("largestGoodInteger() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_waysToSplitArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{nums: []int{2, 3, 1, 0}},
			want: 2,
		},
		{
			name: "",
			args: args{nums: []int{10, 4, -8, 7}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := waysToSplitArray(tt.args.nums); got != tt.want {
				t.Errorf("waysToSplitArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minOperationsII(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				nums: []int{2, 11, 10, 1, 3},
				k:    10,
			},
			want: 2,
		},
		{
			name: "2",
			args: args{
				nums: []int{1, 1, 2, 4, 9},
				k:    20,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minOperationsII(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("minOperationsII() = %v, want %v", got, tt.want)
			}
		})
	}
}
