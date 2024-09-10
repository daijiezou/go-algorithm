package _1_huadongchuangkou

import "math"

func minSubArrayLen(target int, nums []int) int {
	sum := 0
	left := 0
	res := math.MaxInt
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		for sum >= target {
			res = min(res, i-left+1)
			sum -= nums[left]
			left++
		}
	}
	if res == math.MaxInt {
		return 0
	}
	return res
}

// https://leetcode.cn/problems/shortest-subarray-to-be-removed-to-make-array-sorted/
func findLengthOfShortestSubarray(arr []int) int {
	n := len(arr)
	right := n - 1
	for right > 0 && arr[right-1] <= arr[right] {
		right--
	}
	// 已经是非递增数组
	if right == 0 {
		return 0
	}
	ans := right // 删除arr[0:right]
	// 枚举左端点，移动右端点
	for left := 0; left == 0 || arr[left-1] <= arr[left]; left++ {
		for ; right < n && arr[right] < arr[left]; right++ {

		}
		ans = min(right-left-1, ans) // 删除arr[left+1:right]
	}
	return ans
}
