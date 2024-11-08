package _3_monotone_stack

import (
	"fmt"
	"math"
	"sort"
)

/*
给定一个整数数组 temperatures ，表示每天的温度，返回一个数组 answer ，
其中 answer[i] 是指对于第 i 天，下一个更高温度出现在几天后。
如果气温在这之后都不会升高，请在该位置用 0 来代替。
*/
// 从左到右遍历
func dailyTemperatures(temperatures []int) []int {
	length := len(temperatures)
	res := make([]int, length)
	s := make([]int, 0)
	for i := 0; i < length; i++ {
		for len(s) > 0 && temperatures[i] > temperatures[s[len(s)-1]] {
			j := s[len(s)-1]
			s = s[:len(s)-1]
			res[j] = i - j
		}
		s = append(s, i)
	}
	return res
}

// 从右向左
func dailyTemperatures2(temperatures []int) []int {
	length := len(temperatures)
	res := make([]int, length)
	s := make([]int, 0)
	for i := length - 1; i >= 0; i-- {
		for len(s) > 0 && temperatures[s[len(s)-1]] <= temperatures[i] {
			s = s[:len(s)-1]
		}
		if len(s) == 0 {
			res[i] = 0
		} else {
			res[i] = s[len(s)-1] - i
		}
		s = append(s, i)
	}
	return res
}

func finalPrices(prices []int) []int {
	res := make([]int, len(prices))
	stack := make([]int, 0)
	n := len(prices)
	for i := 0; i < n; i++ {
		for len(stack) > 0 && prices[i] <= prices[stack[len(stack)-1]] {
			j := stack[len(stack)-1]
			res[j] = prices[j] - prices[i]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	for i := 0; i < len(stack); i++ {
		res[stack[i]] = prices[stack[i]]
	}
	return res
}

func finalPrices2(prices []int) []int {
	res := make([]int, len(prices))
	stack := make([]int, 0)
	n := len(prices)
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && prices[i] < prices[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			res[i] = prices[i]
		} else {
			res[i] = prices[i] - prices[stack[len(stack)-1]]
		}
		stack = append(stack, i)
	}
	return res
}

/*
nums1 中数字 x 的 下一个更大元素 是指 x 在 nums2 中对应位置 右侧 的 第一个 比 x 大的元素。
给你两个 没有重复元素 的数组 nums1 和 nums2 ，下标从 0 开始计数，其中nums1 是 nums2 的子集。
*/
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	n := len(nums2)
	nextGreaterElementMap := make(map[int]int, len(nums2))
	s := make([]int, 0)
	for i := 0; i < n; i++ {
		for len(s) > 0 && nums2[i] > s[len(s)-1] {
			j := s[len(s)-1]
			s = s[:len(s)-1]
			nextGreaterElementMap[j] = nums2[i]
		}
		s = append(s, nums2[i])
	}
	for i := 0; i < len(s); i++ {
		nextGreaterElementMap[s[i]] = -1
	}
	res := make([]int, len(nums1))
	for i := 0; i < len(nums1); i++ {
		res[i] = nextGreaterElementMap[nums1[i]]
	}
	return res
}

func nextGreaterElements2(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	s := make([]int, 0)
	for i := 0; i < 2*n; i++ {
		x := i % n
		for len(s) > 0 && nums[x] > nums[s[len(s)-1]] {
			j := s[len(s)-1]
			s = s[:len(s)-1]
			res[j] = nums[x]
		}
		if i < n {
			s = append(s, i)
		}

	}
	for i := 0; i < len(s); i++ {
		res[s[i]] = -1
	}
	return res
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 1019. 链表中的下一个更大节点
func nextLargerNodes(head *ListNode) []int {
	nums := make([]int, 0)
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}
	return NextGreaterElement(nums)
}

