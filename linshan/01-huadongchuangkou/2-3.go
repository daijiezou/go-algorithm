package _1_huadongchuangkou

import "slices"

func numSubarrayProductLessThanK(nums []int, k int) int {
	res := 0
	left, right := 0, 0
	cheng := 1
	for right < len(nums) {
		cheng *= nums[right]
		right++
		for cheng >= k && left < right {
			cheng /= nums[left]
			left++
		}

		// 现在必然是一个合法的窗口，但注意思考这个窗口中的子数组个数怎么计算：
		// 比方说 left = 1, right = 4 划定了 [1, 2, 3] 这个窗口（right 是开区间）
		// 但不止 [left..right] 是合法的子数组，[left+1..right], [left+2..right] 等都是合法子数组
		// 所以我们需要把 [3], [2,3], [1,2,3] 这 right - left 个子数组都加上
		res += right - left
	}
	return res
}

func numberOfSubstrings(s string) int {
	res := 0
	window := make(map[byte]int, 0)
	left := 0
	for i := 0; i < len(s); i++ {
		window[s[i]]++
		for len(window) == 3 {
			// 该子数组满足，后面的包含该子数组的都满足
			res += len(s) - i
			window[s[left]]--
			left++
			if window[s[left]] == 0 {
				delete(window, s[left])
			}
		}
	}
	return res
}

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
		res += left
	}
	return int64(res)
}
