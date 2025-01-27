package leetcode

import (
	"container/heap"
	"container/list"
	"fmt"
	"math"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func convertDateToBinary(date string) string {
	dataSlice := strings.Split(date, "-")
	for i := 0; i < 3; i++ {
		dataSlice[i] = convertToBinary(dataSlice[i])
	}
	return strings.Join(dataSlice, "-")
}

func convertToBinary(num string) string {
	numInt64, _ := strconv.Atoi(num)
	res := strings.Builder{}
	resSlice := make([]byte, 0)
	for numInt64 > 0 {
		mod2 := numInt64 % 2
		numInt64 = numInt64 >> 1
		// 将数字 0 或 1 转为字符 '0' 或 '1'
		resSlice = append(resSlice, byte(mod2)+'0')
	}
	// 反转结果
	for i, j := 0, len(resSlice)-1; i < j; i, j = i+1, j-1 {
		resSlice[i], resSlice[j] = resSlice[j], resSlice[i]
	}
	// 将结果写入 Builder
	res.Write(resSlice)
	return res.String()
}

type MyCalendar struct {
	Calendar *list.List
}

func ConstructorCalendar() MyCalendar {
	return MyCalendar{Calendar: list.New()}
}

func (this *MyCalendar) Book(startTime int, endTime int) bool {
	for e := this.Calendar.Front(); e != nil; e = e.Next() {
		event := e.Value.([2]int)
		if event[0] < endTime && startTime < event[1] {
			// 本次日程还没结束，下个日程就开始了
			return false
		}
	}
	// 成功安排日程
	this.Calendar.PushBack([2]int{startTime, endTime})
	return true
}

type ATM struct {
	BanknotesCount  []int
	banknotesAmount []int
}

func ConstructorATM() ATM {
	return ATM{
		BanknotesCount:  make([]int, 5),
		banknotesAmount: []int{20, 50, 100, 200, 500},
	}
}

func (this *ATM) Deposit(banknotesCount []int) {
	for i := 0; i < len(banknotesCount); i++ {
		this.BanknotesCount[i] += banknotesCount[i]
	}
}

func (this *ATM) Withdraw(amount int) []int {
	res := make([]int, 5)
	for i := 4; i >= 0; i-- {
		total := this.BanknotesCount[i]
		curAmount := this.banknotesAmount[i]
		cnt := amount / curAmount
		cnt = min(cnt, total)
		amount -= cnt * curAmount
		res[i] = cnt
		if amount == 0 {
			// 进行扣款
			for k := 0; k < len(res); k++ {
				this.BanknotesCount[k] -= res[k]
			}
			return res
		}
	}
	return []int{-1}
}

func maxConsecutive(bottom int, top int, special []int) int {
	res := 0
	slices.Sort(special)
	n := len(special)
	res = special[0] - bottom
	for i := 1; i < n; i++ {
		res = max(res, special[i]-special[i-1]-1)
	}
	res = max(res, top-special[n-1])
	return res
}

func countKeyChanges(s string) int {
	s = strings.ToLower(s)
	cnt := 0
	for i := 1; i < len(s); i++ {
		if s[i] != s[i-1] {
			cnt++
		}
	}
	return cnt
}

func largestGoodInteger(num string) string {
	cnt := 1
	res := ""
	for i := 1; i < len(num); i++ {
		cnt++
		if num[i] != num[i-1] {
			cnt = 1
		}
		if cnt == 3 {
			//numInt,_ := strconv.Atoi()
			res = max(res, num[i-2:i+1])
		}
	}
	return res
}

func validSubstringCount(s string, t string) int64 {
	need := make(map[byte]int)
	window := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	left, right := 0, 0
	valid := 0
	var res int64
	n := len(s)
	for right < len(s) {
		current := s[right]
		right++
		if _, ok := need[current]; ok {
			window[current]++
			if window[current] == need[current] {
				valid++
			}
		}
		for valid == len(need) {
			res += int64(n - right + 1)
			toDelete := s[left]
			left++
			if _, ok := need[toDelete]; ok {
				if window[toDelete] == need[toDelete] {
					valid--
				}
				window[toDelete]--
			}
		}
	}
	return res
}

func largestCombination(candidates []int) int {
	maxLen := func(x int) int {
		cnt := 0
		for i := 0; i < len(candidates); i++ {
			if (candidates[i])&(1<<x) != 0 {
				cnt++
			}
		}
		return cnt
	}
	res := 0
	for i := 0; i < 24; i++ {
		res = max(res, maxLen(i))
	}
	return res
}

func waysToSplitArray(nums []int) int {
	n := len(nums)
	preSum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
	}
	res := 0
	for i := 1; i < n; i++ {
		left := preSum[i] - preSum[0]
		right := preSum[n] - preSum[i]
		if left >= right {
			res++
		}
	}
	return res
}

func waysToSplitArray2(nums []int) int {
	total := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		total += nums[i]
	}
	res := 0
	curSum := 0
	for i := 0; i < n-1; i++ {
		curSum += nums[i]
		if curSum*2 >= total {
			res++
		}
	}
	return res
}

