package leetcode

import (
	"container/heap"
	"math"
	"sort"
	"strings"
)

func pushDominoes(dominoes string) string {
	bytes := []byte("L" + dominoes + "R")
	pre := 0
	for i := 0; i < len(bytes); i++ {
		if bytes[i] == '.' {
			continue
		}
		if bytes[i] == bytes[pre] {
			fill(bytes[pre:i], bytes[i])
		} else if bytes[i] == 'L' {
			cnt := i - pre - 1
			if cnt%2 == 0 {
				fill(bytes[pre+1:pre+cnt/2+1], 'R') // 前一半变 R
				fill(bytes[pre+cnt/2+1:i], 'L')     // 后一半变 L
			} else {
				fill(bytes[pre+1:pre+cnt/2+1], 'R') // 前一半变 R
				fill(bytes[pre+cnt/2+2:i], 'L')     // 后一半变 L
			}
		}
		pre = i
	}
	return string(bytes[1 : len(bytes)-1])
}

func fill(bytes []byte, s byte) {
	for i := 0; i < len(bytes); i++ {
		bytes[i] = s
	}
}

// https://leetcode.cn/problems/minimum-domino-rotations-for-equal-row/solutions/3042326/du-bian-cheng-tops0-huo-zhe-bottoms0pyth-zvnj/?envType=daily-question&envId=2025-05-03
func minDominoRotations(tops []int, bottoms []int) int {
	n := len(tops)
	topCnt := make(map[int]int)
	botCnt := make(map[int]int)
	dup := make(map[int]int)
	for i := 0; i < n; i++ {
		topCnt[tops[i]]++
		botCnt[bottoms[i]]++
		if tops[i] == bottoms[i] {
			dup[tops[i]]++
		}
	}
	res := -1
	for i := 1; i <= 6; i++ {
		if topCnt[i]+botCnt[i]-dup[i] >= n {
			res = min(n-topCnt[i], n-botCnt[i])
		}
	}
	return res
}

func minDominoRotations1(tops []int, bottoms []int) int {
	var minCnts func(target int) int
	n := len(tops)
	minCnts = func(target int) int {
		topCnt := 0
		botCnt := 0
		for i := 0; i < n; i++ {
			if bottoms[i] != target && tops[i] != target {
				return math.MaxInt
			}
			if bottoms[i] != target {
				botCnt++
			} else if tops[i] != target {
				topCnt++
			}

		}
		return min(topCnt, botCnt)
	}
	// 想要每一行都一样，则必须跟第一列的bot或者top的数据是一致的
	ans := min(minCnts(tops[0]), minCnts(bottoms[0]))
	if ans == math.MaxInt {
		return -1
	}
	return ans
}

func numEquivDominoPairs(dominoes [][]int) int {
	numCnt := make(map[int]int)
	n := len(dominoes)
	res := 0
	for i := 0; i < n; i++ {
		d := dominoes[i]
		key := max(d[0]*10+d[1], d[1]*10+d[0])
		res += numCnt[key]
		numCnt[key]++
	}
	return res
}

// https://leetcode.cn/problems/domino-and-tromino-tiling/?envType=daily-question&envId=2025-05-05

/*
考虑这么一种平铺的方式：在第 i 列前面的正方形都被瓷砖覆盖，在第 i 列后面的正方形都没有被瓷砖覆盖（i 从 1 开始计数）。
那么第 i 列的正方形有四种被覆盖的情况：

一个正方形都没有被覆盖，记为状态 0；
只有上方的正方形被覆盖，记为状态 1；
只有下方的正方形被覆盖，记为状态 2；
上下两个正方形都被覆盖，记为状态 3。
*/
func numTilings(n int) int {
	const mod int = 1e9 + 7
	dp := make([][4]int, n+1)
	dp[0][3] = 1
	for i := 1; i <= n; i++ {
		dp[i][0] = dp[i-1][3]
		dp[i][1] = (dp[i-1][0] + dp[i-1][2]) % mod
		dp[i][2] = (dp[i-1][0] + dp[i-1][1]) % mod
		dp[i][3] = (((dp[i-1][0]+dp[i-1][1])%mod+dp[i-1][2])%mod + dp[i-1][3]) % mod
	}
	return dp[n][3]
}

func numTilings2(n int) int {
	if n == 1 {
		return 1
	}
	f := make([]int, n+1)
	f[0], f[1], f[2] = 1, 1, 2
	for i := 3; i <= n; i++ {
		f[i] = (f[i-1]*2 + f[i-3]) % 1_000_000_007
	}
	return f[n]
}

func buildArray(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = nums[nums[i]]
	}
	return res
}

// https://leetcode.cn/problems/find-minimum-time-to-reach-last-room-i/?envType=daily-question&envId=2025-05-07

