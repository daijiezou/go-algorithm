package leetcode

import (
	"math"
	"sort"
	"strconv"
)

func mincostTickets(days []int, costs []int) int {
	memo := make(map[int]int)

	for i := 0; i < len(days); i++ {
		memo[i] = -1
	}
	return mincostTicketsDP(days, costs, 0, memo)
}

func mincostTicketsDP(days []int, costs []int, startIndex int, memo map[int]int) int {
	if startIndex >= len(days) {
		return 0
	}
	if memo[startIndex] != -1 {
		return memo[startIndex]
	}
	startDay := days[startIndex]
	start1 := startIndex
	start7 := startIndex
	start30 := startIndex
	for i := startIndex; i < len(days); i++ {
		if days[i] < startDay+1 {
			start1++
		}
		if days[i] < startDay+7 {
			start7++
		}
		if days[i] < startDay+30 {
			start30++
		}
	}
	memo[startIndex] = min(mincostTicketsDP(days, costs, start1, memo)+costs[0], mincostTicketsDP(days, costs, start7, memo)+costs[1], mincostTicketsDP(days, costs, start30, memo)+costs[2])
	return memo[startIndex]
}

func minSpeedOnTime(dist []int, hour float64) int {
	if float64(len(dist))-1 > hour {
		return -1
	}

	left := 1
	right := 10000001
	for left < right {
		mid := left + (right-left)/2
		if CostHours(dist, mid, hour) > hour {
			left = mid + 1
		} else {
			right = mid
		}
	}
	if CostHours(dist, left, hour) <= hour {
		return left
	}
	return -1
}

func CostHours(dist []int, speed int, hour float64) float64 {
	sum := 0
	for i := 0; i < len(dist)-1; i++ {
		c1 := dist[i] / speed
		c2 := dist[i] % speed
		if c2 != 0 {
			c1++
		}
		sum += c1
		if float64(sum) > hour {
			return 1e9 + 1
		}
	}
	res := float64(sum) + float64(dist[len(dist)-1])/float64(speed)
	return res
}

func minCost(maxTime int, edges [][]int, passingFees []int) int {
	n := len(passingFees)
	cityMaps := make(map[int][][2]int)
	for _, edge := range edges {
		cityMaps[edge[0]] = append(cityMaps[edge[0]], [2]int{edge[1], edge[2]})
	}
	spend := minCostDp(0, n, maxTime, cityMaps, passingFees, 0)
	if spend == math.MaxInt {
		return -1
	}
	return spend + passingFees[0]

}

func minCostDp(start int, n int, maxTime int, cityMaps map[int][][2]int, passingFees []int, spendTime int) int {
	if start == n-1 {
		return 0
	}
	minSpend := math.MaxInt
	for _, to := range cityMaps[start] {
		spend := to[1]
		if spend > maxTime-spendTime {
			continue
		}
		cost := minCostDp(to[0], n, maxTime, cityMaps, passingFees, spend+spendTime) + passingFees[to[0]]
		if cost < minSpend {
			minSpend = cost
		}
	}
	return minSpend
}

