package _2_array

/*
// f(x) 必须是
func f(x int) int {
	// ...
}

func solution(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	// 问自己：自变量 x 的最小值是多少？
	left := ...
	// 问自己：自变量 x 的最大值是多少？
	right := ... + 1

	for left < right {
		mid := left + (right - left) / 2
		if f(mid) == target {
			// 问自己：题目是求左边界还是右边界？
			// ...
		} else if f(mid) < target {
			// 问自己：怎么让 f(x) 大一点？
			// ...
		} else if f(mid) > target {
			// 问自己：怎么让 f(x) 小一点？
			// ...
		}
	}
	return left
}*/

// https://leetcode.cn/problems/koko-eating-bananas/
/*
珂珂喜欢吃香蕉。这里有 n 堆香蕉，第 i 堆中有 piles[i] 根香蕉。警卫已经离开了，将在 h 小时后回来。
珂珂可以决定她吃香蕉的速度 k （单位：根/小时）。每个小时，她将会选择一堆香蕉，从中吃掉 k 根。如果这堆香蕉少于 k 根，
她将吃掉这堆的所有香蕉，然后这一小时内不会再吃更多的香蕉。
珂珂喜欢慢慢吃，但仍然想在警卫回来前吃掉所有的香蕉。
返回她可以在 h 小时内吃掉所有香蕉的最小速度 k（k 为整数）。
*/
func minEatingSpeed(piles []int, h int) int {
	left, right := 1, 10<<9
	for left <= right {
		mid := left + (right-left)>>1
		// 因为是找最慢的速度，所以实际上是找到一个左侧边界
		if f(piles, mid) == h {
			// 搜索左侧边界，则需要收缩右侧边界
			right = mid - 1
		} else if f(piles, mid) > h {
			// 需要让f的返回值小一点，需要增加入参x的值
			left = mid + 1
		} else if f(piles, mid) < h {
			right = mid - 1
		}
	}
	return left
}

// 每个小时吃x个香蕉的，需要的小时数
func f(piles []int, x int) int {
	hours := 0
	for i := 0; i < len(piles); i++ {
		hours += piles[i] / x
		if piles[i]%x > 0 {
			hours++
		}
	}
	return hours
}

func shipWithinDays(weights []int, days int) int {
	left := 0
	right := 1
	for _, w := range weights {
		left = max(left, w)
		right += w
	}
	for left <= right {
		capacity := left + (right-left)>>1
		if fDays(weights, capacity) == days {
			// 搜索左侧边界，则需要收缩右侧边界
			right = capacity - 1
		} else if fDays(weights, capacity) > days {
			// 需要让f的返回值小一点，需要增加入参x的值
			left = capacity + 1
		} else if fDays(weights, capacity) < days {
			right = capacity - 1
		}
	}
	return left
}

func fDays(weights []int, x int) int {
	days := 1
	weight := 0
	for i := 0; i < len(weights); i++ {
		weight += weights[i]
		if weight > x {
			weight = weights[i]
			days++
		}
	}
	return days
}

// https://leetcode.cn/problems/split-array-largest-sum/description/
func splitArray(nums []int, m int) int {
	left := 0
	right := 1
	for _, w := range nums {
		left = max(left, w)
		right += w
	}
	for left <= right {
		capacity := left + (right-left)>>1
		if maxValue(nums, capacity) == m {
			// 搜索左侧边界，则需要收缩右侧边界
			right = capacity - 1
		} else if maxValue(nums, capacity) > m {
			// 需要让f的返回值小一点，需要增加入参x的值
			left = capacity + 1
		} else if maxValue(nums, capacity) < m {
			right = capacity - 1
		}
	}
	return left
}

// 限制最大值为maxVal的子数组的个数
func maxValue(nums []int, maxVal int) (ArrayCnt int) {
	count := 1
	curSum := 0
	for _, v := range nums {
		curSum += v
		if curSum > maxVal {
			count++
			curSum = v
		}
	}
	return count
}
