package _024

import (
	"container/heap"
	"sort"
)

func busyStudent(startTime []int, endTime []int, queryTime int) int {
	n := len(startTime)
	res := 0
	for i := 0; i < n; i++ {
		if startTime[i] <= queryTime && queryTime <= endTime[i] {
			res++
		}
	}
	return res
}

func maxConsecutiveAnswers(answerKey string, k int) int {
	n := len(answerKey)
	left, right := 0, 0
	maxConsecutive := 0
	countT := 0
	countF := 0
	for right < n {
		if answerKey[right] == 'T' {
			countT++
		} else {
			countF++
		}
		right++
		for countT > k && countF > k {
			if answerKey[left] == 'T' {
				countT--
			} else {
				countF--
			}
			left++
		}
		maxConsecutive = max(maxConsecutive, right-left)
	}
	return maxConsecutive
}

func maxStrength(nums []int) int64 {
	if len(nums) == 1 {
		return int64(nums[0])
	}
	minIndex := 0
	minNum := -10
	count := 0
	zeroCnt := 0
	res := int64(1)
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			zeroCnt++
			continue
		}
		res *= int64(nums[i])
		if nums[i] < 0 {
			count++
			if nums[i] > minNum {
				minNum = nums[i]
				minIndex = i
			}
		}
	}

	// 当数组不包含正数，且负数元素小于等于 1 个时，最大积为 0。
	if zeroCnt+1 == len(nums) && count == 1 || zeroCnt == len(nums) {
		return 0
	}

	if count%2 == 0 {
		return res
	} else {
		return res / int64(nums[minIndex])
	}
}

/*
如果能够满足下述两个条件之一，则认为第 i 位学生将会保持开心：
这位学生被选中，并且被选中的学生人数 严格大于 nums[i] 。
这位学生没有被选中，并且被选中的学生人数 严格小于 nums[i] 。
*/
func countWays(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	res := 0
	// 可以都不选
	if nums[0] > 0 {
		res++
	}
	for i := 1; i < n; i++ {
		// i代表被选中的人数

		if nums[i-1] < i && // 被选中的学生人数 严格大于 nums[i]
			i < nums[i] { // 被选中的学生人数 严格小于 nums[i]
			res++
		}
	}

	// 0 <= nums[i] < nums.length
	// 一定可以都选
	return res + 1
}

/*
https://leetcode.cn/problems/clear-digits/description/
删除 第一个数字字符 以及它左边 最近 的 非数字 字符。
返回删除所有数字字符以后剩下的字符串。
*/
func clearDigits(s string) string {
	res := []byte{}
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			if len(res) > 0 {
				res = res[:len(res)-1]
			}
		} else {
			res = append(res, s[i])
		}
	}
	return string(res)
}

func maximumLength(nums []int, k int) int {
	// dp[i][j] 来表示以 nums[i] 结尾组成的最长合法序列中有 j 个数字与其在序列中的后一个数字不相等。
	//其中 i 的取值为 nums 的长度，j 不超过 k。初始时，有 dp[i][0]=1。
	n := len(nums)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, k+1)
		for j := 0; j <= k; j++ {
			dp[i][j] = -1
		}
	}
	ans := 0
	for i := 0; i < n; i++ {
		dp[i][0] = 1
		for l := 0; l <= k; l++ {
			for j := 0; j < i; j++ {
				add := 0
				if nums[i] != nums[j] {
					add = 1
				}
				if l-add >= 0 && dp[j][l-add] != -1 {
					dp[i][l] = max(dp[i][l], dp[j][l-add]+1)
				}
			}
			ans = max(ans, dp[i][l])
		}
	}
	return ans
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.cn/problems/merge-nodes-in-between-zeros/description/
func mergeNodes(head *ListNode) *ListNode {
	dummy := &ListNode{
		Val:  -1,
		Next: nil,
	}
	p := dummy
	curNum := 0
	head = head.Next
	for head != nil {
		curNum += head.Val
		if head.Val == 0 {
			p.Next = &ListNode{
				Val:  curNum,
				Next: nil,
			}
			p = p.Next
			curNum = 0
		}
		head = head.Next

	}
	return dummy.Next
}

