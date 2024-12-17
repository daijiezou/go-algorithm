package mysort

func mergeSort(sumList []int) []int {
	// 找到数组中点
	length := len(sumList)
	if length < 2 {
		return sumList
	}
	middle := length / 2
	left := sumList[0:middle]
	right := sumList[middle:]
	return merge(mergeSort(left), mergeSort(right))
}

func merge(left, right []int) []int {
	var res []int
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			res = append(res, left[0])
			left = left[1:]
		} else {
			res = append(res, right[0])
			right = right[1:]
		}
	}
	if len(left) > 0 {
		res = append(res, left...)
	}
	if len(right) > 0 {
		res = append(res, right...)
	}
	return res
}

func mergeSort3(sumList []int) []int {
	sort(sumList, 0, len(sumList)-1)
	return sumList
}

func sort(nums []int, lo, hi int) {
	if lo >= hi {
		return
	}
	mid := (lo + hi) / 2
	// 利用定义，排序 nums[lo..mid]
	sort(nums, lo, mid)
	// 利用定义，排序 nums[mid+1..hi]
	sort(nums, mid+1, hi)

	// 此时两部分子数组已经被排好序
	// 合并两个有序数组，使 nums[lo..hi] 有序
	merge2(nums, lo, hi, mid)
}

func merge2(nums []int, lo int, hi int, mid int) {
	temp := make([]int, hi-lo+1)
	for i := lo; i <= hi; i++ {
		temp[i-lo] = nums[i]
	}
	leftIndex := lo
	rightIndex := mid + 1
	for index := lo; index <= hi; index++ {
		if leftIndex == mid+1 { // left数组中已经没有元素
			nums[index] = temp[rightIndex-lo]
			rightIndex++
		} else if rightIndex == hi+1 { // right数组中已经没有元素
			nums[index] = temp[leftIndex-lo]
			leftIndex++
		} else if temp[leftIndex-lo] < temp[rightIndex-lo] {
			nums[index] = temp[leftIndex-lo]
			leftIndex++
		} else {
			nums[index] = temp[rightIndex-lo]
			rightIndex++
		}
	}
}
