package _7game

import "math"

// https://leetcode.cn/problems/minimum-path-sum/
// 最短路径
func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	// dp数组定义为到从0,0到到i，j所需的最短路径为dp[i][j]
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	// 初始化baseCase
	// 第一行和第一列只能走直线
	dp[0][0] = grid[0][0]
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for j := 1; j < n; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[m-1][n-1]
}

// 从 grid[i][j] 到达终点（右下角）所需的最少生命值是 calculateMinimumHPdp(grid, i, j)。
func calculateMinimumHP(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	// 备忘录中都初始化为 -1
	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	return calculateMinimumHPdp(grid, 0, 0, memo)
}

// 备忘录，消除重叠子问题
func calculateMinimumHPdp(grid [][]int, i int, j int, memo [][]int) int {
	m, n := len(grid), len(grid[0])
	// base case
	if i == m-1 && j == n-1 {
		if grid[i][j] >= 0 {
			return 1
		}
		return -grid[i][j] + 1
	}
	if i == m || j == n {
		return math.MaxInt32
	}
	if memo[i][j] != -1 {
		return memo[i][j]
	}
	// 状态转移逻辑
	res := min(
		calculateMinimumHPdp(grid, i, j+1, memo),
		calculateMinimumHPdp(grid, i+1, j, memo),
	) - grid[i][j]
	// 骑士的生命值至少为 1
	if res <= 0 {
		memo[i][j] = 1
	} else {
		memo[i][j] = res
	}
	return memo[i][j]
}

func calculateMinimumHPDP(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
		for j := 0; j < n+1; j++ {
			dp[i][j] = math.MaxInt32
		}
	}

	// 设置baseCase
	if grid[m-1][n-1] <= 0 {
		dp[m-1][n-1] = -grid[m-1][n-1] + 1
	} else {
		dp[m-1][n-1] = 1
	}

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-1 {
				continue
			}
			res := min(dp[i+1][j], dp[i][j+1]) - grid[i][j]
			if res <= 0 {
				res = 1
			}
			dp[i][j] = res
		}
	}
	return dp[0][0]
}
