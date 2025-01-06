package meeting150

import (
	"slices"
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	left := 0
	right := len(s) - 1
	for left < right {
		for left < right && !(s[left] >= 'a' && s[left] <= 'z' || s[left] >= '0' && s[left] <= '9') {
			left++
		}
		for left < right && !(s[right] >= 'a' && s[right] <= 'z' || s[right] >= '0' && s[right] <= '9') {
			right--
		}
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

//	给定字符串 s 和 t ，判断 s 是否为 t 的子序列。
//	字符串的一个子序列是原始字符串删除一些（也可以不删除）字符而不改变剩余字符相对位置形成的新字符串。
//
// （例如，"ace"是"abcde"的一个子序列，而"aec"不是）。
func isSubsequence(s string, t string) bool {
	index := 0
	cnt := 0
loop1:
	for i := 0; i < len(s); i++ {
		for j := index; j < len(t); j++ {
			if s[i] == t[j] {
				index = j
				cnt++
				continue loop1
			}
		}
	}
	return cnt == len(s)
}

// 盛雨水最多的容器

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1

	res := 0
	for left < right {
		area := min(height[left], height[right]) * (right - left)
		res = max(res, area)
		// 此时左边界比较小，
		// 如果此时不移动左边界移动右边界的话，无论右边界的高度是多少，都不会超过当前的容器盛的量
		// 所以此时该边界不会再作为容器的边界了，我们可以移动他找到最佳的容器
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return res
}

// 三数之和
func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	res := [][]int{}
	for i := 0; i < len(nums); i++ {
		twoRes := twoSum(nums, i+1, 0-nums[i])
		for _, two := range twoRes {
			res = append(res, append([]int{nums[i]}, two...))
		}
		for i < len(nums)-1 && nums[i] == nums[i+1] {
			i++
		}
	}
	return res
}

func twoSum(nums []int, start int, target int) [][]int {
	left := start
	right := len(nums) - 1
	res := [][]int{}
	for left < right {
		leftVal := nums[left]
		rightVal := nums[right]
		sum := leftVal + rightVal
		if sum > target {
			right--
		} else if sum < target {
			left++
		} else {
			res = append(res, []int{leftVal, rightVal})
			for left < right && nums[left] == leftVal {
				left++
			}
			for left < right && nums[right] == rightVal {
				right--
			}
		}
	}
	return res
}
