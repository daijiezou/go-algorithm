package _1_huadongchuangkou

import (
	"slices"
)

/*
2.3.1 越短越合法
*/

// https://leetcode.cn/problems/count-subarrays-where-max-element-appears-at-least-k-times/
func countSubarrays(nums []int, k int) int64 {
	res := 0
	xMax := slices.Max(nums)
	left, right := 0, 0
	countMax := 0
	for right < len(nums) {
		if nums[right] == xMax {
			countMax++
		}
		right++
		for countMax >= k {
			if nums[left] == xMax {
				countMax--
			}
			left++
		}
		//该数组满足，包括left之前的数组也都满足

	}
	return int64(res)
}

// https://leetcode.cn/problems/subarray-product-less-than-k/
func numSubarrayProductLessThanK(nums []int, k int) int {
	res := 0
	left, right := 0, 0
	cheng := 1
	for ; right < len(nums); right++ {
		cheng *= nums[right]
		for cheng >= k && left < right {
			cheng /= nums[left]
			left++
		}
		// 现在必然是一个合法的窗口，但注意思考这个窗口中的子数组个数怎么计算：
		// 比方说 left = 1, right = 4 划定了 [1, 2, 3] 这个窗口（right 是开区间）
		// 但不止 [left..right] 是合法的子数组，[left+1..right], [left+2..right] 等都是合法子数组
		// 所以我们需要把 [3], [2,3], [1,2,3] 这 right - left 个子数组都加上
		res += right - left + 1
	}
	return res
}

/*
给你一个整数数组 nums 和一个整数 k ，请你返回 nums 中 好 子数组的数目。
一个子数组 arr 如果有 至少 k 对下标 (i, j) 满足 i < j 且 arr[i] == arr[j] ，那么称它是一个 好 子数组。
子数组 是原数组中一段连续 非空 的元素序列。
*/
func countGood(nums []int, k int) int64 {
	n := len(nums)
	left, right := 0, 1
	res := 0
	count := 0
	window := make(map[int]int)
	window[nums[0]]++
	for ; right < n; right++ {
		window[nums[right]]++
		count += window[nums[right]] - 1
		// 左侧窗口收缩，直到窗口内的元素个数小于 k
		for count >= k {
			count -= window[nums[left]] - 1
			window[nums[left]]--
			left++
		}
		res += left
	}
	return int64(res)
}