func mergeNodes2(head *ListNode) *ListNode {
	tail := head
	cur := head.Next
	for cur.Next != nil {
		if cur.Val != 0 {
			tail.Val += cur.Val
		} else {
			tail = tail.Next
			tail.Val = 0
		}
		cur = cur.Next

	}
	tail.Next = nil
	return head
}

func maximizeWin(prizePositions []int, k int) int {
	n := len(prizePositions)
	mx := make([]int, n+1)
	ans, left := 0, 0
	for right, p := range prizePositions {
		for p-prizePositions[left] > k {
			left++
		}
		ans = max(ans, right-left+1+mx[left])
		mx[right+1] = max(mx[right], right-left+1)
	}
	return ans
}

// https://leetcode.cn/problems/find-the-maximum-number-of-marked-indices/?envType=daily-question&envId=2024-09-12
func maxNumOfMarkedIndices(nums []int) (res int) {
	sort.Ints(nums)
	n := len(nums)
	left, right := 0, (n+1)/2
	for right < n {
		if 2*nums[left] <= nums[right] {
			res += 2
			left++
		}
		right++
	}
	return res
}

func maximumRobots(chargeTimes []int, runningCosts []int, budget int64) (ans int) {
	// 使用单调栈
	q := []int{}
	sum := int64(0)
	left := 0
	for right, t := range chargeTimes {
		// 1. 入
		for len(q) > 0 && t >= chargeTimes[q[len(q)-1]] {
			// 弹出队尾比即将入队小的值，此时队首为最大值
			q = q[:len(q)-1]
		}
		q = append(q, right)
		sum += int64(runningCosts[right])

		// 2. 出
		for len(q) > 0 && int64(chargeTimes[q[0]])+int64(right-left+1)*sum > budget {
			// 最大值出队
			if q[0] == left {
				q = q[1:]
			}
			sum -= int64(runningCosts[left])
			left++
		}

		// 3. 更新答案
		ans = max(ans, right-left+1)
	}
	return
}

func removeStars(s string) string {
	res := []byte{}
	for i := 0; i < len(s); i++ {
		if s[i] != '*' {
			res = append(res, s[i])
		} else {
			if len(res) > 0 {
				res = res[len(res)-i:]
			}

		}
	}
	return string(res)
}

func numberOfPoints(nums [][]int) int {
	n := len(nums)
	maps := make(map[int]struct{})
	for i := 0; i < n; i++ {
		x, y := nums[i][0], nums[i][1]
		for j := x; j <= y; j++ {
			maps[j] = struct{}{}
		}
	}
	return len(maps)
}

func distanceBetweenBusStops(distance []int, start int, destination int) int {
	totao := 0
	zheng := 0
	n := len(distance)
	if start > destination {
		start, destination = destination, start
	}
	for i := 0; i < n; i++ {
		totao += distance[i]
		if start <= i && i < destination {
			zheng += distance[i]
		}
	}

	fan := totao - zheng
	if fan < zheng {
		return fan
	}
	return zheng
}

// https://leetcode.cn/problems/bus-routes/
func numBusesToDestination(routes [][]int, source int, target int) int {
	if source == target {
		return 0
	}
	n := len(routes)
	lineMap := make(map[int][]int) // key:车站,value:公交线路
	for i := 0; i < n; i++ {
		m := len(routes[i])
		for j := 0; j < m; j++ {
			lineMap[routes[i][j]] = append(lineMap[routes[i][j]], i)
		}
	}

	// check 是否无解
	if lineMap[source] == nil || lineMap[target] == nil {
		if source != target {
			return -1
		}
	}
	dis := map[int]int{source: 0}
	q := []int{source}
	visited := make(map[int]bool)
	for len(q) > 0 {
		x := q[0]
		q = q[1:]
		disX := dis[x]
		for _, i := range lineMap[x] {
			if visited[i] {
				continue
			}
			for _, j := range routes[i] {
				if _, ok := dis[j]; !ok {
					dis[j] = disX + 1
					q = append(q, j)
				}
			}
			visited[i] = true
		}
	}
	if d, ok := dis[target]; ok {
		return d
	}
	return -1
}

