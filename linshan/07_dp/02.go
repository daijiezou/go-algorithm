package _7_dp

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
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	n := len(nums)
	dp := make([]int, n)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[n-1]
}
