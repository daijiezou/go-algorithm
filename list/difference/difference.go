package difference

type Difference struct {
	// 差分数组
	diff []int
}

/* 输入一个初始数组，区间操作将在这个数组上进行 */
func NewDifference(nums []int) *Difference {
	diff := make([]int, len(nums))
	diff[0] = 0
	for i := 1; i < len(nums); i++ {
		diff[i] = nums[i] - nums[i-1]
	}
	return &Difference{diff: diff}
}

/* 给闭区间 [i, j] 增加 val（可以是负数）*/
func (d *Difference) Increment(i, j, val int) {
	d.diff[i] += val
	if j+1 <= len(d.diff) {
		d.diff[j+1] -= val
	}
}

/* 返回结果数组 */
func (d *Difference) Result() []int {
	res := make([]int, len(d.diff))
	// 根据差分数组构造结果数组
	res[0] = d.diff[0]
	for i := 1; i < len(d.diff); i++ {
		res[i] = res[i-1] + d.diff[i]
	}
	return res
}

func carPooling(trips [][]int, capacity int) bool {
	nums := make([]int, 1000)
	diff := NewDifference(nums)

	for _, trip := range trips {
		val := trip[0]
		diff.Increment(trip[1], trip[2]-1, val)
	}
	result := diff.Result()
	for i := 0; i < len(result); i++ {
		if result[i] > capacity {
			return false
		}
	}

	return true

}
