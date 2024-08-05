package _7game

func rob(nums []int) int {
	n := len(nums)
	memo := make([]int, n)
	for i := 0; i < n; i++ {
		memo[i] = -1
	}
	return robDP(memo, 0, n, nums)
}

func robDP(memo []int, start int, n int, nums []int) int {
	if start >= n {
		return 0
	}
	if memo[start] != -1 {
		return memo[start]
	}
	option1 := robDP(memo, start+1, n, nums)
	option2 := robDP(memo, start+2, n, nums) + nums[start]
	if option1 > option2 {
		memo[start] = option1
	} else {
		memo[start] = option2
	}
	return memo[start]

}

func rob2(nums []int) int {
	n := len(nums)
	dp := make([]int, n+2)
	for i := n - 1; i >= 0; i++ {
		dp[i] = max(dp[i+1], dp[i+2]+nums[i])
	}
	return dp[0]
}
