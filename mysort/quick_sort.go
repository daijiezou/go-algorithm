package mysort

func QuickSort(arr []int, left, right int) {
	if left >= right {
		return
	}
	part := Part(arr, left, right)
	QuickSort(arr, left, part-1)  // 这里是part-1
	QuickSort(arr, part+1, right) // 这里是part+1
}

func Part(arr []int, left, right int) int {
	key := arr[right]
	index := left
	for left < right {
		if arr[left] < key {
			arr[index], arr[left] = arr[left], arr[index]
			index++
		}
		left++
	}
	arr[index], arr[right] = arr[right], arr[index]
	return index
}
