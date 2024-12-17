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
	for i := left; i < right; i++ {
		if arr[i] < key {
			arr[index], arr[i] = arr[i], arr[index]
			index++
		}

	}
	arr[index], arr[right] = arr[right], arr[index]
	return index
}
