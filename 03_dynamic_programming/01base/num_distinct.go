package _1base

func numDistinct(s string, t string) int {
	m := len(s)
	n := len(t)
	memo := make([][]int, m)
	for i := 0; i < m; i++ {
		memo[i] = make([]int, n)
		for j := 0; j < n; j++ {
			memo[i][j] = -1
		}
	}
	return numDistinctDP(s, 0, t, 0, memo)
}

func numDistinctDP(s string, i int, t string, j int, memo [][]int) int {
	if j == len(t) {
		return 1
	}

	if len(s)-i < len(t)-j {
		return 0
	}

	if memo[i][j] != -1 {
		return memo[i][j]
	}
	res := 0
	for k := i; k < len(s); k++ {
		if s[k] == t[j] {
			// 累加结果
			res += numDistinctDP(s, k+1, t, j+1, memo)
		}
	}
	memo[i][j] = res
	return memo[i][j]
}
