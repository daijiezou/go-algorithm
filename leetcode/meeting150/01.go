package meeting150

import (
	"slices"
	"strings"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	mi, ni, p := m-1, n-1, len(nums1)-1
	for mi >= 0 && ni >= 0 {
		if nums1[mi] > nums2[ni] {
			nums1[p] = nums1[mi]
			p--
			mi--
		} else {
			nums1[p] = nums2[ni]
			p--
			ni--
		}
	}
	for mi >= 0 {
		nums1[p] = nums1[mi]
		p--
		mi--
	}
	for ni >= 0 {
		nums1[p] = nums2[ni]
		p--
		ni--
	}
}

func removeElement(nums []int, val int) int {
	slow := 0
	n := len(nums)
	res := 0
	for i := 0; i < n; i++ {
		if nums[i] != val {
			nums[slow] = nums[i]
			slow++
			res++
		}
	}
	return res
}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
		fast++
	}
	return slow + 1
}

func removeDuplicates2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	slow, fast := 0, 0
	cnt := 1
	for fast < len(nums) {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		} else if slow < fast && cnt < 3 {
			slow++
			nums[slow] = nums[fast]
		}
		fast++
		cnt++
		if fast < len(nums) && nums[fast] != nums[fast-1] {
			cnt = 1
		}
	}
	return slow + 1
}

func majorityElement(nums []int) int {
	numCnt := make(map[int]int)
	n := len(nums)
	targetCnt := n / 2
	for i := 0; i < n; i++ {
		numCnt[nums[i]]++
		if numCnt[nums[i]] > targetCnt {
			return nums[i]
		}
	}
	return -1
}

func maxProfit(prices []int) int {
	s := make([]int, 0)
	res := 0
	revenuePrice := make(map[int]int)
	for i := 0; i < len(prices); i++ {
		revenue := 0
		for len(s) > 0 && prices[i] > prices[s[len(s)-1]] {
			pop := s[len(s)-1]
			s = s[:len(s)-1]
			revenue = max(prices[i]-prices[pop]+revenuePrice[pop], revenue)
			revenuePrice[i] = revenue
			res = max(res, revenue)
		}
		s = append(s, i)
	}
	return res
}

func maxProfit2(prices []int) int {
	minPrice := prices[0]
	res := 0
	for i := 0; i < len(prices); i++ {
		minPrice = min(prices[i], minPrice)
		res = max(prices[i]-minPrice, res)
	}
	return res
}

/*
dp[i][1] 表示第i天 手里有股票
dp[i][0] 表示第i天 手里没股票
*/
func maxProfit_k1(prices []int) int {
	dp := make([][2]int, 0)
	n := len(prices)
	for i := 0; i < n; i++ {
		if i == 0 {
			dp[i][0] = 0
			dp[i][1] = -prices[i]
			continue
		}
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], -prices[i])
	}
	return dp[n-1][0]
}

func maxProfit_kk(prices []int) int {

	n := len(prices)
	dp := make([][2]int, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			dp[i][0] = 0
			dp[i][1] = -prices[i]
			continue
		}
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i]) // 能继承之前的，表示能够操作多次

	}
	return dp[n-1][0]
}

func canJump(nums []int) bool {
	nextRight := 0
	curRight := 0
	n := len(nums)
	for i := 0; i < n-1; i++ {
		nextRight = max(nextRight, nums[i]+i)
		if curRight == i {
			if nextRight == i { // 无法到达i+1
				return false
			}
			curRight = nextRight
		}
	}
	return true
}

func jump(nums []int) int {
	n := len(nums)
	memo := make(map[int]int)
	for i := 0; i < n; i++ {
		memo[i] = n
	}

	return jumpDp(nums, 0, memo)
}

func jumpDp(nums []int, start int, memo map[int]int) int {
	if start >= len(nums)-1 {
		return 0
	}
	if memo[start] != len(nums) {
		return memo[start]
	}
	steps := nums[start]
	for i := 1; i <= steps; i++ {
		// 穷举每一个选择
		// 计算每一个子问题的结果
		step := jumpDp(nums, start+i, memo) + 1
		memo[start] = min(memo[start], step)
	}
	return memo[start]
}

func jump2(nums []int) int {
	n := len(nums)
	curEnd := 0  // 当前终点
	nextEnd := 0 //下一个终点
	step := 0
	for i := 0; i < n; i++ {
		nextEnd = max(nextEnd, nums[i]+i)
		if curEnd == i { //到达当前终点
			curEnd = nextEnd //再造一个新桥
			step++
		}
	}
	return step
}

func minTaps(n int, ranges []int) int {
	rightMost := make([]int, n+1)
	for i, r := range ranges {
		left := max(i-r, 0)
		rightMost[left] = max(rightMost[left], i+r)
	}
	curEnd := 0  // 当前终点
	nextEnd := 0 //下一个终点
	cnt := 0
	for i := 0; i < n; i++ {
		nextEnd = max(nextEnd, rightMost[i]+i)
		if curEnd == i { //到达当前终点
			if nextEnd == i { // 无法到达i+1
				return -1
			}
			curEnd = nextEnd //再造一个新桥
			cnt++
		}
	}
	return cnt
}