func minimumTime(time []int, totalTrips int) int64 {
	left := 1
	right := 100000000000
	for left <= right {
		mid := left + (right-left)>>1
		if getTotalTrips(time, mid) >= int64(totalTrips) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return int64(left)
}

func getTotalTrips(time []int, t int) int64 {
	sum := int64(0)
	for i := 0; i < len(time); i++ {
		sum += int64(t / time[i])
	}
	return sum
}

func minRefuelStops(target int, startFuel int, stations [][]int) int {
	res := minRefuelStopsDP(target, 0, startFuel, stations, 0)
	if res < 0 || res == math.MaxInt32 {
		return -1
	}
	return res
}

func minRefuelStopsDP(target int, cur int, curFuel int, stations [][]int, index int) int {
	if cur+curFuel >= target {
		return 0
	}
	dis := target - cur
	fuel := 0
	if index < len(stations) {
		dis = stations[index][0] - cur
		fuel = stations[index][1]
	}
	if curFuel < dis {
		return math.MaxInt32
	}
	// 加油
	op1 := minRefuelStopsDP(target, cur+dis, curFuel-dis+fuel, stations, index+1) + 1
	// 不加油
	op2 := minRefuelStopsDP(target, cur+dis, curFuel-dis, stations, index+1)
	return min(op1, op2)
}

func destCity(paths [][]string) string {
	city := make(map[string]struct{})
	degree := make(map[string]int)
	for _, path := range paths {
		from := path[0]
		to := path[1]
		city[from] = struct{}{}
		city[to] = struct{}{}
		// 记录出度
		degree[from]++
	}
	for cityName, _ := range city {
		if degree[cityName] == 0 {
			return cityName
		}
	}
	return ""
}

// https://leetcode.cn/problems/find-the-number-of-good-pairs-ii/submissions/571384654/
func numberOfPairs(nums1 []int, nums2 []int, k int) int64 {
	sort.Ints(nums1)
	ans := int64(0)
	cnt1 := make(map[int]int)
	m := 0
	for i := 0; i < len(nums1); i++ {
		if nums1[i]%k == 0 {
			m = max(m, nums1[i]/k)
			cnt1[nums1[i]/k]++
		}
	}
	cnt2 := make(map[int]int)
	for i := 0; i < len(nums2); i++ {
		cnt2[nums2[i]]++
	}
	for x, cnt := range cnt2 {
		sum := 0
		for i := x; i <= m; i += x {
			sum += cnt1[i]
		}
		ans += int64(sum * cnt)
	}

	return ans
}

func duplicateNumbersXOR(nums []int) int {
	numMap := make(map[int]struct{})
	ans := 0
	for i := 0; i < len(nums); i++ {
		if _, ok := numMap[nums[i]]; ok {
			ans ^= nums[i]
		} else {
			numMap[nums[i]] = struct{}{}
		}

	}
	return ans
}

func twoEggDrop(n int) int {
	memo := make([][]int, 3)
	for i := 0; i < 3; i++ {
		memo[i] = make([]int, n+1)
		for j := 0; j < n+1; j++ {
			memo[i][j] = -1
		}
	}
	return superEggDropDP(2, n, memo)
}

func twoEggDrop2(n int) int {
	memo := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		memo[i] = -1
	}
	return twoEggDropDP(n, memo)
}

func twoEggDropDP(n int, memo []int) int {
	if n == 0 {
		return 0
	}

	if memo[n] != -1 {
		return memo[n]
	}
	res := math.MaxInt
	// 从每一层都尝试丢第一个鸡蛋
	for i := 1; i <= n; i++ {
		// 鸡蛋碎了,需要操作i次，只剩一个鸡蛋，需要每层楼都尝试一下
		brokenCnt := i
		// 鸡蛋没碎
		normalCnt := twoEggDropDP(n-i, memo) + 1
		res = min(res, max(brokenCnt, normalCnt))
	}
	memo[n] = res
	return res
}

func superEggDrop(k int, n int) int {
	memo := make([][]int, k+1)
	for i := 0; i < k+1; i++ {
		memo[i] = make([]int, n+1)
		for j := 0; j < n+1; j++ {
			memo[i][j] = -1
		}
	}
	return superEggDropDP(k, n, memo)
}

// 在有k个鸡蛋，n层楼的情况下，最少需要试几次
func superEggDropDP(k int, n int, memo [][]int) int {
	if n == 0 {
		return 0
	}
	// 只有一个鸡蛋需要试n次
	if k == 1 {
		return n
	}
	if memo[k][n] != -1 {
		return memo[k][n]
	}

	/*	// 进行状态转移
		res := math.MaxInt32
		for i := 1; i <= n; i++ {
			// 在第 i 层扔，有两种可能：碎或不碎
			// 根据这两种可能得到两个子问题：
			// 1. 在前 i-1 层楼中扔鸡蛋，鸡蛋没碎，此时鸡蛋个数不会发生变化，问题转化成
			//    了 K 个鸡蛋、n-i 层楼，即 dp(k, n-i)。
			// 2. 在第 i 层楼中扔鸡蛋，鸡蛋碎了，此时应该将鸡蛋个数减一，问题转化成
			//    了 K-1 个鸡蛋、n-1 层楼，即 dp(k-1, N-1)。
			// 因为是求至少需要多少次，所以需要在两者中取较大值
			// 两者取最大值，再加上在第 i 层扔的一次，就是这个状态下的最少实验次数。
			res = min(
				res,
				max(superEggDropDP(k, n-i, memo), superEggDropDP(k-1, i-1, memo))+1,
			)
		}
	*/

	lo := 1
	high := n
	res := math.MaxInt
	for lo <= high {
		// superEggDropDP是一个随着n增加单调递增函数，n越大，需要重试的次数就越多
		mid := lo + (high-lo)/2
		//鸡蛋碎了，楼层-1，鸡蛋-1
		broken := superEggDropDP(k-1, mid-1, memo)
		//鸡蛋没碎，表示下面的楼层都不会碎了，楼层-i
		noBroken := superEggDropDP(k, n-mid, memo)
		if broken > noBroken {
			high = mid - 1
			res = min(res, broken+1)
		} else {
			lo = mid + 1
			res = min(res, noBroken+1)
		}
	}
	memo[k][n] = res
	return res
}

