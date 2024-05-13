package _2_array

type Difference struct {
	diff []int
}

func ConstructorDiffSum(nums []int) Difference {
	diffSum := make([]int, len(nums))
	diffSum[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		diffSum[i] = nums[i] - nums[i-1]
	}
	return Difference{diff: diffSum}
}

/* 给闭区间 [i, j] 增加 val（可以是负数）*/
func (d *Difference) Increment(i, j, val int) {
	d.diff[i] += val

	// 如果j>=length，则说明是对i后面整个数组的值进行修改
	if j < len(d.diff)-1 {
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

// https://leetcode.cn/problems/corporate-flight-bookings/description/
/*
这里有 n 个航班，它们分别从 1 到 n 进行编号。
有一份航班预订表 bookings ，
表中第 i 条预订记录 bookings[i] = [firsti, lasti, seatsi]
意味着在从 firsti 到 lasti （包含 firsti 和 lasti ）的 每个航班 上预订了 seatsi 个座位。
请你返回一个长度为 n 的数组 answer，里面的元素是每个航班预定的座位总数。
*/

func corpFlightBookings(bookings [][]int, n int) []int {
	diffIns := ConstructorDiffSum(make([]int, n))
	for _, booking := range bookings {
		diffIns.Increment(booking[0]-1, booking[1]-1, booking[2])
	}
	return diffIns.Result()
}

/*

车上最初有 capacity 个空座位。车 只能 向一个方向行驶（也就是说，不允许掉头或改变方向）
给定整数 capacity 和一个数组 trips ,
trip[i] = [numPassengersi, fromi, toi] 表示第 i 次旅行有
numPassengersi 乘客，接他们和放他们的位置分别是 fromi 和 toi 。这些位置是从汽车的初始位置向东的公里数。
当且仅当你可以在所有给定的行程中接送所有乘客时，返回 true，否则请返回 false。
*/

func carPooling(trips [][]int, capacity int) bool {
	diffIns := ConstructorDiffSum(make([]int, 1000))
	for _, trip := range trips {
		val := trip[0]
		i := trip[1]
		j := trip[2] - 1
		diffIns.Increment(i, j, val)
	}
	for _, peoples := range diffIns.Result() {
		if peoples > capacity {
			return false
		}
	}
	return true
}
