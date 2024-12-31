package leetcode

import (
	"container/heap"
	"math/rand"
	"slices"
	"sort"
	"strings"
	"time"
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

func minAnagramLength(s string) int {
	n := len(s)
loop1:
	for i := 1; i <= n/2; i++ {
		if n%i != 0 {
			continue
		}
		encodeStr := Encode(s[:i])
		for j := i; j < n; j += i {
			encode := Encode(s[j : j+i])
			if encode != encodeStr {
				continue loop1
			}
		}
		return i
	}
	return n
}

func Encode(s string) [26]int {
	var cnts [26]int
	for _, v := range s {
		cnts[v-'a']++
	}
	//var res strings.Builder
	//for i := 0; i < 26; i++ {
	//	res.WriteByte(byte(cnts[i]))
	//}
	return cnts
}

func sortTheStudents(score [][]int, k int) [][]int {
	sort.Slice(score, func(i, j int) bool { return score[i][k] > score[j][k] })
	return score
}

func getKth2(lo int, hi int, k int) int {
	type pair struct {
		origin int
		weight int
	}
	nums := []pair{}
	for i := lo; i <= hi; i++ {
		weight := getWeight(i)
		nums = append(nums, pair{
			origin: i,
			weight: weight,
		})
	}
	sort.Slice(nums, func(i, j int) bool {
		if nums[i].weight == nums[j].weight {
			return nums[i].origin < nums[j].origin
		}
		return nums[i].weight < nums[j].weight
	})
	return nums[k-1].origin
}

func getWeight(num int) int {
	weight := 0
	for num != 1 {
		if num%2 == 0 {
			num = num / 2
		} else {
			num = 3*num + 1
		}
		weight++
	}
	return weight
}

// https://leetcode.cn/problems/maximum-number-of-eaten-apples/
func eatenApples(apples []int, days []int) int {
	res := 0
	n := len(apples)
	var h appleHp
	heap.Init(&h)
	for i := 0; i < n; i++ {
		for h.Len() > 0 && h[0].rottenDay == i {
			heap.Pop(&h)
		}
		if apples[i] > 0 {
			heap.Push(&h, applePair{
				rottenDay: i + days[i],
				num:       apples[i],
			})
		}
		if h.Len() > 0 {
			res++
			h[0].num--
			if h[0].num == 0 {
				heap.Pop(&h)
			}
		}
	}
	i := len(apples)
	for h.Len() > 0 {
		for h.Len() > 0 && h[0].rottenDay <= i {
			heap.Pop(&h)
		}
		if h.Len() == 0 {
			return res
		}
		p := heap.Pop(&h).(applePair)
		k := min(p.num, p.rottenDay-i)
		res += k
		i += k
	}
	return res
}

type applePair struct {
	rottenDay int
	num       int
}

type appleHp []applePair

func (a appleHp) Len() int {
	return len(a)
}

func (a appleHp) Less(i, j int) bool {
	return a[i].rottenDay < a[j].rottenDay
}

func (a appleHp) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a *appleHp) Push(x any) {
	*a = append(*a, x.(applePair))
}

func (a *appleHp) Pop() any {
	apple := *a
	v := apple[len(apple)-1]
	*a = apple[:len(apple)-1]
	return v
}

// https://leetcode.cn/problems/minimum-cost-for-cutting-cake-i/
func minimumCost(m int, n int, horizontalCut []int, verticalCut []int) int {
	slices.Sort(horizontalCut)
	slices.Sort(verticalCut)
	i := m - 2
	j := n - 2
	res := 0
	cnth := 1
	cntw := 1
	for i >= 0 && j >= 0 {
		if horizontalCut[i] > verticalCut[j] {
			res += horizontalCut[i] * cntw
			cnth++
			i--
		} else {
			res += verticalCut[j] * cnth
			cntw++
			j--
		}
	}
	for i >= 0 {
		res += horizontalCut[i] * cntw
		i--
	}
	for j >= 0 {
		res += verticalCut[j] * cnth
		j--
	}
	return res
}

func isSubstringPresent(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		str := string([]byte{s[i+1], s[i]})
		if strings.Contains(s, str) {
			return true
		}
	}
	return false
}

// https://leetcode.cn/problems/find-occurrences-of-an-element-in-an-array/
func occurrencesOfElement(nums []int, queries []int, x int) []int {
	xIndex := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i] == x {
			xIndex = append(xIndex, i)
		}
	}
	res := make([]int, 0, len(queries))
	for _, query := range queries {
		if query-1 >= len(xIndex) {
			res = append(res, -1)
		} else {
			res = append(res, xIndex[query-1])
		}
	}
	return res
}

func rankTeams(votes []string) string {
	voteRank := make(map[uint8][]int)
	m := len(votes[0])
	voteRankKey := make([]uint8, m)
	for i := 0; i < m; i++ {
		voteRank[votes[0][i]] = make([]int, m)
		voteRankKey[i] = votes[0][i]
	}
	for i := 0; i < len(votes); i++ {
		for j := 0; j < m; j++ {
			voteRank[votes[i][j]][j]++
		}
	}
	sort.Slice(voteRankKey, func(i, j int) bool {
		iKey := voteRankKey[i]
		jKey := voteRankKey[j]
		for index := 0; index < m; index++ {
			if voteRank[iKey][index] != voteRank[jKey][index] {
				return voteRank[iKey][index] > voteRank[jKey][index]
			}
		}
		return voteRankKey[i] < voteRankKey[j]
	})
	res := strings.Builder{}
	for i := 0; i < m; i++ {
		res.WriteByte(voteRankKey[i])
	}
	return res.String()
}

