package leetcode

import "math"

func partition(s string) [][]string {
	n := len(s)
	res := make([][]string, 0)
	curent := make([]string, 0)
	var backtrack func(i int)
	backtrack = func(i int) {
		// 到达终点
		if i == n {
			temp := make([]string, len(curent))
			copy(temp, curent)
			res = append(res, temp)
			return
		}
		for j := i; j < n; j++ {
			if isPalindrome(s, i, j) {
				curent = append(curent, s[i:j+1])
				backtrack(j + 1)
				curent = curent[:len(curent)-1]
			}
		}
	}
	backtrack(0)
	return res
}

func isPalindrome(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// https://leetcode.cn/problems/palindrome-partitioning-ii/
func minCut(s string) int {
	n := len(s)
	palMemo := make([][]int, n)
	for i := 0; i < n; i++ {
		palMemo[i] = make([]int, n)
		for j := 0; j < n; j++ {
			// 表示没有计算过
			palMemo[i][j] = -1
		}
	}

	var isPal func(left, right int) bool
	isPal = func(left, right int) bool {
		if left >= right {
			palMemo[left][right] = 1
		}
		p := &palMemo[left][right]
		if *p != -1 {
			return *p == 1
		}
		res := s[left] == s[right] && isPal(left+1, right-1)
		if res {
			*p = 1
		} else {
			*p = 0
		}
		return res
	}
	memo := make([]int, n)
	for i := 0; i < n; i++ {
		// 表示没有计算过
		memo[i] = -1
	}
	// 把 s[:r+1] 切 i 刀，分成 i+1 个子串，每个子串改成回文串的最小总修改次数
	var dfs func(i int) int
	dfs = func(i int) int {
		if isPal(0, i) {
			return 0
		}
		if memo[i] != -1 {
			return memo[i]
		}
		res := math.MaxInt
		for left := 1; left <= i; left++ {
			if isPal(left, i) {
				res = min(res, dfs(left-1)+1)
			}
		}
		memo[i] = res
		return res
	}
	return dfs(n - 1)
}

func palindromePartition(s string, k int) int {
	n := len(s)
	memoChange := make([][]int, n)
	for i := 0; i < n; i++ {
		memoChange[i] = make([]int, n)
		for j := 0; j < n; j++ {
			memoChange[i][j] = -1 // 表示没有计算过
		}
	}
	var minChange func(i, j int) int //表示s[i:j+1]修改为回文串的最小修改次数
	minChange = func(i, j int) int {
		if i >= j {
			return 0
		}
		if memoChange[i][j] != -1 {
			return memoChange[i][j]
		}
		res := minChange(i+1, j-1)
		if s[i] != s[j] {
			res++
		}
		memoChange[i][j] = res
		return res
	}
	memoDfs := make([][]int, k)
	for i := range memoDfs {
		memoDfs[i] = make([]int, n)
		for j := range memoDfs[i] {
			memoDfs[i][j] = -1 // -1 表示没有计算过
		}
	}
	// i表示还需要切i刀
	// j表示剩余字符串的右端点
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i == 0 {
			return minChange(0, j)
		}
		if memoDfs[i][j] != -1 {
			return memoDfs[i][j]
		}
		res := math.MaxInt
		// 由于不能有空串，所以右端点的初始位置必须>=i
		for l := i; l <= j; l++ {
			res = min(res, dfs(i-1, l-1)+minChange(l, j))
		}
		memoDfs[i][j] = res
		return res
	}
	return dfs(k-1, n-1)
}