func minOperationsI(nums []int, k int) int {
	cnt := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] < k {
			cnt++
		}
	}
	return cnt
}

type mySlice []int

func (m mySlice) Len() int {
	//TODO implement me
	return len(m)
}

func (m mySlice) Less(i, j int) bool {
	return m[i] < m[j]
}

func (m mySlice) Swap(i, j int) {
	//TODO implement me
	m[i], m[j] = m[j], m[i]
}

func (m *mySlice) Push(x any) {
	*m = append(*m, x.(int))
}

func (m *mySlice) Pop() any {

	n := len(*m)
	res := (*m)[n-1]
	*m = (*m)[0 : n-1]
	return res

}

func minOperationsII(nums []int, k int) int {
	nums2 := mySlice(nums)
	heap.Init(&nums2)
	cnt := 0
	for nums2.Len() >= 2 {
		x := heap.Pop(&nums2).(int)
		y := heap.Pop(&nums2).(int)
		if x < k || y < k {
			cnt++
		} else {
			return cnt
		}
		newEle := min(x, y)*2 + max(x, y)
		heap.Push(&nums2, newEle)
	}
	return cnt
}

func rotate(nums []int, k int) {
	n := len(nums)
	k = k % n
	newNums := make([]int, n)
	for i := 0; i < n; i++ {
		newNums[(i+k)%n] = nums[i]
	}
	nums = newNums
	fmt.Println(nums)
}

func minimumSubarrayLength(nums []int, k int) int {
	ans := math.MaxInt
	for i, x := range nums {
		if x >= k {
			return 1
		}
		for j := i - 1; j >= 0 && nums[j]|x != nums[j]; j-- {
			nums[j] = nums[j] | x
			if nums[j] >= k {
				ans = min(ans, i-j+1)
			}
		}
	}
	if ans == math.MaxInt {
		return -1
	}

	return ans
}

func fucn1(nums [][]int) int {
	m := len(nums)
	memo := make([][]int, m)
	for i := 0; i < m; i++ {
		memo[i] = make([]int, m)
	}
	memo[0][0] = nums[0][0]
	return minPathSum(nums, 1, 0, memo) + nums[0][0]
}

func minPathSum(nums [][]int, depth int, row int, memo [][]int) int {
	if depth == len(nums)-1 {
		return min(nums[depth][row], nums[depth][row+1])
	}
	if memo[depth][row] != 0 {
		return memo[depth][row]
	}
	memo[depth][row] = min(minPathSum(nums, depth+1, row, memo)+nums[depth][row],
		minPathSum(nums, depth+1, row+1, memo)+nums[depth][row+1])
	return memo[depth][row]
}

func minimumTotal(triangle [][]int) int {
	m := len(triangle)
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, m)
		for j := 0; j < m; j++ {
			dp[i][j] = math.MaxInt
		}
	}
	dp[0][0] = triangle[0][0]
	for i := 1; i < m; i++ {
		for j := 0; j < len(triangle[m]); j++ {
			if j-1 >= 0 {
				dp[i][j] = min(dp[i-1][j], dp[i-1][j-1]) + triangle[i][j]
			} else {
				dp[i][j] = dp[i-1][j] + triangle[i][j]
			}
		}
	}
	res := math.MaxInt
	for i := 0; i < m; i++ {
		res = min(res, dp[m-1][i])
	}
	return res
}

func findClosestNumber(nums []int) int {
	//slices.Sort(nums)
	//zeroIndex, ok := slices.BinarySearch(nums, 0)
	//if ok {
	//	return 0
	//}
	//var res int
	//curAbs := math.MaxInt
	//if zeroIndex > 0 {
	//	curAbs = abs2(nums[zeroIndex-1])
	//	res = nums[zeroIndex-1]
	//}
	//if zeroIndex < len(nums) {
	//	if nums[zeroIndex] <= curAbs {
	//		res = nums[zeroIndex]
	//	}
	//}
	//return res
	ans := nums[0]
	for i := 0; i < len(nums); i++ {
		if abs2(nums[i]) < abs2(ans) || abs2(nums[i]) == abs2(ans) && nums[i] > ans {
			ans = nums[i]
		}
	}
	return ans

}

func abs2(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func maxValueOfCoins(piles [][]int, k int) int {
	res := 0
	var backtrack func(cnt int, curSum int)
	backtrack = func(cnt int, curSum int) {
		if cnt == 0 {
			res = max(res, curSum)
			return
		}
		for i := 0; i < len(piles); i++ {
			if len(piles[i]) > 0 {
				cur := piles[i][0]
				curSum += cur
				piles[i] = piles[i][1:]
				backtrack(cnt-1, curSum)
				curSum -= cur
				piles[i] = append([]int{cur}, piles[i]...)
			}
		}
	}
	backtrack(k, 0)
	return res
}

func maxValueOfCoins2(piles [][]int, k int) int {
	n := len(piles)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, k+1)
	}
	var dfs func(int, int) int
	dfs = func(i, j int) (res int) {
		if i < 0 {
			return
		}
		if memo[i][j] != 0 { // 之前计算过
			return memo[i][j]
		}
		// 不选这一组中的任何物品
		res = dfs(i-1, j)
		// 枚举选哪个
		v := 0
		for w := range min(j, len(piles[i])) {
			v += piles[i][w]
			// w 从 0 开始，物品体积为 w+1
			res = max(res, dfs(i-1, j-w-1)+v)
		}
		memo[i][j] = res
		return
	}
	return dfs(n-1, k)
}

