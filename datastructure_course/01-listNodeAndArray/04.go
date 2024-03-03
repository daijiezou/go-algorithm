package _1_listNodeAndArray

/*
前缀和技巧经典习题
*/

// https://leetcode.cn/problems/range-sum-query-immutable/
type NumArray struct {
	preSum []int
}

func Constructor1(nums []int) NumArray {
	length := len(nums)
	preSum := make([]int, length+1)
	for i := 1; i < length+1; i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
	}
	return NumArray{preSum: preSum}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.preSum[right+1] - this.preSum[left]
}

type NumMatrix struct {
	PreSum [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	m, n := len(matrix), len(matrix[0])
	preSum := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		preSum[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			// 计算每个矩阵 [0, 0, i, j] 的元素和
			preSum[i][j] = preSum[i-1][j] + preSum[i][j-1] + matrix[i-1][j-1] - preSum[i-1][j-1]
		}
	}
	return NumMatrix{PreSum: preSum}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.PreSum[row2+1][col2+1] - this.PreSum[row1][col2+1] - this.PreSum[row2+1][col1] + this.PreSum[row1][col1]
}

// https://leetcode.cn/problems/find-pivot-index/description/
func pivotIndex(nums []int) int {
	length := len(nums)
	preSum := make([]int, length+1)
	for i := 1; i < length+1; i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
	}

	for i := 0; i < length+1; i++ {
		left := preSum[i]
		right := preSum[length] - preSum[i+1]
		if left == right {
			return i
		}
	}
	return -1
}

func productExceptSelf(nums []int) []int {
	length := len(nums)

	leftPreMul := make([]int, length)
	leftPreMul[0] = nums[0]
	// 下标0到i，存放的为nums[0]*num[1]···*num[i]的乘机
	for i := 1; i < length; i++ {
		leftPreMul[i] = leftPreMul[i-1] * nums[i]
	}

	// 下标0到i。存放的位num[length-1]*nums[length-1]
	rightPreMul := make([]int, length)
	rightPreMul[length-1] = nums[length-1]
	for i := length - 2; i >= 0; i-- {
		rightPreMul[i] = rightPreMul[i+1] * nums[i]
	}

	result := make([]int, length)
	result[0] = rightPreMul[1]
	result[length-1] = leftPreMul[length-2]

	for i := 1; i < length-1; i++ {
		result[i] = leftPreMul[i-1] * rightPreMul[i+1]
	}
	return result
}

// https://leetcode.cn/problems/contiguous-array/
// 找到含有相同数量的 0 和 1 的最长连续子数组，并返回该子数组的长度。
func findMaxLength(nums []int) int {
	length := len(nums)
	preSum := make([]int, length+1)

	// 把0，转为-1，
	// 这样只要发现两个前缀和是一样的，则中间这一段的和为0
	for i := 1; i < length+1; i++ {
		if nums[i-1] == 0 {
			preSum[i] = preSum[i-1] + (-1)
		} else {
			preSum[i] = preSum[i-1] + nums[i-1]
		}
	}
	preSumIdx := make(map[int]int)
	var res int
	for i := 0; i < length+1; i++ {
		if _, ok := preSumIdx[preSum[i]]; !ok {
			preSumIdx[preSum[i]] = i
		} else {
			preIndex := preSumIdx[preSum[i]]
			if i-preIndex > res {
				res = i - preIndex
			}
		}
	}
	return res
}

// https://leetcode.cn/problems/continuous-subarray-sum/
func checkSubarraySum(nums []int, k int) bool {
	length := len(nums)
	preSum := make([]int, length+1)
	for i := 1; i < length+1; i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
	}
	preSumIdx := make(map[int]int)
	for i := 0; i < length+1; i++ {
		if _, ok := preSumIdx[preSum[i]%k]; !ok {
			preSumIdx[preSum[i]] = i
		} else {
			preIndex := preSumIdx[preSum[i]%k]
			if i-preIndex > 1 {
				return true
			}
		}
	}
	return false
}

// https://leetcode.cn/problems/subarray-sum-equals-k/

func subarraySumOrigin(nums []int, k int) int {
	length := len(nums)
	preSum := make([]int, length+1)
	res := 0
	countIdx := make(map[int]int)
	countIdx[0] = 1
	for i := 1; i < length+1; i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
		need := preSum[i] - k
		if _, ok := countIdx[need]; ok {
			res += countIdx[need]
		}
		if _, ok := countIdx[preSum[i]]; ok {
			countIdx[preSum[i]]++
		} else {
			countIdx[preSum[i]] = 1
		}
	}
	return res
}

func subarraySum(nums []int, k int) int {
	length := len(nums)
	currentPresum := 0
	res := 0
	countIdx := make(map[int]int)
	countIdx[0] = 1
	for i := 1; i < length+1; i++ {
		currentPresum = currentPresum + nums[i-1]
		need := currentPresum - k
		if _, ok := countIdx[need]; ok {
			res += countIdx[need]
		}
		if _, ok := countIdx[currentPresum]; ok {
			countIdx[currentPresum]++
		} else {
			countIdx[currentPresum] = 1
		}
	}
	return res
}

// https://leetcode.cn/problems/maximum-size-subarray-sum-equals-k/description/
// 给定一个数组 nums 和一个目标值 k，找到和等于 *k*的最长连续子数组长度。
// 如果不存在任意一个符合要求的子数组，则返回 0。
func MaxSubArrayLen(nums []int, k int) int {
	length := len(nums)
	currentPreSum := 0
	res := 0
	preSumIndex := make(map[int]int)
	preSumIndex[0] = 0
	for i := 1; i < length+1; i++ {
		currentPreSum = currentPreSum + nums[i-1]
		if _, ok := preSumIndex[currentPreSum]; !ok {
			preSumIndex[currentPreSum] = i
		}
		need := currentPreSum - k
		if j, ok := preSumIndex[need]; ok {
			tempLength := i - j
			if tempLength > res {
				res = tempLength
			}
		}
	}
	return res
}

// 给定一个整数数组 nums 和一个整数 k ，返回其中元素之和可被 k 整除的（连续、非空） 子数组 的数目。
// 子数组 是数组的 连续 部分。
// https://leetcode.cn/problems/subarray-sums-divisible-by-k/description/
func subarraysDivByK(nums []int, k int) int {
	length := len(nums)
	preSum := make([]int, length+1)
	divIndex := make(map[int]int)
	for i := 1; i < length+1; i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
	}
	var res int
	for i := 0; i < length+1; i++ {
		div := preSum[i] % k
		if div < 0 {
			div += k
		}

		if count, ok := divIndex[div]; ok {
			res += count
			divIndex[div]++
		} else {
			divIndex[div] = 1
		}
	}
	return res
}

// https://leetcode.cn/problems/longest-well-performing-interval/description/
func longestWPI(hours []int) int {
	length := len(hours)
	preSum := make([]int, length+1)
	valIndex := make(map[int]int)
	res := 0
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	for i := 1; i < length+1; i++ {
		if hours[i-1] <= 8 {
			preSum[i] = preSum[i-1] - 1
		} else {
			preSum[i] = preSum[i-1] + 1
		}
		if _, ok := valIndex[preSum[i]]; !ok {
			valIndex[preSum[i]] = i
		}
		if preSum[i] > 0 {
			res = max(res, i)
		} else {
			// preSum[i] 为负，需要寻找一个 j 使得 preSum[i] - preSum[j] > 0
			// 且 j 应该尽可能小，即寻找 preSum[j] == preSum[i] - 1
			if j, ok := valIndex[preSum[i]-1]; ok {
				res = max(res, i-j)
			}
		}
	}
	return res
}
