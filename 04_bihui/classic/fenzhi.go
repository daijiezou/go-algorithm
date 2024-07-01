package classic

import "strconv"

func diffWaysToCompute(expression string) []int {
	res := make([]int, 0)
	for i := 0; i < len(expression); i++ {
		if expression[i] == '+' || expression[i] == '-' || expression[i] == '*' {
			left := diffWaysToCompute(expression[0:i])
			right := diffWaysToCompute(expression[i+1:])
			for j := 0; j < len(left); j++ {
				for k := 0; k < len(right); k++ {
					switch expression[i] {
					case '+':
						res = append(res, left[j]+right[k])
					case '-':
						res = append(res, left[j]-right[k])
					case '*':
						res = append(res, left[j]*right[k])
					}
				}
			}
		}
	}
	// baseCase 只有数字的情况
	if len(res) == 0 {
		numVal, _ := strconv.Atoi(expression)
		res = append(res, numVal)
	}
	return res
}

func MergeSort(nums []int) []int {
	// base case
	if len(nums) < 2 {
		return nums
	}

	/****** 分 ******/
	// 对数组的两部分分别排序
	length := len(nums)
	mid := length / 2
	left := MergeSort(nums[:mid])
	right := MergeSort(nums[mid:])

	/****** 治 ******/
	// 合并两个排好序的子数组
	newNums := make([]int, length)
	leftLength := len(left)
	rightLength := len(right)
	var lindex, rindex int
	var index int
	for lindex < leftLength && rindex < rightLength {
		if left[lindex] < right[rindex] {
			newNums[index] = left[lindex]
			lindex++
		} else {
			newNums[index] = right[rindex]
			rindex++
		}
		index++
	}
	for lindex < leftLength {
		newNums[index] = left[lindex]
		lindex++
		index++
	}
	for rindex < rightLength {
		newNums[index] = right[rindex]
		rindex++
		index++
	}
	return newNums
}
