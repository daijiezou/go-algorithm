package _3_monotone_stack

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
