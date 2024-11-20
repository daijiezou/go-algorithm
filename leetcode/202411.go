package leetcode

import (
	"math"
	"sort"
)

func minChanges(n int, k int) int {
	if n == k {
		return 0
	}
	res := 0
	for k > 0 || n > 0 {
		modK := k % 2
		modN := n % 2
		if modK != modN {
			if modN == 0 {
				return -1
			}
			res++
		}
		k = k / 2
		n = n / 2
	}
	return res
}

// https://leetcode.cn/problems/shopping-offers/description/
func shoppingOffers(price []int, special [][]int, needs []int) int {
	// 去除那些大礼包还不如单买更便宜的情况
	specials := filterSpecials(price, special)
	return shoppingOffersBackTrack(price, specials, needs, make([]int, len(price)))
}

func filterSpecials(price []int, specials [][]int) [][]int {
	newSpecials := [][]int{}
	for _, special := range specials {
		cost := 0
		for j := 0; j < len(special)-1; j++ {
			cost += special[j] * price[j]
		}
		if cost > special[len(special)-1] {
			newSpecials = append(newSpecials, special)
		}
	}
	return newSpecials
}

func shoppingOffersBackTrack(price []int, special [][]int, needs []int, current []int) int {
	flag := false
	for i := 0; i < len(current); i++ {
		if current[i] != needs[i] {
			flag = true
			break
		}
	}
	// 所有的都已经满足
	if !flag {
		return 0
	}
	res := math.MaxInt
	selectSepecial := false
loop1:
	for i := 0; i < len(special); i++ {
		for j := 0; j < len(special[i])-1; j++ {
			// 超过所需不能选
			if current[j]+special[i][j] > needs[j] {
				continue loop1
			}
		}
		selectSepecial = true
		for j := 0; j < len(special[i])-1; j++ {
			current[j] += special[i][j]
		}
		res = min(shoppingOffersBackTrack(price, special, needs, current)+special[i][len(special[i])-1], res)
		for j := 0; j < len(special[i])-1; j++ {
			current[j] -= special[i][j]
		}
	}
	if selectSepecial {
		return res
	}
	sum := 0
	for i := 0; i < len(needs); i++ {
		sum += (needs[i] - current[i]) * price[i]
	}
	return sum
}

// https://leetcode.cn/problems/sum-of-square-numbers/description/
// 直接使用2分查找解决
func judgeSquareSum(c int) bool {
	// 下界为0
	left := 0
	// 上界为sqrt(c)
	right := int(math.Sqrt(float64(c)))
	for left <= right {
		sum := left*left + right*right
		if sum == c {
			return true
		} else if sum > c {
			right--
		} else {
			left++
		}
	}
	return false
}

func losingPlayer(x int, y int) string {
	res := 1
	for x > 1 && y >= 4 {
		res ^= 1
		x -= 1
		y -= 4
	}
	if res == 0 {
		return "Alice"
	}
	return "Bob"
}

/*
https://leetcode.cn/problems/find-the-power-of-k-size-subarrays-ii/description/
给你一个长度为 n 的整数数组 nums 和一个正整数 k 。
一个数组的 能量值 定义为：
如果 所有 元素都是依次 连续 且 上升 的，那么能量值为 最大 的元素。
否则为 -1 。
你需要求出 nums 中所有长度为 k 的
子数组的能量值。
请你返回一个长度为 n - k + 1 的整数数组 results ，其中 results[i] 是子数组 nums[i..(i + k - 1)] 的能量值。
*/
func resultsArray(nums []int, k int) []int {
	n := len(nums)
	res := make([]int, n-k+1)
	for i := range res {
		res[i] = -1
	}
	cnt := 0
	for i := 0; i < n-1; i++ {
		if i == 0 || nums[i] == nums[i-1]+1 {
			cnt++
		} else {
			cnt = 1
		}
		if cnt >= k {
			res[i-k+1] = nums[i]
		}
	}
	return res
}

type NeighborSum struct {
	grid    [][]int
	n       int
	Address map[int][2]int
}

func Constructor2(grid [][]int) NeighborSum {
	addr := make(map[int][2]int)
	for rowIndex, row := range grid {
		for colIndex, cell := range row {
			addr[cell] = [2]int{rowIndex, colIndex}
		}
	}
	return NeighborSum{grid: grid, Address: addr, n: len(grid)}
}

