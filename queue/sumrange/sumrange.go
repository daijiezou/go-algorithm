package sumrange

type NumArray struct {
	preSum []int
}

func Constructor(nums []int) NumArray {
	length := len(nums)
	preSum := make([]int, length+1)
	preSum[0] = 0
	for i := 1; i < length+1; i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
	}
	return NumArray{preSum: preSum}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.preSum[right+1] - this.preSum[left]
}
