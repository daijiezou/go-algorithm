package _7_dp

import "sort"

func rob(nums []int) int {

	var dfs func(n int) int
	memo := make([]int, len(nums))
	for i := range memo {
		memo[i] = -1
	}
	dfs = func(n int) int {
		if n < 0 {
			return 0
		}
		if memo[n] != -1 {
			return memo[n]
		}
		memo[n] = max(dfs(n-1), dfs(n-2)+nums[n])
		return memo[n]
	}
	return dfs(len(nums) - 1)
}

func rob2(nums []int) int {

	n := len(nums)
	dp := make([]int, n+2)

	for i := n - 1; i >= 0; i-- {
		dp[i] = max(dp[i+1], dp[i+2]+nums[i])
	}
	return dp[0]
}

// https://leetcode.cn/problems/delete-and-earn/
func deleteAndEarn(nums []int) int {
	maxVal := 0
	for _, val := range nums {
		maxVal = max(maxVal, val)
	}
	sum := make([]int, maxVal+1)
	for _, val := range nums {
		sum[val] += val
	}
	return rob2(sum)

}

func countHousePlacements(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1 // 不放
	dp[1] = 2 // 放或者不放
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	const mod int = 1e9 + 7
	return dp[n] * dp[n] % mod
}

func rob3(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	return max(robRange(nums, 2, len(nums)-1)+nums[0], robRange(nums, 1, len(nums)))
}

// 定义：返回[start,end) 能抢到的最大值
func robRange(nums []int, start, end int) int {
	if start >= end {
		return 0
	}
	if end-start == 1 {
		return nums[start]
	}
	if end-start == 2 {
		return max(nums[start], nums[start+1])
	}
	dp := make([]int, end-start)
	dp[0] = nums[start]
	dp[1] = max(nums[start], nums[start+1])

	for i := start + 2; i < end; i++ {
		dp[i-start] = max(dp[i-1-start], dp[i-2-start]+nums[i])
	}
	return dp[end-start-1]
}

func maximumTotalDamage(power []int) int64 {
	cnt := make(map[int]int)
	for _, val := range power {
		cnt[val]++
	}
	a := make([]int, 0, len(power))
	for x := range cnt {
		a = append(a, x)
	}
	sort.Ints(a)
	memo := make([]int, len(cnt))
	for i := range memo {
		memo[i] = -1
	}

	/*

		不选：问题变成从 a[0] 到 a[i−1] 中选择，可以得到的伤害值之和的最大值，即 dfs(i)=dfs(i−1)。
		选：那么伤害值等于 a[i]−2 和 a[i]−1 的数不能选，问题变成从 a[0] 到 a[j−1] 中选择，
		可以得到的伤害值之和的最大值，其中 j 是最小的满足 a[j]≥a[i]−2 的数。那么 dfs(i)=dfs(j−1)+a[i]⋅cnt[a[i]]。
	*/

	var dfs func(n int) int
	dfs = func(i int) int {
		if i < 0 {
			return 0
		}
		if memo[i] != -1 {
			return memo[i]
		}
		x := a[i]
		j := i
		for j > 0 && a[j-1] >= x-2 {
			j--
		}
		memo[i] = max(dfs(i-1), dfs(j-1)+x*cnt[x])
		return memo[i]
	}
	return int64(dfs(len(cnt) - 1))
}
