package _024

import (
	"math"
	"math/bits"
	"slices"
	"sort"
	"strconv"
	"strings"
)

// https://leetcode.cn/problems/check-if-move-is-legal/
func checkMove(board [][]byte, rMove int, cMove int, color byte) bool {
	// 从y轴正方向开始遍历
	//上、右上、右、右下、下、左下、左、左上
	dxs := []int{0, 1, 1, 1, 0, -1, -1, -1}
	dys := []int{-1, -1, 0, 1, 1, 1, 0, -1}
	for i := 0; i < 8; i++ {
		// 检查8个方向
		if checkGood(board, rMove, cMove, color, dxs[i], dys[i]) {
			return true
		}
	}
	return false
}

func checkGood(board [][]byte, rMove int, cMove int, color byte, dx, dy int) bool {
	x := rMove + dx
	y := cMove + dy
	step := 1
	for x >= 0 && x < 8 && y >= 0 && y < 8 {
		//第一步必须是其他颜色
		if step == 1 {
			if board[x][y] == color || board[x][y] == '.' {
				return false
			}
		} else {
			//中间不能有空棋盘
			if board[x][y] == '.' {
				return false
			}
			// 遍历到了终点
			if board[x][y] == color {
				return true
			}
		}
		x += dx
		y += dy
		step++
	}
	return false
}

// https://leetcode.cn/problems/find-pivot-index/description/?envType=daily-question&envId=2024-07-08
func pivotIndex(nums []int) int {
	length := len(nums)
	presum := make([]int, length+1)

	for i := 1; i < length+1; i++ {
		presum[i] = presum[i-1] + nums[i-1]
	}
	for i := 1; i < length+1; i++ {
		// 计算 nums[i-1] 左侧和右侧的元素和
		left := presum[i-1]
		right := presum[length] - presum[i]
		if left == right {
			return i
		}
	}
	return -1
}

// https://leetcode.cn/problems/count-the-number-of-incremovable-subarrays-i/?envType=daily-question&envId=2024-07-10
// 找到有几个递增子数组
func incremovableSubarrayCount(nums []int) int {
	current := []int{}
	count := 0
	currentIndex := make([]bool, len(nums))
	incremovableSubarrayCountBacktack(nums, 0, current, currentIndex, &count)
	return count + 1
}

func incremovableSubarrayCountBacktack(nums []int, start int, zijihe []int, index []bool, count *int) {
	for i := start; i < len(nums); i++ {
		zijihe = append(zijihe, nums[i])
		index[i] = true
		// 判断是否为递增子数组
		if heckIncremovable(index, zijihe) {
			*count++
		}
		incremovableSubarrayCountBacktack(nums, i+1, zijihe, index, count)
		zijihe = zijihe[:len(zijihe)-1]
		index[i] = false
	}
}

func heckIncremovable(index []bool, nums []int) bool {
	if len(nums) == len(index) {
		return false
	}
	leftIndex := make([]int, 0)
	for i := 0; i < len(index); i++ {
		if !index[i] {
			leftIndex = append(leftIndex, i)
		}
	}

	// 判断index是否连续
	for i := 0; i < len(leftIndex)-1; i++ {
		if leftIndex[i]+1 != leftIndex[i+1] {
			return false
		}
	}
	if len(nums) == 1 {
		return true
	}
	length := len(nums)
	for i := 0; i < length-1; i++ {
		if nums[i] >= nums[i+1] {
			return false
		}
	}

	return true
}

func numberGame(nums []int) []int {
	sort.Ints(nums)
	for i := 0; i < len(nums); i += 2 {
		nums[i], nums[i+1] = nums[i+1], nums[i]
	}
	return nums
}

