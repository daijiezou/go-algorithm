package _1_base

import (
	"sort"
)

// https://leetcode.cn/problems/coin-change/description/
// 凑零钱
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	// 将备忘录初始化为 -666，代表还未被计算
	for i := range dp {
		dp[i] = -666
	}

	// base case
	// amount=0时，需要0枚硬币
	dp[0] = 0
	for i := 1; i < len(dp); i++ {
		for coin := range coins {
			// amount<0时，无解
			if i-coin < 0 {
				continue
			}
			dp[i] = min(dp[i], dp[i-coin]+1)
		}
	}
	if dp[amount] == -666 {
		return -1
	}
	return dp[amount]
}

// 最长递增子数组
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

func lengthOfLIS2(nums []int) int {
	top := make([]int, len(nums))
	piles := 0
	for i := 0; i < len(nums); i++ {
		// 要处理的扑克票
		poker := nums[i]
		left := 0
		right := piles
		for left < right {
			mid := left + (right-left)/2
			if poker == top[mid] {
				right = mid
			} else if poker > top[mid] {
				left = mid + 1
			} else if poker < top[mid] {
				right = mid
			}
		}
		if left < piles {
			top[left] = poker
		} else {
			piles++
		}
	}
	return piles
}

// https://leetcode.cn/problems/russian-doll-envelopes/
func maxEnvelopes(envelopes [][]int) int {
	// 先按照weight升序
	// 在按照height倒序
	// 再对height使用lengthOfLIS算法
	sort.Slice(envelopes, func(i, j int) bool {
		return envelopes[i][0] < envelopes[j][0] ||
			(envelopes[i][0] == envelopes[j][0] &&
				envelopes[i][1] > envelopes[j][1])
	})
	n := len(envelopes)
	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1
	}
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		poker := envelopes[i][1]
		left := sort.SearchInts(nums, poker)
		if left < len(nums) {
			nums[left] = poker
		} else {
			nums = append(nums, poker)
		}
		//for j := 0; j < i; j++ {
		//	if envelopes[j][1] < envelopes[i][1] {
		//		minFallingPathSumDp[i] = max(minFallingPathSumDp[i], minFallingPathSumDp[j]+1)
		//	}
		//}
	}
	//maxRes := 1
	//for i := 0; i < n; i++ {
	//	if minFallingPathSumDp[i] > maxRes {
	//		maxRes = minFallingPathSumDp[i]
	//	}
	//}
	return len(nums)
}