// h指数
func hIndex(citations []int) int {
	slices.Sort(citations)
	h := len(citations)
	n := len(citations)
	for h >= 1 {
		res := leftBound(citations, h)
		if n-res >= h {
			return h
		}
		h--
	}
	return 0
}

func leftBound(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] >= target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

// https://leetcode.cn/problems/product-of-array-except-self/?envType=study-plan-v2&envId=top-interview-150
func productExceptSelf(nums []int) []int {
	n := len(nums)

	// pre[i] 为nums[0]到nums[i-1]的乘积
	pre := make([]int, n)
	pre[0] = 1
	for i := 1; i < n; i++ {
		pre[i] = pre[i-1] * nums[i-1]
	}
	suf := make([]int, n)
	suf[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		suf[i] = suf[i+1] * nums[i+1]
	}
	res := make([]int, n)
	//res[0] = suf[1]
	//res[n-1] = pre[n-2]
	for i := 0; i < n; i++ {
		res[i] = pre[i] * suf[i]
	}
	return res
}

func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	for i := 0; i < n; i++ {
		totalCost := 0
		totalGas := 0
		step := 0
		for step < n {
			index := (i + step) % n
			totalGas += gas[index]
			totalCost += cost[index]
			if totalCost > totalGas {
				i += step
				break
			}
			step++
		}
		if step == n {
			return i
		}
	}
	return -1
}

// 困难,依赖了题解
func candy(ratings []int) int {
	n := len(ratings)
	candys := make([]int, n)
	for i := 0; i < n; i++ {
		if i > 0 && ratings[i] > ratings[i-1] {
			candys[i] = candys[i-1] + 1
		} else {
			candys[i] = 1
		}
	}
	for i := n - 1; i >= 0; i-- {
		if i < n-1 && ratings[i] > ratings[i+1] {
			candys[i] = max(candys[i+1]+1, ratings[i])
		}
	}
	total := 0
	for i := 0; i < n; i++ {
		total += candys[i]
	}
	return total
}

func romanToInt(s string) int {
	res := 0
	n := len(s)
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'I':
			if i < n-1 && s[i+1] == 'V' {
				res += 4
				i++
				continue
			}
			if i < n-1 && s[i+1] == 'X' {
				res += 9
				i++
				continue
			}
			res += 1

		case 'V':
			res += 5
		case 'X':
			if i < n-1 && s[i+1] == 'L' {
				res += 40
				i++
				continue
			}
			if i < n-1 && s[i+1] == 'C' {
				res += 90
				i++
				continue
			}
			res += 10
		case 'L':
			res += 50
		case 'C':
			if i < n-1 && s[i+1] == 'D' {
				res += 400
				i++
				continue
			}
			if i < n-1 && s[i+1] == 'M' {
				res += 900
				i++
				continue
			}
			res += 100
		case 'D':
			res += 500
		case 'M':
			res += 1000
		}
	}
	return res
}

func romanToInt2(s string) int {
	romanMap := map[uint8]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	n := len(s)
	res := 0
	for i := 0; i < n; i++ {
		if i < n-1 && romanMap[s[i]] < romanMap[s[i+1]] {
			res -= romanMap[s[i]]
		} else {
			res += romanMap[s[i]]
		}
	}
	return res
}

func intToRoman(num int) string {
	var valueSymbols = []struct {
		value  int
		symbol string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}
	res := []byte{}
	for _, vs := range valueSymbols {
		for num >= vs.value {
			num -= vs.value
			res = append(res, vs.symbol...)
		}
		if num == 0 {
			break
		}
	}
	return string(res)
}

/*
给你一个字符串 s，由若干单词组成，单词前后用一些空格字符隔开。返回字符串中 最后一个 单词的长度。

单词 是指仅由字母组成、不包含任何空格字符的最大
子字符串
。
*/
func lengthOfLastWord(s string) int {
	words := strings.Split(s, " ")

	for i := len(words) - 1; i >= 0; i-- {
		if words[i] != "" {
			return len(words[i])
		}
	}
	return -1
}

func lengthOfLastWord2(s string) int {
	i := len(s) - 1
	for i >= 0 {
		if s[i] == ' ' {
			i--
		} else {
			break
		}
	}
	j := i - 1
	for j >= 0 {
		if s[j] != ' ' {
			j--
		} else {
			break
		}
	}
	return i - j
}

// 最长公共前缀
func longestCommonPrefix(strs []string) string {
	minLengthVal := strs[0]
	length := len(minLengthVal)
	for length > 0 {

		var i int
	loop:
		for i = 0; i < len(strs); i++ {
			if !strings.HasPrefix(strs[i], minLengthVal) {
				minLengthVal = minLengthVal[:length-1]
				length--
				break loop
			}
		}
		if i == len(strs) {
			return minLengthVal
		}

	}

	return ""
}

func strStr(haystack string, needle string) int {
	n := len(needle)
	for i := 0; i < len(haystack)-n; i++ {
		if haystack[i:i+n] == needle {
			return i
		}
	}
	return -1
}
