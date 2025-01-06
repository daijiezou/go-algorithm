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
