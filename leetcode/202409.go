package leetcode

import "sort"

func busyStudent(startTime []int, endTime []int, queryTime int) int {
	n := len(startTime)
	res := 0
	for i := 0; i < n; i++ {
		if startTime[i] <= queryTime && queryTime <= endTime[i] {
			res++
		}
	}
	return res
}

func maxConsecutiveAnswers(answerKey string, k int) int {
	n := len(answerKey)
	left, right := 0, 0
	maxConsecutive := 0
	countT := 0
	countF := 0
	for right < n {
		if answerKey[right] == 'T' {
			countT++
		} else {
			countF++
		}
		right++
		for countT > k && countF > k {
			if answerKey[left] == 'T' {
				countT--
			} else {
				countF--
			}
			left++
		}
		maxConsecutive = max(maxConsecutive, right-left)
	}
	return maxConsecutive
}

func maxStrength(nums []int) int64 {
	if len(nums) == 1 {
		return int64(nums[0])
	}
	minIndex := 0
	minNum := -10
	count := 0
	zeroCnt := 0
	res := int64(1)
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			zeroCnt++
			continue
		}
		res *= int64(nums[i])
		if nums[i] < 0 {
			count++
			if nums[i] > minNum {
				minNum = nums[i]
				minIndex = i
			}
		}
	}

	// 当数组不包含正数，且负数元素小于等于 1 个时，最大积为 0。
	if zeroCnt+1 == len(nums) && count == 1 || zeroCnt == len(nums) {
		return 0
	}

	if count%2 == 0 {
		return res
	} else {
		return res / int64(nums[minIndex])
	}
}

/*
如果能够满足下述两个条件之一，则认为第 i 位学生将会保持开心：
这位学生被选中，并且被选中的学生人数 严格大于 nums[i] 。
这位学生没有被选中，并且被选中的学生人数 严格小于 nums[i] 。
*/
func countWays(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	res := 0
	// 可以都不选
	if nums[0] > 0 {
		res++
	}
	for i := 1; i < n; i++ {
		// i代表被选中的人数

		if nums[i-1] < i && // 被选中的学生人数 严格大于 nums[i]
			i < nums[i] { // 被选中的学生人数 严格小于 nums[i]
			res++
		}
	}

	// 0 <= nums[i] < nums.length
	// 一定可以都选
	return res + 1
}
