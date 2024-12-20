package top100

import (
	"slices"
)

// https://leetcode.cn/problems/move-zeroes/?envType=study-plan-v2&envId=top-100-liked
func moveZeroes(nums []int) {
	var slow, fast int
	for ; fast < len(nums); fast++ {
		if nums[fast] != 0 {
			nums[slow] = nums[fast]
			slow++
		}
	}
	for i := slow; i < len(nums); i++ {
		nums[i] = 0
	}

}

// https://leetcode.cn/problems/container-with-most-water/?envType=study-plan-v2&envId=top-100-liked
func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	res := 0
	for left < right {
		h := min(height[left], height[right])
		area := (right - left) * h
		res = max(area, res)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return res
}

func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	slices.Sort(nums)
	for i := 0; i < len(nums); i++ {
		target := 0 - nums[i]
		twos := twoSum2(nums, i+1, target)
		for _, v := range twos {
			res = append(res, append([]int{nums[i]}, v...))
		}
		for i < len(nums)-1 && nums[i] == nums[i+1] {
			i++
		}
	}
	return res
}

func twoSum2(nums []int, start int, target int) [][]int {
	res := make([][]int, 0)
	left := start
	right := len(nums) - 1
	for left < right {
		leftVal := nums[left]
		rightVal := nums[right]
		sum := nums[left] + nums[right]
		if sum == target {
			res = append(res, []int{leftVal, rightVal})
			for left < right && nums[left] == leftVal {
				left++
			}
			for left < right && nums[right] == rightVal {
				right--
			}
		} else if sum < target {
			left++

		} else {
			right--

		}
	}
	return res
}

// https://leetcode.cn/problems/trapping-rain-water/?envType=study-plan-v2&envId=top-100-liked
// 接雨水
func trap(height []int) int {
	n := len(height)
	leftMax := make([]int, n)
	rightMax := make([]int, n)
	lmax := height[0]
	for i := 0; i < n; i++ {
		lmax = max(lmax, height[i])
		leftMax[i] = lmax
	}
	rMax := height[n-1]
	for i := n - 1; i >= 0; i-- {
		rMax = max(rMax, height[i])
		rightMax[i] = rMax
	}
	res := 0
	for i := 0; i < len(height); i++ {
		h := min(leftMax[i], rightMax[n-i-1])
		res += (h - height[i])
	}
	return res
}

func lengthOfLongestSubstring(s string) int {
	window := make(map[byte]int, 0)
	res := 0
	left := 0
	for right := 0; right < len(s); right++ {
		cur := s[right]
		window[cur]++
		for window[cur] > 1 {
			leftVal := s[left]
			window[leftVal]--
			left++
		}
		res = max(res, right-left+1)
	}
	return res
}