// https://leetcode.cn/problems/maximum-height-of-a-triangle/
func maxHeightOfTriangle(red int, blue int) int {
	left := 1
	right := max(red, blue)
	for left <= right {
		mid := left + (right-left)/2
		if maxHeightOfTriangleB(red, blue, mid) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if maxHeightOfTriangleB(red, blue, left) {
		return left
	}
	return left - 1
}

func maxHeightOfTriangleB(red int, blue int, n int) bool {
	flagou := false
	flagji := false
	if n%2 == 0 {
		firstCnt := sumn(n - 1)
		secondCnt := sumn(n)
		if (red >= firstCnt && blue >= secondCnt) || (red >= secondCnt && blue >= firstCnt) {
			flagou = true
		}
	} else {
		firstCnt := sumn(n)
		secondCnt := sumn(n - 1)
		if (red >= firstCnt && blue >= secondCnt) || (red >= secondCnt && blue >= firstCnt) {
			flagji = true
		}
	}
	return flagou || flagji
}

func sumn(n int) int {
	sum := 0
	for n > 0 {
		sum += n
		n -= 2
	}
	return sum
}

func maxHeightOfTriangle2(red int, blue int) int {
	return max(maxHeight(red, blue), maxHeight(blue, red))
}

func maxHeight(x, y int) int {
	for i := 1; ; i++ {
		if i%2 == 0 {
			x -= i
			if x < 0 {
				return i - 1
			}
		} else {
			y -= i
			if y < 0 {
				return i - 1
			}
		}
	}
}

func minimumAverage(nums []int) float64 {
	res := math.MaxFloat64
	sort.Ints(nums)
	n := len(nums)
	for i := 0; i < n/2; i++ {
		temp := (float64(nums[i] + nums[n-1-i])) / 2
		if temp < res {
			res = temp
		}
	}
	return res
}

// https://leetcode.cn/problems/minimum-operations-to-make-binary-array-elements-equal-to-one-i/description/
func minOperations(nums []int) int {
	n := len(nums)
	ans := 0
	for i := 0; i < n-2; i++ {
		if nums[i] == 0 {
			ans++
			// 异或运算符
			// 1^1 = 0
			// 1^0 = 1
			nums[i+1] ^= 1
			nums[i+2] ^= 1
		}
	}
	if nums[n-1] == 0 || nums[n-2] == 0 {
		return -1
	}
	return ans
}

// 0 1 1 0 1
// 1 0 0 1 0
// 1 1 1 0 1
// 1 1 1 1 0
// 1 1 1 1 1

// 0 0 1 0 0 1
// 1 1 0 1 1 0
// 1 1 1 0 0 1
// 1 1 1 1 1 0
// 1 1 1 1 1 1
func minOperations2(nums []int) int {
	n := len(nums)
	ans := 0
	flag := 0
	for i := 0; i < n; i++ {
		if nums[i] == flag {
			ans++
			j := i + 1
			for ; j < n; j++ {
				if nums[j] != flag {
					break
				}
			}
			flag ^= 1
			i = j - 1
		}
	}
	return ans
}

func minOperations3(nums []int) int {
	n := len(nums)
	op := 0
	for i := 0; i < n; i++ {
		// 纪录之前操作的次数
		if nums[i] == 0 && op%2 == 0 {
			op++
		}
		if nums[i] == 1 && op%2 == 1 {
			op++
		}
	}
	return op
}

func smallestRangeI(nums []int, k int) int {
	min1 := math.MaxInt
	max1 := math.MinInt
	for i := 0; i < len(nums); i++ {
		if nums[i] < min1 {
			min1 = nums[i]
		}
		if nums[i] > max1 {
			max1 = nums[i]
		}
		if max1-min1 <= 2*k {
			return 0
		}
	}
	return max1 - min1 - 2*k
}

func smallestRangeII(nums []int, k int) int {
	sort.Ints(nums)
	n := len(nums)
	if n == 0 {
		return 0
	}
	ans := nums[n-1] - nums[0]
	// 前i个变大，后n-i个变小
	// nums[0] 到 nums[i] 都变大 k，把 nums[i+1] 到 nums[n−1] 都变小 k
	for i := 0; i < n-1; i++ {
		maxn := max(nums[i]+k, nums[n-1]-k)
		minn := min(nums[0]+k, nums[i+1]-k)
		ans = min(ans, maxn-minn)
	}
	return ans
}

func countCompleteDayPairs(hours []int) (res int) {
	n := len(hours)
	cnt := [24]int{}
	for i := 0; i < n; i++ {
		res += cnt[24-(cnt[hours[i]%24])]
		cnt[hours[i]%24]++
	}
	return res
}

func countCompleteDayPairs2(hours []int) (res int64) {
	n := len(hours)
	cnt := [24]int64{}
	for i := 0; i < n; i++ {
		res += cnt[(24-hours[i]%24)%24]
		cnt[hours[i]%24]++
	}
	return res
}

func findWinningPlayer(skills []int, k int) int {
	maxIndex := 0
	winCnt := 0
	for i := 1; i < len(skills); i++ {
		if skills[i] > skills[maxIndex] {
			maxIndex = i
			winCnt = 0
		}
		winCnt++
		if winCnt == k {
			return maxIndex
		}
	}
	return maxIndex
}

// https://leetcode.cn/problems/maximum-total-reward-using-operations-i/
func maxTotalReward(rewardValues []int) int {
	sort.Ints(rewardValues)
	m := rewardValues[len(rewardValues)-1]
	dp := make([]bool, 2*m)
	dp[0] = true
	/*
		v为当前的值
		动态转移方程为：f[i][j] = f[i-1][j] || f[i-1][j-v]，
		其中 v 是当前的 rewardValue， v>j-v => j<2v
		并且 v <= j < 2v。
	*/
	for _, x := range rewardValues {
		for k := 2*x - 1; k >= x; k-- {
			dp[k] = dp[k] || dp[k-x]
		}
	}
	n := 2 * m
	for i := n - 1; i >= 0; i-- {
		if dp[i] {
			return i
		}
	}
	return 0
}

func maxTotalRewardDp(start int, reward int, rewardValues []int, memo map[string]int) int {
	if start == len(rewardValues) {
		return 0
	}
	key := strconv.Itoa(start) + ":" + strconv.Itoa(reward)
	if memo[key] != 0 {
		return memo[key]
	}
	if rewardValues[start] > reward {
		memo[key] = max(maxTotalRewardDp(start+1, reward+rewardValues[start], rewardValues, memo)+rewardValues[start],
			maxTotalRewardDp(start+1, reward, rewardValues, memo))
		return memo[key]
	}
	memo[key] = maxTotalRewardDp(start+1, reward, rewardValues, memo)
	return memo[key]
}

func validStrings(n int) []string {
	res := make([]string, 0)
	trackBack("", n, &res, false)
	return res
}

func trackBack(temp string, n int, res *[]string, preZero bool) {
	if len(temp) == n {
		*res = append(*res, temp)
		return
	}
	// 做选择
	if preZero {
		temp += "1"
		trackBack(temp, n, res, false)
	} else {
		// 做选择
		temp += "0"
		trackBack(temp, n, res, true)

		// 取消选择
		temp = temp[:len(temp)-1]
		temp += "1"
		trackBack(temp, n, res, false)
	}
}
