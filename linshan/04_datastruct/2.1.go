package _4_datastruct

func carPooling(trips [][]int, capacity int) bool {
	diff := make([]int, 1001)
	for i := 0; i < len(trips); i++ {
		diff[trips[i][1]] += trips[i][0]
		diff[trips[i][2]] -= trips[i][0]
	}
	sum := 0
	for _, x := range diff {
		sum += x
		if sum > capacity {
			return false
		}
	}
	return true
}

// 2848
/*
给你一个下标从 0 开始的二维整数数组 nums 表示汽车停放在数轴上的坐标。
对于任意下标 i，nums[i] = [starti, endi] ，其中 starti 是第 i 辆车的起点，endi 是第 i 辆车的终点。
返回数轴上被车 任意部分 覆盖的整数点的数目。
*/
func numberOfPoints(nums [][]int) int {
	diff := make([]int, 102)
	for _, t := range nums {
		diff[t[0]] += 1
		diff[t[1]+1] -= 1
	}
	sum := 0
	res := 0
	for _, x := range diff {
		sum += x
		if sum > 0 {
			res++
		}
	}
	return res
}

/*
给你一个二维整数数组 ranges 和两个整数 left 和 right 。每个 ranges[i] = [starti, endi] 表示一个从 starti 到 endi 的 闭区间 。

如果闭区间 [left, right] 内每个整数都被 ranges 中 至少一个 区间覆盖，那么请你返回 true ，否则返回 false 。

已知区间 ranges[i] = [starti, endi] ，如果整数 x 满足 starti <= x <= endi ，那么我们称整数x 被覆盖了。
*/
func isCovered(ranges [][]int, left int, right int) bool {
	diff := make([]int, 52)
	for _, t := range ranges {
		diff[t[0]] += 1
		diff[t[1]+1] -= 1
	}
	sum := 0
	for i, x := range diff {
		sum += x
		if i >= left && i <= right && sum == 0 {
			return false
		}
		if i > right {
			return true
		}
	}
	return true
}

// https://leetcode.cn/problems/maximum-population-year/description/
func maximumPopulation(logs [][]int) int {
	diff := make([]int, 101)
	for _, t := range logs {
		diff[t[0]-1950] += 1
		diff[t[1]-1950] -= 1
	}
	sum := 0
	curMax := 0
	res := 0
	for i, x := range diff {
		sum += x
		if sum > curMax {
			curMax = sum
			res = i + 1950
		}
	}
	return res
}

/*
这里有 n 个航班，它们分别从 1 到 n 进行编号。

有一份航班预订表 bookings ，表中第 i 条预订记录 bookings[i] = [firsti, lasti, seatsi]
意味着在从 firsti 到 lasti （包含 firsti 和 lasti ）的 每个航班 上预订了 seatsi 个座位。
请你返回一个长度为 n 的数组 answer，里面的元素是每个航班预定的座位总数。
*/
func corpFlightBookings(bookings [][]int, n int) []int {
	diff := make([]int, n)
	for _, t := range bookings {
		diff[t[0]-1] += t[2]
		if t[1] < n {
			diff[t[1]-1] -= t[2]
		}
	}
	res := make([]int, n)
	sum := 0
	for i, x := range diff {
		sum += x
		res[i] = sum
	}
	return res
}
