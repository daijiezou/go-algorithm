package leetcode

import (
	"math"
	"sort"
	"strconv"
)

// 20040801
// https://leetcode.cn/problems/uOAnQW/
func maxmiumScore3(cards []int, cnt int) int {
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

func numberOfRightTriangles(grid [][]int) int64 {
	m := len(grid)
	n := len(grid[0])
	hangMap := make(map[int]int, m)
	lieMap := make(map[int]int, n)
	oneList := make([][]int, 0)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				hangMap[i] += 1
				lieMap[j] += 1
				oneList = append(oneList, []int{i, j})
			}
		}
	}
	ans := 0
	for _, v := range oneList {
		i, j := v[0], v[1]
		hangCount := hangMap[i]
		lieCount := lieMap[j]
		if hangCount < 2 {
			continue
		}
		if lieCount < 2 {
			continue
		}
		ans += (hangCount - 1) * (lieCount - 1)
	}
	return int64(ans)
}

type pointstru struct {
	x, y  int
	label byte
}

// https://leetcode.cn/problems/maximum-points-inside-the-square/description/
func maxPointsInsideSquare(points [][]int, s string) int {
	pointList := make([]pointstru, 0)
	for i := 0; i < len(points); i++ {
		if points[i][0] < 0 {
			points[i][0] = -points[i][0]
		}
		if points[i][1] < 0 {
			points[i][1] = -points[i][1]
		}
		pointList = append(pointList, pointstru{
			x:     points[i][0],
			y:     points[i][1],
			label: s[i],
		})
	}
	sort.Slice(pointList, func(i, j int) bool {
		return max(pointList[i].x, pointList[i].y) < max(pointList[j].x, pointList[j].y)
	})
	mymap := make(map[byte]struct{})
	count := 0
	tempCnt := 0
	startIndex := 0
loop1:
	for i := 0; i <= 1e9; i++ {
		tempCnt = 0
		if startIndex >= len(points) {
			break loop1
		}
		for j := startIndex; j < len(pointList); j++ {
			cur := pointList[j]
			if cur.x > i || cur.y > i {
				count += tempCnt
				i = max(cur.x, cur.y) - 1
				continue loop1
			}
			startIndex++
			if _, ok := mymap[cur.label]; !ok {
				mymap[cur.label] = struct{}{}
				tempCnt++
			} else {
				break loop1
			}
		}
		count += tempCnt
	}
	return count
}

func maxPointsInsideSquare2(points [][]int, s string) int {
	// 维护每个字符的最小距离
	min1 := make([]int, 26)
	for i := 0; i < 26; i++ {
		min1[i] = 1e10 + 1
	}

	// 维护所有字符的次小距离
	var min2 int = 1e10 + 1
	for i := 0; i < len(points); i++ {
		x, y := points[i][0], points[i][1]
		label := s[i] - 'a'
		d := getd(x, y)
		if d < min1[label] {
			min2 = min(min2, min1[label])
			min1[label] = d
		} else {
			min2 = min(min2, d)
		}
	}
	count := 0
	for i := 0; i < len(min1); i++ {
		if min1[i] < min2 {
			count++
		}
	}
	return count
}

func getd(x, y int) int {
	if x < 0 {
		x = -x

	}
	if y < 0 {
		y = -y
	}
	return max(x, y)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {

	if isSame(root, subRoot) {
		return true
	}
	if root == nil {
		return subRoot == nil
	}
	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func isSame(t1, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil {
		return false
	}
	if t1.Val != t2.Val {
		return false
	}
	if t1.Left == nil && t2.Left != nil {
		return false
	}
	if t1.Right == nil && t2.Right != nil {
		return false
	}
	return isSame(t1.Left, t2.Left) && isSame(t1.Right, t2.Right)
}

func findIntegers(n int) int {
	nstr := strconv.FormatInt(int64(n), 2)

	length := len(nstr) - 1
	cnt := 0
	for i := 2; i <= length; i++ {
		cnt += getCnt(length, i)
	}
	return n + 1 - cnt
}

func getCnt(n, sub int) int {
	cnt := 0
	for sub > n {
		cnt += n - sub + 1
		n--
	}
	return cnt
}

// https://leetcode.cn/problems/find-all-possible-stable-binary-arrays-i/
func numberOfStableArrays(zero int, one int, limit int) int {
	return 0
}

func addedInteger(nums1 []int, nums2 []int) int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	length := len(nums1)
	if length == 0 {
		return 0
	}
	return nums1[0] - nums2[0]
}

// https://leetcode.cn/problems/find-the-integer-added-to-array-ii/
func minimumAddedInteger(nums1 []int, nums2 []int) int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	back := make([]int, 0)
	res := math.MaxInt32
	diff := make([]int, len(nums2)-1)
	for i := 1; i < len(nums2); i++ {
		diff[i-1] = nums2[i] - nums2[i-1]
	}

	minimumAddedIntegerBackTack(nums1, nums2, back, 0, diff, &res)
	return res
}

func minimumAddedIntegerBackTack(nums1 []int, nums2 []int, back []int, start int, diff []int, res *int) {
	if len(back) == len(nums2) {
		sub := nums2[0] - back[0]
		*res = min(*res, sub)
		return
	}

	for i := start; i < len(nums1); i++ {
		if len(back)+len(nums1)-start < len(nums2) {
			return
		}
		if len(back)+1 > len(nums2) {
			return
		}
		if len(back) >= 1 {
			if nums1[i]-back[len(back)-1] < diff[len(back)-1] {
				continue
			}
			if nums1[i]-back[len(back)-1] > diff[len(back)-1] {
				return
			}
		}
		back = append(back, nums1[i])
		minimumAddedIntegerBackTack(nums1, nums2, back, i+1, diff, res)
		back = back[:len(back)-1]
		minimumAddedIntegerBackTack(nums1, nums2, back, i+1, diff, res)
	}
}