func NextGreaterElement(nums []int) []int {
	length := len(nums)
	res := make([]int, length)
	s := make([]int, 0)
	for i := 0; i < length; i++ {
		for len(s) > 0 && nums[i] > nums[s[len(s)-1]] {
			j := s[len(s)-1]
			s = s[:len(s)-1]
			res[j] = nums[i]
		}
		s = append(s, i)
	}
	for i := 0; i < len(s); i++ {
		res[s[i]] = 0
	}
	return res
}

// https://leetcode.cn/problems/car-fleet/
/*
在一条单行道上，有 n 辆车开往同一目的地。目的地是几英里以外的 target 。
给定两个整数数组 position 和 speed ，长度都是 n ，其中 position[i] 是第 i 辆车的位置， speed[i] 是第 i 辆车的速度(单位是英里/小时)。
一辆车永远不会超过前面的另一辆车，但它可以追上去，并以较慢车的速度在另一辆车旁边行驶。
车队 是指并排行的一辆或几辆汽车。车队的速度是车队中 最慢 的车的速度。
即便一辆车在 target 才赶上了一个车队，它们仍然会被视作是同一个车队。
返回到达目的地的车队数量 。
*/
func carFleet(target int, position []int, speed []int) int {
	n := len(position)
	cars := make([][2]int, n)
	for i := 0; i < n; i++ {
		cars[i][0] = position[i]
		cars[i][1] = speed[i]
	}
	// 按照初始位置，从小到大排序
	sort.Slice(cars, func(i, j int) bool {
		return cars[i][0] < cars[j][0]
	})
	// 计算每辆车到达终点的时间
	time := make([]float64, n)
	for i := 0; i < n; i++ {
		car := cars[i]
		time[i] = float64(target-car[0]) / float64(car[1])
	}
	// 如果位置再后面，需要的时间却比前面的车长，则肯定会形成一个车队
	s := make([]int, 0)
	for i := 0; i < n; i++ {
		for len(s) > 0 && time[i] >= time[s[len(s)-1]] {
			// 说明栈顶的车子被合并了
			s = s[:len(s)-1]
		}
		s = append(s, i)
	}
	return len(s)
}

/*
https://leetcode.cn/problems/online-stock-span/description/
设计一个算法收集某些股票的每日报价，并返回该股票当日价格的 跨度 。
当日股票价格的 跨度 被定义为股票价格小于或等于今天价格的最大连续日数（从今天开始往回数，包括今天）。
例如，如果未来 7 天股票的价格是 [100,80,60,70,60,75,85]，那么股票跨度将是 [1,1,1,2,1,4,6] 。
*/
type StockSpanner struct {
	s        []pair
	curDay   int
	daysSpan map[int]int
}

type pair struct {
	day   int
	price int
}

func Constructor() StockSpanner {
	return StockSpanner{
		s:        []pair{{-1, math.MaxInt}},
		curDay:   -1,
		daysSpan: map[int]int{},
	}
}

func (this *StockSpanner) Next(price int) int {
	this.curDay++
	for len(this.s) != 0 && price >= this.s[len(this.s)-1].price {
		this.s = this.s[:len(this.s)-1]
	}
	this.s = append(this.s, pair{this.curDay, price})
	return this.curDay - this.s[len(this.s)-2].day
}

// https://leetcode.cn/problems/132-pattern/

func find132pattern(nums []int) bool {
	n := len(nums)
	leftMin := make([]int, n)
	leftMin[0] = math.MaxInt
	for i := 1; i < n; i++ {
		leftMin[i] = min(leftMin[i-1], nums[i-1])
	}
	fmt.Println(leftMin)
	s := make([]int, 0)
	for i := 0; i < n; i++ {
		rightMax := -1
		for len(s) > 0 && nums[i] > nums[s[len(s)-1]] {
			rightMax = nums[s[len(s)-1]]
			s = s[:len(s)-1]
		}
		if leftMin[i] < rightMax {
			return true
		}
		s = append(s, i)
	}
	return false
}

