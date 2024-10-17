package _2_binary_search

// https://leetcode.cn/problems/kth-smallest-number-in-multiplication-table/description/
func findKthNumber(m int, n int, k int) int {
	left := 1
	right := m * n
	for left <= right {
		mid := left + (right-left)/2
		if checkK(m, n, k, mid) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

// 计算出在整个乘法口诀表中比小于等于x的数
func checkK(m int, n int, k int, x int) bool {
	cnt := 0
	for i := 1; i <= m; i++ {
		cnt += min(x/i, n)
	}
	return cnt >= k
}
