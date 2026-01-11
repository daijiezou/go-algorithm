package _4_datastruct

import "math"

// https://leetcode.cn/problems/subarray-sum-equals-k/
/*
给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的子数组的个数 。

子数组是数组中元素的连续非空序列。
*/
func subarraySum2(nums []int, k int) int {
	n := len(nums)
	preSum := make([]int, n+1)
	preSum[0] = 0
	for i := 0; i < n; i++ {
		preSum[i+1] = preSum[i] + nums[i]
	}
	sum := 0
	cnts := make(map[int]int)
	for _, sj := range preSum {
		target := sj - k
		// 计算有多少个到当前的index的和为target
		sum += cnts[target]
		cnts[sj]++
	}
	return sum
}

func subarraySum3(nums []int, k int) int {
	cnts := make(map[int]int)
	sum := 0
	cnts[0] = 0
	res := 0
	for _, sj := range nums {
		sum += sj
		res += cnts[sum-k]
		cnts[sum]++
	}
	return res
}

// https://leetcode.cn/problems/binary-subarrays-with-sum/

/*
给你一个二元数组 nums ，和一个整数 goal ，请你统计并返回有多少个和为 goal 的 非空 子数组。

子数组 是数组的一段连续部分。
*/
func numSubarraysWithSum(nums []int, goal int) int {
	n := len(nums)
	preSum := make([]int, n+1)
	for i := 0; i < n; i++ {
		preSum[i+1] = preSum[i] + nums[i]
	}
	cnts := make(map[int]int)
	res := 0
	for _, v := range preSum {
		target := v - goal
		res += cnts[target]
		cnts[v]++
	}
	return res
}

// 恰好型的滑动窗口
func numSubarraysWithSum2(nums []int, goal int) int {
	goalCnt := 0
	goal1Cnt := 0
	left := 0
	sum := 0
	sum1 := 0
	length := len(nums)
	for right, x := range nums {
		sum += x
		for sum >= goal && left <= right {
			goalCnt += length - right
			sum -= nums[left]
			left++
		}
	}

	left = 0
	for right, x := range nums {
		sum1 += x
		for sum1 >= goal+1 && left <= right {
			goal1Cnt += length - right
			sum1 -= nums[left]
			left++
		}
	}
	return goalCnt - goal1Cnt
}

// 1524.和为奇数的子数组数目
func numOfSubarrays(arr []int) int {
	n := len(arr)
	preSum := make([]int, n+1)
	for i := 0; i < n; i++ {
		preSum[i+1] = preSum[i] + arr[i]
	}
	cnts := make(map[int]int)
	res := 0
	for _, v := range preSum {
		if v%2 == 0 {
			res += cnts[1] % mod
			cnts[0]++
		} else {
			res += cnts[0] % mod
			cnts[1]++
		}

	}
	return res
}

// 974. 和可被 K 整除的子数组
func subarraysDivByK(nums []int, k int) int {
	n := len(nums)
	preSum := make([]int, n+1)
	for i := 0; i < n; i++ {
		preSum[i+1] = preSum[i] + nums[i]
	}
	cnts := make(map[int]int)
	res := 0
	for _, v := range preSum {
		target := v % k
		if target < 0 {
			target += k
		}
		res += cnts[target]
		cnts[target]++
	}
	return res
}

func subarraysDivByK2(nums []int, k int) int {
	cnts := make(map[int]int)
	res := 0
	s := 0
	cnts[0] = 1
	for _, v := range nums {
		s = (s + v%k + k) % k
		res += cnts[s]
		cnts[s]++
	}
	return res
}

