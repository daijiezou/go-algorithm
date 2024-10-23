package _3_monotone_stack

import (
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

func maxWidthRamp(nums []int) int {
	stack := []int{}
	ans := 0
	// First pass: populate the stack
	for i, x := range nums {
		if len(stack) == 0 || nums[stack[len(stack)-1]] > x {
			stack = append(stack, i)
		}
	}
	// Second pass: calculate the maximum width ramp

	//j := len(nums) - 1
	for j := len(nums) - 1; j >= 0; j-- {
		for len(stack) > 0 && nums[j] >= nums[stack[len(stack)-1]] {
			ans = max(ans, j-stack[len(stack)-1])
			stack = stack[:len(stack)-1] // pop the last element
		}
	}
	//for len(stack) > 0 && j >= 0 {
	//	for j >= 0 && nums[j] < nums[stack[len(stack)-1]] {
	//		j--
	//	}
	//	if j >= 0 {
	//		ans = max(ans, j-stack[len(stack)-1])
	//		stack = stack[:len(stack)-1] // pop the last element
	//	}
	//}

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
