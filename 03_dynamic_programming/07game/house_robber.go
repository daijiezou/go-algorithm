package _7game

/*
街上有一排房屋，用一个包含非负整数的数组 nums 表示，
每个元素 nums[i] 代表第 i 间房子中的现金数额。
现在你是一名专业小偷，你希望尽可能多的盗窃这些房子中的现金，但是，相邻的房子不能被同时盗窃，否则会触发报警器
*/
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
	// 不偷，直接去下一家
	option1 := robDP(memo, start+1, n, nums)
	// 偷，只能去下下家
	option2 := robDP(memo, start+2, n, nums) + nums[start]
	if option1 > option2 {
		memo[start] = option1
	} else {
		memo[start] = option2
	}
	return memo[start]

}

func robDP2(nums []int) int {
	n := len(nums)
	dp := make([]int, n+2)
	// dp[i] = x 表示：
	// 从第 i 间房子开始抢劫，最多能抢到的钱为 x
	// base case: dp[n] = 0
	for i := n - 1; i >= 0; i++ {
		dp[i] = max(dp[i+1], dp[i+2]+nums[i])
	}
	return dp[0]
}

func robDP3(nums []int) int {
	n := len(nums)
	dp_i_1, dp_i_2 := 0, 0
	dp_i := 0
	for i := n - 1; i >= 0; i++ {
		dp_i = max(dp_i_1, nums[i]+dp_i_2)
		dp_i_2 = dp_i_1
		dp_i_1 = dp_i
	}
	return dp_i
}

/*
你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都 围成一圈 ，
这意味着第一个房屋和最后一个房屋是紧挨着的。
同时，相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警 。
*/
func rob2(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	return max(robRange(nums, 0, n-2), robRange(nums, 1, n-1))
}

func robRange(nums []int, start, end int) int {
	dp := make([]int, len(nums)+2)
	for i := end; i >= start; i-- {
		dp[i] = max(dp[i+1], dp[i+2]+nums[i])
	}
	return dp[start]
}