// https://leetcode.cn/problems/maximum-width-ramp/description/
/*
给定一个整数数组 A，坡是元组 (i, j)，其中  i < j 且 A[i] <= A[j]。这样的坡的宽度为 j - i。

找出 A 中的坡的最大宽度，如果不存在，返回 0 。
*/
// 暴力做法超时了
func maxWidthRamp2(nums []int) int {
	n := len(nums)
	res := 0
	//s := make([]int, 0)
	for i := 0; i < n; i++ {
		if n-i < res {
			break
		}
		for j := n - 1; j >= 0 && j-i > res; j-- {
			if nums[i] <= nums[j] {
				res = max(res, j-i)
				break
			}
		}
	}
	return res
}

// 单调栈的第二种模板
func maxWidthRamp(nums []int) int {
	stack := []int{0}
	ans := 0
	n := len(nums)
	// First pass: populate the stack
	// s中的元素是递减的
	for i := 1; i < n; i++ {
		if nums[stack[len(stack)-1]] > nums[i] {
			stack = append(stack, i)
		}
	}

	// Second pass: calculate the maximum width ramp
	for j := n - 1; j >= 0; j-- {
		for len(stack) > 0 && nums[j] >= nums[stack[len(stack)-1]] {
			ans = max(ans, j-stack[len(stack)-1])
			stack = stack[:len(stack)-1] // pop the last element
		}
	}
	return ans
}

// https://leetcode.cn/problems/longest-well-performing-interval/
/*
给你一份工作时间表 hours，上面记录着某一位员工每天的工作小时数。
我们认为当员工一天中的工作小时数大于 8 小时的时候，那么这一天就是「劳累的一天」。
所谓「表现良好的时间段」，意味在这段时间内，「劳累的天数」是严格 大于「不劳累的天数」。
请你返回「表现良好时间段」的最大长度。
*/
func longestWPI(hours []int) int {
	n := len(hours)
	s := []int{0}

	preSum := make([]int, n+1)
	for i := 0; i < n; i++ {
		preSum[i+1] = preSum[i]
		if hours[i] > 8 {
			preSum[i+1]++
		} else {
			preSum[i+1]--
		}
		if preSum[i+1] < preSum[s[len(s)-1]] {
			s = append(s, i+1)
		}
	}
	ans := 0
	for i := n; i > 0; i-- {
		for len(s) > 0 && preSum[i] > preSum[s[len(s)-1]] {
			ans = max(ans, i-s[len(s)-1]) // [栈顶,i) 可能是最长子数组
			s = s[:len(s)-1]
		}
	}
	return ans
}

// https://leetcode.cn/problems/maximum-score-of-a-good-subarray/
func maximumScore(nums []int, k int) int {
	n := len(nums)
	s1 := []int{k}
	// 纪录最小值
	s1Map := map[int]int{
		k: nums[k],
	}
	s1Min := nums[k]
	for i := k - 1; i >= 0; i-- {
		if nums[i] < s1Min {
			s1Min = nums[i]
		}
		for len(s1) > 0 && nums[i] > nums[s1[len(s1)-1]] && s1Min >= s1Map[s1[len(s1)-1]] {
			s1 = s1[:len(s1)-1]
		}
		s1Map[i] = s1Min
		s1 = append(s1, i)
	}

	// 逐渐递减
	s2 := []int{k}
	s2Map := map[int]int{
		k: nums[k],
	}
	s2Min := nums[k]
	for i := k + 1; i < n; i++ {
		if nums[i] < s2Min {
			s2Min = nums[i]
		}
		for len(s2) > 0 && nums[i] > nums[s2[len(s2)-1]] && s2Min >= s2Map[s2[len(s2)-1]] {
			s2 = s2[:len(s2)-1]
		}
		s2Map[i] = s2Min
		s2 = append(s2, i)
	}
	//fmt.Println(s1, s1Map)
	//fmt.Println(s2, s2Map)
	ans := 0
	for i := 0; i < len(s1); i++ {
		for j := 0; j < len(s2); j++ {
			minNum := s1Map[s1[i]]
			if s2Map[s2[j]] < minNum {
				minNum = s2Map[s2[j]]
			}
			ans = max(ans, minNum*(s2[j]-s1[i]+1))
		}
	}
	return ans
}

