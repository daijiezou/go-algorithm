package _2_queueAndStack

import "math"

// https://leetcode.cn/problems/longest-continuous-subarray-with-absolute-diff-less-than-or-equal-to-limit/description/
// 绝对差不超过限制的最长连续子数组
func longestSubarray(nums []int, limit int) int {
	length := len(nums)
	var minQ, maxQ []int
	var res int
	left := 0
	for right := 0; right < length; right++ {
		for len(minQ) > 0 && nums[right] < minQ[len(minQ)-1] {
			minQ = minQ[:len(minQ)-1]
		}
		minQ = append(minQ, nums[right])
		for len(maxQ) > 0 && nums[right] > maxQ[len(maxQ)-1] {
			maxQ = maxQ[:len(maxQ)-1]
		}
		maxQ = append(maxQ, nums[right])
		for len(minQ) > 0 && len(maxQ) > 0 && maxQ[0]-minQ[0] > limit {
			if nums[left] == minQ[0] {
				minQ = minQ[1:]
			}
			if nums[left] == maxQ[0] {
				maxQ = maxQ[1:]
			}
			left++
		}
		res = max(res, right-left+1)
	}
	return res
}

func shortestSubarray(nums []int, k int) int {
	n := len(nums)
	preSum := make([]int, n+1)
	preSum[0] = 0
	// 计算 nums 的前缀和数组
	for i := 1; i <= n; i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
	}
	// 单调队列结构辅助滑动窗口算法,window是一个单调递增的队列，队首是最小的
	window := []int{}
	right := 0
	length := math.MaxInt64
	// 开始执行滑动窗口算法框架
	for right < len(preSum) {
		// 若新进入窗口的元素和窗口中的最小值之差大于等于 k，
		// 说明得到了符合条件的子数组，缩小窗口，使子数组长度尽可能小
		for len(window) > 0 && preSum[right]-preSum[window[0]] >= k {
			// 更新答案
			length = min(length, right-window[0])
			// 缩小窗口
			window = window[1:]
		}
		for len(window) > 0 && preSum[window[len(window)-1]] >= preSum[right] {
			window = window[:len(window)-1]
		}
		// 扩大窗口，元素入队
		window = append(window, right)
		right++
	}
	if length == math.MaxInt64 {
		return -1
	}
	return length
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxSubarraySumCircular(nums []int) int {
	n := len(nums)
	// 模拟环状的 nums 数组
	preSum := make([]int, 2*n+1)
	preSum[0] = 0
	// 计算环状 nums 的前缀和
	for i := 1; i < len(preSum); i++ {
		preSum[i] = preSum[i-1] + nums[(i-1)%n]
	}
	window := make([]int, 1) //保证队首永远都是最小的单调递增队列
	window[0] = 0
	maxSum := math.MinInt32
	for right := 1; right < len(preSum); right++ {
		// 超过一个循环数组，将其移除
		for len(window) > 0 && window[0] < right-n {
			window = window[1:]
		}
		maxSum = max(maxSum, preSum[right]-preSum[window[0]])
		for len(window) > 0 && preSum[right] <= preSum[window[len(window)-1]] {
			window = window[:len(window)-1]
		}
		window = append(window, right)
	}
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
