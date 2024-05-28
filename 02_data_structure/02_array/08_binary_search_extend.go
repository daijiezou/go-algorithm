package _2_array

// 搜索二维矩阵
func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	// 把二维数组映射到一维
	left, right := 0, m*n-1
	for left <= right {
		mid := left + (right-left)/2
		if target < getValueFromMatrix(matrix, mid) {
			right = mid - 1
		} else if target > getValueFromMatrix(matrix, mid) {
			left = mid + 1
		} else if target == getValueFromMatrix(matrix, mid) {
			return true
		}
	}
	return false
}

func getValueFromMatrix(matrix [][]int, index int) int {
	n := len(matrix[0])
	// 计算二维中的横纵坐标
	i, j := index/n, index%n
	return matrix[i][j]
}

// https://leetcode.cn/problems/search-a-2d-matrix-ii/description/
func searchMatrix2(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	i := 0
	j := n - 1
	for i < m && j >= 0 {
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] > target {
			j--
		} else if matrix[i][j] < target {
			i++
		}
	}
	return false

}

func findPeakElement(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		// mid是峰值，或者mid左边是峰值
		if nums[mid] > nums[mid+1] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func peakIndexInMountainArray(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		// mid是峰值，或者mid左边是峰值
		if nums[mid] > nums[mid+1] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

// https://leetcode.cn/problems/zai-pai-xu-shu-zu-zhong-cha-zhao-shu-zi-lcof/
//某班级考试成绩按非严格递增顺序记录于整数数组 scores，请返回目标成绩 target 的出现次数。

func countTarget(nums []int, target int) int {
	// 元素第一次出现的位置
	leftIndex := leftBound(nums, target)
	if leftIndex == -1 {
		return 0
	}
	count := 0
	for i := leftIndex; i < len(nums); i++ {
		if nums[i] == target {
			count++
		}
	}
	return count
}

func leftBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	// 搜索区间为 [left, right]
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			// 搜索区间变为 [mid+1, right]
			left = mid + 1
		} else if nums[mid] > target {
			// 搜索区间变为 [left, mid-1]
			right = mid - 1
		} else if nums[mid] == target {
			// 收缩右侧边界
			right = mid - 1
		}
	}
	// 检查出界情况
	if left >= len(nums) || nums[left] != target {
		// 该元素在数组中不存在的情况
		return -1
	}
	return left
}
