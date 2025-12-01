package _4_datastruct

import (
	"math"
	"strings"
)

// https://leetcode.cn/problems/minimum-sum-of-mountain-triplets-ii/
func minimumSum(nums []int) int {
	n := len(nums)
	mn := nums[0]
	res := math.MaxInt
	sufMin := make([]int, n)
	sufMin[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		sufMin[i] = min(sufMin[i+1], nums[i])
	}
	for i := 1; i < n-1; i++ {
		if mn < nums[i] && sufMin[i+1] < nums[i] {
			res = min(res, sufMin[i+1]+mn+nums[i])
		}
		mn = min(mn, nums[i])
	}
	if res == math.MaxInt {
		return -1
	}
	return res
}

/*
特殊三元组 定义为满足以下条件的下标三元组 (i, j, k)：

0 <= i < j < k < n，其中 n = nums.length
nums[i] == nums[j] * 2
nums[k] == nums[j] * 2
返回数组中 特殊三元组 的总数。

由于答案可能非常大，请返回结果对 109 + 7 取余数后的值。
*/
const mod = 1e9 + 7

func specialTriplets(nums []int) int {
	posMap := make(map[int]int)
	for _, num := range nums {
		posMap[num]++
	}

	curMap := make(map[int]int)
	total := 0
	for _, num := range nums {
		posMap[num]--
		// 左边x/2的个数 * 右边x*2的个数
		leftCnt := curMap[num*2]
		rightCnt := posMap[num*2]
		total = (total + leftCnt*rightCnt) % mod
		curMap[num]++
	}
	return total
}

/*
给你一个字符串 s ，返回 s 中 长度为 3 的不同回文子序列 的个数。

即便存在多种方法来构建相同的子序列，但相同的子序列只计数一次。

回文 是正着读和反着读一样的字符串。

子序列 是由原字符串删除其中部分字符（也可以不删除）且不改变剩余字符之间相对顺序形成的一个新字符串。

例如，"ace" 是 "abcde" 的一个子序列。
*/

// 枚举中间
func countPalindromicSubsequence(s string) int {
	sufCnt := [26]int{}
	for _, ch := range s {
		sufCnt[ch-'a']++
	}
	preHas := [26]bool{}
	has := [26][26]bool{}
	res := 0
	for i := 0; i < len(s); i++ { // 枚举中间字母 mid
		mid := s[i] - 'a'
		sufCnt[mid]--           // 撤销 mid 的计数，suf_cnt 剩下的就是后缀 [i+1,n-1] 每个字母的个数
		for alpha := range 26 { // 枚举两侧字母 alpha
			// 判断 mid 的左右两侧是否都有字母 alpha
			if preHas[alpha] && sufCnt[alpha] > 0 && !has[mid][alpha] {
				has[mid][alpha] = true
				res++
			}
		}
		preHas[s[i]-'a'] = true // 记录前缀 [0,i-1] 有哪些字母
	}
	return res
}

// 枚举两边
func countPalindromicSubsequence2(s string) (ans int) {
	for alpha := byte('a'); alpha <= 'z'; alpha++ { // 枚举两侧字母 alpha
		i := strings.IndexByte(s, alpha) // 最左边的 alpha 的下标
		if i < 0 {                       // s 中没有 alpha
			continue
		}
		j := strings.LastIndexByte(s, alpha) // 最右边的 alpha 的下标
		if i+1 >= j {                        // 长度不足 3
			continue
		}

		has := [26]bool{}
		for _, mid := range s[i+1 : j] { // 枚举中间字母 mid
			if !has[mid-'a'] {
				has[mid-'a'] = true // 避免重复统计
				ans++
			}
		}
	}
	return
}

/*
给你一个整数数组 nums ，数组中共有 n 个整数。
132 模式的子序列 由三个整数 nums[i]、nums[j] 和 nums[k] 组成，并同时满足：i < j < k 和 nums[i] < nums[k] < nums[j] 。

如果 nums 中存在 132 模式的子序列 ，返回 true ；否则，返回 false 。
*/
func find132pattern(nums []int) bool {
	n := len(nums)
	leftMin := make([]int, n)
	leftMin[0] = nums[0]
	for i := 1; i < n; i++ {
		leftMin[i] = min(leftMin[i-1], nums[i])
	}
	s := make([]int, 0)
	for i := n - 1; i > 0; i-- {
		rightMax := math.MinInt
		for len(s) > 0 && s[len(s)-1] < nums[i] {
			rightMax = s[len(s)-1]
			s = s[:len(s)-1]
		}
		s = append(s, nums[i])
		if rightMax > leftMin[i-1] {
			return true
		}
	}
	return false
}

/*
给你一个二维 boolean 矩阵 grid 。

如果 grid 的 3 个元素的集合中，一个元素与另一个元素在 同一行，并且与第三个元素在 同一列，则该集合是一个 直角三角形。3 个元素 不必 彼此相邻。

请你返回使用 grid 中的 3 个元素可以构建的 直角三角形 数目，且满足 3 个元素值 都 为 1 。
*/
func numberOfRightTriangles(grid [][]int) int64 {
	oneLocation := make([][]int, 0)
	rowMap := make(map[int]int)
	colMap := make(map[int]int)
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				oneLocation = append(oneLocation, []int{i, j})
				rowMap[i]++
				colMap[j]++
			}
		}
	}
	res := 0
	for _, loc := range oneLocation {
		row := loc[0]
		col := loc[1]
		if rowMap[row] < 2 {
			continue
		}
		if colMap[col] < 2 {
			continue
		}
		res += (rowMap[row]-1) * (colMap[col]-1)
	}
	return int64(res)
}
