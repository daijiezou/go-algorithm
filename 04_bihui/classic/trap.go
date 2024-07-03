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

// https://leetcode.cn/problems/trapping-rain-water/

/*
water[i] = min(

	# 左边最高的柱子
	max(height[0..i]),
	# 右边最高的柱子
	max(height[i..end])

) - height[i]
*/
func trap(height []int) int {
	length := len(height)
	leftMax := make([]int, length)
	rightMax := make([]int, length)
	leftMax[0] = height[0]
	rightMax[length-1] = height[length-1]
	for i := 1; i < length; i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}
	for i := length - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}
	res := 0
	for i := 1; i < length-1; i++ {
		minHeight := min(leftMax[i], rightMax[i])
		res += minHeight - height[i]
	}
	return res
}

func trap2(height []int) int {
	left, right := 0, len(height)-1
	l_max, r_max := 0, 0

	res := 0
	for left < right {
		l_max = max(l_max, height[left])
		r_max = max(r_max, height[right])

		if l_max < r_max {
			res += l_max - height[left]

			left++
		} else {
			res += r_max - height[right]
			right--
		}
	}
	return res
}
