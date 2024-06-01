package _1_base

import "math"

func minFallingPathSum(matrix [][]int) int {
	minRes := math.MaxInt32
	m := len(matrix)
	n := len(matrix[0])
	// 备忘录里的值初始化为 66666
	memo := make([][]int, m)
	for i := 0; i < n; i++ {
		memo[i] = make([]int, n)
		for j := 0; j < n; j++ {
			memo[i][j] = 66666
		}
	}
	// 终点可能在最后一行的任意一列
	for j := 0; j < n; j++ {
		minRes = min(dp(matrix, m-1, j, memo), minRes)
	}
	return minRes
}

// 从第一行（matrix[0][..]）向下落，
// 落到位置 matrix[i][j] 的最小路径和为 dp(matrix, i, j)。
func dp(matrix [][]int, i, j int, memo [][]int) int {
	if i < 0 || j < 0 || i >= len(matrix) || j >= len(matrix[0]) {
		return 99999
	}
	if i == 0 {
		return matrix[i][j]
	}
	if memo[i][j] != 66666 {
		return memo[i][j]
	}
	memo[i][j] = matrix[i][j] +
		threeMin(dp(matrix, i-1, j-1, memo), dp(matrix, i-1, j, memo), dp(matrix, i-1, j+1, memo))
	return memo[i][j]
}

func threeMin(a int, b int, c int) int {
	if a < b {
		if a < c {
			return a
		} else {
			return c
		}
	} else if b < c {
		return b
	} else {
		return c
	}
}
