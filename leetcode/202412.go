package leetcode

import (
	"container/heap"
	"sort"
)

func minMovesToCaptureTheQueen(a int, b int, c int, d int, e int, f int) int {

	// 先判断车和后是否在一条直线
	if a == e {
		if c != a {
			return 1
		} else {
			if !inBetween(b, d, f) {
				return 1
			}
		}

	}
	if b == f {
		if d != f {
			return 1
		} else {
			if !inBetween(a, c, e) {
				return 1
			}
		}

	}

	// 判断皇后和象是否在一条斜线
	if (c - e) == (d - f) {
		if (c - a) != (d - b) {
			return 1
		} else {
			if !inBetween(c, a, e) {
				return 1
			}
		}
	}

	if (c + d) == (e + f) {
		if (a + b) != (e + f) {
			return 1
		} else {
			if !inBetween(c, a, e) {
				return 1
			}
		}
	}
	return 2
}

func inBetween(l, m, r int) bool {
	return min(l, r) < m && m < max(l, r)
}

func numRookCaptures(board [][]byte) int {
	cnt := 0
	rRow := 0
	rCol := 0
loop1:
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 'R' {
				rRow = i
				rCol = j
				break loop1
			}
		}
	}

	for i := rRow + 1; i < 8; i++ {
		if board[i][rCol] == 'p' {
			cnt++
			break
		}
		if board[i][rCol] == 'B' {
			break
		}
	}
	for i := rRow - 1; i >= 0; i-- {
		if board[i][rCol] == 'p' {
			cnt++
			break
		}
		if board[i][rCol] == 'B' {
			break
		}
	}

	for j := rCol - 1; j >= 0; j-- {
		if board[rRow][j] == 'p' {
			cnt++
			break
		}
		if board[rRow][j] == 'B' {
			break
		}
	}
	for j := rCol + 1; j < 8; j++ {
		if board[rRow][j] == 'p' {
			cnt++
			break
		}
		if board[rRow][j] == 'B' {
			break
		}
	}
	return cnt
}

// https://leetcode.cn/problems/determine-color-of-a-chessboard-square/
func squareIsWhite(coordinates string) bool {
	bytes := []byte(coordinates)
	row := bytes[0] - 'a'
	col := bytes[1] - '0'
	if row%2 == 0 {
		if col%2 == 0 {
			return true
		} else {
			return false
		}
	} else {
		if col%2 == 0 {
			return false
		} else {
			return true
		}
	}
}

var memo1 [5000][10]int

// https://leetcode.cn/problems/knight-dialer/
func knightDialer(n int) int {
	next := [][]int{{4, 6}, {6, 8}, {7, 9}, {4, 8}, {0, 3, 9}, {}, {0, 1, 7}, {2, 6}, {1, 3}, {2, 4}}

	const mod = 1_000_000_007
	var dfs func(start int, n int) int
	dfs = func(start int, n int) int {
		if n == 0 {
			return 1
		}
		if memo1[n][start] != 0 {
			return memo1[n][start]
		}
		cnt := 0
		for _, i2 := range next[start] {
			cnt += dfs(i2, n-1) % mod
		}
		memo1[n][start] = cnt
		return cnt
	}
	if n == 1 {
		return 10
	}
	ans := 0
	for j := range 10 {
		ans += dfs(j, n-1)
	}
	return ans % mod
}

func semiOrderedPermutation(nums []int) int {
	res := 0
	minIndex := 0
	maxIndex := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] >= nums[maxIndex] {
			maxIndex = i
		}
		if nums[i] <= nums[minIndex] {
			minIndex = i
		}
	}
	res = minIndex + len(nums) - 1 - maxIndex
	if maxIndex < minIndex {
		res -= 1
	}
	return res
}

// https://leetcode.cn/problems/maximum-spending-after-buying-items/
func maxSpending(values [][]int) int64 {
	m := len(values)
	n := len(values[0])
	nums := make([]int, 0, m*n)
	for i := range values {
		nums = append(nums, values[i]...)
	}
	sort.Ints(nums)
	res := 0
	for i := range nums {
		res += nums[i] * (i + 1)
	}
	return int64(res)
}

type pair2 struct {
	first  int
	second int
}

type minHeap []pair2

func (h minHeap) Len() int {
	return len(h)
}
func (h minHeap) Less(i, j int) bool {
	return h[i].first < h[j].first || h[i].first == h[j].first && h[i].second < h[j].second
}
func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(pair2))
}

func (h *minHeap) Pop() interface{} {
	n := len(*h)
	res := (*h)[n-1]
	*h = (*h)[0 : n-1]
	return res
}
func getFinalState(nums []int, k int, multiplier int) []int {
	var newNums minHeap
	heap.Init(&newNums)

	for i := 0; i < len(nums); i++ {
		heap.Push(&newNums, pair2{
			first:  nums[i],
			second: i,
		})
	}
	for i := 0; i < k; i++ {
		numPair := &newNums[0]      // 直接获取堆顶元素
		numPair.first *= multiplier // 修改堆顶元素的值
		nums[numPair.second] = numPair.first
		heap.Fix(&newNums, 0)
	}
	return nums
}

func getFinalState2(nums []int, k int, multiplier int) []int {
	for i := 0; i < k; i++ {
		minIndex := 0
		for j := 0; j < len(nums); j++ {
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}
		nums[minIndex] *= multiplier
	}
	return nums
}
