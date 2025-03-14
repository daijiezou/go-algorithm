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
	// 选择最左边的元素作为基准值
	pivot := arr[right]
	// 初始化指向小于基准值区域的末尾
	i := left

	// 遍历数组，将小于基准值的元素移到左侧
	for j := left; j < right; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	// 将基准值放到正确的位置
	arr[right], arr[i] = arr[i], arr[right]
	return i
}
