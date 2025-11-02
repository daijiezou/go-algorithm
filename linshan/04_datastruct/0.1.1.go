package _4_datastruct

import (
	"math"
)

/*
§0.1枚举右，维护左
对于 双变量问题，例如两数之和ai +a；=t，可以枚举右边的aj，转换成 单变量问题，也就是在a；左边查找是否有a= Qg
，这可以用哈希表维护。
我把这个技巧叫做枚举右，维护左。
*/

func twoSum(nums []int, target int) []int {
	maps := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if v, ok := maps[target-nums[i]]; ok {
			return []int{v, i}
		}
		maps[nums[i]] = i
	}
	return []int{-1, -1}
}

func findMaxK(nums []int) int {
	maps := make(map[int]struct{})
	res := -1
	for i := 0; i < len(nums); i++ {
		if _, ok := maps[0-nums[i]]; ok {
			if nums[i] > 0 {
				res = max(res, nums[i])
			} else {
				res = max(res, -nums[i])
			}
		}
		maps[nums[i]] = struct{}{}
	}
	return res
}

func interchangeableRectangles(rectangles [][]int) int64 {
	res := 0
	maps := make(map[float64]int)
	for i := 0; i < len(rectangles); i++ {
		x := float64(rectangles[i][0]) / float64(rectangles[i][1])
		res += maps[x]
		maps[x]++
	}
	return int64(res)
}

func maxProfit(prices []int) int {
	profit := 0
	minVal := prices[0]
	for i := 1; i < len(prices); i++ {
		profit = max(profit, prices[i]-minVal)
		minVal = min(minVal, prices[i])
	}
	return profit
}

func getLargestOutlier(nums []int) int {
	cnts := make(map[int]int)
	total := 0
	for _, x := range nums {
		cnts[x]++
		total += x
	}
	res := math.MinInt
	for _, num := range nums {
		x := total - 2*num
		if cnts[x] > 1 || (cnts[x] > 0 && x != num) {
			res = max(res, x)
		}
	}
	return res
}

func countBadPairs(nums []int) int64 {

	cnts := make(map[int]int)
	n := len(nums)
	total := n * (n - 1) / 2
	for i := 0; i < n; i++ {
		x := nums[i] - i
		total -= cnts[x]
		cnts[x]++
	}
	return int64(total)
}

func maxScoreSightseeingPair(values []int) int {
	res := math.MinInt
	maxV := values[0]
	for i := 1; i < len(values); i++ {
		res = max(res, values[i]-i+maxV)
		maxV = max(maxV, values[i]+i)
	}
	return res
}

func countNicePairs(nums []int) int {
	cnts := make(map[int]int)
	n := len(nums)
	total := 0
	for i := 0; i < n; i++ {
		x := nums[i] - reverseNum(nums[i])
		total += cnts[x]
		cnts[x]++
	}
	mod := int(1e9 + 7)
	return total % mod
}

func reverseNum(num int) int {
	rev := 0
	for x := num; x > 0; x /= 10 {
		rev = rev*10 + x%10
	}
	return rev
}

// 子序列首尾元素的最大乘积
func maximumProduct(nums []int, m int) int64 {
	minVal := math.MaxInt32
	maxVal := math.MinInt32
	res := math.MinInt
	for i := m - 1; i < len(nums); i++ {
		y := nums[i-m+1]
		minVal = min(y, minVal)
		maxVal = max(y, maxVal)
		x := nums[i]
		res = max(x*minVal, x*maxVal, res)
	}
	return int64(res)
}

/*
给你一个下标从 0 开始、长度为 n 的整数数组 nums ，以及整数 indexDifference 和整数 valueDifference 。

你的任务是从范围 [0, n - 1] 内找出  2 个满足下述所有条件的下标 i 和 j ：
*/
func findIndices(nums []int, indexDifference int, valueDifference int) []int {
	minIndex := 0
	maxIndex := 0

	for i := indexDifference; i < len(nums); i++ {
		j := i - indexDifference
		if nums[j] > nums[maxIndex] {
			maxIndex = j
		}
		if nums[j] < nums[minIndex] {
			minIndex = j
		}
		if nums[maxIndex]-nums[i] >= valueDifference {
			return []int{i, maxIndex}
		}
		if nums[i]-nums[minIndex] >= valueDifference {
			return []int{i, minIndex}
		}
	}
	return []int{-1, -1}
}

func MyAbs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
