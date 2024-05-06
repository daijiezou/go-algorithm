package _2_array

import "sort"

func removeDuplicates(nums []int) int {
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != nums[slow] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow + 1
}

func removeElement(nums []int, val int) int {
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}

// 两数之和
func twoSum(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		sum := nums[left] + nums[right]
		if sum == target {
			return []int{left + 1, right + 1}
		}
		if sum > target {
			right--
		}
		if sum < target {
			left++
		}
	}
	return []int{-1, -1}
}

func twoSumTarget(nums []int, start int, target int) [][]int {
	left, right := start, len(nums)-1
	res := make([][]int, 0)
	for left < right {
		sum := nums[left] + nums[right]
		leftVal := nums[left]
		rightVal := nums[right]
		if sum == target {
			res = append(res, []int{nums[left], nums[right]})
			for left < right && leftVal == nums[left] {
				left++
			}
			for left < right && rightVal == nums[right] {
				right--
			}
		}
		if sum > target {
			right--
		}
		if sum < target {
			left++
		}
	}
	return res
}

func threeSumTarget(nums []int, target int) [][]int {
	// 输入数组 nums，返回所有和为 target 的三元组
	sort.Ints(nums)
	length := len(nums)
	res := make([][]int, 0)
	for i := 0; i < length; i++ {

		// 这里需要指定起始，避免重复的答案
		tuples := twoSumTarget(nums, i+1, target-nums[i])
		for _, tuple := range tuples {
			tuple = append(tuple, nums[i])
			res = append(res, tuple)
		}

		// 这里针对首个数字的去重必须放在后面，避免遗漏答案。
		for i < length-1 && nums[i] == nums[i+1] {
			i++
		}
	}
	return res
}
