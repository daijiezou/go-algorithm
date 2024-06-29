package _7game

// https://leetcode.cn/problems/burst-balloons/description/
/*
有 n 个气球，编号为0 到 n - 1，每个气球上都标有一个数字，这些数字存在数组 nums 中。

现在要求你戳破所有的气球。戳破第 i 个气球，你可以获得 nums[i - 1] * nums[i] * nums[i + 1] 枚硬币。 这里的 i - 1 和 i + 1 代表和 i 相邻的两个气球的序号。如果 i - 1或 i + 1 超出了数组的边界，那么就当它是一个数字为 1 的气球。

求所能获得硬币的最大数量。
*/
var maxScore int

func maxCoins(nums []int) int {
	n := len(nums)
	// 添加两侧的虚拟气球
	scores := make([]int, n+2)
	scores[0] = 1
	scores[n+1] = 1
	copy(scores[1:n+1], nums)

	// 初始化dp数组
	dp := make([][]int, n+2)
	for i := 0; i <= n+1; i++ {
		dp[i] = make([]int, n+2)
	}
	// 从下到上
	for i := n + 1; i >= 0; i-- {
		// 从左到右
		for j := i + 1; j <= n+1; j++ {
			// 最后戳破的气球是哪个？
			// 在(i,j)中任意一个气球都有可能最后一个被戳破，在这里进行便利
			for k := i + 1; k < j; k++ {
				dp[i][j] = max(dp[i][j], dp[i][k]+dp[k][j]+scores[k]*scores[i]*scores[j])
			}
		}
	}
	return dp[0][n+1]
}

func maxCoins2(nums []int) int {
	n := len(nums)
	// 添加两侧的虚拟气球
	scores := make([]int, n+2)
	scores[0] = 1
	scores[n+1] = 1
	copy(scores[1:n+1], nums)
	memo := make([][]int, n+2)
	for i := 0; i <= n+1; i++ {
		memo[i] = make([]int, n+2)
		for j := 0; j <= n+1; j++ {
			memo[i][j] = -1
		}
	}
	return maxCoins2Dp(scores, 0, n+1, memo)
}

func maxCoins2Dp(scores []int, left int, right int, memo [][]int) int {
	if left >= right {
		return 0
	}
	if memo[left][right] != -1 {
		return memo[left][right]
	}
	for i := left + 1; i < right; i++ {
		sum := scores[i] * scores[left] * scores[right]
		sum += maxCoins2Dp(scores, left, i, memo) + maxCoins2Dp(scores, i, right, memo)
		memo[left][right] = max(memo[left][right], sum)
	}
	return memo[left][right]
}