// https://leetcode.cn/problems/find-if-array-can-be-sorted/description/?envType=daily-question&envId=2024-07-13
func canSortArray(nums []int) bool {
	oneCount := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		res := strconv.FormatInt(int64(nums[i]), 2)
		oneCount[i] = strings.Count(res, "1")
	}
	temp := make([][2]int, 0)
	left, right := 0, 0
	for right < len(nums) {
		if oneCount[right] == oneCount[left] {
			right++
			if right < len(nums) && oneCount[right] != oneCount[left] {
				temp = append(temp, [2]int{left, right})
			}
		} else {
			left++
		}
	}
	temp = append(temp, [2]int{left, right})
	numMax, _ := getNumsMaxAndMin(nums, temp[0])
	for i := 1; i < len(temp); i++ {
		tempMax, tempMin := getNumsMaxAndMin(nums, temp[i])
		if tempMin < numMax {
			return false
		}
		numMax = tempMax
	}
	return true
}

func getNumsMaxAndMin(nums []int, index2 [2]int) (max, min int) {

	left := index2[0]
	right := index2[1]
	numMax := nums[left]
	numMin := nums[left]
	for i := left + 1; i < right; i++ {
		if nums[i] > numMax {
			numMax = nums[i]
		}
		if nums[i] < numMin {
			numMin = nums[i]
		}
	}
	return numMax, numMin
}

// https://leetcode.cn/problems/find-if-array-can-be-sorted/description/?envType=daily-question&envId=2024-07-13
func canSortArray2(nums []int) bool {
	//当前组的最大值
	currentGroupMax := 0

	// 当前1的个数
	latestOneCnt := 0

	// 上一组的最大值
	lastGroupMax := 0
	for i := 0; i < len(nums); i++ {
		if bits.OnesCount(uint(nums[i])) == latestOneCnt {
			currentGroupMax = max(currentGroupMax, nums[i])
		} else {
			// 更新最新的1的个数
			latestOneCnt = bits.OnesCount(uint(nums[i]))

			// 将当前组的最大值赋予上一组
			lastGroupMax = currentGroupMax

			// 更新当前组的最大值
			currentGroupMax = nums[i]
		}

		// 后面组的每个都必须大于上个组的最大值，否则无法排序
		if nums[i] < lastGroupMax {
			return false
		}
	}
	return true
}

// https://leetcode.cn/problems/max-increase-to-keep-city-skyline/?envType=daily-question&envId=2024-07-14
func maxIncreaseKeepingSkyline(grid [][]int) int {
	// 找到每一个点所在行列的最大值，
	// 这个最多比这个点所在的行列最大值要小
	m := len(grid)
	n := len(grid[0])
	hangMax := make([]int, m)
	lieMax := make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			hangMax[i] = max(hangMax[i], grid[i][j])
			lieMax[j] = max(lieMax[j], grid[i][j])
		}
	}
	totalCount := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			minHeight := min(hangMax[i], lieMax[j])
			if grid[i][j] < minHeight {
				totalCount += minHeight - grid[i][j]
			}
		}
	}
	return totalCount
}

type UF struct {
	// 记录连通分量
	Count int
	// 节点 x 的父节点是 Parent[x]
	Parent []int
}

// NewUF /* 构造函数，n 为图的节点总数 */
func NewUF(n int) *UF {
	// 一开始互不连通
	uf := &UF{Count: n, Parent: make([]int, n)}
	// 父节点指针初始指向自己
	for i := 0; i < n; i++ {
		uf.Parent[i] = i
	}
	return uf
}
func (uf *UF) find(x int) int {
	// 根节点的 Parent[x] == x
	if uf.Parent[x] != x {
		uf.Parent[x] = uf.find(uf.Parent[x])
	}
	return uf.Parent[x]
}

/* 返回当前的连通分量个数 */
func (uf *UF) count() int {
	return uf.Count
}

func (uf *UF) union(p int, q int) {
	rootP := uf.find(p)
	rootQ := uf.find(q)
	if rootP == rootQ {
		return
	}
	uf.Parent[rootQ] = rootP
	uf.Count--
}

func (uf *UF) connected(p int, q int) bool {
	rootP := uf.find(p)
	rootQ := uf.find(q)
	return rootP == rootQ
}

