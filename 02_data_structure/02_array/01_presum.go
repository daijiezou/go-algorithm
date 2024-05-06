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