var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func minTimeToReach(moveTime [][]int) (ans int) {
	n, m := len(moveTime), len(moveTime[0])
	dis := make([][]int, n)
	for i := range dis {
		dis[i] = make([]int, m)
		for j := range dis[i] {
			dis[i][j] = math.MaxInt
		}
	}
	dis[0][0] = 0

	h := hp{{}}
	for {
		top := heap.Pop(&h).(tuple)
		i, j := top.x, top.y
		if i == n-1 && j == m-1 {
			return top.dis
		}
		if top.dis > dis[i][j] {
			continue
		}
		time := (i+j)%2 + 1
		for _, d := range dirs {
			x, y := i+d.x, j+d.y
			if 0 <= x && x < n && 0 <= y && y < m {
				newD := max(top.dis, moveTime[x][y]) + time
				if newD < dis[x][y] {
					dis[x][y] = newD
					heap.Push(&h, tuple{newD, x, y})
				}
			}
		}
	}
}

type tuple struct{ dis, x, y int }
type hp []tuple

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].dis < h[j].dis }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(tuple)) }
func (h *hp) Pop() (v any)      { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }

func minSum(nums1 []int, nums2 []int) int64 {
	sum1 := 0
	sum2 := 0
	zero1 := 0
	zero2 := 0
	for i := 0; i < len(nums1); i++ {
		sum1 += nums1[i]
		if nums1[i] == 0 {
			zero1++
			sum1++
		}
	}
	for i := 0; i < len(nums2); i++ {
		sum2 += nums2[i]
		if nums2[i] == 0 {
			zero2++
			sum2++
		}
	}
	if zero1 == 0 && sum1 < sum2 {
		return -1
	}
	if zero2 == 0 && sum2 < sum1 {
		return -1
	}
	return int64(max(sum1, sum2))
}

func threeConsecutiveOdds(arr []int) bool {
	n := len(arr)
	oddIndex := make([]int, n)
	for i := 0; i < len(arr); i++ {
		if arr[i]%2 == 1 {
			oddIndex[i] = 1
			if i > 1 {
				if oddIndex[i-1] == 1 && oddIndex[i-2] == 1 {
					return true
				}
			}
		}
	}
	return false
}

func findEvenNumbers(digits []int) []int {
	n := len(digits)
	res := make([]int, 0)
	used := make([]bool, n)
	cur := make([]int, 0)

	var backtrack func(int)
	backtrack = func(pos int) {
		if pos == 3 {
			x := cur[0]*100 + cur[1]*10 + cur[2]
			if x >= 100 && x%2 == 0 {
				res = append(res, x)
			}
			return
		}

		seen := make(map[int]bool) // 避免重复数字
		for i := 0; i < n; i++ {
			if used[i] || seen[digits[i]] {
				continue
			}
			if pos == 0 && digits[i] == 0 { // 第一位不能为0
				continue
			}
			seen[digits[i]] = true
			used[i] = true
			cur = append(cur, digits[i])
			backtrack(pos + 1)
			cur = cur[:len(cur)-1]
			used[i] = false
		}
	}

	backtrack(0)
	sort.Ints(res)
	return res
}

func findEvenNumbers2(digits []int) []int {
	n := len(digits)
	unique := make(map[int]struct{})
	used := make([]bool, n)

	var backtrack func(cur []int)
	backtrack = func(cur []int) {
		if len(cur) == 3 {
			x := cur[0]*100 + cur[1]*10 + cur[2]
			if x%2 == 0 {
				unique[x] = struct{}{}
			}
			return
		}

		for i := 0; i < n; i++ {
			if !used[i] {
				if len(cur) == 0 && digits[i] == 0 { // 提前终止前导零
					continue
				}
				used[i] = true
				backtrack(append(cur, digits[i]))
				used[i] = false
			}
		}
	}

	backtrack([]int{})

	res := make([]int, 0, len(unique))
	for k := range unique {
		res = append(res, k)
	}
	sort.Ints(res)
	return res
}

const mod = 1000000007

func lengthAfterTransformations(s string, t int) int {
	cnt := make([]int, 26)
	for _, ch := range s {
		cnt[ch-'a']++
	}
	for i := 0; i < t; i++ {
		newCnt := make([]int, 26)
		newCnt[0] = cnt[25]
		newCnt[1] = (cnt[0] + cnt[25]) % mod
		for j := 2; j < 26; j++ {
			newCnt[j] = cnt[j-1]
		}
		cnt = newCnt
	}
	ans := 0
	for _, total := range cnt {
		ans += total % mod
	}
	return ans % mod
}

// https://leetcode.cn/problems/sort-colors/?envType=daily-question&envId=2025-05-17
func sortColors(nums []int) {
	quickSort(nums, 0, len(nums)-1)
}

func quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	p := Part(nums, left, right)
	quickSort(nums, left, p-1)
	quickSort(nums, p+1, right)
}

