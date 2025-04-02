package leetcode

func mostPoints(questions [][]int) int64 {
	n := len(questions)
	memo := make([]int, len(questions))
	var dfs func(i int) int
	dfs = func(i int) int {
		if i >= n {
			return 0
		}
		if memo[i] != 0 {
			return memo[i]
		}
		q := questions[i]
		memo[i] = max(q[0]+dfs(i+q[1]+1), dfs(i+1))
		return memo[i]
	}
	return int64(dfs(0))
}

func mostPoints2(questions [][]int) int64 {
	n := len(questions)
	dp := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		q := questions[i]
		j := min(n, i+q[1]+1)
		dp[i] = max(dp[i+1], dp[j]+q[0])
	}
	return int64(dp[0])
}

func maximumTripletValue(nums []int) int64 {
	n := len(nums)
	res := 0
	//preMax := make([]int, n+1)
	sufMax := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		sufMax[i] = max(sufMax[i+1], nums[i])
	}

	preMax := 0
	for i := 0; i < n; i++ {
		res = max(res, (preMax-nums[i])*sufMax[i+1])
		preMax = max(preMax, nums[i])

	}
	return int64(res)
}
