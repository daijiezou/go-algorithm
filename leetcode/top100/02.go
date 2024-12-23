package top100

import (
	"fmt"
	"math"
	"slices"
	"sort"
)

// https://leetcode.cn/problems/move-zeroes/?envType=study-plan-v2&envId=top-100-liked
func moveZeroes(nums []int) {
	var slow, fast int
	for ; fast < len(nums); fast++ {
		if nums[fast] != 0 {
			nums[slow] = nums[fast]
			slow++
		}
	}
	for i := slow; i < len(nums); i++ {
		nums[i] = 0
	}

}

// https://leetcode.cn/problems/container-with-most-water/?envType=study-plan-v2&envId=top-100-liked
func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	res := 0
	for left < right {
		h := min(height[left], height[right])
		area := (right - left) * h
		res = max(area, res)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return res
}

func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	slices.Sort(nums)
	for i := 0; i < len(nums); i++ {
		target := 0 - nums[i]
		twos := twoSum2(nums, i+1, target)
		for _, v := range twos {
			res = append(res, append([]int{nums[i]}, v...))
		}
		for i < len(nums)-1 && nums[i] == nums[i+1] {
			i++
		}
	}
	return res
}

func twoSum2(nums []int, start int, target int) [][]int {
	res := make([][]int, 0)
	left := start
	right := len(nums) - 1
	for left < right {
		leftVal := nums[left]
		rightVal := nums[right]
		sum := nums[left] + nums[right]
		if sum == target {
			res = append(res, []int{leftVal, rightVal})
			for left < right && nums[left] == leftVal {
				left++
			}
			for left < right && nums[right] == rightVal {
				right--
			}
		} else if sum < target {
			left++

		} else {
			right--

		}
	}
	return res
}

// https://leetcode.cn/problems/trapping-rain-water/?envType=study-plan-v2&envId=top-100-liked
// 接雨水
func trap(height []int) int {
	n := len(height)
	leftMax := make([]int, n)
	rightMax := make([]int, n)
	lmax := height[0]
	for i := 0; i < n; i++ {
		lmax = max(lmax, height[i])
		leftMax[i] = lmax
	}
	rMax := height[n-1]
	for i := n - 1; i >= 0; i-- {
		rMax = max(rMax, height[i])
		rightMax[i] = rMax
	}
	res := 0
	for i := 0; i < len(height); i++ {
		h := min(leftMax[i], rightMax[n-i-1])
		res += (h - height[i])
	}
	return res
}

func lengthOfLongestSubstring(s string) int {
	window := make(map[byte]int, 0)
	res := 0
	left := 0
	for right := 0; right < len(s); right++ {
		cur := s[right]
		window[cur]++
		for window[cur] > 1 {
			leftVal := s[left]
			window[leftVal]--
			left++
		}
		res = max(res, right-left+1)
	}
	return res
}

// https://leetcode.cn/problems/find-all-anagrams-in-a-string/?envType=study-plan-v2&envId=top-100-liked
func findAnagrams(s string, p string) []int {
	cnts := [26]int{}
	for _, v := range p {
		cnts[v-'a']++
	}
	win := [26]int{}
	n := len(p)
	res := make([]int, 0)
	left := 0
	for i := 0; i < len(s); i++ {
		cur := s[i]
		win[cur-'a']++
		if i-left+1 == n {
			if win == cnts {
				res = append(res, i+1-n)
			}
			win[s[left]-'a']--
			left++
		}
	}
	return res
}

// 和为k的子数组
func subarraySum(nums []int, k int) int {
	n := len(nums)
	preSum := make([]int, n+1)
	// 前缀和到该前缀和出现次数的映射，方便快速查找所需的前缀和
	count := make(map[int]int)
	count[0] = 1
	// 记录和为 k 的子数组个数
	res := 0
	for i := 1; i <= n; i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
		needPreSum := preSum[i] - k
		res += count[needPreSum]
		count[preSum[i]]++
	}
	return res
}

