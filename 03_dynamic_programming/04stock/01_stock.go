package _4stock

import "math"

/*
dp[-1][...][0] = 0
解释：因为 i 是从 0 开始的，所以 i = -1 意味着还没有开始，这时候的利润当然是 0。

dp[-1][...][1] = -infinity
解释：还没开始的时候，是不可能持有股票的。
因为我们的算法要求一个最大值，所以初始值设为一个最小值，方便取最大值。

dp[...][0][0] = 0
解释：因为 k 是从 1 开始的，所以 k = 0 意味着根本不允许交易，这时候利润当然是 0。

dp[...][0][1] = -infinity
解释：不允许交易的情况下，是不可能持有股票的。
因为我们的算法要求一个最大值，所以初始值设为一个最小值，方便取最大值。
*/

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/
func maxProfit(prices []int) int {
	n := len(prices)
	dp := make([][2]int, n)
	for i := 0; i < n; i++ {
		if i-1 == -1 {
			dp[i][0] = 0
			dp[i][1] = -prices[i]
			continue
		}

		/*
			dp[i][1][0] = max(dp[i-1][1][0], dp[i-1][1][1] + prices[i])
			dp[i][1][1] = max(dp[i-1][1][1], dp[i-1][0][0] - prices[i])
			            = max(dp[i-1][1][1], -prices[i])
			解释：k = 0 的 base case，所以 dp[i-1][0][0] = 0。

			现在发现 k 都是 1，不会改变，即 k 对状态转移已经没有影响了。
			可以进行进一步化简去掉所有 k：
			dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])
			dp[i][1] = max(dp[i-1][1], -prices[i])
		*/

		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], -prices[i])
	}
	return dp[n-1][0]
}

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/
func maxProfit2(prices []int) int {
	minPrice := math.MaxInt
	maxprofit := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		}
		if prices[i]-minPrice > maxprofit {
			maxprofit = prices[i] - minPrice
		}
	}
	return maxprofit
}
