package _1_core_platform

import "math"

func fib(n int) int {
	note := make([]int, n+1)
	return fibHelp(n, note)
}

func fibHelp(n int, note []int) int {
	if n == 0 || n == 1 {
		return n
	}
	if note[n] != 0 {
		return note[n]
	}
	note[n] = fibHelp(n-1, note) + fibHelp(n-2, note)
	return note[n]
}

func fib2(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func fib3(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	dp_i_1 := 1
	dp_i_2 := 0
	for i := 2; i <= n; i++ {
		dp_i := dp_i_1 + dp_i_2
		dp_i_2 = dp_i_1
		dp_i_1 = dp_i
	}
	return dp_i_1
}

/*
	凑零钱
*/

func coinChange(coins []int, amount int) int {
	note := make([]int, amount+1)
	for i := 0; i < amount+1; i++ {
		note[i] = -666
	}
	return dp2(coins, amount, note)
}

// 暴利解法
func db(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	res := math.MaxInt32
	for _, coin := range coins {
		subProblem := db(coins, amount-coin)
		if subProblem == -1 {
			continue
		}
		res = mymin(res, subProblem+1)
	}
	if res == math.MaxInt32 {
		return -1
	}
	return res
}

func dp2(coins []int, amount int, note []int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	if note[amount] != -666 {
		return note[amount]
	}
	res := math.MaxInt32
	for _, coin := range coins {
		subProblem := dp2(coins, amount-coin, note)
		if subProblem == -1 {
			continue
		}
		res = mymin(res, subProblem+1)
	}
	if res == math.MaxInt32 {
		note[amount] = -1
	} else {
		note[amount] = res
	}
	return note[amount]
}

func mymin(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func coinChange2(coins []int, amount int) int {
	dp := make([]int, amount+1)
	//数组大小为 amount + 1，初始值也为 amount + 1
	for i := 0; i < amount+1; i++ {
		dp[i] = amount + 1
	}

	// base case
	dp[0] = 0
	for i := 0; i < amount+1; i++ {
		for _, coin := range coins {
			// 无解，直接跳过
			if i-coin < 0 {
				continue
			}
			dp[i] = mymin(dp[i], dp[i-coin]+1)
		}
	}
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = myMax(dp[j]+1, dp[i])
			}

		}
	}
	res := math.MinInt32
	for i := 0; i < len(dp); i++ {
		res = myMax(dp[i], res)
	}
	return res
}

func myMax(x, y int) int {
	if x < y {
		return y
	}
	return x
}
