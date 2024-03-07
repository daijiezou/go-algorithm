package _1_listNodeAndArray

import "fmt"

func leftBound(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			right = mid - 1
		}
		if nums[mid] > target {
			right = mid - 1
		}
		if nums[mid] < target {
			left = mid + 1
		}
	}
	return left
}

func rightBound(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			left = mid + 1
		}
		if nums[mid] > target {
			right = mid - 1
		}
		if nums[mid] < target {
			left = mid + 1
		}
	}
	return right
}

// https://leetcode.cn/problems/search-a-2d-matrix/description/
// 搜索二维矩阵
func searchMatrix(matrix [][]int, target int) bool {
	left := 0
	right := len(matrix)*len(matrix[0]) - 1
	for left <= right {
		mid := left + (right-left)/2
		if getMatrix(matrix, mid) == target {
			return true
		}
		if getMatrix(matrix, mid) > target {
			right = mid - 1
		}
		if getMatrix(matrix, mid) < target {
			left = mid + 1
		}
	}
	return false
}

func getMatrix(matrix [][]int, index int) int {
	row := index / len(matrix[0])
	col := index % len(matrix[0])
	return matrix[row][col]
}

func searchMatrix2(matrix [][]int, target int) bool {
	rowLength := len(matrix)
	colLength := len(matrix[0])
	fmt.Println(rowLength)
	// 把指针初始化到右上角
	i := 0
	j := colLength - 1
	for i < rowLength && j >= 0 {
		if matrix[i][j] == target {
			return true
		}
		if matrix[i][j] < target {
			i++
		} else {
			j--
		}
	}
	return false
}

// https://leetcode.cn/problems/find-k-closest-elements/
// 找到K个最接近的元素
func findClosestElements(arr []int, k int, x int) []int {
	leftIndex := leftBound(arr, x)

	left := leftIndex - 1
	right := leftIndex
	res := make([]int, 0, k)
	for right-left-1 < k {
		if left == -1 {
			res = append(res, arr[right])
			right++
		} else if right == len(arr) {
			res = append([]int{arr[left]}, res...)
			left--
		} else if x-arr[left] > arr[right]-x {
			res = append(res, arr[right])
			right++
		} else {
			res = append([]int{arr[left]}, res...)
			left--
		}
	}
	return res
}

// https://leetcode.cn/problems/find-peak-element/
// 寻找峰值
func findPeakElement(nums []int) int {
	left := 0
	right := len(nums) - 1
	mid := left + (right-left)/2
	//
	for left < right {
		if nums[mid] > nums[mid+1] {
			// mid可能为峰值或其左侧为峰值
			// 收缩右边
			right = mid
		} else {
			// mid的右侧有峰值
			left = mid + 1
		}
	}
	return left
}