func Part(nums []int, left, right int) int {
	flag := nums[right]
	i := left
	for j := left; j < right; j++ {
		if nums[j] < flag {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	nums[i], nums[right] = nums[right], nums[i]
	return i
}

func triangleType(nums []int) string {
	sort.Ints(nums)
	if nums[0]+nums[1] <= nums[2] {
		return "none"
	}
	if nums[0] == nums[2] {
		return "equilateral"
	}
	if nums[0] == nums[1] || nums[1] == nums[2] {
		return "isosceles"
	}
	return "scalene"
}

func isZeroArray(nums []int, queries [][]int) bool {
	n := len(nums)
	diff := make([]int, n+1)
	diff[0] = nums[0]
	for i := 1; i < n; i++ {
		diff[i] = nums[i] - nums[i-1]
	}

	for i := 0; i < len(queries); i++ {
		left := queries[i][0]
		right := queries[i][1]
		diff[left] -= 1
		if right < n-1 {
			diff[right+1] += 1
		}

	}
	if diff[0] > 0 {
		return false
	}
	result := make([]int, n)
	result[0] = diff[0]

	for i := 1; i < n; i++ {
		result[i] = result[i-1] + diff[i]
		if result[i] > 0 {
			return false
		}
	}
	return true
}

func isZeroArray2(nums []int, queries [][]int) bool {
	n := len(nums)
	diff := make([]int, n+1)
	for i := 0; i < len(queries); i++ {
		l := queries[i][0]
		r := queries[i][1]
		diff[l] += 1
		diff[r+1] -= 1

	}
	sumD := 0
	for i := 0; i < len(nums); i++ {
		sumD += diff[i]
		if sumD < nums[i] {
			return false
		}
	}
	return true
}

// https://leetcode.cn/problems/zero-array-transformation-ii/?envType=daily-question&envId=2025-05-21
func minZeroArray(nums []int, queries [][]int) int {
	left := 0
	right := len(queries)
	for left <= right {
		mid := left + (right-left)/2
		if check(nums, queries, mid) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	if left > len(queries) {
		return -1
	}
	return left
}

func check(nums []int, queries [][]int, k int) bool {
	diff := make([]int, len(nums)+1)
	for _, q := range queries[:k] {
		l, r, val := q[0], q[1], q[2]
		diff[l] += val
		diff[r+1] -= val
	}
	sumD := 0
	for i := 0; i < len(nums); i++ {
		sumD += diff[i]
		if sumD < nums[i] {
			return false
		}
	}
	return true
}

// https://leetcode.cn/problems/zero-array-transformation-iii/?envType=daily-question&envId=2025-05-22
//func maxRemoval(nums []int, queries [][]int) int {
//
//}

func findWordsContaining(words []string, x byte) []int {
	res := make([]int, 0)
	for i := 0; i < len(words); i++ {
		if strings.Contains(words[i], string(x)) {
			res = append(res, i)
		}
	}
	return res
}

// https://leetcode.cn/problems/longest-palindrome-by-concatenating-two-letter-words/?envType=daily-question&envId=2025-05-25
func longestPalindrome(words []string) int {
	res := 0
	letters := make(map[string]int)
	dupLetters := make(map[string]int)
	for i := 0; i < len(words); i++ {
		letter1 := words[i][0]
		letter2 := words[i][1]
		if letter1 == letter2 {
			dupLetters[words[i]]++
		} else {
			if _, ok := letters[words[i]]; ok {
				letters[words[i]] += 1
			} else {
				key := string([]byte{letter2, letter1})
				if _, ok := letters[key]; ok {
					res += 4
					letters[key]--
					if letters[key] == 0 {
						delete(letters, key)
					}
				} else {
					letters[words[i]] = 1
				}
			}
		}
	}
	flag := false
	for _, v := range dupLetters {
		res += (v / 2) * 4
		if v%2 != 0 && !flag {
			res += 2
			flag = true
		}
	}
	return res
}

func longestPalindrome2(words []string) int {
	wordsMap := make(map[string]int)
	for i := 0; i < len(words); i++ {
		wordsMap[words[i]]++
	}
	res := 0
	mid := false
	for k, cnt := range wordsMap {
		rk := string(k[1]) + string(k[0])
		if rk == k {
			res += (cnt / 2) * 4
			if cnt%2 == 1 {
				mid = true
			}
		} else {
			res += min(cnt, wordsMap[rk]) * 2
		}
	}
	if mid {
		res += 2
	}
	return res
}

// https://leetcode.cn/problems/divisible-and-non-divisible-sums-difference/description/?envType=daily-question&envId=2025-05-27
func differenceOfSums(n int, m int) int {
	k := n / m
	nums2 := (m + k*m) * k / 2
	nums1 := (1+n)*n/2 - nums2
	return nums1 - nums2
}
