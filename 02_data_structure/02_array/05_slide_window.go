package _2_array

/*
给你一个整数数组 nums 和一个整数 x 。
每一次操作时，你应当移除数组 nums 最左边或最右边的元素，
然后从 x 中减去该元素的值。请注意，需要 修改 数组以供接下来的操作使用。

如果可以将 x 恰好 减到 0 ，返回 最小操作数 ；否则，返回 -1 。
*/
func minOperations(nums []int, x int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	// 相当于找到最长的子数组之为target的子数组
	target := sum - x
	manLen := -1
	left, right := 0, 0
	windowSum := 0
	for right < len(nums) {
		windowSum += nums[right]
		right++
		for windowSum > target && left < right {
			windowSum -= nums[left]
			left++
		}
		if windowSum == target {
			// 左闭又开区间[left,right)
			if right-left > manLen {
				manLen = right - left
			}
		}
	}
	if manLen == -1 {
		return -1
	}
	// 操作数=数组总长度-最长子数组长度
	return len(nums) - manLen
}

func numSubarrayProductLessThanK(nums []int, k int) int {
	n := len(nums)
	if len(nums) < 1 {
		return 0
	}
	count := 0
	left, right := 0, 0
	windowProduct := 1
	for right < n {
		windowProduct *= nums[right]
		right++
		for windowProduct >= k && left < right {
			windowProduct /= nums[left]
			left++
		}
		// 这个地方还是有点吊，不得不说，算法还是需要一点灵感的
		// 现在必然是一个合法的窗口，但注意思考这个窗口中的子数组个数怎么计算：
		// 比方说 left = 1, right = 4 划定了 [1, 2, 3] 这个窗口（right 是开区间）
		// 但不止 [left..right] 是合法的子数组，[left+1..right], [left+2..right] 等都是合法子数组
		// 所以我们需要把 [3], [2,3], [1,2,3] 这 right - left 个子数组都加上
		count += right - left
	}
	return count
}
