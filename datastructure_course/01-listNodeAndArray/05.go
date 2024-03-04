package _1_listNodeAndArray

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
	if j+1 < len(d.diff) {
		d.diff[j] -= val
	}
}

func (d *Difference) Result() []int {
	lenth := len(d.diff)
	res := make([]int, len(d.diff))
	res[0] = d.diff[0]
	for i := 1; i < lenth; i++ {
		res[i] = res[i-1] + d.diff[i]
	}
	return res
}

func corpFlightBookings(bookings [][]int, n int) []int {
	diff := NewDifference(make([]int, n))
	for _, book := range bookings {
		diff.Increment(book[0]-1, book[1]+1, book[2])
	}
	return diff.Result()
}