func (this *NeighborSum) AdjacentSum(value int) int {
	sum := 0
	addr := this.Address[value]
	row := addr[0]
	col := addr[1]
	if row > 0 {
		sum += this.grid[row-1][col]
	}
	if row < this.n-1 {
		sum += this.grid[row+1][col]
	}
	if col > 0 {
		sum += this.grid[row][col-1]
	}
	if col < this.n-1 {
		sum += this.grid[row][col+1]
	}
	return sum
}

func (this *NeighborSum) DiagonalSum(value int) int {
	sum := 0
	addr := this.Address[value]
	row := addr[0]
	col := addr[1]
	if row > 0 && col > 0 {
		sum += this.grid[row-1][col-1]
	}
	if row > 0 && col < this.n-1 {
		sum += this.grid[row-1][col+1]
	}
	if row < this.n-1 && col > 0 {
		sum += this.grid[row+1][col-1]
	}
	if row < this.n-1 && col < this.n-1 {
		sum += this.grid[row+1][col+1]
	}
	return sum
}

/*
有序数组中的单一元素
https://leetcode.cn/problems/single-element-in-a-sorted-array/description/
*/
func singleNonDuplicate(nums []int) int {
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		res ^= nums[i]
	}
	return res
}

/*
https://leetcode.cn/problems/minimum-cost-to-cut-a-stick/
*/
func minCost2(n int, cuts []int) int {
	cuts = append(cuts, 0, n)
	sort.Ints(cuts)
	memo := make([][]int, len(cuts))
	for i := range memo {
		memo[i] = make([]int, len(cuts))
	}
	res := getCost(0, len(cuts)-1, cuts, memo)
	return res
}

func getCost(left, right int, cuts []int, memo [][]int) int {
	if right-left <= 1 {
		return 0
	}
	if memo[left][right] > 0 {
		return memo[left][right]
	}
	res := math.MaxInt
	for i := left + 1; i < right; i++ {
		res = min(res, getCost(left, i, cuts, memo)+getCost(i, right, cuts, memo))
	}
	memo[left][right] = res + cuts[right] - cuts[left]
	// 切割之前的木棍长度
	return memo[left][right]
}

/*
给你一个 二进制 字符串 s 和一个整数 k。

如果一个 二进制字符串 满足以下任一条件，则认为该字符串满足 k 约束：

字符串中 0 的数量最多为 k。
字符串中 1 的数量最多为 k。
返回一个整数，表示 s 的所有满足 k 约束 的
子字符串
的数量。
*/
// https://leetcode.cn/problems/count-substrings-that-satisfy-k-constraint-i/description/
func countKConstraintSubstrings(s string, k int) int {
	res := 0
	left, right := 0, 0
	n := len(s)
	zeroCnt := 0
	oneCnt := 0
	for right < n {
		if s[right] == '0' {
			zeroCnt++
		} else {
			oneCnt++
		}

		for oneCnt > k && zeroCnt > k {
			if s[left] == '0' {
				zeroCnt--
			} else {
				oneCnt--
			}
			left++
		}
		res += right - left + 1
		right++
	}
	return res
}

// https://leetcode.cn/problems/count-substrings-that-satisfy-k-constraint-ii/
func countKConstraintSubstrings2(s string, k int, queries [][]int) []int64 {
	left, right := 0, 0
	n := len(s)
	zeroCnt := 0
	oneCnt := 0
	res := make([]int64, len(queries))
	sum := make([]int, n+1)
	leftBound := make([]int, n)
	for right < n {
		if s[right] == '0' {
			zeroCnt++
		} else {
			oneCnt++
		}

		for oneCnt > k && zeroCnt > k {
			if s[left] == '0' {
				zeroCnt--
			} else {
				oneCnt--
			}
			left++
		}
		leftBound[right] = left //  记录合法子串右端点 i 对应的最小左端点 l
		sum[right+1] = sum[right] + right - left + 1
		right++
	}
	for i, q := range queries {
		l, r := q[0], q[1]
		// [l:j]内的所有子串均满足要求
		j := l + sort.SearchInts(leftBound[l:r+1], l) // 如果区间内所有数都小于 l，结果是 j=r+1
		res[i] = int64(sum[r+1] - sum[j] + (j-l+1)*(j-l)/2)
	}
	return res

}

