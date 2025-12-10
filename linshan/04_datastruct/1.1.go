package _4_datastruct

import (
	"math"
	"strings"
)

func subarraySum(nums []int) int {
	n := len(nums)
	preSum := make([]int, n+1)
	preSum[0] = 0
	for i := 0; i < n; i++ {
		preSum[i+1] = preSum[i] + nums[i]
	}
	total := 0
	for i := 0; i < n; i++ {
		start := max(0, i-nums[i])
		total += (preSum[i+1] - preSum[start])
	}
	return total
}

// https://leetcode.cn/problems/count-vowel-strings-in-ranges/
func vowelStrings(words []string, queries [][]int) []int {
	n := len(words)
	preSum := make([]int, n+1)
	preSum[0] = 0
	for i := 0; i < n; i++ {
		word := words[i]
		add := 0
		if strings.Contains("aeiou", string(word[0])) &&
			strings.Contains("aeiou", string(word[len(word)-1])) {
			add = 1
		}
		preSum[i+1] = preSum[i] + add
	}
	res := make([]int, len(queries))
	for i := 0; i < len(queries); i++ {
		res[i] = preSum[queries[i][1]+1] - preSum[queries[i][0]]
	}
	return res
}

// https://leetcode.cn/problems/special-array-ii/
func isArraySpecial(nums []int, queries [][]int) []bool {
	n := len(nums)
	preSum := make([]int, n)
	preSum[0] = 0
	for i := 1; i < n; i++ {
		preSum[i] = preSum[i-1]
		if nums[i]%2 == nums[i-1]%2 {
			preSum[i]++
		}
	}
	res := make([]bool, len(queries))
	for i := 0; i < len(queries); i++ {
		start := queries[i][0]
		end := queries[i][1]
		// 说明中间全部符合要求
		if preSum[start] == preSum[end] {
			res[i] = true
		}
	}
	return res
}

// https://leetcode.cn/problems/maximum-absolute-sum-of-any-subarray/
/*
给你一个整数数组 nums 。一个子数组 [numsl, numsl+1, ..., numsr-1, numsr] 的 和的绝对值 为 abs(numsl + numsl+1 + ... + numsr-1 + numsr) 。

请你找出 nums 中 和的绝对值 最大的任意子数组（可能为空），并返回该 最大值 。

abs(x) 定义如下：

如果 x 是负整数，那么 abs(x) = -x 。
如果 x 是非负整数，那么 abs(x) = x 。
*/
func maxAbsoluteSum(nums []int) int {
	n := len(nums)
	sum := 0
	sum2 := 0
	res := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
		if sum < 0 {
			sum = 0
		}
		sum2 += nums[i]
		if sum2 > 0 {
			sum2 = 0
		}
		res = max(sum, -sum2, res)
	}
	return res

}

//https://leetcode.cn/problems/maximum-subarray/description/
/*
给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

子数组是数组中的一个连续部分。
*/
func maxSubArray(nums []int) int {
	sum := 0
	res := math.MinInt
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if sum < 0 {
			sum = 0
		}
		res = max(res, sum)
	}
	if res == 0 {
		res = math.MinInt
		for i := 0; i < len(nums); i++ {
			res = max(res, nums[i])
		}
	}
	return res
}

func maxSubArray2(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	n := len(nums)
	dp := make([]int, n)
	res := nums[0]
	dp[0] = nums[0]
	for i := 1; i < n; i++ {
		if dp[i-1] > 0 {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}
		res = max(res, dp[i])
	}
	return res
}