// https://leetcode.cn/problems/accounts-merge/
func accountsMerge(accounts [][]string) [][]string {
	emailId := make(map[string]int)
	emailName := make(map[string]string)
	for i := 0; i < len(accounts); i++ {
		name := accounts[i][0]
		for j := 1; j < len(accounts[i]); j++ {
			// 在这里只纪录第一次出现的emailName，相当于做了一次去重
			if _, ok := emailId[accounts[i][j]]; !ok {
				emailId[accounts[i][j]] = i
				emailName[accounts[i][j]] = name
			}
		}
	}
	uf := NewUF(len(emailId))

	// 同一个账户下的邮箱先将他们连接起来
	for _, account := range accounts {
		firstIndex := emailId[account[1]]
		for _, email := range account[2:] {
			uf.union(firstIndex, emailId[email])
		}
	}

	resMap := make(map[int][]string)
	for email, index := range emailId {
		parent := uf.find(index)
		resMap[parent] = append(resMap[parent], email)
	}

	res := make([][]string, 0)
	for _, emails := range resMap {
		sort.Strings(emails)
		account := append([]string{emailName[emails[0]]}, emails...)
		res = append(res, account)
	}
	return res
}

// https://leetcode.cn/problems/minimum-moves-to-spread-stones-over-grid/description/
func minimumMoves(grid [][]int) int {
	zero := make([][2]int, 0)
	more := make([][2]int, 0)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 0 {
				zero = append(zero, [2]int{i, j})
			}
			if grid[i][j] > 1 {
				for k := 2; k <= grid[i][j]; k++ {
					more = append(more, [2]int{i, j})
				}
			}
		}
	}
	length := len(zero)
	res := permutation(length)
	ans := math.MaxInt32
	for i := 0; i < len(res); i++ {
		step := 0
		for j := 0; j < len(res[i]); j++ {
			// more的排列
			moreOrder := res[i][j]
			step += manhadunDIstance(more[moreOrder], zero[j])
		}
		if step < ans {
			ans = step
		}
	}
	return ans
}

func permutation(n int) [][]int {
	res := make([][]int, 0)

	used := make([]bool, n)
	permutationBacktrck(n, &res, []int{}, used)
	return res
}

func permutationBacktrck(n int, res *[][]int, list []int, used []bool) {
	if len(list) == n {
		temp := make([]int, len(list))
		copy(temp, list)
		*res = append(*res, temp)
	}
	for i := 0; i < n; i++ {
		if used[i] {
			continue
		}
		// 做选择
		used[i] = true
		list = append(list, i)
		permutationBacktrck(n, res, list, used)
		list = list[:len(list)-1]
		used[i] = false
		//取消选择
	}
}

func manhadunDIstance(from [2]int, to [2]int) int {
	var step int
	x1, y1 := from[0], from[1]
	x2, y2 := to[0], to[1]
	if x1 > x2 {
		step += (x1 - x2)
	} else {
		step += (x2 - x1)
	}
	if y1 > y2 {
		step += (y1 - y2)
	} else {
		step += (y2 - y1)
	}
	return step
}

// https://leetcode.cn/problems/maximum-subarray-sum-with-one-deletion/description/
func maximumSum(arr []int) int {
	dp := make([][]int, len(arr))
	for i := 0; i < len(arr); i++ {
		dp[i] = make([]int, 2)
		dp[i][0] = math.MinInt32
		dp[i][1] = math.MinInt32
	}
	dp[0][0] = arr[0]
	dp[0][1] = 0
	res := arr[0]
	for i := 1; i < len(arr); i++ {
		// 不执行删除
		dp[i][0] = max(dp[i-1][0]+arr[i], arr[i])
		// 删除一次
		dp[i][1] = max(dp[i-1][0], dp[i-1][1]+arr[i])
		res = max(res, dp[i][1], dp[i][0])
	}
	return res
}

// https://leetcode.cn/problems/detonate-the-maximum-bombs/description/?envType=daily-question&envId=2024-07-22
// 引爆最多的炸弹

func maximumDetonation(bombs [][]int) int {
	length := len(bombs)
	graph := make([][]int, length)
	for i := 0; i < length; i++ {
		graph[i] = make([]int, 0)
	}
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if j == i {
				continue
			}
			if connect(bombs, i, j) {
				graph[i] = append(graph[i], j)
			}
		}
	}
	res := 0
	for i := 0; i < len(graph); i++ {
		visited := make([]bool, length)
		queue := []int{i}
		total := 0
		visited[i] = true
		for len(queue) > 0 {
			current := queue[0]
			queue = queue[1:]
			total++
			for j := 0; j < len(graph[current]); j++ {
				if !visited[graph[current][j]] {
					queue = append(queue, graph[current][j])
					visited[graph[current][j]] = true
				}
			}
		}
		res = max(res, total)
	}

	return res
}

