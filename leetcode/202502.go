package leetcode

import "slices"

func maxCount(m int, n int, ops [][]int) int {
	minRow := m
	minCol := n
	for i := 0; i < len(ops); i++ {
		row := ops[i][0]
		col := ops[i][1]
		minRow = min(minRow, row)
		minCol = min(minCol, col)
	}
	return minCol * minRow
}

func validPalindrome(s string) bool {
	left := 0
	right := len(s) - 1
	for left < right {
		if s[left] != s[right] {
			return checkPalindrome(left+1, right, s) || checkPalindrome(left, right-1, s)
		}
		left++
		right--
	}
	return true
}

func checkPalindrome(left, right int, s string) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func sortArrayByParityII(nums []int) []int {
	n := len(nums)
	//nums2 := make([]int, n)
	jIndex := 1
	oIndex := 0
	for oIndex < n && jIndex < n {
		if nums[jIndex]%2 == 1 {
			jIndex += 2
		} else if nums[oIndex]%2 == 0 {
			oIndex += 2
		} else {
			nums[jIndex], nums[oIndex] = nums[oIndex], nums[jIndex]
			jIndex += 2
			oIndex += 2
		}
	}
	return nums
}

func subsetsWithDup(nums []int) [][]int {
	slices.Sort(nums)
	res := [][]int{}
	var backtrack func(start int, cur []int)
	backtrack = func(start int, cur []int) {
		temp := make([]int, len(cur))
		copy(temp, cur)
		res = append(res, temp)
		for i := start; i < len(nums); i++ {
			if i > start && nums[i] == nums[i-1] {
				continue
			}
			cur = append(cur, nums[i])
			backtrack(i+1, cur)
			cur = cur[:len(cur)-1]
		}
	}
	backtrack(0, []int{})
	return res
}

func permuteUnique(nums []int) [][]int {
	slices.Sort(nums)
	n := len(nums)
	res := [][]int{}
	var backtrack func(cur []int)
	used := make([]bool, n)
	backtrack = func(cur []int) {
		if len(cur) == n {
			temp := make([]int, len(cur))
			copy(temp, cur)
			res = append(res, temp)
			return
		}

		for i := 0; i < len(nums); i++ {
			// 当出现重复元素时，比如输入 nums = [1,2,2',2'']，
			// 2' 只有在 2 已经被使用的情况下才会被选择，
			// 同理，2'' 只有在 2' 已经被使用的情况下才会被选择，这就保证了相同元素在排列中的相对位置保证固定。
			if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}
			if used[i] {
				continue
			}
			used[i] = true
			cur = append(cur, nums[i])
			backtrack(cur)
			cur = cur[:len(cur)-1]
			used[i] = false
		}
	}
	backtrack([]int{})
	return res
}

func removeDuplicates(nums []int) int {
	fast := 0
	slow := 0
	cnt := 0
	n := len(nums)
	for fast < n {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		} else if slow < fast && cnt < 3 {
			slow++
			nums[slow] = nums[fast]
		}
		cnt++
		fast++
		if fast < n && nums[fast] != nums[fast-1] {
			cnt = 0
		}
	}
	return slow + 1
}

func removeDuplicates_lin(nums []int) int {
	stackSize := 2 //数组的前两个元素默认保留
	for i := 2; i < len(nums); i++ {
		if nums[i] != nums[stackSize-2] {
			nums[stackSize] = nums[i]
			stackSize++
		}
	}
	return min(len(nums), stackSize)
}
