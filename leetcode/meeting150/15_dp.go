package meeting150

func climbStairs(n int) int {
	dp := make([]int, n+2)
	dp[1] = 1
	dp[2] = 2
	for i := 0; i < n; i++ {
		dp[i+2] = dp[i] + dp[i+1]
	}
	return dp[n+1]
}

func rob(nums []int) int {
	n := len(nums)
	dp := make([]int, n+2)
	for i := 0; i < n; i++ {
		dp[i+2] = max(dp[i+1], dp[i]+nums[i])
	}
	return dp[n+1]
}

func wordBreak(s string, wordDict []string) bool {
	words := make(map[string]bool, len(wordDict))
	for _, w := range wordDict {
		words[w] = true
	}
	var dfs func(n int) int
	n := len(s)
	memo := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		memo[i] = -1 // 表示没有计算过
	}
	dfs = func(i int) int {
		if i == 0 {
			return 1
		}
		if memo[i] != -1 {
			return memo[i]
		}
		for j := i - 1; j >= 0; j-- {
			if words[s[j:i]] && dfs(j) == 1 {
				memo[i] = 1
				return 1
			}
		}
		memo[i] = 0
		return 0
	}
	return dfs(n) == 1
}

func wordBreak2(s string, wordDict []string) bool {
	words := make(map[string]bool, len(wordDict))
	for _, w := range wordDict {
		words[w] = true
	}
	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := i - 1; j >= 0; j-- {
			if dp[j] && words[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[n]
}

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 0; i <= amount; i++ {
		dp[i] = amount + 1
	}
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}
			dp[i] = min(dp[i], dp[i-coin]+1)
		}
	}
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}

// 最长递增子序列
func lengthOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	res := 1
	for i := 1; i < n; i++ {
		for j := i - 1; j >= 0; j-- {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		res = max(res, dp[i])
	}
	return res
}
