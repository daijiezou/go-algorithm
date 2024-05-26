package _2_array

import (
	"fmt"
)

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

/*
1、当可替换次数大于等于 0 时，扩大窗口，让进入窗口的 0 都变成 1，使得连续的 1 的长度尽可能大。
2、当可替换次数小于 0 时，缩小窗口，空余出可替换次数，以便继续扩大窗口。
3、只要可替换次数大于等于 0，窗口中的元素都会被替换成 1，也就是连续为 1 的子数组，我们想求的就是最大窗口长度。
*/
func longestOnes(nums []int, k int) int {
	n := len(nums)
	zeroCnt := 0
	left, right := 0, 0
	maxLength := -1
	for right < n {
		current := nums[right]
		right++
		if current == 0 {
			zeroCnt++
		}
		for zeroCnt > k {
			leave := nums[left]
			left++
			if leave == 0 {
				zeroCnt--
			}
		}
		if right-left > maxLength {
			maxLength = right - left
		}
	}
	return maxLength
}

// 给你一个字符串 s 和一个整数 k 。你可以选择字符串中的任一字符，并将其更改为任何其他大写英文字符。该操作最多可执行 k 次。
//
// 在执行上述操作后，返回 包含相同字母的最长子字符串的长度。
func characterReplacement(s string, k int) int {
	window := make(map[byte]int)
	left, right := 0, 0
	maxLength := -1
	for right < len(s) {
		current := s[right]
		right++
		if _, ok := window[current]; ok {
			window[current]++
		} else {
			window[current] = 1
		}

		//
		for len(window) > 1 && needReplaceCnt(window) > k && left < right {
			le := s[left]
			left++
			window[le]--
			if window[le] == 0 {
				delete(window, le)
			}
		}
		fmt.Println(left, right)
		fmt.Println(window)
		if right-left > maxLength {
			maxLength = right - left
		}
	}
	return maxLength
}

func needReplaceCnt(window map[byte]int) int {
	maxCount := 0
	allCnt := 0
	for _, v := range window {
		if v > maxCount {
			maxCount = v

		}
		allCnt += v
	}
	return allCnt - maxCount
}

func characterReplacement2(s string, k int) int {
	window := make(map[byte]int)
	left, right := 0, 0
	maxLength := -1
	maxCount := 0
	for right < len(s) {
		current := s[right]
		right++
		if _, ok := window[current]; ok {
			window[current]++
		} else {
			window[current] = 1
		}
		if window[current] > maxCount {
			maxCount = window[current]
		}

		for right-left-maxCount > k {
			le := s[left]
			left++
			window[le]--
			if window[le] == 0 {
				delete(window, le)
			}
		}
		if right-left > maxLength {
			maxLength = right - left
		}
	}
	return maxLength
}

/*
给你一个整数数组 nums 和一个整数 k ，判断数组中是否存在两个 不同的索引 i 和 j ，
满足 nums[i] == nums[j] 且 abs(i - j) <= k 。如果存在，返回 true ；否则，返回 false 。
*/
func containsNearbyDuplicate(nums []int, k int) bool {
	n := len(nums)
	left, right := 0, 0
	window := make(map[int]bool)
	for right < n {
		if window[nums[right]] {
			return true
		}
		window[nums[right]] = true
		right++
		for right-left > k && left < right {
			delete(window, nums[left])
			left++
		}
	}
	return false
}

/*
 */
func containsNearbyAlmostDuplicate(nums []int, indexDiff int, valueDiff int) bool {
	n := len(nums)
	left, right := 0, 0
	window := make([]int, 0)
	for right < n {
		window = append(window, nums[right])
		// check valueDiff

		right++
		for right-left > indexDiff && left < right {
			window = window[1:]
			left++
		}
	}
	return false
}

func longestSubstring(s string, k int) int {
	res := 0
	for i := 1; i <= 26; i++ {
		tempRes := logestKLetterSubstr(s, k, i)
		if tempRes > res {
			res = tempRes
		}
	}
	return res
}

// 在 s 中寻找仅含有 count 种字符，且每种字符出现次数都大于 k 的最长子串
func logestKLetterSubstr(s string, k int, count int) int {
	n := len(s)
	left, right := 0, 0
	window := [26]int{}
	validCnt := 0
	uniqueCnt := 0
	res := 0
	for right < n {
		current := s[right] - 'a'
		right++
		if window[current] == 0 {
			uniqueCnt++
		}
		window[current]++
		if window[current] == k {
			validCnt++
		}

		// 当窗口中字符种类大于K时，缩小窗口
		for uniqueCnt > count {
			d := s[left] - 'a'
			left++

			if window[d] == k {
				validCnt--
			}
			window[d]--
			if window[d] == 0 {
				uniqueCnt--
			}
		}
		if validCnt == count {
			if right-left > res {
				res = right - left
			}
		}
	}
	return res
}