func connect(bombs [][]int, x, y int) bool {
	originX := bombs[x][0]
	originY := bombs[x][1]
	r := bombs[x][2]
	targetX := bombs[y][0]
	targetY := bombs[y][1]
	x1 := targetX - originX
	y1 := targetY - originY
	return r*r >= (x1*x1)+(y1*y1)
}

// https://leetcode.cn/problems/find-the-sum-of-subsequence-powers/
func sumOfPowers(nums []int, k int) int {
	sort.Ints(nums)
	length := len(nums)
	used := make([]bool, length)
	tack := make([]int, 0)
	var res int
	sumOfPowersBacktack(nums, tack, k, 0, used, &res)
	return res
}

func sumOfPowersBacktack(nums []int, tack []int, k int, start int, used []bool, total *int) {
	if len(tack) == k {
		minEnergy := math.MaxInt64
		for i := 0; i < k; i++ {
			for j := i + 1; j < k; j++ {
				diff := abs(nums[tack[i]] - nums[tack[j]])
				if diff < minEnergy {
					minEnergy = diff
				}
			}
		}
		*total += minEnergy
	}
	for i := start; i < len(nums); i++ {
		tack = append(tack, i)
		sumOfPowersBacktack(nums, tack, k, i+1, used, total)
		tack = tack[:len(tack)-1]
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func sumOfPowers2(nums []int, k int) int {
	const mod int = 1e9 + 7
	slices.Sort(nums)
	memo, inf := map[string]int{}, math.MaxInt/2
	var dfs func(int, int, int, int) int
	// rest == 还需要选的元素数量, pre == 上一个元素的值, mn == 当前方案的最小值
	dfs = func(i, rest, pre, mn int) int {
		if i+1 < rest { // 剩下的元素不够选择了
			return 0
		}
		if rest == 0 {
			return mn
		}
		key := strconv.Itoa(i) + "#" + strconv.Itoa(rest) + "#" + strconv.Itoa(pre) + "#" + strconv.Itoa(mn)
		if v, ok := memo[key]; ok {
			return v
		}
		option1 := dfs(i-1, rest, pre, mn) % mod                         // 不选当前的数
		option2 := dfs(i-1, rest-1, nums[i], min(mn, pre-nums[i])) % mod // 选择当前的数
		ans := (option1 + option2) % mod
		memo[key] = ans
		return memo[key]
	}
	return dfs(len(nums)-1, k, inf, inf)
}

// https://leetcode.cn/problems/relocate-marbles/
func relocateMarbles(nums []int, moveFrom []int, moveTo []int) []int {
	myMap := make(map[int]struct{}, len(nums))
	for i := 0; i < len(nums); i++ {
		myMap[nums[i]] = struct{}{}
	}
	length := len(moveFrom)
	for i := 0; i < length; i++ {
		if _, ok := myMap[moveFrom[i]]; ok {
			delete(myMap, moveFrom[i])
		}
		myMap[moveTo[i]] = struct{}{}
	}
	res := make([]int, 0)
	for i, _ := range myMap {
		res = append(res, i)
	}
	sort.Ints(res)
	return res
}

// https://leetcode.cn/problems/minimum-operations-to-make-a-special-number/description/
/*
给你一个下标从 0 开始的字符串 num ，表示一个非负整数。
在一次操作中，您可以选择 num 的任意一位数字并将其删除。请注意，如果你删除 num 中的所有数字，则 num 变为 0。
返回最少需要多少次操作可以使 num 变成特殊数字。
如果整数 x 能被 25 整除，则该整数 x 被认为是特殊数字。
*/
func minimumOperations2(num string) int {
	atoi, _ := strconv.Atoi(num)
	if atoi%25 == 0 {
		return 0
	}

	length := len(num)
	res := length
	for i := 0; i < length; i++ {
		if num[i] == '0' {
			res = length - 1
		}
	}

	for i := length - 1; i >= 1; i-- {
		if num[i] != '0' && num[i] != '5' {
			continue
		}
		for j := i - 1; j >= 0; j-- {
			if length-j-2 > res {
				break
			}
			temp := string(num[j]) + string(num[i])
			atoi, _ = strconv.Atoi(temp)
			if atoi%25 == 0 {
				res = min(res, length-j-2)
				break
			}
		}
	}
	return res
}

func minimumOperations(num string) int {
	atoi, _ := strconv.Atoi(num)
	if atoi%25 == 0 {
		return 0
	}

	length := len(num)
	flag := false
	for i := length - 1; i > 0; i-- {
		minimumOperationsBackTrack(num, []byte{}, 0, i, &flag)
		if flag {
			return length - i
		}
	}
	return length
}

func minimumOperationsBackTrack(num string, back []byte, start int, length int, flag *bool) {
	if len(back) == length {
		numStr := string(back)
		atoi, _ := strconv.Atoi(numStr)
		if atoi%25 == 0 {
			*flag = true
			return
		}
	}
	for i := start; i < len(num); i++ {
		back = append(back, num[i])
		minimumOperationsBackTrack(num, back, i+1, length, flag)
		back = back[:len(back)-1]
	}
}

// https://leetcode.cn/problems/lexicographically-smallest-string-after-operations-with-constraint/description/
func getSmallestString(s string, k int) string {
	s1 := []rune(s)
	for i := 0; i < len(s); i++ {
		dis := distance(s1[i], 'a')
		if dis <= k {
			s1[i] = 'a'
			k -= dis
		} else {
			s1[i] = s1[i] - int32(k)
			break
		}
	}
	return string(s1)
}
func distance(s1, s2 rune) int {
	if s2 > s1 {
		return int(min(s2-s1, s2-s2+26))
	} else {
		return int(min(s1-s2, s2-s1+26))
	}
}

// https://leetcode.cn/problems/falling-squares/

func fallingSquares2(positions [][]int) []int {
	length := len(positions)
	ans := make([]int, 0, length)
	height := make([]int, 1e8)
	maxHeight := 0
	for i := 0; i < length; i++ {
		pos := positions[i]
		xstart, xend := pos[0], pos[0]+pos[1]
		preHeight := max(height[xstart+1], height[xend-1])
		for j := xstart + 1; j <= xend-1; j++ {
			preHeight = max(preHeight, height[j])
		}
		curHeight := preHeight + pos[1]
		maxHeight = max(maxHeight, curHeight)
		for j := xstart; j <= xend; j++ {
			height[j] = curHeight
		}

		ans = append(ans, maxHeight)

	}
	return ans
}

func fallingSquares(positions [][]int) []int {
	length := len(positions)
	height := make([]int, length)
	for i := 0; i < length; i++ {
		pos := positions[i]
		xStart, xEnd := pos[0], pos[0]+pos[1]-1
		height[i] = pos[1]
		for j := 0; j < i; j++ {
			xStart2, xEnd2 := positions[j][0], positions[j][0]+positions[j][1]-1
			// 保证线段重叠，一定能落在上一个方块的上面
			if xEnd >= xStart2 && xEnd2 >= xStart {
				height[i] = max(height[i], height[j]+pos[1])
			}
		}
	}
	for i := 1; i < length; i++ {
		height[i] = max(height[i], height[i-1])
	}

	return height
}

func findValueOfPartition(nums []int) int {
	sort.Ints(nums)
	ans := math.MaxInt
	for i := 0; i < len(nums)-1; i++ {
		ans = min(ans, nums[i+1]-nums[i])
	}
	return ans
}

// https://leetcode.cn/problems/baseball-game/
func calPoints(operations []string) int {
	nums := make([]int, 0, len(operations))
	total := 0
	for i := 0; i < len(operations); i++ {
		op := operations[i]
		switch op {
		case "+":
			length := len(nums)
			num := nums[length-1] + nums[length-2]
			nums = append(nums, num)
		case "C":
			nums = nums[:len(nums)-1]
		case "D":
			length := len(nums)
			num := nums[length-1] * 2
			nums = append(nums, num)
		default:
			num, _ := strconv.Atoi(op)
			nums = append(nums, num)
		}
	}
	for i := 0; i < len(nums); i++ {
		total += nums[i]
	}
	return total
}

// https://leetcode.cn/problems/double-modular-exponentiation/?envType=daily-question&envId=2024-07-30
func getGoodIndices(variables [][]int, target int) []int {
	ans := make([]int, 0, len(variables))
	for i := 0; i < len(variables); i++ {
		vari := variables[i]
		a, b, c, m := vari[0], vari[1], vari[2], vari[3]
		if target >= m {
			continue
		}
		num := 1
		for j := 0; j < b; j++ {
			num = num * a
			num = num % 10
		}
		num2 := 1
		for j := 0; j < c; j++ {
			num2 = num * num2
			num2 = num2 % m
		}
		num2 = num2 % m
		if num2 == target {
			ans = append(ans, i)
		}
	}
	return ans
}

// 快速幂算法
func pow_mod(x, y, mod int) int {
	res := 1
	for y > 0 {
		if (y & 1) == 1 {
			res = res * x % mod
		}
		x = x * x % mod
		y >>= 1
	}
	return res
}

// https://leetcode.cn/problems/minimum-rectangles-to-cover-points/description/
func minRectanglesToCoverPoints(points [][]int, w int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})
	ans := 1
	x := points[0][0] + w
	for i := 0; i < len(points); i++ {
		if points[i][0] > x {
			x = points[i][0] + w
			ans++
		}

	}
	return ans
}

