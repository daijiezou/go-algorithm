package _1_huadongchuangkou

import "math"

func minSubArrayLen(target int, nums []int) int {
	sum := 0
	left := 0
	res := math.MaxInt
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		for sum >= target {
			res = min(res, i-left+1)
			sum -= nums[left]
			left++
		}
	}
	if res == math.MaxInt {
		return 0
	}
	return res
}

// 给你一个字符串 s 、一个字符串 t 。
// 返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
func minWindow(s string, t string) string {
	// 用map存t的字符
	need := make(map[byte]int, len(t))
	window := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	left, right := 0, 0
	valid := 0
	res := s
	flag := false
	for right < len(s) {
		window[s[right]]++
		if window[s[right]] == need[s[right]] {
			valid++
		}
		for valid == len(need) {
			flag = true
			if right-left+1 < len(res) {
				res = s[left : right+1]
			}
			window[s[left]]--
			if window[s[left]] < need[s[left]] {
				valid--
			}
			left++
		}
		right++
	}
	if flag {
		return res
	}
	return ""
}

// https://leetcode.cn/problems/smallest-range-covering-elements-from-k-lists/
func smallestRange(nums [][]int) []int {
	size := len(nums)
	// 纪录一个数字在哪些组里有
	indices := make(map[int][]int)
	xMin, XMax := math.MaxInt32, math.MinInt32
	for i := 0; i < size; i++ {
		for _, v := range nums[i] {
			indices[v] = append(indices[v], i)
			xMin = min(v, xMin)
			XMax = max(v, XMax)
		}
	}
	left, right := xMin, xMin
	bestLeft, bestRight := xMin, XMax
	freq := make(map[int]int)
	for right <= XMax {
		// 该数字至少在一个组内，计算该数字在几个组内
		if len(indices[right]) > 0 {
			for _, i := range indices[right] {
				freq[i]++
			}
			for len(freq) == size {
				// 更新答案
				if right-left < bestRight-bestLeft {
					bestLeft, bestRight = left, right
				}
				for _, i := range indices[left] {
					freq[i]--
					if freq[i] == 0 {
						delete(freq, i)
					}
				}
				left++
			}
		}
		right++
	}
	return []int{bestLeft, bestRight}
}

//
/*
https://leetcode.cn/problems/shortest-and-lexicographically-smallest-beautiful-string/description/
*/
func shortestBeautifulSubstring(s string, k int) string {
	n := len(s)
	left, right := 0, 0
	res := s + "a"
	oneCnt := 0
	for right < n {
		if s[right] == '1' {
			oneCnt++
		}
		for oneCnt == k {
			if right-left+1 <= len(res) {
				tmp := s[left : right+1]
				if len(tmp) == len(res) {
					res = min(res, tmp)
				} else {
					res = tmp
				}
			}
			if s[left] == '1' {
				oneCnt--
			}
			left++
		}
		right++
	}
	if len(res) > n {
		return ""
	}
	return res
}

// https://leetcode.cn/problems/replace-the-substring-for-balanced-string/
func balancedString(s string) int {
	letterCnt := make(map[byte]int, 4)
	n := len(s)
	needCnt := n / 4
	need := make(map[byte]int)
	flag := false
	for i := 0; i < n; i++ {
		letterCnt[s[i]]++
		if letterCnt[s[i]] > needCnt {
			need[s[i]] = letterCnt[s[i]] - needCnt
			flag = true
		}
	}
	if !flag {
		return 0
	}
	left := 0
	right := 0
	res := n
	for right < n {
		letterCnt[s[right]]--
		// 当除子串外，其他字符数量均不超过平均，则该子串是符合要求的
		for letterCnt['Q'] <= needCnt && letterCnt['W'] <= needCnt && letterCnt['E'] <= needCnt && letterCnt['R'] <= needCnt {
			res = min(res, right-left+1)
			letterCnt[s[left]]++
			left++
		}
		right++
	}
	return res
}

// https://leetcode.cn/problems/minimum-size-subarray-in-infinite-array/
func minSizeSubarray(nums []int, target int) int {
	total := 0
	length := len(nums)
	for i := 0; i < length; i++ {
		total += nums[i]
	}
	ans := math.MaxInt
	left := 0
	ctarget := target % total
	sum := 0
	for i := 0; i < length*2; i++ {
		sum += nums[i%length]
		for sum > ctarget {
			sum -= nums[left%length]
			left++
		}
		if sum == ctarget {
			ans = min(ans, i-left+1)
		}
	}
	if ans == math.MaxInt {
		return -1
	}
	return ans + (target/total)*length
}
