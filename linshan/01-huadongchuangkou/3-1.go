package _1_huadongchuangkou

// https://leetcode.cn/problems/container-with-most-water/
func maxArea(height []int) int {
	left, right := 0, len(height)-1

	area := 0
	for left <= right {
		if height[left] < height[right] {
			area = max(area, height[left]*(right-left))
			left++
		} else {
			area = max(area, height[left]*(right-left))
			right--
		}
	}
	return area
}

func trap(height []int) int {
	n := len(height)
	leftMemo := make([]int, n)
	rightMemo := make([]int, n)
	leftMemo[0] = height[0]
	rightMemo[n-1] = height[n-1]
	for i := 1; i < len(height); i++ {
		leftMemo[i] = max(leftMemo[i-1], height[i])
	}
	for i := len(height) - 2; i >= 0; i-- {
		rightMemo[i] = max(rightMemo[i+1], height[i])
	}
	res := 0
	for i := 0; i < n-1; i++ {
		res += min(leftMemo[i], rightMemo[i]) - height[i]
	}
	return res
}