func maximumScore2(nums []int, k int) int {
	n := len(nums)
	left := k - 1
	right := k + 1
	ans := 0
	for i := nums[k]; ; i-- {
		for left >= 0 && nums[left] >= i {
			left--
		}
		for right < n && nums[right] >= i {
			right++
		}
		ans = max(ans, (right-left-1)*i)
		if left == -1 && right == n {
			break
		}
	}
	return ans
}

func maximumScore3(nums []int, k int) (ans int) {
	n := len(nums)
	left := make([]int, n)
	st := []int{}
	for i, x := range nums {
		for len(st) > 0 && x <= nums[st[len(st)-1]] {
			st = st[:len(st)-1]
		}
		if len(st) > 0 {
			left[i] = st[len(st)-1]
		} else {
			left[i] = -1
		}
		st = append(st, i)
	}

	right := make([]int, n)
	st = st[:0]
	for i := n - 1; i >= 0; i-- {
		for len(st) > 0 && nums[i] <= nums[st[len(st)-1]] {
			st = st[:len(st)-1]
		}
		if len(st) > 0 {
			right[i] = st[len(st)-1]
		} else {
			right[i] = n
		}
		st = append(st, i)
	}
	// 枚举当前元素的左边的下一个更小元素和右边的下一个更小元素，则其上一个就是>=当前元素的
	for i, h := range nums {
		l, r := left[i], right[i]
		if l < k && k < r { // 相比 84 题多了个 if 判断
			ans = max(ans, h*((r-1)-(l+1)+1))
		}
	}
	return ans
}

// https://leetcode.cn/problems/largest-rectangle-in-histogram/
func largestRectangleArea(heights []int) int {
	n := len(heights)
	Right := make([]int, n)
	Left := make([]int, n)

	s := []int{}
	for i := 0; i < n; i++ {
		for len(s) != 0 && heights[i] < heights[s[len(s)-1]] {
			x := s[len(s)-1]
			s = s[:len(s)-1]
			Right[x] = i
		}
		s = append(s, i)
	}
	for _, h := range s {
		Right[h] = n
	}

	s = []int{}
	for i := n - 1; i >= 0; i-- {
		for len(s) != 0 && heights[i] < heights[s[len(s)-1]] {
			x := s[len(s)-1]
			s = s[:len(s)-1]
			Left[x] = i
		}
		if len(s) == 0 {
			Left[i] = -1
		} else {
			Left[i] = s[len(s)-1]
		}
		s = append(s, i)
	}
	for _, h := range s {
		Left[h] = -1
	}
	//fmt.Println(Left)
	//fmt.Println(Right)
	res := 0
	for i := 0; i < len(heights); i++ {
		width := Right[i] - 1 - (Left[i] + 1) + 1
		res = max(res, heights[i]*width)
	}
	return res
}

func trap(height []int) int {
	res := 0
	leftMaxList := make([]int, len(height))
	rightMaxList := make([]int, len(height))

	leftMaxList[0] = height[0]
	for i := 1; i < len(height); i++ {
		leftMaxList[i] = max(leftMaxList[i-1], height[i])
	}

	rightMaxList[len(height)-1] = height[len(height)-1]
	for i := len(height) - 2; i >= 0; i-- {
		rightMaxList[i] = max(rightMaxList[i+1], height[i])
	}
	for i := 0; i < len(height); i++ {
		leftMax := leftMaxList[i]
		rightMax := rightMaxList[i]
		if min(leftMax, rightMax) > height[i] {
			res += min(leftMax, rightMax) - height[i]
		}
	}
	return res
}

func trap2(height []int) int {
	res := 0
	s := make([]int, 0)
	for i := 0; i < len(height); i++ {
		for len(s) > 0 && height[s[len(s)-1]] < height[i] {
			bottom := s[len(s)-1]
			s = s[:len(s)-1]
			if len(s) == 0 {
				break
			}
			left := s[len(s)-1]
			width := i - left - 1
			h := min(height[left], height[i]) - height[bottom]
			res += width * h
		}
		s = append(s, i)
	}
	return res
}

