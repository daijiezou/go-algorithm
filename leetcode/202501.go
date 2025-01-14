package leetcode

import (
	"container/heap"
	"container/list"
	"slices"
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

func maxConsecutive(bottom int, top int, special []int) int {
	res := 0
	slices.Sort(special)
	n := len(special)
	res = special[0] - bottom
	for i := 1; i < n; i++ {
		res = max(res, special[i]-special[i-1]-1)
	}
	res = max(res, top-special[n-1])
	return res
}

func countKeyChanges(s string) int {
	s = strings.ToLower(s)
	cnt := 0
	for i := 1; i < len(s); i++ {
		if s[i] != s[i-1] {
			cnt++
		}
	}
	return cnt
}

func largestGoodInteger(num string) string {
	cnt := 1
	res := ""
	for i := 1; i < len(num); i++ {
		cnt++
		if num[i] != num[i-1] {
			cnt = 1
		}
		if cnt == 3 {
			//numInt,_ := strconv.Atoi()
			res = max(res, num[i-2:i+1])
		}
	}
	return res
}

func validSubstringCount(s string, t string) int64 {
	need := make(map[byte]int)
	window := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	left, right := 0, 0
	valid := 0
	var res int64
	n := len(s)
	for right < len(s) {
		current := s[right]
		right++
		if _, ok := need[current]; ok {
			window[current]++
			if window[current] == need[current] {
				valid++
			}
		}
		for valid == len(need) {
			res += int64(n - right + 1)
			toDelete := s[left]
			left++
			if _, ok := need[toDelete]; ok {
				if window[toDelete] == need[toDelete] {
					valid--
				}
				window[toDelete]--
			}
		}
	}
	return res
}

func largestCombination(candidates []int) int {
	maxLen := func(x int) int {
		cnt := 0
		for i := 0; i < len(candidates); i++ {
			if (candidates[i])&(1<<x) != 0 {
				cnt++
			}
		}
		return cnt
	}
	res := 0
	for i := 0; i < 24; i++ {
		res = max(res, maxLen(i))
	}
	return res
}

func waysToSplitArray(nums []int) int {
	n := len(nums)
	preSum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
	}
	res := 0
	for i := 1; i < n; i++ {
		left := preSum[i] - preSum[0]
		right := preSum[n] - preSum[i]
		if left >= right {
			res++
		}
	}
	return res
}

func waysToSplitArray2(nums []int) int {
	total := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		total += nums[i]
	}
	res := 0
	curSum := 0
	for i := 0; i < n-1; i++ {
		curSum += nums[i]
		if curSum*2 >= total {
			res++
		}
	}
	return res
}

func minOperationsI(nums []int, k int) int {
	cnt := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] < k {
			cnt++
		}
	}
	return cnt
}

type mySlice []int

func (m mySlice) Len() int {
	//TODO implement me
	return len(m)
}

func (m mySlice) Less(i, j int) bool {
	return m[i] < m[j]
}

func (m mySlice) Swap(i, j int) {
	//TODO implement me
	m[i], m[j] = m[j], m[i]
}

func (m *mySlice) Push(x any) {
	*m = append(*m, x.(int))
}

func (m *mySlice) Pop() any {

	n := len(*m)
	res := (*m)[n-1]
	*m = (*m)[0 : n-1]
	return res

}

func minOperationsII(nums []int, k int) int {
	nums2 := mySlice(nums)
	heap.Init(&nums2)
	cnt := 0
	for nums2.Len() >= 2 {
		x := heap.Pop(&nums2).(int)
		y := heap.Pop(&nums2).(int)
		if x < k || y < k {
			cnt++
		} else {
			return cnt
		}
		newEle := min(x, y)*2 + max(x, y)
		heap.Push(&nums2, newEle)
	}
	return cnt
}
