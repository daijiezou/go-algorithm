package _1_huadongchuangkou

import "math"

func minSubArrayLen(target int, nums []int) int {
	sum := 0
	left := 0
	res := math.MaxInt
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		for sum >= target {
			res = min(res, i-left+1)
			sum -= nums[left]
			left++
		}
	}
	if res == math.MaxInt {
		return 0
	}
	return res
}

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
		for ; right < n && arr[right] < arr[left]; right++ {

		}
		ans = min(right-left-1, ans) // 删除arr[left+1:right]
	}
	return ans
}

// https://leetcode.cn/problems/smallest-range-covering-elements-from-k-lists/
func smallestRange(nums [][]int) []int {
	size := len(nums)
	// 纪录一个数字在哪些组里有
	indices := make(map[int][]int)
	xMin, XMax := math.MaxInt32, math.MinInt32
	for i := 0; i < size; i++ {
		for _, v := range nums[i] {
			indices[v] = append(indices[v], i)
			xMin = min(v, xMin)
			XMax = max(v, XMax)
		}
	}
	left, right := xMin, xMin-1
	bestLeft, bestRight := xMin, XMax
	freq := make(map[int]int)
	for right < XMax {
		right++
		// 该数字至少在一个组内，计算该数字在几个组内
		if len(indices[right]) > 0 {
			for _, i := range indices[right] {
				freq[i]++
			}
			for len(freq) == size {
				// 更新答案
				if right-left < bestRight-bestLeft {
					bestLeft, bestRight = left, right
				}
				for _, i := range indices[left] {
					freq[i]--
					if freq[i] == 0 {
						delete(freq, i)
					}
					left++
				}
			}
		}
	}
	return []int{bestLeft, bestRight}
}
