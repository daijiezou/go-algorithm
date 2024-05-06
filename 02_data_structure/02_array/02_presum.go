package _2_array

type NumArray struct {
	Presum []int
}

func Constructor(nums []int) NumArray {
	presum := make([]int, len(nums)+1)
	for i := 1; i <= len(nums); i++ {
		presum[i] = presum[i-1] + nums[i-1]
	}
	return NumArray{Presum: presum}
}

// 查询闭区间 [left, right] 的累加和
func (this *NumArray) SumRange(left int, right int) int {
	return this.Presum[right+1] - this.Presum[left]
}

// https://leetcode.cn/problems/find-pivot-index/
// 寻找数组的中心下标
func pivotIndex(nums []int) int {
	length := len(nums)
	presum := make([]int, len(nums)+1)
	for i := 1; i <= length; i++ {
		presum[i] = presum[i-1] + nums[i-1]
	}
	for i := 1; i <= len(nums); i++ {
		leftSum := presum[i-1]
		rightSum := presum[length] - presum[i]
		if leftSum == rightSum {
			return i - 1
		}
	}
	return -1
}

// 除自身以外数组的乘积
func productExceptSelf(nums []int) []int {
	length := len(nums)

	// 前缀积 [0...i]之间的乘积
	prefix := make([]int, length)
	prefix[0] = nums[0]
	for i := 1; i < length; i++ {
		prefix[i] = prefix[i-1] * nums[i]
	}

	// 后缀积 [i...length]之间的乘积
	suffix := make([]int, length)
	suffix[length-1] = nums[length-1]
	for i := length - 2; i >= 0; i-- {
		suffix[i] = suffix[i+1] * nums[i]
	}

	res := make([]int, length)
	res[0] = suffix[1]
	res[length-1] = prefix[length-2]
	for i := 1; i < length-1; i++ {
		res[i] = prefix[i-1] * suffix[i+1]
	}
	return res
}
