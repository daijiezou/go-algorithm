package _2sub_seq

import "math"

// https://leetcode.cn/problems/maximum-subarray/submissions/548374430/
// 最大子数组的和
func maxSubArray(nums []int) int {
	// 表示以i为结尾的最大子数组和
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	res := dp[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		res = max(res, dp[i])
	}
	return res
}

func maxSubArray2(nums []int) int {
	// 表示以i为结尾的最大子数组和
	dp_0 := nums[0]
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		dp_0 = max(dp_0+nums[i], nums[i])
		res = max(res, dp_0)
	}
	return res
}

// https://leetcode.cn/problems/maximum-subarray-sum-with-one-deletion/
// 删除一个元素能得到的最大子数组的和
func maximumSum(arr []int) int {
	dp := make([][]int, len(arr))
	for i := 0; i < len(arr); i++ {
		dp[i] = make([]int, 2)
		dp[i][0] = math.MinInt32
		dp[i][1] = math.MinInt32
	}
	dp[0][0] = arr[0]
	dp[0][1] = 0
	res := arr[0]
	for i := 1; i < len(arr); i++ {
		// 不执行删除
		dp[i][0] = max(dp[i-1][0]+arr[i], arr[i])
		// 删除一次
		dp[i][1] = max(dp[i-1][0], dp[i-1][1]+arr[i])
		res = max(res, dp[i][1], dp[i][0])
	}
	return res
}

func longestCommonSubsequence(s1 string, s2 string) int {
	memo := make([][]int, len(s1))
	for i := 0; i < len(s1); i++ {
		memo[i] = make([]int, len(s2))
		for j := 0; j < len(s2); j++ {
			memo[i][j] = -1
		}
	}
	return longestCommonSubsequenceDP(s1, 0, s2, 0, memo)
}

func longestCommonSubsequenceDP(s1 string, i int, s2 string, j int, memo [][]int) int {
	// base case
	if i == len(s1) || j == len(s2) {
		return 0
	}
	if memo[i][j] != -1 {
		return memo[i][j]
	}
	// 当前字符一样的话，两个指针都往后移一位
	if s1[i] == s2[j] {
		memo[i][j] = 1 + longestCommonSubsequenceDP(s1, i+1, s2, j+1, memo)
	} else {
		// 当前字符不一样的话，两个分别往后移一位，取较大值
		memo[i][j] = max(longestCommonSubsequenceDP(s1, i+1, s2, j, memo), longestCommonSubsequenceDP(s1, i, s2, j+1, memo))
	}
	return memo[i][j]
}

// https://leetcode.cn/problems/minimum-ascii-delete-sum-for-two-strings/description/
func minimumDeleteSum(s1 string, s2 string) int {
	memo := make([][]int, len(s1))
	for i := 0; i < len(s1); i++ {
		memo[i] = make([]int, len(s2))
		for j := 0; j < len(s2); j++ {
			memo[i][j] = -1
		}
	}
	return minimumDeleteSumDP(s1, 0, s2, 0, memo)
}

func minimumDeleteSumDP(s1 string, i int, s2 string, j int, memo [][]int) int {
	if i == len(s1) {
		sum := 0
		for ; j < len(s2); j++ {
			sum += int(s2[j])
		}
		return sum
	}
	if j == len(s2) {
		sum := 0
		for ; i < len(s1); i++ {
			sum += int(s1[i])
		}
		return sum
	}
	if memo[i][j] != -1 {
		return memo[i][j]
	}
	if s1[i] == s2[j] {
		memo[i][j] = minimumDeleteSumDP(s1, i+1, s2, j+1, memo)
	} else {
		memo[i][j] = min(minimumDeleteSumDP(s1, i+1, s2, j, memo)+int(s1[i]), minimumDeleteSumDP(s1, i, s2, j+1, memo)+int(s2[j]))
	}
	return memo[i][j]
}

func longestPalindromeSubseq2(s string) int {
	n := len(s)
	// dp 数组全部初始化为 0
	// dp数组的含义是dp[i][j],代表s[i:j]的最长回文子序列
	// 最终需要返回dp[0][n-1]
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	// base case
	for i := 0; i < n; i++ {
		dp[i][i] = 1
	}

	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			// 状态转移方程
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][n-1]
}

// 让字符串变成回文串的最小插入次数
// https://leetcode.cn/problems/minimum-insertion-steps-to-make-a-string-palindrome/
func minInsertions(s string) int {
	// 对字符串 s[i..j]，最少需要进行 dp[i][j] 次插入才能变成回文串。
	n := len(s)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	// base case
	for i := 0; i < n; i++ {
		dp[i][i] = 1
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			// 状态转移方程
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1]
			} else {
				// 可以选择从右边插入或者从左边插入
				// 选择需要次数较少的插入次数
				dp[i][j] = min(dp[i+1][j], dp[i][j-1]) + 1
			}
		}
	}
	return dp[0][n-1]
}
