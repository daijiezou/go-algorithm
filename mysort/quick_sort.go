package mysort

func QuickSort(arr []int, left, right int) {
	if left >= right {
		return
	}
	p := partition(arr, left, right)
	QuickSort(arr, left, p-1)
	QuickSort(arr, p+1, right)
}

func partition(nums []int, left, right int) int {
	key := nums[right]
	i := left
	for j := left; j < right; j++ {
		if nums[j] < key {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	nums[i], nums[right] = nums[right], nums[i]
	return i
}

func QuickSort2(arr []int, left, right int) {
	if left >= right {
		return
	}
	p := partition2(arr, left, right)
	QuickSort2(arr, left, p-1)
	QuickSort2(arr, p+1, right)

}

func partition2(arr []int, left, right int) int {
	pivot := arr[right]
	i := left
	for j := left; j < right; j++ {
		if arr[j] > pivot {
			arr[j], arr[i] = arr[i], arr[j]
			i++
		}
	}
	arr[i], arr[right] = arr[right], arr[i]
	return i
}
