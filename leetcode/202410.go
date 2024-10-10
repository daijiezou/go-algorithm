package leetcode

import (
	"math"
	"sort"
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