// 滑动窗口的最大值
func maxSlidingWindow(nums []int, k int) []int {
	left, right := 0, 0
	n := len(nums)
	stack := []int{}
	res := []int{}
	for ; right < n; right++ {
		cur := nums[right]
		for len(stack) > 0 && cur > stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, cur)
		if right-left+1 == k {
			res = append(res, stack[0])
			leave := nums[left]
			if leave == stack[0] {
				stack = stack[1:]
			}
			left++
		}
	}
	return res
}

// 最小覆盖子串
func minWindow(s string, t string) string {
	targetMap := make(map[uint8]int)
	for i := 0; i < len(t); i++ {
		targetMap[t[i]]++
	}
	targetCnt := len(targetMap)
	win := make(map[uint8]int)
	left, right := 0, 0
	valid := 0
	res := ""
	length := math.MaxInt
	for ; right < len(s); right++ {
		cur := s[right]
		win[cur]++
		if win[cur] == targetMap[cur] {
			valid++
		}
		for valid == targetCnt {
			if right-left+1 < length {
				res = s[left : right+1]
				length = len(res)
			}
			leave := s[left]
			if win[leave] == targetMap[leave] {
				valid--
			}
			win[leave]--
			left++
		}
	}
	if length == math.MaxInt32 {
		return ""
	}
	return res
}

// 最大子数组和
func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	n := len(nums)
	preSum := make([]int, n+1)
	minPreSum := 0
	res := math.MinInt32
	for i := 1; i <= n; i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
		res = max(res, preSum[i]-minPreSum)
		minPreSum = min(preSum[i], minPreSum)

	}
	return res
}

// https://leetcode.cn/problems/merge-intervals/?envType=study-plan-v2&envId=top-100-liked
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(a, b int) bool {
		if intervals[a][0] == intervals[b][0] {
			return intervals[a][1] < intervals[b][1]
		}
		return intervals[a][0] < intervals[b][0]
	})
	res := [][]int{}
	res = append(res, intervals[0])
	preRight := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= preRight {
			if intervals[i][1] > preRight {
				preRight = intervals[i][1]
				res[len(res)-1][1] = preRight
			}

		} else {
			res = append(res, intervals[i])
			preRight = intervals[i][1]
		}

	}
	return res
}

func rotate(nums []int, k int) {
	n := len(nums)
	k = k % n
	newNums := make([]int, n)
	copy(newNums, nums)
	for i := 0; i < n; i++ {
		nums[(i+k)%n] = newNums[i]
	}
	fmt.Println(nums)
}

func rotate2(nums []int, k int) {
	n := len(nums)
	k = k % n
	slices.Reverse(nums)
	slices.Reverse(nums[0:k])
	slices.Reverse(nums[k:])
}

// https://leetcode.cn/problems/product-of-array-except-self/description/?envType=study-plan-v2&envId=top-100-liked
// 思路：使用前缀积和后缀积
func productExceptSelf(nums []int) []int {
	n := len(nums)
	pre := make([]int, n)
	pre[0] = nums[0]
	for i := 1; i < n; i++ {
		pre[i] = pre[i-1] * nums[i]
	}
	suf := make([]int, n)
	suf[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		suf[i] = nums[i] * suf[i+1]
	}
	res := make([]int, n)
	res[0] = suf[1]
	res[n-1] = pre[n-2]
	for i := 1; i < n-1; i++ {
		res[i] = suf[i+1] * pre[i-1]
	}
	return res
}

// https://leetcode.cn/problems/first-missing-positive/?envType=study-plan-v2&envId=top-100-liked
/*
在恢复后，数组应当有 [1, 2, ..., N] 的形式，但其中有若干个位置上的数是错误的，
每一个错误的位置就代表了一个缺失的正数。以题目中的示例二 [3, 4, -1, 1] 为例，
恢复后的数组应当为 [1, -1, 3, 4]，我们就可以知道缺失的数为 2。
*/
func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		// 这里是for
		for nums[i] > 0 && nums[i] <= n && nums[nums[i]-1] != nums[i] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return n + 1
}
