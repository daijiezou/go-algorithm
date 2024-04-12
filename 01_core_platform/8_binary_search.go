package _1_core_platform

// 最普通的二分查找
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}
	return -1
}

func searchRange(nums []int, target int) []int {
	return []int{leftBound(nums, target), rightBound(nums, target)}
}

func rightBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target { // 注意
			left = mid + 1 // 注意
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}
	if right < 0 || right >= len(nums) {
		return -1
	}
	if nums[right] == target {
		return right
	}
	return -1
}

func leftBound(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target { // 注意
			right = mid - 1 // 注意
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}
	if left < 0 || left >= len(nums) {
		return -1
	}
	if nums[left] == target {
		return left
	}
	return -1
}
