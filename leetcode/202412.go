package leetcode

import (
	"container/heap"
	"slices"
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
	n := len(nums)
	const mod = 1_000_000_007
	numMax := 0
	for i := 0; i < len(nums); i++ {
		heap.Push(&newNums, pair2{
			first:  nums[i],
			second: i,
		})
		numMax = max(numMax, nums[i])
	}
	for ; k > 0 && newNums[0].first < numMax; k-- {
		newNums[0].first *= multiplier
		heap.Fix(&newNums, 0)
	}
	sort.Slice(newNums, func(i, j int) bool {
		return newNums[i].first < newNums[j].first || newNums[i].first == newNums[j].first && newNums[i].second < newNums[j].second
	})
	for i, p := range newNums {
		e := k / n
		if i < k%n {
			e++
		}
		nums[p.second] = p.first % mod * pow(multiplier, e) % mod
	}
	return nums

}

func pow(x, n int) int {
	res := 1
	for n > 0 {
		if n&1 > 0 { // 这个比特位是1，需要乘上
			res = res * x % mod
		}
		x = x * x % mod
		n >>= 1
	}
	return res
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

func minSetSize(arr []int) int {
	n := len(arr)
	maxN := slices.Max(arr)
	letterCnt := make([]int, maxN+1)
	for i := 0; i < n; i++ {
		letterCnt[arr[i]]++
	}
	sort.Slice(letterCnt, func(i, j int) bool {
		return letterCnt[i] > letterCnt[j]
	})
	numTotal := 0
	res := 0
	for i := range letterCnt {
		numTotal += letterCnt[i]
		res++
		if numTotal > n/2 {
			return res
		}
	}

	freq := map[int]int{}
	for _, x := range arr {
		freq[x]++
	}

	return 1
}

func closestRoom(rooms [][]int, queries [][]int) []int {
	ans := make([]int, len(queries))
	n := len(rooms)
	// 按照房子面积从小到大排序
	sort.Slice(rooms, func(i, j int) bool {
		return rooms[i][1] < rooms[j][1] || (rooms[i][1] == rooms[j][1] && rooms[i][0] < rooms[j][0])
	})
	for i := 0; i < len(queries); i++ {
		perferredId := queries[i][0]
		leftBound := RoomSizeLeftBound(rooms, queries[i][1])
		if leftBound >= n {
			ans[i] = -1
		} else {
			minAbs := myAbs(perferredId, rooms[leftBound][0])
			ans[i] = rooms[leftBound][0]
			minId := rooms[leftBound][0]
			for j := leftBound + 1; j < n; j++ {
				newAbs := myAbs(perferredId, rooms[j][0])
				newMinId := rooms[j][0]
				if newAbs <= minAbs {
					if newAbs == minAbs {
						if newMinId < minId {
							minAbs = newAbs
							minId = newMinId
							ans[i] = minId
						}
					} else {
						minAbs = newAbs
						minId = newMinId
						ans[i] = minId

					}
					if minAbs == 0 {
						break
					}

				}
			}
		}
	}
	return ans
}

func RoomSizeLeftBound(rooms [][]int, target int) int {
	left := 0
	right := len(rooms) - 1
	for left <= right {
		mid := (left + right) / 2
		if rooms[mid][1] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

func myAbs(x, y int) int {
	if x > y {
		return x - y
	} else {
		return y - x
	}
}

func minValidStrings(words []string, target string) int {
	targets := []byte(target)

	memo := make([]int, len(targets))
	for i := range targets {
		memo[i] = -1
	}
	res := minValDp(words, targets, 0, memo)

	if res >= 50000 {
		return -1
	}
	return res
}

func minValDp(words []string, target []byte, start int, memo []int) int {
	if start >= len(target) {
		return 0
	}
	if memo[start] != -1 {
		return memo[start]
	}
	cnt := 50000
	for _, word := range words {
		var end int
		for end = start; end < len(target) && end-start < len(word); end++ {
			if word[end-start] != target[end] {
				break
			} else {
				cnt = min(cnt, minValDp(words, target, end+1, memo)+1)
			}
		}
	}
	memo[start] = cnt
	return cnt
}

// https://leetcode.cn/problems/jump-game-ii/
func jumpGame(nums []int) int {
	var step, forest, end int
	for i := 0; i < len(nums)-1; i++ {
		forest = max(forest, nums[i]+i)
		if end == i { // 到达了桥的终点，需要重新架一座桥
			end = forest
			step++
		}
	}
	return step
}

func canJump(nums []int) bool {
	var farthest int
	for i := 0; i < len(nums)-1; i++ {
		farthest = max(farthest, nums[i]+i)
		if farthest <= i {
			return false
		}
	}
	return farthest >= (len(nums) - 1)
}

func stableMountains(height []int, threshold int) []int {
	res := make([]int, 0)
	n := len(height)
	for i := 1; i < n; i++ {
		if height[i-1] > threshold {
			res = append(res, i)
		}
	}
	return res
}