/*
由于只能移除两个元素，所以 nums1
的前三小元素必定有一个是保留下来的，我们可以枚举保留下来的最小元素是 nums
*/
func minimumAddedInteger2(nums1 []int, nums2 []int) int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	for i := 2; i >= 0; i-- {
		x := nums2[0] - nums1[i]
		j := 0
		for _, v := range nums1[i:] {
			if nums2[j] == v+x {
				j++
				// nums2 是 {nums1[i] + x} 的子序列
				if j == len(nums2) {
					return x
				}
			}
		}

	}
	return nums2[0] - nums1[0]
}

func maxUncrossedLines(nums1 []int, nums2 []int) int {
	m := len(nums1)
	n := len(nums2)
	//dp := make([][]int, m)
	//for i := 0; i < m; i++ {
	//	dp[i] = make([]int, n)
	//}
	memo := make([][]int, m)
	for i := 0; i < m; i++ {
		memo[i] = make([]int, n)
		for j := 0; j < n; j++ {
			memo[i][j] = -1
		}
	}
	return maxUncrossedLinesDp(memo, 0, 0, nums1, nums2)
}

func maxUncrossedLinesDp(memo [][]int, i, j int, nums1 []int, nums2 []int) int {
	if i == len(nums1) {
		return 0
	}
	if j == len(nums2) {
		return 0
	}
	if memo[i][j] != -1 {
		return memo[i][j]
	}
	if nums1[i] == nums2[j] {
		option1 := maxUncrossedLinesDp(memo, i+1, j, nums1, nums2)
		option2 := maxUncrossedLinesDp(memo, i, j+1, nums1, nums2)
		option3 := maxUncrossedLinesDp(memo, i+1, j+1, nums1, nums2) + 1
		memo[i][j] = max(option1, option2, option3)
	} else {
		option1 := maxUncrossedLinesDp(memo, i+1, j, nums1, nums2)
		option2 := maxUncrossedLinesDp(memo, i, j+1, nums1, nums2)
		memo[i][j] = max(option1, option2)
	}
	return memo[i][j]
}

func maxUncrossedLines2(nums1 []int, nums2 []int) int {
	m := len(nums1)
	n := len(nums2)
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[m][n]
}

// https://leetcode.cn/problems/implement-magic-dictionary/description/
type MagicDictionary struct {
	v map[int][]string
}

func Constructor() MagicDictionary {
	v := make(map[int][]string, 100)
	return MagicDictionary{
		v: v,
	}
}

func (this *MagicDictionary) BuildDict(dictionary []string) {
	for _, word := range dictionary {
		n := len(word)
		if _, ok := this.v[n]; ok {
			this.v[n] = append(this.v[n], word)
		} else {
			this.v[n] = []string{word}
		}
	}
}

func (this *MagicDictionary) Search(searchWord string) bool {
	n := len(searchWord)
	if _, ok := this.v[n]; !ok {
		return false
	}
	for _, word := range this.v[n] {
		diff := 0
		for i := 0; i < n; i++ {
			if word[i] != searchWord[i] {
				diff++
			}
		}
		if diff == 1 {
			return true
		}
	}
	return false
}

func MagicDict(req1 []string, req2 string) bool {
	m := Constructor()
	m.BuildDict(req1)
	return m.Search(req2)
}

// https://leetcode.cn/problems/special-array-i/description/
func arraySpecial(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	for i := 1; i < len(nums); i++ {
		if nums[i]%2 == nums[i-1]%2 {
			return false
		}
	}
	return true
}

// https://leetcode.cn/problems/special-array-ii/description/
func isArraySpecial2(nums []int, queries [][]int) []bool {
	preSpe := make([]int, len(nums))
	for i := 1; i < len(nums); i++ {
		preSpe[i] = preSpe[i-1]
		if nums[i]%2 == nums[i-1]%2 {
			preSpe[i]++
		}
	}
	res := make([]bool, len(queries))
	for i := 0; i < len(queries); i++ {
		start, end := queries[i][0], queries[i][1]
		res[i] = preSpe[start] == preSpe[end]
	}
	return res
}

// https://leetcode.cn/problems/maximum-difference-score-in-a-grid/
func maxScore(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
		// 第一列初始化为最小
		dp[i][0] = math.MaxInt32
		dp[0][1] = math.MaxInt32
	}
	return ans
}

func maxScoreDp(i, j int, grid [][]int, move int) int {
	if i == len(grid) && j == len(grid[0]) {
		if move > 0 {
			return 0
		} else {
			return math.MinInt32
		}

	}
	right := math.MinInt32
	if move > 0 {

	} else {
		for i < len(grid)-1 {
			op := maxScoreDp(i+1, j, grid, move+1) + grid[i+1][j] - grid[i][j]
			right = max(right, op)
			i++
		}
	}

	down := math.MinInt32
	for j < len(grid[0])-1 {
		op := maxScoreDp(i, j+1, grid, move+1) + grid[i][j+1] - grid[i][j]
		down = max(down, op)
		j++
	}
	return max(right, down)
}
