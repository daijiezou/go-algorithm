package meeting150

func searchInsert(nums []int, target int) int {
	//res, _ := slices.BinarySearch(nums, target)
	return leftB(nums, target)
}

func leftB(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] == target {
			right = mid - 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}
	return left
}

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])
	left := 0
	right := m*n - 1
	for left <= right {
		mid := left + (right-left)/2
		row := mid / n
		col := mid % n
		if matrix[row][col] > target {
			right = mid - 1
		} else if matrix[row][col] < target {
			left = mid + 1
		} else if matrix[row][col] == target {
			return true
		}
	}
	return false
}

func findPeakElement(nums []int) int {
	return -1
}

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] >= nums[left] {
			//left 到 mid有序
			if target < nums[mid] && target >= nums[left] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

func findMin(nums []int) int {
	n := len(nums)
	left, right := 0, n-1
	for left <= right {
		mid := left + (right-left)/2
		x := nums[mid]
		if x < nums[n-1] { // 最小值在mid的左边，缩减右边界
			right = mid - 1
		} else if x > nums[n-1] { // 最小值在mid的右边，缩减左边界
			left = mid + 1
		} else if x == nums[n-1] {
			right = mid - 1
		}
	}
	return nums[left]
}

func searchRange(nums []int, target int) []int {
	n := len(nums)
	first := leftBound(nums, target)
	if first == n || nums[first] != target {
		return []int{-1, -1}
	}
	last := leftBound(nums, target+1) - 1
	return []int{first, last}
}
