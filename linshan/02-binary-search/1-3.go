package _2_binary_search

/*
二分答案：求最大
*/

/*

给你一个 下标从 0 开始 的整数数组 candies 。数组中的每个元素表示大小为 candies[i] 的一堆糖果。
你可以将每堆糖果分成任意数量的 子堆 ，但 无法 再将两堆合并到一起。
另给你一个整数 k 。你需要将这些糖果分配给 k 个小孩，
使每个小孩分到 相同 数量的糖果。每个小孩可以拿走 至多一堆 糖果，有些糖果可能会不被分配。
返回每个小孩可以拿走的 最大糖果数目 。
*/

// https://leetcode.cn/problems/maximum-candies-allocated-to-k-children/description/
func maximumCandies(candies []int, k int64) int {
	sum := 0
	for i := 0; i < len(candies); i++ {
		sum += candies[i]
	}
	if int64(sum) < k {
		return 0
	}
	right := int64(sum) / k
	left := int64(1)
	for left <= right {
		mid := left + (right-left)/2
		if canSplit(candies, k, int(mid)) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return int(right)
}

func canSplit(candies []int, k int64, perCnt int) bool {
	sum := int64(0)
	for i := 0; i < len(candies); i++ {
		sum += int64(candies[i] / perCnt)
	}
	return sum >= k
}

// https://leetcode.cn/problems/h-index-ii/
func hIndex(citations []int) int {
	n := len(citations)
	right := citations[n-1]
	left := 1
	for left <= right {
		mid := left + (right-left)/2
		if H(citations, mid, n) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return right
}

func H(citations []int, h int, n int) bool {
	if h > n {
		return false
	}
	start := n - h
	if citations[start] >= h {
		return true
	}
	return false
}
