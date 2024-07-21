package _2sub_seq

// 解决两个字符串的动态规划问题，一般都是用两个指针 i, j 分别指向两个字符串的最后，然后一步步往前移动，缩小问题的规模。
/*
给你两个单词 word1 和 word2， 请返回将 word1 转换成 word2 所使用的最少操作数  。
你可以对一个单词进行如下三种操作：
插入一个字符
删除一个字符
替换一个字符
*/
func minDistance(s1 string, s2 string) int {
	m := len(s1)
	n := len(s2)
	memo := make([][]int, m)
	for i := 0; i < m; i++ {
		memo[i] = make([]int, n)
		for j := 0; j < n; j++ {
			memo[i][j] = -1
		}
	}
	return minDistanceDp(s1, m-1, s2, n-1, memo)
}

func minDistanceDp(s1 string, i int, s2 string, j int, memo [][]int) int {
	// 返回index+1,则为剩余的操作数
	if i < 0 {
		return j + 1
	}
	if j < 0 {
		return i + 1
	}
	if memo[i][j] != -1 {
		return memo[i][j]
	}
	if s1[i] == s2[j] {
		memo[i][j] = minDistanceDp(s1, i-1, s2, j-1, memo)
	} else {
		memo[i][j] = min(
			minDistanceDp(s1, i, s2, j-1, memo)+1,   //插入
			minDistanceDp(s1, i-1, s2, j, memo)+1,   //删除
			minDistanceDp(s1, i-1, s2, j-1, memo)+1, //替换
		)
	}
	return memo[i][j]
}

func longestPalindromeSubseq(s string) int {
	m := len(s)
	memo := make([][]int, m)
	for i := 0; i < m; i++ {
		memo[i] = make([]int, m)
		for j := 0; j < m; j++ {
			memo[i][j] = -1
		}
	}
	return longestPalindromeSubseqDp(s, 0, len(s)-1, memo)
}

func longestPalindromeSubseqDp(s string, left, right int, memo [][]int) int {
	if left > right {
		return 0
	}
	if left == right {
		return 1
	}
	if memo[left][right] != -1 {
		return memo[left][right]
	}
	if s[left] == s[right] {
		memo[left][right] = longestPalindromeSubseqDp(s, left+1, right-1, memo) + 2
	} else {
		memo[left][right] = max(longestPalindromeSubseqDp(s, left+1, right, memo), longestPalindromeSubseqDp(s, left, right-1, memo))
	}
	return memo[left][right]
}