func maxValueOfCoins3(piles [][]int, k int) int {
	n := len(piles)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, k+1)
	}
	for i, pile := range piles {
		for j := 0; j <= k; j++ {
			dp[i+1][j] = dp[i][j]
			v := 0
			for w := 0; w < min(j, len(pile)); w++ {
				v += pile[w]
				dp[i+1][j] = max(dp[i+1][j], dp[i][j-w-1]+v)
			}
		}
	}

	return dp[n][k]
}

func maxCoins(piles []int) int {
	slices.Sort(piles)
	res := 0
	cnt := 0
	n := len(piles)
	for i := len(piles) - 2; i > 0 && cnt < n/3; i -= 2 {
		res += piles[i]
		cnt++
	}
	return res
}

func minimumCoins(prices []int) int {
	n := len(prices)
	var dfs func(i int) int
	dfs = func(i int) int {
		if i*2 >= n {
			return prices[i-1]
		}
		res := math.MaxInt32
		for j := i + 1; j <= 2*i+1; j++ {
			res = min(res, dfs(j))
		}
		return res + prices[i-1]
	}
	return dfs(1)
}

func minimumCoins_dp(prices []int) int {
	n := len(prices)
	dp := make([]int, n+1)
	// dp[i] 表示在购买第 i 个水果的前提下，
	// 获得第 i 个及其后面的水果所需要的最少金币数。注意 i 从 1 开始。
	for i := n; i >= 1; i-- {
		if i*2 >= n {
			dp[i] = prices[i-1]
		} else {
			res := math.MaxInt32
			// 如果[i+1,2*i] 都不购买的话，则必须购买2*i+1，比较哪种方式花费的金币少
			for j := i + 1; j <= 2*i+1; j++ {
				res = min(res, dp[j])
			}
			dp[i] = res + prices[i-1]
		}
	}
	return dp[1]
}

func minimumMoney(transactions [][]int) int64 {
	allLose := 0
	maxFirstrLose := 0
	for i := 0; i < len(transactions); i++ {
		cost := transactions[i][0]
		back := transactions[i][1]
		allLose += max(0, cost-back)
		if cost < back {
			// 赚钱的
			maxFirstrLose = max(maxFirstrLose, cost)
		} else {
			// 亏钱的 因为之前计算总亏损的时候将刚back减去了，需要加回来
			maxFirstrLose = max(maxFirstrLose, back)
		}

	}
	return int64(allLose + maxFirstrLose)
}

func combinationSum2(candidates []int, target int) [][]int {
	backtack := func(precombina []int, presum int, start int) {
	}
	// 需要去重，所以先排序
	sort.Ints(candidates)
	res := [][]int{}
	//redup := make(map[[51]int]struct{})
	backtack = func(precombina []int, presum int, start int) {
		if presum == target {
			temp := make([]int, len(precombina))
			copy(temp, precombina)
			//key := genKey(temp)
			//if _, ok := redup[key]; ok {
			//	return
			//}
			//redup[key] = struct{}{}
			res = append(res, temp)
		}
		if presum > target {
			return
		}
		for i := start; i < len(candidates); i++ {
			// 说明没有选candidates[i-1]，则所有等于该数的都不选，防止重复
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}
			precombina = append(precombina, candidates[i])
			backtack(precombina, presum+candidates[i], i+1)
			precombina = precombina[:len(precombina)-1]
		}
	}
	backtack([]int{}, 0, 0)
	return res
}

func genKey(nums []int) [51]int {
	keys := [51]int{}
	for i := 0; i < len(nums); i++ {
		keys[nums[i]]++
	}
	return keys
}

func jumpjump(nums []int) int {
	step := 0
	end := 0
	formest := 0
	for i := 0; i < len(nums); i++ {
		formest = max(nums[i]+i, formest)
		if end == i {
			step++
			end = formest
		}
	}
	return step
}

func minTaps(n int, ranges []int) int {
	nums := make([]int, n+1)
	for i := 0; i < len(ranges); i++ {
		start := max(0, i-ranges[i])
		nums[start] = max(nums[start], i+ranges[i])
	}
	step := 0
	end := 0
	nextRight := 0
	for i := 0; i <= n-1; i++ {
		nextRight = max(nums[i], nextRight)
		if end == i { //需要新造一座桥
			if nextRight == i {
				return -1
			}
			end = nextRight
			step++
		}
	}
	return step
}
