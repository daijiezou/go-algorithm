package _1_huadongchuangkou

import "sort"

/*
相向双指针
*/

func reverseString(s []byte) {
	left := 0
	right := len(s) - 1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

// https://leetcode.cn/problems/minimum-length-of-string-after-deleting-similar-ends/
/*
给你一个只包含字符 'a'，'b' 和 'c' 的字符串 s ，你可以执行下面这个操作（5 个步骤）任意次：
选择字符串 s 一个 非空 的前缀，这个前缀的所有字符都相同。
选择字符串 s 一个 非空 的后缀，这个后缀的所有字符都相同。
前缀和后缀在字符串中任意位置都不能有交集。
前缀和后缀包含的所有字符都要相同。
同时删除前缀和后缀。
请你返回对字符串 s 执行上面操作任意次以后（可能 0 次），能得到的 最短长度 。
*/
func minimumLength(s string) int {
	n := len(s)
	left := 0
	right := n - 1
	for left < right && s[left] == s[right] {
		x := s[left]
		for left+1 < right && s[left+1] == x {
			left++
		}
		for right-1 > left && s[right-1] == x {
			right--
		}
		left++
		right--
	}

	// 如果能全部删除完毕,right=left+1
	return right - left + 1
}

/*
https://leetcode.cn/problems/watering-plants-ii/
*/
func minimumRefill(plants []int, capacityA int, capacityB int) int {
	res := 0
	left := 0
	right := len(plants) - 1
	ca := capacityA
	cb := capacityB
	for left < right {
		if ca >= plants[left] {
			ca -= plants[left]
		} else {
			res++
			ca = capacityA - plants[left]
		}

		if cb >= plants[right] {
			cb -= plants[right]
		} else {
			res++
			cb = capacityB - plants[right]
		}
		left++
		right--
	}

	if left == right {
		mx := max(ca, cb)
		if mx < plants[left] {
			res++

		}
	}
	return res
}

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

// https://leetcode.cn/problems/the-k-strongest-values-in-an-array/
func getStrongest(arr []int, k int) []int {
	n := len(arr)
	sort.Ints(arr)
	mid := arr[(n-1)/2]
	left := 0
	right := n - 1
	res := make([]int, 0)
	for left <= right && len(res) < k {
		if myAbs2(arr[left], mid) > myAbs2(arr[right], mid) {
			res = append(res, arr[left])
			left++
		} else {
			res = append(res, arr[right])
			right--
		}
	}
	return res
}
func myAbs2(a, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}

const mod = 1e9 + 7

// https://leetcode.cn/problems/4xy4Wx/description/
func purchasePlans(nums []int, target int) int {
	sort.Ints(nums)
	left := 0

	right := len(nums) - 1
	res := 0

	// 枚举左维护右
	for ; left < right; left++ {
		for right > left && nums[left]+nums[right] > target {
			right--
		}
		/*
			该范围的边界值满足要求，则该范围的所有的值都满足要求
		*/
		res += right - left
	}
	return res % mod
}
