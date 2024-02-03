package search

func Bsearch(array []int, target int) int {
	var low int
	var high = len(array) - 1
	for low <= high {
		mid := (low + high) / 2
		if array[mid] == target {
			return mid
		} else if array[mid] < target {
			low = mid + 1
		} else if array[mid] > target {
			high = mid - 1
		}
	}
	return -1
}

// Bsearch2 二分查找的递归实现
func Bsearch2(array []int, low, high, target int) int {
	if low > high {
		return -1
	}
	mid := low + ((high - low) >> 1)
	if array[mid] == target {
		return mid
	} else if array[mid] < target {
		return Bsearch2(array, mid+1, high, target)
	} else {
		return Bsearch2(array, low, mid-1, target)
	}
}

// Bsearch3 查找第一个值等于给定值的元素
func Bsearch3(array []int, target int) int {
	var low int
	var high = len(array) - 1
	for low <= high {
		mid := (low + high) / 2
		if array[mid] == target {
			if mid == 0 || array[mid-1] != target {
				return mid
			} else {
				high = mid - 1
			}
		} else if array[mid] < target {
			low = mid + 1
		} else if array[mid] > target {
			high = mid - 1
		}
	}
	return -1
}

// Bsearch4 查找第一个值大于等于给定值的元素
func Bsearch4(array []int, target int) int {
	var low int
	var high = len(array) - 1
	for low <= high {
		mid := (low + high) / 2
		if array[mid] >= target {
			if mid == 0 || array[mid-1] < target {
				return mid
			} else {
				high = mid - 1
			}
		} else {
			low = mid + 1
		}
	}
	return -1
}