func latestTimeCatchTheBus(buses []int, passengers []int, capacity int) int {
	sort.Ints(buses)
	sort.Ints(passengers)

	n := len(buses)
	m := len(passengers)

	j := 0
	cc := 0
	for i := 0; i < n; i++ {
		for cc = capacity; cc > 0 && j < m && passengers[j] <= buses[i]; cc-- {
			j++
		}
	}
	timeMap := make(map[int]struct{})
	for i := 0; i < j; i++ {
		timeMap[passengers[i]] = struct{}{}
	}
	res := 0
	if cc > 0 {
		res = buses[len(buses)-1] //说明最后一班公交车没有坐满
	} else {
		res = passengers[j-1] // 最后一个上车的乘客的时间,之前j多加了一次
	}

	// 寻找插队的时机
	for ; res >= passengers[0]; res-- {
		if _, ok := timeMap[res]; !ok {
			break
		}
	}

	return res
}

func longestContinuousSubstring(s string) int {
	n := len(s)
	res := 1
	cur := 1
	for right := 0; right < n-1; right++ {
		if s[right+1]-s[right] == 1 {
			cur++
		} else {
			cur = 1
		}
		res = max(res, cur)
	}
	return res
}

func edgeScore(edges []int) int {
	n := len(edges)
	score := make([]int, n)
	for i := 0; i < n; i++ {
		score[edges[i]] += i
	}
	res := 0
	maxScore := 0
	for i := 0; i < n; i++ {
		if score[i] > maxScore {
			maxScore = score[i]
			res = i
		}
	}
	return res
}

// https://leetcode.cn/problems/find-the-town-judge/description/
func findJudge(n int, trust [][]int) int {
	if n == 1 {
		return 1
	}
	inDegrees := make([]int, n+1)
	outDegrees := make([]int, n+1)
	for _, v := range trust {
		inDegrees[v[1]]++
		outDegrees[v[0]]++
	}
	for i := 1; i <= n; i++ {
		if inDegrees[i] == n-1 && outDegrees[i] == 0 {
			return i
		}
	}
	return -1
}

/*
给你一个正整数数组 values，其中 values[i] 表示第 i 个观光景点的评分，并且两个景点 i 和 j 之间的 距离 为 j - i。
一对景点（i < j）组成的观光组合的得分为 values[i] + values[j] + i - j ，
也就是景点的评分之和 减去 它们两者之间的距离。
返回一对观光景点能取得的最高分。
*/
// https://leetcode.cn/problems/best-sightseeing-pair/description/
func maxScoreSightseeingPair(values []int) int {

	n := len(values)
	res := 0
	for i := 0; i < n; i++ {
		jMax := min(n, i+values[i]+1)
		for j := i + 1; j < jMax; j++ {
			cur := values[i] + i + values[j] - j
			res = max(res, cur)
		}

	}
	return res
}

func OfficialMaxScoreSightseeingPair(values []int) int {
	ans, mx := 0, values[0]+0
	for j := 1; j < len(values); j++ {
		ans = max(ans, mx+values[j]-j)
		// 边遍历边维护
		mx = max(mx, values[j]+j)
	}
	return ans
}

func maximumSubsequenceCount(text string, pattern string) int64 {
	x, y := pattern[0], pattern[1]
	cntX := 0
	cntY := 0
	res := int64(0)
	for i := range text {
		c := text[i]
		// x==y 的情况下 要先更新答案再更新Y
		if c == y {
			res += int64(cntX)
			cntY++
		}
		if c == x {
			cntX++
		}
	}
	return res + int64(max(cntY, cntX))
}

// https://leetcode.cn/problems/naming-a-company/
func distinctNames(ideas []string) int64 {
	group := [26]map[string]bool{}
	for i := range group {
		group[i] = map[string]bool{}
	}
	for _, s := range ideas {
		group[s[0]-'a'][s[1:]] = true // 按照首字母分组
	}
	res := int64(0)
	for i, a := range group { // 枚举所有名字
		for _, b := range group[i:] {
			same := 0
			for mm := range a {
				if b[mm] {
					same++
				}
			}
			res += int64(len(a)-same) * int64(len(b)-same)
		}
	}
	return res * 2
}

func takeCharacters(s string, k int) int {
	if k == 0 {
		return 0
	}
	step := 0
	lettersCnt := make(map[byte]int)
	return dp(s, k, lettersCnt, step, 0, len(s)-1)
}

