package meeting150

import "math"

/*
多维动态数组
*/

func minimumTotal(triangle [][]int) int {
	m := len(triangle)
	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, m)
		for j := range memo[i] {
			memo[i][j] = math.MinInt32 // math.MinInt 表示没有计算过
		}
	}
	var dfs func(depth int, index int) int
	dfs = func(depth int, index int) int {
		if depth == m-1 {
			return triangle[depth][index]
		}
		if memo[depth][index] != math.MinInt32 {
			return memo[depth][index]
		}
		memo[depth][index] = min(dfs(depth+1, index), dfs(depth+1, index+1)) + triangle[depth][index]
		return memo[depth][index]

	}
	return dfs(0, 0)
}

func minimumTotal2(triangle [][]int) int {
	m := len(triangle)
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, m)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt64
		}
	}
	dp[0][0] = triangle[0][0]
	for i := 1; i < m; i++ {
		for j := 0; j < len(triangle[i]); j++ {
			if j-1 >= 0 {
				dp[i][j] = min(dp[i-1][j], dp[i-1][j-1]) + triangle[i][j]
			} else {
				dp[i][j] = dp[i-1][j] + triangle[i][j]
			}
		}
	}
	res := math.MaxInt64
	for i := 0; i < m; i++ {
		res = min(res, dp[m-1][i])
	}
	return res
}

func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	memo := make([][]int, m)
	for i := 0; i < m; i++ {
		memo[i] = make([]int, n)
		for j := 0; j < n; j++ {
			memo[i][j] = math.MaxInt
		}
	}
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i == m {
			return math.MaxInt32 / 2
		}
		if j == n {
			return math.MaxInt32 / 2
		}
		if i == m-1 && j == n-1 {
			return grid[i][j]
		}
		if memo[i][j] != math.MaxInt {
			return memo[i][j]
		}
		memo[i][j] = grid[i][j] + min(dfs(i+1, j), dfs(i, j+1))
		return memo[i][j]
	}
	return dfs(0, 0)
}

func minPathSum2(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = math.MaxInt
		}
	}
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
