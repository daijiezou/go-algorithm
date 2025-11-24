package _4_datastruct

import (
	"math"
)

// https://leetcode.cn/problems/minimum-sum-of-mountain-triplets-ii/
func minimumSum(nums []int) int {
	n := len(nums)
	mn := nums[0]
	res := math.MaxInt
	sufMin := make([]int, n)
	sufMin[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		sufMin[i] = min(sufMin[i+1], nums[i])
	}
	for i := 1; i < n-1; i++ {
		if mn < nums[i] && sufMin[i+1] < nums[i] {
			res = min(res, sufMin[i+1]+mn+nums[i])
		}
		mn = min(mn, nums[i])
	}
	if res == math.MaxInt {
		return -1
	}
	return res
}

/*
特殊三元组 定义为满足以下条件的下标三元组 (i, j, k)：

0 <= i < j < k < n，其中 n = nums.length
nums[i] == nums[j] * 2
nums[k] == nums[j] * 2
返回数组中 特殊三元组 的总数。

由于答案可能非常大，请返回结果对 109 + 7 取余数后的值。
*/
const mod = 1e9 + 7

func specialTriplets(nums []int) int {
	posMap := make(map[int]int)
	for _, num := range nums {
		posMap[num]++
	}

	curMap := make(map[int]int)
	total := 0
	for _, num := range nums {
		posMap[num]--
		// 左边x/2的个数 * 右边x*2的个数
		leftCnt := curMap[num*2]
		rightCnt := posMap[num*2]
		total = (total + leftCnt*rightCnt) % mod
		curMap[num]++
	}
	return total
}
