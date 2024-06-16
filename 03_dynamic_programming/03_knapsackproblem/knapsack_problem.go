package _3_knapsackproblem

import "fmt"

// wt 每件物品的重量，
func knapsack(W, N int,
	wt []int, // 每件物品的重量
	val []int, // 每件物品的价值
) int {

	dp := make([][]int, N+1)
	for i, _ := range dp {
		dp[i] = make([]int, W+1)
	}
	for i := 1; i <= N; i++ {
		for w := 1; w <= W; w++ {
			if w-wt[i-1] < 0 {
				// 这种情况下只能选择不装入背包
				dp[i][w] = dp[i-1][w]
			} else {
				// 装入或者不装入背包，择优
				dp[i][w] = max(
					dp[i-1][w-wt[i-1]]+val[i-1],
					dp[i-1][w],
				)
			}
		}
	}
	return dp[N][W]
}

// https://leetcode.cn/problems/partition-equal-subset-sum/submissions/539717647/
/*
给你一个 只包含正整数 的 非空 数组 nums 。请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。
*/
func canPartition(nums []int) bool {
	totalSum := 0
	for i := 0; i < len(nums); i++ {
		totalSum += nums[i]
	}
	if totalSum%2 != 0 {
		return false
	}
	target := totalSum / 2
	n := len(nums)
	dp := make([][]bool, n+1)
	for i, _ := range dp {
		dp[i] = make([]bool, target+1)
		//
		dp[i][0] = true
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= target; j++ {
			// note 注意这里是i-1
			if j-nums[i-1] < 0 {
				// 背包容量不足，不能装入第 i 个物品
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j-nums[i-1]] || dp[i-1][j]
			}
		}
	}
	return dp[n][target]
}

/*
有一个背包，最大容量为 amount，有一系列物品 coins，每个物品的重量为 coins[i]，
每个物品的数量无限。请问有多少种方法，能够把背包恰好装满？
*/
// https://leetcode.cn/problems/coin-change-ii/description/
func change(amount int, coins []int) int {
	/*
		dp[i][j] 的定义如下：
		若只使用前 i 个物品（可以重复使用），当背包容量为 j 时，有 dp[i][j] 种方法可以装满背包。
	*/
	n := len(coins)
	dp := make([][]int, n+1)
	for i, _ := range dp {
		dp[i] = make([]int, amount+1)
		// 目标金额为0，有一种方法可以凑
		dp[i][0] = 1
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= amount; j++ {
			if j-coins[i-1] < 0 {
				// 只能不选该面额的金币
				dp[i][j] = dp[i-1][j]
			} else {
				option1 := dp[i-1][j] // 不使用该面额的金币

				// note:区别在于因为每种金额的硬币是无限，所以这里是dp[i][j-coins[i-1]]
				// note：				   而不是0-1背包问题中的dp[i-1][j-coins[i-1]]
				option2 := dp[i][j-coins[i-1]] // 使用该面额的金币
				dp[i][j] = option1 + option2
			}
		}
	}
	return dp[n][amount]
}

func findTargetSumWays(nums []int, target int) int {
	operation := []int{1, -1}
	resust := 0
	var backtack func(nums []int, start int, target int, sum int)
	backtack = func(nums []int, start int, target int, sum int) {
		if start == len(nums) {
			if sum == target {
				resust++
			}
			return
		}
		// 遍历每种可以的选择
		for _, op := range operation {
			// 做选择
			sum += nums[start] * op
			backtack(nums, start+1, target, sum)
			// 撤销选择
			sum -= nums[start] * op
		}
	}
	backtack(nums, 0, target, 0)
	return resust
}

func findTargetSumWays2(nums []int, target int) int {
	memo := make(map[string]int)
	return findTargetSumWaysdp(nums, target, memo, 0)
}

// 使用动态规划算法
func findTargetSumWaysdp(nums []int, target int, memo map[string]int, start int) int {
	key := fmt.Sprintf("%d:%d", start, target)
	if v, ok := memo[key]; ok {
		return v
	}
	if start == len(nums) {
		if target == 0 {
			return 1
		}
		return 0
	}
	result := findTargetSumWaysdp(nums, target-nums[start], memo, start+1) + findTargetSumWaysdp(nums, target+nums[start], memo, start+1)
	// 将结果加入备忘录
	memo[key] = result
	return result
}
