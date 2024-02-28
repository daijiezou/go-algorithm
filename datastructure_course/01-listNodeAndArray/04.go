package _1_listNodeAndArray

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