// 2588. 统计美丽子数组数目
/*
推广到多个列（多个比特位）：

如果每一列都有偶数个 1，那么所有数的异或和必然等于 0。
如果某一列有奇数个 1，那么所有数的异或和必然不等于 0。
所以美丽子数组等价于：
子数组的异或和等于 0。

*/
func beautifulSubarrays(nums []int) int64 {
	cnts := make(map[int]int)
	s := 0
	res := int64(0)
	cnts[0] = 1 // 如果当前s正好异或和为0，应该被正确的记入
	for _, x := range nums {
		s ^= x
		res += int64(cnts[s]) // 自己异或自己肯定为0
		cnts[s]++
	}
	return res
}

// 525. 连续数组
func findMaxLength(nums []int) int {
	firstIndexMap := make(map[int]int)
	sum := 0
	res := -1
	// 前缀和数组的首项 0 相当于在 -1 下标
	firstIndexMap[0] = -1
	for cutIndex, x := range nums {
		if x == 0 {
			x = -1
		}
		sum += x
		if index, ok := firstIndexMap[sum]; ok {
			res = max(cutIndex-index, res)
		} else {
			firstIndexMap[sum] = cutIndex
		}

	}
	return res
}

// https://leetcode.cn/problems/maximum-good-subarray-sum/
/*
给你一个长度为 n 的数组 nums 和一个 正 整数 k 。
如果 nums 的一个子数组中，第一个元素和最后一个元素 差的绝对值恰好 为 k ，
我们称这个子数组为 好 的。换句话说，如果子数组 nums[i..j] 满足 |nums[i] - nums[j]| == k ，那么它是一个好子数组。
请你返回 nums 中 好 子数组的 最大 和，如果没有好子数组，返回 0 。


*/
func maximumSubarraySum(nums []int, k int) int64 {
	type pair struct {
		minPreSum int64
		index     int
	}
	minPreSumMap := make(map[int]pair)
	res := int64(math.MinInt64)
	preSum := int64(0)

	for i, x := range nums {
		preSum += int64(x)

		// 检查 x+k 是否存在，计算以当前位置结尾的子数组和
		if p, ok := minPreSumMap[x+k]; ok {
			sum := preSum - p.minPreSum
			res = max(sum, res)
		}

		// 检查 x-k 是否存在，计算以当前位置结尾的子数组和
		if p, ok := minPreSumMap[x-k]; ok {
			sum := preSum - p.minPreSum
			res = max(sum, res)
		}

		// 更新当前值 x 对应的最小前缀和
		if p, ok := minPreSumMap[x]; !ok || preSum-int64(x) < p.minPreSum {
			minPreSumMap[x] = pair{
				minPreSum: preSum - int64(x),
				index:     i,
			}
		}
	}

	if res != int64(math.MinInt64) {
		return res
	}
	return 0
}

/*
给你一个整数数组 arr 和一个整数值 target 。

请你在 arr 中找 两个互不重叠的子数组 且它们的和都等于 target 。可能会有多种方案，请你返回满足要求的两个子数组长度和的 最小值 。

请返回满足要求的最小长度和，如果无法找到这样的两个子数组，请返回 -1 。
*/
func minSumOfLengths(arr []int, target int) int {
	n := len(arr)
	//  minLen[i] 的定义是：在 [0, i] 范围内，和为 target 的最短子数组长度
	minLen := make([]int, n)
	for i := range minLen {
		minLen[i] = math.MaxInt32
	}

	preSum := 0
	index := make(map[int]int)
	index[0] = -1
	res := math.MaxInt32

	for i, x := range arr {
		preSum += x

		if leftIdx, ok := index[preSum-target]; ok {
			curLen := i - leftIdx

			if leftIdx >= 0 && minLen[leftIdx] != math.MaxInt32 {
				res = min(res, minLen[leftIdx]+curLen)
			}

			if i > 0 {
				minLen[i] = min(minLen[i-1], curLen)
			} else {
				minLen[i] = curLen
			}
		} else {
			if i > 0 {
				minLen[i] = minLen[i-1]
			}
		}

		index[preSum] = i
	}

	if res == math.MaxInt32 {
		return -1
	}
	return res
}
