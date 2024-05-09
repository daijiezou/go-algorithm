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
