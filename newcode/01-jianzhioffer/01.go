package _1_jianzhioffer

func duplicate(numbers []int) int {
	// write code here
	cnt := make(map[int]struct{})
	for i := 0; i < len(numbers); i++ {
		if _, ok := cnt[numbers[i]]; ok {
			return numbers[i]
		}
		cnt[numbers[i]] = struct{}{}
	}
	return -1
}

func duplicate2(numbers []int) int {
	// write code here
	for i := 0; i < len(numbers); i++ {
		for numbers[i] != i {
			if numbers[i] == numbers[numbers[i]] {
				return numbers[i]
			}
			numbers[i], numbers[numbers[i]] = numbers[numbers[i]], numbers[i]
		}
	}
	return -1
}

/*
在一个长度为n+1的数组里的所有数字都在1～n的范围内，所以数组
中至少有一个数字是重复的。请找出数组中任意一个重复的数字，但不能
修改输入的数组。例如，如果输入长度为8的数组 2,3,5,4,3,2,6,7那
么对应的输出是重复的数字2或者3。
*/

func getDup(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	left := 1
	right := len(nums) - 1
	for right >= left {
		mid := left + (right-left)/2
		cnt := countRange(left, mid, nums)
		if left == right {
			if cnt > 1 {
				return left
			} else {
				return -1
			}
		}
		// 说明重复数字在left到mid之间
		if cnt > mid-left+1 {
			// 这里right不能等于mid-1，因为有可能mid是重复的数字
			right = mid
		} else {
			// 说明重复数字在mid+1到right之间
			left = mid + 1
		}
	}
	return -1
}

func countRange(left, right int, nums []int) int {
	cnt := 0
	for i := 0; i < len(nums); i++ {
		if left <= nums[i] && nums[i] <= right {
			cnt++
		}
	}
	return cnt
}
