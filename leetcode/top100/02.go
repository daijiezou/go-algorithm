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
