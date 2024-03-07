package _1_listNodeAndArray

import "fmt"

/*
	经典数组技巧：差分数组
*/

type Difference struct {
	// 差分数组
	diff []int
}

func NewDifference(nums []int) *Difference {
	length := len(nums)
	diff := make([]int, length)
	diff[0] = nums[0]
	for i := 1; i < length; i++ {
		diff[i] = nums[i] - nums[i-1]
	}
	return &Difference{diff: diff}
}

func (d *Difference) Increment(i, j, val int) {
	d.diff[i] += val
	// 当 j+1 >= diff.length 时，说明是对 nums[i] 及以后的整个数组都进行修改，
	// 那么就不需要再给 diff 数组减 val 了。
	if j+1 < len(d.diff) {
		d.diff[j+1] -= val
	}
}

func (d *Difference) Result() []int {
	length := len(d.diff)
	res := make([]int, len(d.diff))
	res[0] = d.diff[0]
	for i := 1; i < length; i++ {
		res[i] = res[i-1] + d.diff[i]
	}
	return res
}

// https://leetcode.cn/problems/corporate-flight-bookings/description/
// 这里有 n 个航班，它们分别从 1 到 n 进行编号。
func corpFlightBookings(bookings [][]int, n int) []int {
	nums := make([]int, n)
	diff := NewDifference(nums)
	for _, book := range bookings {
		i := book[0] - 1
		j := book[1] - 1
		val := book[2]
		diff.Increment(i, j, val)
	}
	return diff.Result()
}

// https://leetcode.cn/problems/car-pooling/
func carPooling(trips [][]int, capacity int) bool {
	diff := NewDifference(make([]int, 10))
	for _, trip := range trips {
		val := trip[0]
		i := trip[1]
		j := trip[2] - 1
		diff.Increment(i, j, val)
	}
	fmt.Println(diff.Result())
	for _, peoples := range diff.Result() {
		if peoples > capacity {
			return false
		}
	}
	return true
}
