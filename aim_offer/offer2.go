package aim_offer

import (
	"math"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func PrintMinNumber(numbers []int) string {
	if len(numbers) == 0 {
		return ""
	}

	// 将数字转换为字符串
	strs := make([]string, len(numbers))
	for i, num := range numbers {
		strs[i] = strconv.Itoa(num)
	}

	// 自定义排序
	sort.Slice(strs, func(i, j int) bool {
		s1 := strs[i] + strs[j]
		s2 := strs[j] + strs[i]
		return s1 < s2
	})

	return strings.Join(strs, "")
}

func solve(nums string) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	dp := make([]int, n+1)
	// 空字符串有一种解码方式
	dp[n] = 1

	// 处理最后一个字符
	if nums[n-1] != '0' {
		dp[n-1] = 1
	}

	for i := n - 2; i >= 0; i-- {
		// 处理单个数字
		if nums[i] != '0' {
			dp[i] = dp[i+1]
		}

		// 处理两个数字
		if nums[i] == '1' || (nums[i] == '2' && nums[i+1] <= '6') {
			dp[i] += dp[i+2]
		}
	}

	return dp[0]
}

func solve2(nums string) int {
	n := len(nums)
	if n == 0 {
		return 0 // 空字符串没有解码方式
	}

	// dp[i] 表示从第i个字符到末尾的子串的解码方式数
	dp := make([]int, n)

	// 初始化最后一个字符
	// 如果最后一个字符不是'0'，则有一种解码方式（A-Z对应1-26）
	if nums[n-1] != '0' {
		dp[n-1] = 1
	}

	// 从倒数第二个字符开始向前遍历
	for i := n - 2; i >= 0; i-- {
		// 情况1：当前字符单独解码
		// 如果当前字符是'0'，不能单独解码，保持dp[i] = 0
		if nums[i] != '0' {
			dp[i] = dp[i+1] // 当前字符单独解码的方式数等于后面子串的方式数
		}

		// 情况2：当前字符与下一个字符组合解码
		// 计算两位数的值
		num := int(nums[i]-'0')*10 + int(nums[i+1]-'0')
		// 检查是否在有效范围内（10-26）
		if num >= 10 && num <= 26 {
			if i+2 < n {
				// 如果组合有效，则加上dp[i+2]的方式数
				dp[i] += dp[i+2]
			} else {
				// 如果i+2越界，说明当前是最后两个字符，组合解码增加1种方式
				dp[i] += 1
			}
		}
	}

	// dp[0]表示整个字符串的解码方式数
	return dp[0]
}

func maxValue(grid [][]int) int {
	// write code here
	m := len(grid)
	n := len(grid[0])
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			dp[i][j] = grid[i-1][j-1] + max(dp[i-1][j], dp[i][j-1])
		}
	}
	return dp[m][n]
}

func maxValue2(grid [][]int) int {
	// write code here
	m := len(grid)
	n := len(grid[0])
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			dp[i][j] = grid[i-1][j-1] + max(dp[i-1][j], dp[i][j-1])
		}
	}
	return dp[m][n]
}

func lengthOfLongestSubstring(s string) int {
	left := 0
	n := len(s)
	numCnt := make(map[byte]int)
	res := 0
	for i := 0; i < n; i++ {
		x := s[i]
		numCnt[x]++
		for numCnt[x] > 1 {
			leave := s[left]
			numCnt[leave]--
			left++
		}
		res = max(res, i-left+1)
	}
	return res
}

func GetUglyNumber_Solution(index int) int {
	ugly2 := 1
	p2 := 0
	ugly3 := 1
	p3 := 0
	ugly5 := 1
	p5 := 0
	p := 0
	ugly := make([]int, index)
	for p < index {
		x := myMin(myMin(ugly2, ugly3), ugly5)
		ugly[p] = x
		p++
		if x >= ugly2 {
			ugly2 = ugly[p2] * 2
			p2++
		}
		if x >= ugly3 {
			ugly3 = ugly[p3] * 3
			p3++
		}
		if x >= ugly5 {
			ugly5 = ugly[p5] * 5
			p5++
		}
	}
	return ugly[index-1]
}

func myMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func FirstNotRepeatingChar(str string) int {
	// write code here
	numCnts := make(map[byte][]int)
	for i := range str {
		numCnts[str[i]] = append(numCnts[str[i]], i)
	}
	res := math.MaxInt
	for _, v := range numCnts {
		if len(v) == 1 {
			res = min(res, v[0])
		}
	}
	if res == math.MaxInt {
		return -1
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// InversePairs 计算数组中的逆序对数量
// 使用归并排序的思想，在排序过程中统计逆序对数量
// 时间复杂度：O(n log n)
// 空间复杂度：O(n)
func InversePairs(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	// 创建临时数组，避免修改原数组
	tmp := make([]int, len(nums))
	copy(tmp, nums)
	return mergeSortAndCount(nums, tmp, 0, len(nums)-1)
}

// mergeSortAndCount 对数组进行归并排序并统计逆序对数量
// nums: 原始数组
// tmp: 临时数组，用于归并过程
// left: 当前处理区间的左边界
// right: 当前处理区间的右边界
// 返回值：当前区间的逆序对数量
func mergeSortAndCount(nums, tmp []int, left, right int) int {
	// 区间只有一个元素时，逆序对数量为0
	if left >= right {
		return 0
	}

	// 计算中间位置
	mid := left + (right-left)/2
	// 递归计算左右子数组的逆序对数量
	count := mergeSortAndCount(tmp, nums, left, mid) +
		mergeSortAndCount(tmp, nums, mid+1, right)

	// 归并过程，同时统计跨左右子数组的逆序对数量
	i, j, k := left, mid+1, left
	for i <= mid && j <= right {
		if tmp[i] <= tmp[j] {
			nums[k] = tmp[i]
			i++
			// 当左半部分的元素小于右半部分的元素时
			// 说明右半部分中已经处理过的元素(j - (mid+1)个)都小于当前左半部分元素
			// 这些元素与当前左半部分元素构成逆序对
			count += j - (mid + 1)
		} else {
			nums[k] = tmp[j]
			j++
		}
		k++
	}

	// 处理左半部分剩余元素
	for i <= mid {
		nums[k] = tmp[i]
		i++
		k++
		// 右半部分所有元素都与当前左半部分元素构成逆序对
		count += j - (mid + 1)
	}

	// 处理右半部分剩余元素
	for j <= right {
		nums[k] = tmp[j]
		j++
		k++
	}

	return count
}

func FindFirstCommonNode(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	if pHead1 == nil || pHead2 == nil {
		return nil
	}
	a := pHead2
	b := pHead1
	cnt := 0
	for pHead1 != pHead2 {
		if pHead1 == nil {
			cnt++
			pHead1 = a
		} else {
			pHead1 = pHead1.Next
		}
		if pHead2 == nil {
			cnt++
			pHead2 = b
		} else {
			pHead2 = pHead2.Next
		}
		if cnt > 1 {
			return nil
		}
	}
	return pHead2
}

func FindFirstCommonNode2(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	if pHead1 == nil || pHead2 == nil {
		return nil
	}
	p1length := 0
	p2length := 0
	a := pHead1
	b := pHead2
	for a != nil {
		a = a.Next
		p1length++
	}
	for b != nil {
		b = b.Next
		p2length++
	}
	long, short := pHead1, pHead2
	diff := p1length - p2length
	if p2length > p1length {
		long, short = pHead2, pHead1
		diff = p2length - p1length
	}
	for i := 0; i < diff; i++ {
		long = long.Next
	}
	for long != nil && short != nil {
		if long == short {
			return long
		}
		long = long.Next
		short = short.Next
	}

	return nil

}

// 53
func GetNumberOfK(nums []int, k int) int {
	first := lowerBound(nums, k)
	if first == len(nums) {
		return -1
	}
	last := lowerBound(nums, k+1) - 1
	return last - first + 1
}

func lowerBound(nums []int, k int) int {
	left := -1
	right := len(nums)
	for left+1 < right {
		mid := left + (right-left)/2
		if nums[mid] < k {
			left = mid
		} else {
			right = mid
		}
	}
	return right
}

func GetMissNumber(nums []int) int {
	left := -1
	right := len(nums)
	for left+1 < right {
		mid := left + (right-left)/2
		if nums[mid] != mid {
			right = mid
		} else {
			left = mid
		}
	}
	return right
}

func GetNumSameAsIndex(nums []int) int {
	left := -1
	right := len(nums)
	for left+1 < right {
		mid := left + (right-left)/2
		if nums[mid] == mid {
			return mid
		} else if nums[mid] > mid {
			right = mid
		} else {
			left = mid
		}
	}
	return -1
}

func KthNode(proot *TreeNode, k int) int {
	if proot == nil || k == 0 {
		return -1
	}
	res := -1
	var dfs = func(root *TreeNode) {}
	dfs = func(root *TreeNode) {
		if root == nil || res != -1 {
			return
		}
		dfs(root.Left)
		if res != -1 {
			return
		}
		k--
		if k == 0 {
			res = root.Val
			return
		}
		dfs(root.Right)
	}
	dfs(proot)
	return res
}

// 55
func TreeDepth(pRoot *TreeNode) int {
	if pRoot == nil {
		return 0
	}
	left := TreeDepth(pRoot.Left)
	right := TreeDepth(pRoot.Right)
	if left > right {
		return left + 1
	} else {
		return right + 1
	}
}

func getHeight(node *TreeNode) int {
	if node == nil {
		return 0
	}
	leftH := getHeight(node.Left)
	if leftH == -1 {
		return -1 // 提前退出，不再递归
	}
	rightH := getHeight(node.Right)
	if rightH == -1 || abs(leftH-rightH) > 1 {
		return -1
	}
	return max(leftH, rightH) + 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func isBalanced(root *TreeNode) bool {
	return getHeight(root) != -1
}

// 56.数组中只出现一次的两个数字
func FindNumsAppearOnce(nums []int) []int {
	res := 0
	for _, v := range nums {
		res ^= v
	}
	// mask := res & -res 取出的是 res 的最低位 1。
	// 设它对应第 p 位，则 mask 只有第 p 位为 1，其余为 0。
	mask := res & -res

	x, y := 0, 0
	for _, v := range nums {
		if v&mask == 0 {
			x ^= v
		} else {
			y ^= v
		}
	}
	// 确保固定返回顺序，便于测试
	if x < y {
		return []int{x, y}
	}
	return []int{y, x}
}

/*
在一个数组中除一个数字只出现一次之外，其他数字都出现了三次。
请找出那个只出现一次的数字。
*/
// FindNumsAppearOnce3 返回只出现一次的数字，其他数字都出现三次
// 思路：位运算状态机（ones、twos）
// 对于每一位，出现次数在模3意义下循环：
// ones 记录出现一次的位，twos 记录出现两次的位；第三次出现时这位会被同时从 ones、twos 清除。
// 该方法时间 O(n)，空间 O(1)，且天然支持负数。
func FindNumsAppearOnce3(nums []int) int {
	ones, twos := 0, 0
	for _, v := range nums {
		ones = (ones ^ v) & ^twos
		twos = (twos ^ v) & ^ones
	}
	return ones
}

// FindNumsAppearOnce3ByCount 方法2：逐位计数，对 3 取模后还原结果
// 复杂度：O(n * w)，w 为整型位宽；空间 O(1)。支持负数。
func FindNumsAppearOnce3ByCount(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	// 使用当前平台的整型位宽，确保兼容 32/64 位
	width := strconv.IntSize
	ans := 0
	for i := 0; i < width; i++ {
		sum := 0
		mask := 1 << i
		for _, v := range nums {
			if v&mask != 0 {
				sum++
			}
		}
		if sum%3 != 0 {
			ans |= mask
		}
	}
	return ans
}

// 57. 和为S的两个数字
func FindNumbersWithSum(array []int, sum int) []int {
	targetMap := make(map[int]struct{})
	for _, x := range array {
		need := sum - x
		if _, ok := targetMap[need]; ok {
			return []int{need, x}
		}
		targetMap[x] = struct{}{}
	}
	return []int{-1, -1}
}

func FindNumbersWithSum2(array []int, sum int) []int {
	left := 0
	right := len(array) - 1
	for left < right {
		csum := array[left] + array[right]
		if csum == sum {
			return []int{array[left], array[right]}
		} else if csum < sum {
			left++
		} else {
			right--
		}
	}
	return []int{}
}

func FindContinuousSeq(sum int) [][]int {
	res := make([][]int, 0)
	small := 1
	big := 2
	windows := []int{small, big}
	mid := (1 + sum) / 2
	curSum := small + big
	for small < mid {
		if curSum == sum {
			temp := make([]int, len(windows))
			copy(temp, windows)
			res = append(res, temp)
		}
		for curSum > sum && small < mid {
			curSum -= small
			small++
			windows = windows[1:]
			if curSum == sum {
				temp := make([]int, len(windows))
				copy(temp, windows)
				res = append(res, temp)
			}
		}
		big++
		curSum += big
		windows = append(windows, big)
	}
	return res
}

// 58. 循环左移N位
func LeftRotateString(str string, n int) string {

	length := len(str)
	if length == 0 {
		return ""
	}
	n = length - (n % length)
	res := make([]byte, length)
	for i := 0; i < length; i++ {
		res[(i+n)%length] = str[i]
	}
	return string(res)
}

func LeftRotateString2(str string, n int) string {
	length := len(str)
	if length == 0 {
		return ""
	}
	n = n % length
	bytes := []byte(str)
	slices.Reverse(bytes[:n])
	slices.Reverse(bytes[n:])
	slices.Reverse(bytes)
	return string(bytes)
}

// 59.滑动窗口的最大值
func maxInWindows(num []int, size int) []int {
	n := len(num)
	if size == 0 || size > n {
		return []int{}
	}
	res := make([]int, 0, n-size+1)
	pq := []int{}
	for i := 0; i < n; i++ {
		for len(pq) > 0 && num[i] > num[pq[len(pq)-1]] {
			pq = pq[:len(pq)-1]
		}

		// 存放元素的索引
		pq = append(pq, i)
		// 弹出过期的元素
		if pq[0] <= i-size {
			pq = pq[1:]
		}
		if i >= size-1 {
			res = append(res, num[pq[0]])
		}
	}
	return res
}

// 61.扑克牌顺子
func IsContinuous(numbers []int) bool {
	// write code here
	zeroCnt := 0
	nums := []int{}
	for i := 0; i < len(numbers); i++ {
		if numbers[i] == 0 {
			zeroCnt++
		} else {
			nums = append(nums, numbers[i])
		}
	}
	sort.Ints(nums)
	x := nums[0]
	need := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] == x {
			return false
		}
		need += nums[i] - x - 1
		x = nums[i]
	}
	return need <= zeroCnt
}

func LastRemaining_Solution(n int, m int) int {
	// write code here
	if n < 1 || m < 1 {
		return -1
	}
	// 只剩一个元素，直接返回
	if n == 1 {
		return 0
	}
	x1 := LastRemaining_Solution(n-1, m)
	/*
		第一次淘汰的是编号 2（因为从 0 开始数：0→1→2）。
		剩下 0,1,3,4。若把 3 作为新的 0 开始重新编号：3→0, 4→1, 0→2, 1→3。
		在这个“只剩 4 人的新编号”下，递归计算得到的存活位置是 x；映射回老编号就要把它“逆偏移”回来，即 (x + m) % n。
	*/
	return (m + x1) % n
}

func maxProfit(prices []int) int {
	// write code here
	m := prices[0]
	res := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > m {
			res = max(prices[i]-m, res)
		} else {
			m = min(prices[i], m)
		}
	}
	return res
}
