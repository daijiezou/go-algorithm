package _1_listNodeAndArray

import (
	"math"
)

/*
	滑动窗口算法框架
*/

func minWindow(s string, t string) string {
	sByte := []byte(s)
	tByte := []byte(t)
	need := make(map[byte]int)
	window := make(map[byte]int)
	for _, b := range tByte {
		need[b]++
	}
	left, right := 0, 0 // 滑动窗口
	var start int
	var valid int
	length := math.MaxInt32
	for right < len(sByte) {
		c := sByte[right] // c 是要加入窗口中的字符
		right++
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}
		for valid == len(need) {
			if right-left < length {
				start = left
				length = right - left
			}
			c = sByte[left]
			left++
			if _, ok := need[c]; ok {
				if window[c] == need[c] {
					valid--
				}
				window[c]--
			}
		}
	}
	if length == math.MaxInt32 { // 如果最小子串长度没有更新，则返回空格
		return ""
	}
	return string(sByte[start : start+length+1])
}

func checkInclusion(s1 string, s2 string) bool {
	sByte := []byte(s2)
	tByte := []byte(s1)
	need := make(map[byte]int)
	window := make(map[byte]int)
	for _, b := range tByte {
		need[b]++
	}
	left, right := 0, 0 // 滑动窗口
	var valid int
	for ; right < len(sByte); right++ {
		c := sByte[right] // c 是要加入窗口中的字符
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}
		for valid == len(need) {
			if right-left == len(tByte)-1 {
				return true
			}
			c = sByte[left]
			left++
			if _, ok := need[c]; ok {
				if window[c] == need[c] {
					valid--
				}
				window[c]--
			}
		}
	}
	return false
}

// https://leetcode.cn/problems/longest-substring-without-repeating-characters/description/
// 无重复的最长子串
func lengthOfLongestSubstring(s string) int {
	sByte := []byte(s)
	window := make(map[byte]int)
	left, right := 0, 0 // 滑动窗口
	maxLength := 0
	for ; right < len(sByte); right++ {
		c := sByte[right]
		window[c] += 1
		for window[c] > 1 {
			d := sByte[left]
			window[d]--
			left += 1
		}
		if right-left+1 > maxLength {
			maxLength = right - left + 1
		}
	}
	return maxLength
}