// https://leetcode.cn/problems/linked-list-in-binary-tree/
func isSubPath(head *ListNode, root *TreeNode) bool {
	// 表示已经将链表遍历完
	if head == nil {
		return true
	}
	// 链表没有遍历完，但是二叉树已经到底
	if root == nil {
		return false
	}

	if root.Val == head.Val {
		if check(head, root) {
			return true
		}
	}
	return isSubPath(head, root.Left) || isSubPath(head, root.Right)
}

// 检查是否能够将链表嵌入二叉树
func check(head *ListNode, root *TreeNode) bool {
	if head == nil {
		return true
	}
	if root == nil {
		return false
	}

	if head.Val == root.Val {
		// 在子树上嵌入子链表
		return check(head.Next, root.Left) || check(head.Next, root.Right)
	}

	return false
}

func findMid(nums []int) float64 {
	n := len(nums)
	if n%2 == 1 {
		return float64(findk(nums, n/2))
	}
	// 偶数个元素，返回中间两个值的平均值
	left := findk(nums, n/2-1)
	right := findk(nums, n/2)
	return float64(left+right) / 2.0
}

func findk(nums []int, k int) int {
	keyValue := nums[0]
	equalNums := make([]int, 0)
	leftNums := make([]int, 0)
	rightNums := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i] == keyValue {
			equalNums = append(equalNums, nums[i])
		} else if nums[i] < keyValue {
			leftNums = append(leftNums, nums[i])
		} else {
			rightNums = append(rightNums, nums[i])
		}
	}
	if k < len(leftNums) {
		return findk(leftNums, k)
	} else if k <= len(leftNums)+len(equalNums) {
		return keyValue
	} else {
		return findk(rightNums, k-len(leftNums)-len(equalNums))
	}
}

func findMid2(nums []int) float64 {
	n := len(nums)
	if n%2 == 1 {
		return float64(findk2(nums, n/2, 0, len(nums)-1))
	}
	// 偶数个元素，返回中间两个值的平均值
	left := findk2(nums, n/2-1, 0, len(nums)-1)
	right := findk2(nums, n/2, 0, len(nums)-1)
	return float64(left+right) / 2.0
}

func findk2(nums []int, k int, left, right int) int {
	if left >= right {
		return nums[left]
	}
	keyValue := nums[right]
	equalNums := 0
	leftNums := 0
	rightNums := 0
	part := left
	for i := left; i < right; i++ {
		if nums[i] == keyValue {
			equalNums++
		} else if nums[i] < keyValue {
			nums[part], nums[i] = nums[i], nums[part]
			part++
			leftNums++
		} else {
			rightNums++
		}
	}
	nums[part], nums[right] = nums[right], nums[part]
	if k < leftNums {
		return findk2(nums, k, left, part-1)
	} else if k < leftNums+equalNums {
		return keyValue
	} else {
		return findk2(nums, k-leftNums-equalNums, part+1, right)
	}
}

func findMedian(nums []int) float64 {
	n := len(nums)
	if n == 0 {
		panic("数组不能为空")
	}

	if n%2 == 1 {
		// 奇数个元素，中位数是第 n/2 小的元素
		return float64(quickSelect(nums, n/2))
	} else {
		// 偶数个元素，中位数是第 n/2-1 和第 n/2 小的元素的平均值
		left := quickSelect(nums, n/2-1)
		right := quickSelect(nums, n/2)
		return float64(left+right) / 2.0
	}
}

func quickSelect(nums []int, k int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		pivotIndex := partition(nums, left, right)
		if pivotIndex == k {
			return nums[pivotIndex]
		} else if pivotIndex < k {
			left = pivotIndex + 1
		} else {
			right = pivotIndex - 1
		}
	}
	panic("无法找到中位数")
}

func partition(nums []int, left, right int) int {
	// 随机选择基准值，避免退化
	rand.Seed(time.Now().UnixNano())
	pivotIndex := left + rand.Intn(right-left+1)
	nums[pivotIndex], nums[right] = nums[right], nums[pivotIndex]

	pivot := nums[right]
	i := left
	for j := left; j < right; j++ {
		if nums[j] < pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	nums[i], nums[right] = nums[right], nums[i]
	return i
}

// https://leetcode.cn/problems/minimum-cost-for-cutting-cake-ii/description/
func minimumCost2(m int, n int, horizontalCut []int, verticalCut []int) int64 {
	res := 0
	slices.Sort(horizontalCut)
	slices.Sort(verticalCut)
	i := m - 2
	j := n - 2
	hcnt := 1
	vcnt := 1

	for i >= 0 && j >= 0 {
		if horizontalCut[i] > verticalCut[j] {
			res += horizontalCut[i] * vcnt
			i--
			hcnt++
		} else {
			res += verticalCut[j] * hcnt
			j--
			vcnt++
		}
	}
	for i >= 0 {
		res += horizontalCut[i] * vcnt
		i--
		hcnt++
	}
	for j >= 0 {
		res += verticalCut[j] * hcnt
		j--
		vcnt++
	}
	return int64(res)
}
