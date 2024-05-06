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
