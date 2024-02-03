package mysort

func MergeSort(sumList []int) []int {
	length := len(sumList)
	if length < 2 {
		return sumList
	}
	middle := length / 2
	left := sumList[0:middle]
	right := sumList[middle:]
	return Merge(MergeSort(left), MergeSort(right))
}

func Merge(left, right []int) []int {
	var result []int
	for len(left) != 0 && len(right) != 0 {
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}
	for len(left) != 0 {
		result = append(result, left[0])
		left = left[1:]
	}

	for len(right) != 0 {
		result = append(result, right[0])
		right = right[1:]
	}
	return result
}
