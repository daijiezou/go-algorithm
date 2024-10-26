package _3_monotone_stack

import (
	"fmt"
	"math"
	"sort"
)

// 从左到右遍历
func dailyTemperatures(temperatures []int) []int {
	length := len(temperatures)
	res := make([]int, length)
	s := make([]int, 0)
	for i := 0; i < length; i++ {
		for len(s) > 0 && temperatures[i] > temperatures[s[len(s)-1]] {
			j := s[len(s)-1]
			s = s[:len(s)-1]
			res[j] = i - j
		}
		s = append(s, i)
	}
	return res
}

// 从右向左
func dailyTemperatures2(temperatures []int) []int {
	length := len(temperatures)
	res := make([]int, length)
	s := make([]int, 0)
	for i := length - 1; i >= 0; i-- {
		for len(s) > 0 && temperatures[s[len(s)-1]] <= temperatures[i] {
			s = s[:len(s)-1]
		}
		if len(s) == 0 {
			res[i] = 0
		} else {
			res[i] = s[len(s)-1] - i
		}
		s = append(s, i)
	}
	return res
}

func finalPrices(prices []int) []int {
	res := make([]int, len(prices))
	stack := make([]int, 0)
	n := len(prices)
	for i := 0; i < n; i++ {
		for len(stack) > 0 && prices[i] <= prices[stack[len(stack)-1]] {
			j := stack[len(stack)-1]
			res[j] = prices[j] - prices[i]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	for i := 0; i < len(stack); i++ {
		res[stack[i]] = prices[stack[i]]
	}
	return res
}

func finalPrices2(prices []int) []int {
	res := make([]int, len(prices))
	stack := make([]int, 0)
	n := len(prices)
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && prices[i] < prices[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			res[i] = prices[i]
		} else {
			res[i] = prices[i] - prices[stack[len(stack)-1]]
		}
		stack = append(stack, i)
	}
	return res
}

/*
nums1 中数字 x 的 下一个更大元素 是指 x 在 nums2 中对应位置 右侧 的 第一个 比 x 大的元素。
给你两个 没有重复元素 的数组 nums1 和 nums2 ，下标从 0 开始计数，其中nums1 是 nums2 的子集。
*/
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	n := len(nums2)
	nextGreaterElementMap := make(map[int]int, len(nums2))
	s := make([]int, 0)
	for i := 0; i < n; i++ {
		for len(s) > 0 && nums2[i] > s[len(s)-1] {
			j := s[len(s)-1]
			s = s[:len(s)-1]
			nextGreaterElementMap[j] = nums2[i]
		}
		s = append(s, nums2[i])
	}
	for i := 0; i < len(s); i++ {
		nextGreaterElementMap[s[i]] = -1
	}
	res := make([]int, len(nums1))
	for i := 0; i < len(nums1); i++ {
		res[i] = nextGreaterElementMap[nums1[i]]
	}
	return res
}