// https://leetcode.cn/problems/uOAnQW/
func maxmiumScore(cards []int, cnt int) int {
	sort.Slice(cards, func(i, j int) bool {
		return cards[i] > cards[j]
	})
	ans := 0
	sum := 0
	minODD := math.MaxInt32
	minEven := math.MaxInt32
	for i := 0; i < cnt; i++ {
		sum += cards[i]
		if cards[i]%2 == 0 {
			minEven = min(minEven, cards[i])
		} else {
			minODD = min(minODD, cards[i])
		}
	}
	if sum%2 == 0 {
		return sum
	}
	nextOdd, nextEven := -1, -1
	for i := cnt; i < len(cards); i++ {
		if (nextOdd != -1) && (nextEven != -1) {
			break
		}
		if cards[i]%2 == 0 {
			if nextEven == -1 {
				nextEven = cards[i]
			}

		} else {
			if nextOdd == -1 {
				nextOdd = cards[i]
			}

		}
	}
	if minEven != math.MinInt32 && nextOdd != -1 {
		ans = max(ans, sum-minEven+nextOdd)
	}
	if minODD != math.MinInt32 && nextEven != -1 {
		ans = max(ans, sum-minODD+nextEven)
	}
	return ans
}

func maxmiumScore2(cards []int, cnt int) int {
	sort.Slice(cards, func(i, j int) bool {
		return cards[i] > cards[j]
	})
	ans := 0
	back := make([]int, 0, cnt)
	maxmiumScoreBacktrack(cards, back, 0, cnt, &ans, 0, math.MinInt32)
	return ans
}

func maxmiumScoreBacktrack(cards []int, back []int, start int, cnt int, ans *int, sum int, minNum int) int {
	if len(back) == cnt {
		if sum%2 == 0 {
			*ans = max(*ans, sum)
			minNum = min(minNum, back[cnt-1])
			return *ans
		}
	}
	for i := start; i < len(cards); i++ {
		if len(back)+len(cards)-start < cnt {
			break
		}
		if cards[i] < minNum {
			break
		}
		sum += cards[i]
		back = append(back, cards[i])
		ans1 := maxmiumScoreBacktrack(cards, back, i+1, cnt, ans, sum, minNum)
		sum -= cards[i]
		back = back[:len(back)-1]
		ans2 := maxmiumScoreBacktrack(cards, back, i+1, cnt, ans, sum, minNum)
		*ans = max(ans1, ans2)
		break

	}
	return *ans
}
