package leetcode

import (
	"container/list"
	"strconv"
	"strings"
)

func convertDateToBinary(date string) string {
	dataSlice := strings.Split(date, "-")
	for i := 0; i < 3; i++ {
		dataSlice[i] = convertToBinary(dataSlice[i])
	}
	return strings.Join(dataSlice, "-")
}

func convertToBinary(num string) string {
	numInt64, _ := strconv.Atoi(num)
	res := strings.Builder{}
	resSlice := make([]byte, 0)
	for numInt64 > 0 {
		mod2 := numInt64 % 2
		numInt64 = numInt64 >> 1
		// 将数字 0 或 1 转为字符 '0' 或 '1'
		resSlice = append(resSlice, byte(mod2)+'0')
	}
	// 反转结果
	for i, j := 0, len(resSlice)-1; i < j; i, j = i+1, j-1 {
		resSlice[i], resSlice[j] = resSlice[j], resSlice[i]
	}
	// 将结果写入 Builder
	res.Write(resSlice)
	return res.String()
}

type MyCalendar struct {
	Calendar *list.List
}

func ConstructorCalendar() MyCalendar {
	return MyCalendar{Calendar: list.New()}
}

func (this *MyCalendar) Book(startTime int, endTime int) bool {
	for e := this.Calendar.Front(); e != nil; e = e.Next() {
		event := e.Value.([2]int)
		if event[0] < endTime && startTime < event[1] {
			// 本次日程还没结束，下个日程就开始了
			return false
		}
	}
	// 成功安排日程
	this.Calendar.PushBack([2]int{startTime, endTime})
	return true
}

type ATM struct {
	BanknotesCount  []int
	banknotesAmount []int
}

func ConstructorATM() ATM {
	return ATM{
		BanknotesCount:  make([]int, 5),
		banknotesAmount: []int{20, 50, 100, 200, 500},
	}
}

func (this *ATM) Deposit(banknotesCount []int) {
	for i := 0; i < len(banknotesCount); i++ {
		this.BanknotesCount[i] += banknotesCount[i]
	}
}

func (this *ATM) Withdraw(amount int) []int {
	res := make([]int, 5)
	for i := 4; i >= 0; i-- {
		total := this.BanknotesCount[i]
		curAmount := this.banknotesAmount[i]
		cnt := amount / curAmount
		cnt = min(cnt, total)
		amount -= cnt * curAmount
		res[i] = cnt
		if amount == 0 {
			// 进行扣款
			for k := 0; k < len(res); k++ {
				this.BanknotesCount[k] -= res[k]
			}
			return res
		}
	}
	return []int{-1}
}