func dp(s string, k int, cnts map[byte]int, steps int, leftIndex, rightIndex int) int {
	// 遍历完成没有找到满足条件的，返回-1
	if leftIndex > rightIndex {
		return -1
	}
	leftMap := make(map[byte]int)
	rightMap := make(map[byte]int)
	for b, i := range cnts {
		leftMap[b] = i
		rightMap[b] = i
	}
	leftMap[s[leftIndex]]++
	rightMap[s[rightIndex]]++
	leftValid := 0
	for _, v := range leftMap {
		if v >= k {
			leftValid++
		}
	}
	if leftValid == 3 {
		return steps + 1
	}
	rightValid := 0
	for _, v := range rightMap {
		if v >= k {
			rightValid++
		}
	}
	if rightValid == 3 {
		return steps + 1
	}
	left := dp(s, k, leftMap, steps+1, leftIndex+1, rightIndex)
	right := dp(s, k, rightMap, steps+1, leftIndex, rightIndex-1)
	return min(left, right)
}

func takeCharacters2(s string, k int) int {
	if k == 0 {
		return 0
	}
	cnt := make(map[uint8]int)
	for i := 0; i < len(s); i++ {
		cnt[s[i]-'a']++
	}
	for _, i := range cnt {
		if i < k {
			return -1
		}
	}
	// 滑动窗口，找到满足条件的最大长度
	left, right := 0, 0
	ans := 0
	for right < len(s) {
		c := s[right] - 'a'
		cnt[c]--         // 相当于加入窗口
		for cnt[c] < k { // 窗口之外的 c 不足 k
			cnt[s[left]-'a']++ // 移出窗口，相当于取走 s[left]
			left++
		}
		ans = max(ans, right-left+1)
		right++
	}
	return len(s) - ans
}

func timeRequiredToBuy(tickets []int, k int) int {
	cntK := tickets[k]
	res := cntK
	//for ; cntK > 0; cntK-- {
	//	for i := 0; i < len(tickets); i++ {
	//		if i < k && tickets[i] > 0 {
	//			tickets[i]--
	//			res++
	//		}
	//		if i > k && tickets[i] > 0 && cntK > 1 {
	//			tickets[i]--
	//			res++
	//		}
	//	}
	//}
	for i := 0; i < len(tickets); i++ {
		if i < k {
			if tickets[i] > tickets[k] {
				res += cntK
			} else {
				res += tickets[i]
			}
		}
		if i > k {
			if tickets[i] > tickets[k]-1 {
				res += tickets[k] - 1
			} else {
				res += tickets[i]
			}
		}
	}
	return res
}

/*
请你设计一个管理 n 个座位预约的系统，座位编号从 1 到 n 。
请你实现 SeatManager 类：
SeatManager(int n) 初始化一个 SeatManager 对象，它管理从 1 到 n 编号的 n 个座位。所有座位初始都是可预约的。
int reserve() 返回可以预约座位的 最小编号 ，此座位变为不可预约。
void unreserve(int seatNumber) 将给定编号 seatNumber 对应的座位变成可以预约。
*/

// https://leetcode.cn/problems/seat-reservation-manager/description/
type SeatManager struct {
	SeatList []int
}

func Constructor(n int) SeatManager {
	list := make([]int, n)
	for i := 0; i < n; i++ {
		list[i] = i + 1
	}
	s := SeatManager{list}
	return s
}

func (this *SeatManager) Reserve() int {
	return heap.Pop(this).(int)
}

func (this *SeatManager) Unreserve(seatNumber int) {
	heap.Push(this, seatNumber)
}

func (this *SeatManager) Len() int {
	return len(this.SeatList)
}

func (this *SeatManager) Less(i, j int) bool {
	return this.SeatList[i] < this.SeatList[j]
}

func (this *SeatManager) Swap(i, j int) {
	this.SeatList[i], this.SeatList[j] = this.SeatList[j], this.SeatList[i]
}

func (this *SeatManager) Push(x any) {
	this.SeatList = append(this.SeatList, x.(int))
}

func (this *SeatManager) Pop() any {
	x := this.SeatList[len(this.SeatList)-1]
	this.SeatList = this.SeatList[0 : len(this.SeatList)-1]
	return x
}
