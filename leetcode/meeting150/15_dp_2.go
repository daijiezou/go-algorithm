package meeting150

import (
	"math"
)

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

func uniquePathsWithObstacles(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = -1
		}
	}
	dp[0][0] = grid[0][0] ^ 1

	for i := 1; i < m; i++ {
		if grid[i][0] == 1 {
			dp[i][0] = 0
		} else {
			dp[i][0] = dp[i-1][0]
		}
	}
	for i := 1; i < n; i++ {
		if grid[0][i] == 1 {
			dp[0][i] = 0
		} else {
			dp[0][i] = dp[0][i-1]
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if grid[i][j] == 1 {
				dp[i][j] = 0
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}

func longestPalindrome(s string) string {
	n := len(s)
	res := ""
	for i := 0; i < n; i++ {
		res1 := longest(s, i, i)
		if len(res1) > len(res) {
			res = res1
		}
		res2 := longest(s, i, i+1)
		if len(res2) > len(res) {
			res = res2
		}
	}
	return res
}

func longest(s string, left, right int) string {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return s[left+1 : right]
}

func isInterleave(s1 string, s2 string, s3 string) bool {
	m, n := len(s1), len(s2)
	// 如果长度对不上，必然不可能
	if m+n != len(s3) {
		return false
	}
	memo := make([][]int, len(s1))
	for i := 0; i < len(s1); i++ {
		memo[i] = make([]int, len(s2))
	}
	var dfs func(s1Start int, s2Start int, s3Start int) bool
	dfs = func(s1Start int, s2Start int, s3Start int) bool {
		if s3Start == len(s3) {
			return true
		}
		res := false
		if s1Start < len(s1) && s1[s1Start] == s3[s3Start] {
			res = dfs(s1Start+1, s2Start, s3Start+1)
		}
		if s2Start < len(s2) && s2[s2Start] == s3[s3Start] {
			res = res || dfs(s1Start, s2Start+1, s3Start+1)
		}
		return res
	}
	return dfs(0, 0, 0)
}

func isInterleave2(s1 string, s2 string, s3 string) bool {
	m, n := len(s1), len(s2)
	// 如果长度对不上，必然不可能
	if m+n != len(s3) {
		return false
	}
	dp := make([][]bool, len(s1)+1)
	for i := 0; i < len(s1)+1; i++ {
		dp[i] = make([]bool, len(s2)+1)
	}
	/*
		我们定义 f(i,j) 表示 s1的前 i 个元素和 s2的前 j 个元素是否能交错组成 s3的前 i+j 个元素。
	*/
	dp[0][0] = true
	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			p := i + j - 1
			if i > 0 {
				dp[i][j] = dp[i][j] || dp[i-1][j] && s1[i-1] == s3[p]
			}
			if j > 0 {
				dp[i][j] = dp[i][j] || dp[i][j-1] && s2[j-1] == s3[p]
			}
		}
	}
	return dp[m][n]
}

func minDistance(word1 string, word2 string) int {
	if word1 == word2 {
		return 0
	}
	m := len(word1)
	n := len(word2)
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		dp[i][0] = i
	}
	for j := 1; j <= n; j++ {
		dp[0][j] = j
	}
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i][j-1]+1, dp[i-1][j-1]+1, dp[i-1][j]+1)
			}
		}
	}
	return dp[m][n]
}

// 买卖股票
func maxProfit_1(prices []int) int {
	n := len(prices)
	dp := make([][2]int, n+1)
	dp[0][1] = math.MinInt32
	for i := 0; i < n; i++ {
		dp[i+1][0] = max(dp[i][0], dp[i][1]+prices[i])
		dp[i+1][1] = max(dp[i][1], -prices[i])
	}
	return dp[n][0]
}

func maxProfit_2(prices []int) int {
	n := len(prices)
	dp := make([][2]int, n+1)
	dp[0][1] = math.MinInt32
	dp[0][0] = 0
	for i := 0; i < n; i++ {
		dp[i+1][0] = max(dp[i][0], dp[i][1]+prices[i])
		dp[i+1][1] = max(dp[i][1], dp[i][0]-prices[i])
	}
	return dp[n][0]
}

func maxProfit_leng(prices []int) int {
	n := len(prices)
	dp := make([][2]int, n+2)
	dp[0][1] = math.MinInt32
	dp[1][1] = math.MinInt32
	for i := 0; i < n; i++ {
		dp[i+2][0] = max(dp[i+1][0], dp[i+1][1]+prices[i])
		dp[i+2][1] = max(dp[i+1][1], dp[i][0]-prices[i])
	}
	return dp[n][0]
}

func maxProfit_4(max_k int, prices []int) int {
	n := len(prices)
	dp := make([][][2]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([][2]int, max_k+1)
	}
	for i := 0; i < n+1; i++ {
		dp[i][0][1] = math.MinInt32
	}
	for k := max_k; k >= 1; k-- {
		dp[0][k][1] = math.MinInt32
	}
	for i := 0; i < n; i++ {
		for k := max_k; k >= 1; k-- {
			dp[i+1][k][0] = max(dp[i][k][0], dp[i][k][1]+prices[i])
			dp[i+1][k][1] = max(dp[i][k][1], dp[i][k-1][0]-prices[i])
		}

	}
	return dp[n][max_k][0]
}

func maxProfit_3(prices []int) int {
	max_k := 2
	n := len(prices)
	dp := make([][][2]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([][2]int, max_k+1)
	}

	for i := 0; i < n+1; i++ {
		dp[i][0][1] = math.MinInt32
	}
	for k := max_k; k >= 1; k-- {
		dp[0][k][1] = math.MinInt32
	}
	for i := 0; i < n; i++ {
		for k := max_k; k >= 1; k-- {
			dp[i+1][k][0] = max(dp[i][k][0], dp[i][k][1]+prices[i])
			dp[i+1][k][1] = max(dp[i][k][1], dp[i][k-1][0]-prices[i])
		}

	}
	return dp[n][max_k][0]
}

func maximalSquare(matrix [][]byte) int {
	m := len(matrix)
	n := len(matrix[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		if matrix[i][0] == '1' {
			dp[i][0] = 1
		}
	}
	for j := 0; j < n; j++ {
		if matrix[0][j] == '1' {
			dp[0][j] = 1
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == '1' {
				dp[i][j] = min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
			}
		}
	}
	lenth := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			lenth = max(lenth, dp[i][j])
		}
	}
	return lenth * lenth
}