func nextGreaterElements2(nums []int) []int {
	n := len(nums)
	//nums = append(nums, nums...)
	res := make([]int, n)
	s := make([]int, 0)
	for i := 0; i < 2*n; i++ {
		x := i % n
		for len(s) > 0 && nums[x] > nums[s[len(s)-1]] {
			j := s[len(s)-1]
			s = s[:len(s)-1]
			res[j] = nums[x]
		}
		if i < n {
			s = append(s, i)
		}

	}
	for i := 0; i < len(s); i++ {
		res[s[i]] = -1
	}
	return res
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func maxWidthRamp2(nums []int) int {
	n := len(nums)
	res := 0
	//s := make([]int, 0)
	for i := 0; i < n; i++ {
		if n-i < res {
			break
		}
		for j := n - 1; j >= 0 && j-i > res; j-- {
			if nums[i] <= nums[j] {
				res = max(res, j-i)
				break
			}
		}
	}
	return res
}

// https://leetcode.cn/problems/maximum-width-ramp/description/
/*
给定一个整数数组 A，坡是元组 (i, j)，其中  i < j 且 A[i] <= A[j]。这样的坡的宽度为 j - i。

找出 A 中的坡的最大宽度，如果不存在，返回 0 。
*/
func maxWidthRamp(nums []int) int {
	stack := []int{0}
	ans := 0

	// First pass: populate the stack
	for i := 1; i < len(nums); i++ {
		if nums[stack[len(stack)-1]] > nums[i] {
			stack = append(stack, i)
		}
	}
	fmt.Println(stack)
	// Second pass: calculate the maximum width ramp
	for j := len(nums) - 1; j >= 0; j-- {
		for len(stack) > 0 && nums[j] >= nums[stack[len(stack)-1]] {
			ans = max(ans, j-stack[len(stack)-1])
			stack = stack[:len(stack)-1] // pop the last element
		}
	}
	return ans
}

// https://leetcode.cn/problems/car-fleet/
func carFleet(target int, position []int, speed []int) int {
	n := len(position)
	cars := make([][2]int, n)
	for i := 0; i < n; i++ {
		cars[i][0] = position[i]
		cars[i][1] = speed[i]
	}
	// 按照初始位置，从小到大排序
	sort.Slice(cars, func(i, j int) bool {
		return cars[i][0] < cars[j][0]
	})
	// 计算每辆车到达终点的时间
	time := make([]float64, n)
	for i := 0; i < n; i++ {
		car := cars[i]
		time[i] = float64(target-car[0]) / float64(car[1])
	}
	// 如果位置再后面，需要的时间却比前面的车长，则肯定会形成一个车队
	s := make([]int, n)
	for i := 0; i < n; i++ {
		for len(s) > 0 && time[i] >= time[s[len(s)-1]] {
			s = s[:len(s)-1]
		}
		s = append(s, i)
	}
	return len(s)
}

type StockSpanner struct {
	s        []pair
	curDay   int
	daysSpan map[int]int
}

type pair struct {
	day   int
	price int
}

func Constructor() StockSpanner {
	return StockSpanner{
		s:        []pair{{-1, math.MaxInt}},
		curDay:   -1,
		daysSpan: map[int]int{},
	}
}

func (this *StockSpanner) Next(price int) int {
	this.curDay++
	//cnt := 1
	for len(this.s) != 0 && price >= this.s[len(this.s)-1].price {
		//cnt += this.daysSpan[this.s[len(this.s)-1].day]
		this.s = this.s[:len(this.s)-1]
	}
	this.s = append(this.s, pair{this.curDay, price})
	//this.daysSpan[this.curDay] = cnt
	return this.curDay - this.s[len(this.s)-2].day
}

// https://leetcode.cn/problems/longest-well-performing-interval/
/*
给你一份工作时间表 hours，上面记录着某一位员工每天的工作小时数。
我们认为当员工一天中的工作小时数大于 8 小时的时候，那么这一天就是「劳累的一天」。
所谓「表现良好的时间段」，意味在这段时间内，「劳累的天数」是严格 大于「不劳累的天数」。
请你返回「表现良好时间段」的最大长度。
*/
func longestWPI(hours []int) int {
	window := make([]int, 0)
	left := 0
	laolei := 0
	bulaolei := 0
	res := 0
	for i := 0; i < len(hours); i++ {
		window = append(window, hours[i])
		if hours[i] > 8 {
			laolei++
		} else {
			bulaolei++
		}
		for bulaolei >= laolei && left <= i {
			window = window[1:]
			if hours[left] > 8 {
				laolei--
			} else {
				bulaolei--
			}
			left++
		}
		res = max(res, len(window))
	}
	return res
}

// https://leetcode.cn/problems/maximum-score-of-a-good-subarray/
func maximumScore(nums []int, k int) int {
	n := len(nums)
	s1 := []int{k}
	// 纪录最小值
	s1Map := map[int]int{
		k: nums[k],
	}
	s1Min := nums[k]
	for i := k - 1; i >= 0; i-- {
		if nums[i] < s1Min {
			s1Min = nums[i]
		}
		for len(s1) > 0 && nums[i] > nums[s1[len(s1)-1]] && s1Min >= s1Map[s1[len(s1)-1]] {
			s1 = s1[:len(s1)-1]
		}
		s1Map[i] = s1Min
		s1 = append(s1, i)
	}

	// 逐渐递减
	s2 := []int{k}
	s2Map := map[int]int{
		k: nums[k],
	}
	s2Min := nums[k]
	for i := k + 1; i < n; i++ {
		if nums[i] < s2Min {
			s2Min = nums[i]
		}
		for len(s2) > 0 && nums[i] > nums[s2[len(s2)-1]] && s2Min >= s2Map[s2[len(s2)-1]] {
			s2 = s2[:len(s2)-1]
		}
		s2Map[i] = s2Min
		s2 = append(s2, i)
	}
	//fmt.Println(s1, s1Map)
	//fmt.Println(s2, s2Map)
	ans := 0
	for i := 0; i < len(s1); i++ {
		for j := 0; j < len(s2); j++ {
			minNum := s1Map[s1[i]]
			if s2Map[s2[j]] < minNum {
				minNum = s2Map[s2[j]]
			}
			ans = max(ans, minNum*(s2[j]-s1[i]+1))
		}
	}
	return ans
}

func maximumScore2(nums []int, k int) int {
	n := len(nums)
	left := k - 1
	right := k + 1
	ans := 0
	for i := nums[k]; ; i-- {
		for left >= 0 && nums[left] >= i {
			left--
		}
		for right < n && nums[right] >= i {
			right++
		}
		ans = max(ans, (right-left-1)*i)
		if left == -1 && right == n {
			break
		}
	}
	return ans
}

func maximumScore3(nums []int, k int) (ans int) {
	n := len(nums)
	left := make([]int, n)
	st := []int{}
	for i, x := range nums {
		for len(st) > 0 && x <= nums[st[len(st)-1]] {
			st = st[:len(st)-1]
		}
		if len(st) > 0 {
			left[i] = st[len(st)-1]
		} else {
			left[i] = -1
		}
		st = append(st, i)
	}

	right := make([]int, n)
	st = st[:0]
	for i := n - 1; i >= 0; i-- {
		for len(st) > 0 && nums[i] <= nums[st[len(st)-1]] {
			st = st[:len(st)-1]
		}
		if len(st) > 0 {
			right[i] = st[len(st)-1]
		} else {
			right[i] = n
		}
		st = append(st, i)
	}
	// 枚举当前元素的左边的下一个更小元素和右边的下一个更小元素，则其上一个就是>=当前元素的
	for i, h := range nums {
		l, r := left[i], right[i]
		if l < k && k < r { // 相比 84 题多了个 if 判断
			ans = max(ans, h*((r-1)-(l+1)+1))
		}
	}
	return ans
}
