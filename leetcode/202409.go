package leetcode

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
