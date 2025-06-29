package leetcode

import (
	"math"
	"slices"
	"sort"
	"strconv"
	"strings"
)

// https://leetcode.cn/problems/distribute-candies-among-children-ii/?envType=daily-question&envId=2025-06-01
/*
给你两个正整数 n 和 limit 。
请你将 n 颗糖果分给 3 位小朋友，确保没有任何小朋友得到超过 limit 颗糖果，
请你返回满足此条件下的 总方案数 。
*/
func distributeCandies(n int, limit int) int64 {
	ans := int64(0)
	for i := 0; i <= min(n, limit); i++ {

		// 剩下的糖果超过2*limit,不满足条件
		if n-i > 2*limit {
			// 不存在合法的方案
			continue
		}
		// 第二个小孩最多可以分
		secondMax := min(limit, n-i)

		// 第二个小孩最少得分
		secondMin := max(0, n-i-limit)
		ans += int64(secondMax - secondMin + 1)
	}
	return ans
}

func candy(ratings []int) int {
	n := len(ratings)
	ans := 0
	for i := 0; i < n; i++ {
		start := i
		if i > 0 && ratings[i-1] < ratings[i] {
			start--
		}
		for i+1 < n && ratings[i+1] > ratings[i] {
			i++
		}
		top := i
		for i+1 < n && ratings[i] > ratings[i+1] {
			i++
		}
		inc := top - start
		desc := i - top
		ans += max(inc, desc) + ((inc-1)*inc+(desc-1)*desc)/2
	}
	return ans + n
}

func maxCandies(status []int, candies []int, keys [][]int, containedBoxes [][]int, initialBoxes []int) int {
	total := 0
	n := len(status)
	visited := make([]int, n)
	var bfs func(boxs []int)
	bfs = func(boxs []int) {
		for len(boxs) > 0 {
			length := len(boxs)
			for _, box := range boxs {
				if status[box] == 0 {
					continue
				}

				for _, key := range keys[box] {
					status[key] = 1
				}
			}
			for i := 0; i < length; i++ {
				cur := boxs[0]
				boxs = boxs[1:]
				//if visited[cur] == 1 {
				//	continue
				//}

				if status[cur] == 1 {
					visited[cur] = 1
					total += candies[cur]
					boxs = append(boxs, containedBoxes[cur]...)
				}
			}
		}
	}
	bfs(initialBoxes)
	return total
}

func answerString1(word string, numFriends int) string {
	if numFriends == 1 {
		return word
	}
	res := ""
	n := len(word)
	maxLength := n - numFriends + 1
	maxIndex := [26][]int{}
	maxByte := uint8(0)
	for i := 0; i < n; i++ {
		b := word[i] - 'a'
		maxByte = max(maxByte, b)
		maxIndex[b] = append(maxIndex[b], i)
	}
	for _, i := range maxIndex[maxByte] {
		end := min(n, i+maxLength)
		res = max(res, word[i:end])
	}
	return res
}

// https://leetcode.cn/problems/find-the-lexicographically-largest-string-from-the-box-i/?envType=daily-question&envId=2025-06-04
func answerString(word string, numFriends int) string {
	if numFriends == 1 {
		return word
	}
	res := ""
	n := len(word)
	maxLength := n - numFriends + 1
	for i := 0; i < n; i++ {
		res = max(res, word[i:min(i+maxLength, n)])
	}
	return res
}

func smallestEquivalentString(s1 string, s2 string, baseStr string) string {
	parent := [26]byte{}
	for i := range parent {
		parent[i] = byte(i)
	}
	var find func(byte) byte
	find = func(b byte) byte {
		if parent[b] != b {
			b = find(parent[b])
		}
		return b
	}
	union := func(x, y byte) {
		small, big := find(x), find(y)
		if small > big {
			small, big = big, small
		}
		parent[big] = small
	}
	for i := 0; i < len(s1); i++ {
		union(s1[i]-'a', s2[i]-'a')
	}

	s := make([]byte, len(baseStr))
	for i, c := range baseStr {
		s[i] = find(byte(c)-'a') + 'a'
	}
	return string(s)

}

// https://leetcode.cn/problems/lexicographically-minimum-string-after-removing-stars/?envType=daily-question&envId=2025-06-07
func clearStars(s string) string {
	stack := make([][]int, 26)
	sByte := []byte(s)
	for i := 0; i < len(s); i++ {
		if s[i] != '*' {
			stack[s[i]-'a'] = append(stack[s[i]-'a'], i)
			continue
		}
		for j, sItem := range stack {
			if len(sItem) > 0 {
				x := sItem[len(sItem)-1]
				sByte[x] = '*'
				stack[j] = sItem[:len(sItem)-1]
				break
			}
		}
	}
	res := []byte{}
	for i := 0; i < len(sByte); i++ {
		if sByte[i] != '*' {
			res = append(res, sByte[i])
		}
	}
	return string(res)
}

// https://leetcode.cn/problems/lexicographical-numbers/?envType=daily-question&envId=2025-06-08
// 给你一个整数 n ，按字典序返回范围 [1, n] 内所有整数。
// 你必须设计一个时间复杂度为 O(n) 且使用 O(1) 额外空间的算法。
func lexicalOrder(n int) []int {
	res := make([]int, n)
	num := 1
	for i := 0; i < n; i++ {
		res[i] = num
		if num*10 <= n {
			num *= 10
		} else {
			for num%10 == 9 || num+1 > n {
				num /= 10
			}
			num++
		}
	}
	return res
}

func findKthNumber(n int, k int) int {
	num := 1
	for i := 0; i < k; i++ {
		if num*10 <= n {
			num *= 10
		} else {
			for num%10 == 9 || num+1 > n {
				num /= 10
			}
			num++
		}
	}
	return num
}

