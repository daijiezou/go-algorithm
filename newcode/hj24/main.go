package main

import "fmt"

func main() {
	var count int
	fmt.Scan(&count)
	list := make([]int, count)
	//scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < count; i++ {
		fmt.Scan(&list[i])
	}
	for i := 0; i < count; i++ {

	}
	fmt.Println(count - dp(0, count-1, list))
}

// 表示从list 从下标i,j 能组成最长合唱队长度
func dp(i, j int, list []int) int {
	if i == j {
		return 1
	}
	if i > j {
		return 0
	}
	return max(dp(i+1, j, list), dp(i, j-1, list))
}

func lengthOfLIS(nums []int) int {
	// 定义dp[i]为以i结尾的最长子数组
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	maxLen := 0
	for i := 0; i < len(dp); i++ {
		if dp[i] > maxLen {
			maxLen = dp[i]
		}
	}
	return maxLen
}
