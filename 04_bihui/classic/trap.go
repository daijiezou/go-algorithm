package classic

// https://leetcode.cn/problems/container-with-most-water/description/
// 给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。
//
// 找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	res := 0
	for left <= right {
		// [left, right] 之间的矩形面积
		res = max(res, height[left]*(right-left))

		// 双指针技巧，移动较低的一边
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return res
}
