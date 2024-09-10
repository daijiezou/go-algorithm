package _1_huadongchuangkou

import "math"

func maximumLengthSubstring(s string) int {
	window := make(map[byte]int)
	ans := math.MinInt32

	left, right := 0, 0
	for right < len(s) {
		window[s[right]]++
		for window[s[right]] > 2 {
			window[s[left]]--
			left++
		}
		ans = max(ans, right-left+1)
		right++
	}
	return ans
}

// https://leetcode.cn/problems/longest-subarray-of-1s-after-deleting-one-element/
/*
给你一个二进制数组 nums ，你需要从中删掉一个元素。

请你在删掉元素的结果数组中，返回最长的且只包含 1 的非空子数组的长度。

如果不存在这样的子数组，请返回 0 。
*/
func longestSubarray(nums []int) int {
	left, right := 0, 0
	ans := 0
	zeroCnt := 0
	for right < len(nums) {
		if nums[right] == 0 {
			zeroCnt++
		}
		for zeroCnt > 1 {
			if nums[left] == 0 {
				zeroCnt--
			}
			left++
		}
		ans = max(ans, right-left)
		right++
	}
	return ans
}

func equalSubstring(s string, t string, maxCost int) int {
	left, right := 0, 0
	curCost := 0
	ans := 0
	for right < len(s) {
		curCost += myAbs(s[right], t[right])
		for curCost > maxCost {
			curCost -= myAbs(s[left], t[left])
			left++
		}
		ans = max(ans, right-left+1)
		right++
	}
	return ans
}

func myAbs(a, b uint8) int {
	if a > b {
		return int(a - b)
	}
	return (int)(b - a)
}

func longestSemiRepetitiveSubstring(s string) int {
	ans := 1
	left, right := 0, 1
	moreThanOne := 0
	for right < len(s) {
		if s[right] == s[right-1] {
			moreThanOne++
		}
		for moreThanOne > 1 {
			if s[left] == s[left+1] {
				moreThanOne--
			}
			left++
		}
		ans = max(ans, right-left+1)
		right++
	}
	return ans
}

// https://leetcode.cn/problems/fruit-into-baskets/
func totalFruit(fruits []int) int {
	left, right := 0, 0
	ans := 0
	window := make(map[int]int)
	for right < len(fruits) {
		window[fruits[right]]++
		for len(window) > 2 {
			window[fruits[left]]--
			if window[fruits[left]] == 0 {
				delete(window, fruits[left])
			}
			left++
		}
		ans = max(ans, right-left+1)
		right++
	}
	return ans
}

func maximumUniqueSubarray(nums []int) int {
	left, right := 0, 0
	ans := 0
	window := make(map[int]int)
	cur := 0
	for right < len(nums) {
		window[nums[right]]++
		cur += nums[right]
		for window[nums[right]] > 1 {
			cur -= nums[left]
			window[nums[left]]--
			left++
		}
		ans = max(ans, cur)
		right++
	}
	return ans
}
