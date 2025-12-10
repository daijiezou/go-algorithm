package _4_datastruct

// https://leetcode.cn/problems/subarray-sum-equals-k/
/*
给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的子数组的个数 。

子数组是数组中元素的连续非空序列。
*/
func subarraySum2(nums []int, k int) int {
	n := len(nums)
	preSum := make([]int, n+1)
	preSum[0] = 0
	for i := 0; i < n; i++ {
		preSum[i+1] = preSum[i] + nums[i]
	}
	sum := 0
	cnts := make(map[int]int)
	for _, sj := range preSum {
		target := sj - k
		// 计算有多少个到当前的index的和为target
		sum += cnts[target]
		cnts[sj]++
	}
	return sum
}
