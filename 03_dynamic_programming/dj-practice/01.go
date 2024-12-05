package dj_practice

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1
	}
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			} else {
				dp[i] = min(dp[i-coin]+1, dp[i])
			}
		}
	}
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}

// 0-1 背包问题
func knapsack(W, N int, wt, val []int) int {
	// base case 已初始化
	dp := make([][]int, N+1)
	for i := range dp {
		dp[i] = make([]int, W+1)
	}
	for i := 1; i <= N; i++ {
		for w := 1; w <= W; w++ {
			if w-wt[i-1] < 0 {
				// 这种情况下只能选择不装入背包
				dp[i][w] = dp[i-1][w]
			} else {
				// 装入或者不装入背包，择优
				dp[i][w] = max(
					dp[i-1][w-wt[i-1]]+val[i-1],
					dp[i-1][w],
				)
			}
		}
	}

	return dp[N][W]
}

// 完全背包问题
func change(amount int, coins []int) int {
	m := len(coins)
	/*
		若只使用前 i 个物品（可以重复使用），当背包容量为 j 时，有 dp[i][j] 种方法可以装满背包。
	*/
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, amount+1)
	}
	for i := 1; i <= m; i++ {
		dp[i][0] = 1
	}
	for i := 1; i <= len(coins); i++ {
		for j := 1; j <= amount; j++ {
			if j >= coins[i-1] {
				// 				 注意这里是dp[i]而不是dp[i-1]表示是可以重复使用该硬币
				dp[i][j] = dp[i-1][j] + dp[i][j-coins[i-1]]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[m][amount]
}
