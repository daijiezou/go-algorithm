package _7_dp

func climbStairs(n int) int {
	memo := make([]int, n+1)
	return climbStairsDp(memo, n)
}

func climbStairsDp(memo []int, n int) int {
	if n == 2 {
		return 2
	}
	if n == 1 {
		return 1
	}
	if memo[n] != 0 {
		return memo[n]
	}
	memo[n] = climbStairsDp(memo, n-1) + climbStairsDp(memo, n-2)
	return memo[n]
}

// https://leetcode.cn/problems/count-ways-to-build-good-strings/
func countGoodStrings(low int, high int, zero int, one int) int {
	memo := make([]int, high+1)
	for i := 0; i <= high; i++ {
		memo[i] = -1
	}
	mod1 := 1000000007
	var dfs func(i int) int
	dfs = func(i int) int {
		if i < 0 {
			return 0
		}
		if i == 0 {
			return 1
		}
		if memo[i] != -1 {
			return memo[i]
		}
		memo[i] = (dfs(i-one) + dfs(i-zero)) % mod1
		return memo[i]
	}
	res := 0
	for i := low; i <= high; i++ {
		res = (res + dfs(i)) % mod1
	}
	return res
}

func countGoodStrings2(low int, high int, zero int, one int) int {
	memo := make([]int, high+1)
	for i := 0; i <= high; i++ {
		memo[i] = -1
	}
	mod1 := 1000000007
	dp := make([]int, high+1)
	res := 0
	dp[0] = 1
	for i := 1; i <= high; i++ {
		temp := 0
		if i >= one {
			dp[i] = dp[i-one] % mod1
		}
		if i >= zero {
			dp[i] = temp + dp[i-zero]%mod1
		}
		if i >= low {
			res += dp[i] % mod1
		}
	}

	return res
}
