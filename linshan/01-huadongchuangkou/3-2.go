package _1_huadongchuangkou

// https://leetcode.cn/problems/shortest-subarray-to-be-removed-to-make-array-sorted/
// 1574. 删除最短的子数组使剩余数组有序
func findLengthOfShortestSubarray(arr []int) int {
	n := len(arr)
	right := n - 1
	for right > 0 && arr[right-1] <= arr[right] {
		right--
	}
	// 已经是非递增数组
	if right == 0 {
		return 0
	}
	ans := right // 删除arr[0:right]
	// 枚举左端点，移动右端点
	for left := 0; left == 0 || arr[left-1] <= arr[left]; left++ {

		// 直到arr[right] > arr[left]
		for right < n && arr[right] < arr[left] {
			right++
		}
		ans = min(right-left-1, ans) // 删除arr[left+1:right]
	}
	return ans
}

func findLengthOfShortestSubarray_2(arr []int) int {
	n := len(arr)
	right := n - 1
	for right > 0 && arr[right-1] <= arr[right] {
		right--
	}
	// 已经是非递增数组
	if right == 0 {
		return 0
	}
	ans := right // 删除arr[0:right]
	left := 0
	// 枚举右维护左
	for ; ; right++ {
		// right==n，
		for ; right == n || arr[left] <= arr[right]; left++ {
			ans = min(ans, right-left-1)
			if arr[left+1] < arr[left] {
				return ans
			}
		}

	}
	return ans
}
