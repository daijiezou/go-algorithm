package _2_queueAndStack

/*
	特殊数据结构：单调栈
*/

// 单调栈的基础模板
func NextGreaterElement(nums []int) []int {
	length := len(nums)
	res := make([]int, length)
	s := make([]int, 0)
	for i := length - 1; i >= 0; i-- {
		for len(s) > 0 && s[len(s)-1] <= nums[i] {
			s = s[:len(s)-1]
		}
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
