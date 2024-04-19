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
	for len(left) > 0 {
		res = append(res, left[0])
		left = left[1:]
	}
	for len(right) > 0 {
		res = append(res, right[0])
		right = right[1:]
	}
	return res
}