func maxDifference(s string) int {
	numsCnt := [26]int{}
	for i := range s {
		numsCnt[s[i]-'a']++
	}
	maxOdd := 0
	minEven := math.MaxInt32
	for _, i := range numsCnt {
		if i == 0 {
			continue
		}
		if i%2 == 1 {
			maxOdd = max(maxOdd, i)
		} else {
			minEven = min(minEven, i)
		}
	}
	return maxOdd - minEven
}

func minMaxDifference(num int) int {
	numstr := strconv.Itoa(num)
	mx := num
	for i := 0; i < len(numstr); i++ {
		if numstr[i] != '9' {
			mx, _ = strconv.Atoi(strings.ReplaceAll(numstr, string(numstr[i]), "9"))
			break
		}
	}
	mn, _ := strconv.Atoi(strings.ReplaceAll(numstr, string(numstr[0]), "0"))
	return mx - mn
}

func maxDiff(num int) int {
	numstr := strconv.Itoa(num)
	mx := num
	mn := num
	for i := 0; i < len(numstr); i++ {
		if numstr[i] != '9' {
			mx, _ = strconv.Atoi(strings.ReplaceAll(numstr, string(numstr[i]), "9"))
			break
		}
	}
	if numstr[0] != '1' {
		mn, _ = strconv.Atoi(strings.ReplaceAll(numstr, string(numstr[0]), "1"))
	} else {
		for i := 1; i < len(numstr); i++ {
			if numstr[i] != '1' && numstr[i] != '0' {
				mn, _ = strconv.Atoi(strings.ReplaceAll(numstr, string(numstr[i]), "0"))
				break
			}
		}
	}
	return mx - mn
}

func maximumDifference(nums []int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}
	preMin := math.MaxInt
	ans := 0
	for _, x := range nums {
		preMin = min(x, preMin)
		ans = max(ans, x-preMin)
	}
	if ans == 0 {
		return -1
	}
	return ans
}

// https://leetcode.cn/problems/count-the-number-of-arrays-with-k-matching-adjacent-elements/?envType=daily-question&envId=2025-06-17
func countGoodArrays(n int, m int, k int) int {
	return 0
}

// https://leetcode.cn/problems/partition-array-such-that-maximum-difference-is-k/?envType=daily-question&envId=2025-06-19
func partitionArray(nums []int, k int) int {
	slices.Sort(nums)
	res := 1
	start := nums[0]
	for i := 1; i < len(nums); i++ {
		x := nums[i]
		if x-start > k {
			start = x
			res++
		}
	}
	return res
}

// https://leetcode.cn/problems/maximum-manhattan-distance-after-k-changes/?envType=daily-question&envId=2025-06-20
func maxDistance2(s string, k int) int {
	res := 0
	x, y := 0, 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'N':
			y++
		case 'S':
			y--
		case 'W':
			x++
		case 'E':
			x--
		}
		res = max(res, min(abs2(x)+abs2(y)+2*k, i+1))
	}

	return res
}

func minimumDeletions(word string, k int) int {
	cnts := make([]int, 26)
	for _, b := range word {
		cnts[b-'a']++
	}
	slices.Sort(cnts)
	maxRemain := 0
	for i, base := range cnts {
		sum := 0
		for _, c := range cnts[i:] {
			sum += min(c, base+k)
		}
		maxRemain = max(maxRemain, sum)
	}
	return len(word) - maxRemain
}

func divideString(s string, k int, fill byte) []string {
	res := make([]string, 0)
	n := len(s)
	if n%k != 0 {
		for i := 0; i < k-(n%k); i++ {
			s += string(fill)
		}
	}

	for i := 0; i < len(s); i += k {
		res = append(res, s[i:i+k])
	}
	return res
}

func findKDistantIndices(nums []int, key int, k int) []int {
	res := make([]int, 0)
	n := len(nums)
	end := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] == key {
			start := max(0, i-k, end+1)
			end = min(n-1, i+k)
			for j := start; j <= end; j++ {
				res = append(res, j)
			}
			if end == n-1 {
				break
			}
		}
	}
	return res
}

// https://leetcode.cn/problems/find-subsequence-of-length-k-with-the-largest-sum/?envType=daily-question&envId=2025-06-28
func maxSubsequence(nums []int, k int) []int {
	type VI struct {
		v int
		i int
	}
	vis := make([]VI, 0, len(nums))
	for i := 0; i < len(nums); i++ {
		vis = append(vis, VI{
			v: nums[i],
			i: i,
		})
	}
	sort.Slice(vis, func(i, j int) bool {
		return vis[i].v > vis[j].v
	})
	vis = vis[:k]
	sort.Slice(vis, func(i, j int) bool {
		return vis[i].i < vis[j].i
	})
	res := []int{}
	for i := 0; i < k; i++ {
		res = append(res, vis[i].v)
	}
	return res
}

// https://leetcode.cn/problems/count-pairs-whose-sum-is-less-than-target/
func countPairs2824(nums []int, target int) int {
	n := len(nums)
	res := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i]+nums[j] < target {
				res++
			}
		}
	}
	return res
}

func countPairs2824_2(nums []int, target int) int {
	slices.Sort(nums)
	n := len(nums)
	left := 0
	right := n - 1
	res := 0
	for left < right {
		if nums[left]+nums[right] < target {
			// 以left为主，和right,right-1,一直到left+1都是符合要求的
			res += right - left
			left++
		} else {
			// 和最小的left相加都不符合
			right--
		}
	}
	return res

}

// https://leetcode.cn/problems/number-of-subsequences-that-satisfy-the-given-sum-condition/?envType=daily-question&envId=2025-06-29
func numSubseq(nums []int, target int) int {
	return 0
}