// https://leetcode.cn/problems/sum-of-subarray-minimums/
// 子数组的最小值之和
func sumSubarrayMins(arr []int) int {
	n := len(arr)
	left := make([]int, n)
	right := make([]int, n)
	for i := 0; i < n; i++ {
		right[i] = n
	}
	s := make([]int, 0)
	for i := 0; i < len(arr); i++ {
		for len(s) > 0 && arr[s[len(s)-1]] >= arr[i] {
			right[s[len(s)-1]] = i
			s = s[:len(s)-1]
		}
		// 剩下的就是小于当前数字的，可以作为左边界，没有的话左边界就是-1
		if len(s) == 0 {
			left[i] = -1
		} else {
			left[i] = s[len(s)-1]
		}
		s = append(s, i)
	}
	res := 0
	for i := 0; i < n; i++ {
		res += (i - left[i]) * (right[i] - i) * arr[i]
	}
	return res
}

// https://leetcode.cn/problems/sum-of-subarray-ranges/
// 子数组的最小值和最大值的差值之和
// 实际就是求子数组的最大值之和减去子数组的最小值之和
func subArrayRanges(nums []int) int64 {
	ans := sumSubarrayMins(nums)
	for i, v := range nums { // 小技巧：所有元素取反后算的就是最大值的贡献
		nums[i] = -v
	}
	return int64(-ans - sumSubarrayMins(nums))
}

// https://leetcode.cn/problems/maximum-subarray-min-product/
/*
一个数组的 最小乘积 定义为这个数组中 最小值 乘以 数组的 和 。
比方说，数组 [3,2,5] （最小值是 2）的最小乘积为 2 * (3+2+5) = 2 * 10 = 20 。
给你一个正整数数组 nums ，请你返回 nums 任意 非空子数组 的最小乘积 的 最大值 。由于答案可能很大，请你返回答案对  109 + 7 取余 的结果。
请注意，最小乘积的最大值考虑的是取余操作 之前 的结果。题目保证最小乘积的最大值在 不取余 的情况下可以用 64 位有符号整数 保存。
子数组 定义为一个数组的 连续 部分。

*/
func maxSumMinProduct(nums []int) int {
	n := len(nums)
	Right := make([]int, n)
	for i := 0; i < n; i++ {
		Right[i] = n
	}
	Left := make([]int, n)
	for i := 0; i < n; i++ {
		Left[i] = -1
	}
	s := make([]int, 0)
	preSum := make([]int, n+1)
	for i := 0; i < len(nums); i++ {
		preSum[i+1] = preSum[i] + nums[i]
		for len(s) > 0 && nums[s[len(s)-1]] >= nums[i] { // 剩下的那个值就是比当前小的
			// 栈顶即将被弹出的这个值的右边最大值就是当前的i
			Right[s[len(s)-1]] = i
			s = s[:len(s)-1]
		}
		if len(s) != 0 {
			Left[i] = s[len(s)-1]
		}
		s = append(s, i)
	}
	res := 0
	for i := 0; i < len(nums); i++ {
		left := Left[i] + 1
		right := Right[i] - 1
		sum := preSum[right+1] - preSum[left]
		res = max(res, sum*nums[i])
	}
	return res % (1e9 + 7)
}

func removeKdigits(num string, k int) string {
	if len(num) == k {
		return "0"
	}
	s := []byte{}
	n := len(num)
	cnt := 0
	for i := 0; i < n; i++ {
		for len(s) > 0 && s[len(s)-1] > num[i] && cnt < k {
			cnt++
			s = s[:len(s)-1]
		}
		s = append(s, num[i])
	}
	for i := cnt; i < k && len(s) > 1; i++ {
		s = s[:len(s)-1]
	}
	for len(s) > 1 && s[0] == '0' {
		s = s[1:]
	}
	return string(s)
}
