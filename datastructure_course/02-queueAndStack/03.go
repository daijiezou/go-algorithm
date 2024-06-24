package _2_queueAndStack

/*
特殊数据结构：单调栈
单调栈实际上就是栈，
只是利用了一些巧妙的逻辑，使得每次新元素入栈后，栈内的元素都保持有序
*/

// 单调栈的基础模板
// 下一个更大的元素
func NextGreaterElement(nums []int) []int {
	length := len(nums)
	res := make([]int, length)
	stack := make([]int, 0)
	// 倒着往栈里放
	for i := length - 1; i >= 0; i-- {
		// 删掉 nums[i] 后面较小的元素
		for len(stack) > 0 && stack[len(stack)-1] <= nums[i] {
			stack = stack[:len(stack)-1]
		}
		// 现在栈顶就是 nums[i] 后面的更大元素
		if len(stack) == 0 {
			res[i] = 0
		} else {
			res[i] = stack[len(stack)-1]
		}
		stack = append(stack, nums[i])
	}
	return res
}

// 上一个更大的元素
// 之前我们的 for 循环都是从数组的尾部开始往栈里添加元素，
// 这样栈顶元素就是 nums[i] 之后的元素。所以只要我们从数组的头部开始往栈里添加元素，栈顶的元素就是 nums[i] 之前的元素，即可计算 nums[i] 的上一个更大元素。
func PreGreaterElement(nums []int) []int {
	length := len(nums)
	res := make([]int, length)
	s := make([]int, 0)
	for i := 0; i < length; i++ {
		// 删掉 nums[i] 前面较小的元素
		for len(s) > 0 && s[len(s)-1] <= nums[i] {
			s = s[:len(s)-1]
		}
		// 现在栈顶就是 nums[i] 前面的更大元素
		if len(s) == 0 {
			res[i] = 0
		} else {
			res[i] = s[len(s)-1]
		}
		s = append(s, nums[i])
	}
	return res
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	num2Res := make([]int, len(nums2))
	nums2GreaterMap := make(map[int]int, len(nums2))
	s := make([]int, 0)
	length := len(nums2)
	for i := length - 1; i >= 0; i-- {
		for 0 < len(s) && s[len(s)-1] <= nums2[i] {
			s = s[:len(s)-1]
		}
		if len(s) == 0 {
			num2Res[i] = -1
		} else {
			num2Res[i] = s[len(s)-1]
		}
		s = append(s, nums2[i])
		nums2GreaterMap[nums2[i]] = num2Res[i]
	}
	res := make([]int, len(nums1))
	for i := 0; i < len(nums1); i++ {
		res[i] = nums2GreaterMap[nums1[i]]
	}
	return res
}

// https://leetcode.cn/problems/daily-temperatures/description/
/*
给定一个整数数组 temperatures ，表示每天的温度，返回一个数组 answer，
其中 answer[i] 是指对于第 i 天，下一个更高温度出现在几天后。
如果气温在这之后都不会升高，请在该位置用 0 来代替。
*/
func dailyTemperatures(temperatures []int) []int {
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
		// 这里存放的是元素的下标
		s = append(s, i)
	}
	return res
}