func minFlips(grid [][]int) int {
	cowCnt := 0
	colCnt := 0
	m := len(grid)
	n := len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n/2; j++ {
			if grid[i][j] != grid[i][n-j-1] {
				cowCnt++
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m/2; j++ {
			if grid[j][i] != grid[m-j-1][i] {
				colCnt++
				if colCnt >= cowCnt {
					return cowCnt
				}
			}
		}
	}
	return min(cowCnt, colCnt)
}

/*
在社交媒体网站上有 n 个用户。给你一个整数数组 ages ，其中 ages[i] 是第 i 个用户的年龄。

如果下述任意一个条件为真，那么用户 x 将不会向用户 y（x != y）发送好友请求：

ages[y] <= 0.5 * ages[x] + 7  只能向年龄大于自己一半+7岁
ages[y] > ages[x]  只能向年龄比自己小的人发送请求
ages[y] > 100 && ages[x] < 100
*/
func numFriendRequests(ages []int) int {
	sort.Ints(ages)
	n := len(ages)
	ans := 0
	for i := 0; i < n; i++ {
		leftBound := sort.SearchInts(ages, ages[i]/2+8)
		rightBound := sort.SearchInts(ages, ages[i]+1) - 1
		if rightBound >= leftBound {
			if rightBound >= i {
				rightBound--
			}
			ans += rightBound - leftBound + 1
		}
	}
	return ans
}

// https://leetcode.cn/problems/image-smoother/
func imageSmoother(img [][]int) [][]int {

	m := len(img)
	n := len(img[0])
	ans := make([][]int, m)
	preSum := make([][]int, m+10)
	for i := 0; i < m+10; i++ {
		preSum[i] = make([]int, n+10)

	}
	for i := 0; i < m; i++ {
		ans[i] = make([]int, n)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			preSum[i][j] = preSum[i-1][j] + preSum[i][j-1] - preSum[i-1][j-1] + img[i-1][j-1]
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			a := max(0, i-1)
			b := min(m-1, i+1)
			c := max(0, j-1)
			d := min(n-1, j+1)
			cnt := (d - c + 1) * (b - a + 1)
			ans[i][j] = (preSum[b+1][d+1] - preSum[a][d+1] - preSum[b+1][c] + preSum[a][c]) / cnt
		}
	}
	return ans
}

// https://leetcode.cn/problems/shortest-distance-after-road-addition-queries-i/description/
func shortestDistanceAfterQueries(n int, queries [][]int) []int {

	m := len(queries)
	ans := make([]int, m)
	graph := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		graph[i] = append(graph[i], i+1)
	}
	for i := 0; i < m; i++ {
		source := queries[i][0]
		target := queries[i][1]
		graph[source] = append(graph[source], target)
		ans[i] = MinimumDistance(graph)
	}
	return ans
}

func MinimumDistance(graph [][]int) int {
	step := 1
	queue := []int{0}
	visited := make([]bool, len(graph)+1)
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]
			for _, neighbor := range graph[cur] {
				if neighbor == len(graph) {
					return step
				}
				if !visited[neighbor] {
					queue = append(queue, neighbor)
					visited[neighbor] = true
				}
			}

		}
		step++
	}
	return step
}

func shortestDistanceAfterQueries2(n int, queries [][]int) []int {
	parent := make([]int, n-1)
	for i := 0; i < n-1; i++ {
		parent[i] = i
	}
	// 找到自己的父节点
	find := func(x int) int {
		root := x
		for parent[root] != root {
			root = parent[root]
		}
		oloParent := parent[x]
		// 然后把 x 到根节点之间的所有节点直接接到根节点下面
		for x != root {
			parent[x] = root
			x = oloParent
			oloParent = parent[oloParent]
		}
		return root
	}
	ans := make([]int, len(queries))
	cnt := n - 1
	for qi, q := range queries {
		l, r := q[0], q[1]-1
		fr := find(r)
		for i := find(l); i < r; i = find(i + 1) {
			parent[i] = fr
			cnt--
		}
		ans[qi] = cnt
	}
	return ans
}

func shortestDistanceAfterQueries3(n int, queries [][]int) []int {
	nxt := make([]int, n-1)
	for i := range nxt {
		nxt[i] = i + 1 // 表示i指向的最右节点编号
	}

	ans := make([]int, len(queries))
	cnt := n - 1
	for qi, q := range queries {
		source := q[0]
		target := q[1]
		for i := source; nxt[i] < target; i, nxt[i] = nxt[i], target {
			cnt--
		}
		ans[qi] = cnt
	}
	return ans
}
