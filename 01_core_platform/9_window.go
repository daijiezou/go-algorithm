package _1_core_platform

import (
	"math"
)

// 最小覆盖子串
// https://leetcode.cn/problems/minimum-window-substring/
// 给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。
// 如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
func minWindow(s string, t string) string {
	need := make(map[byte]int)
	window := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	left, right := 0, 0
	valid := 0
	length := math.MaxInt
	start := 0
	for right < len(s) {
		current := s[right]
		right++
		if _, ok := need[current]; ok {
			window[current]++
			if window[current] == need[current] {
				valid++
			}
		}
		for valid == len(need) {
			if right-left < length {
				start = left
				length = right - left
			}
			toDelete := s[left]
			left++
			if _, ok := need[toDelete]; ok {
				if window[toDelete] == need[toDelete] {
					valid--
				}
				window[toDelete]--
			}
		}
	}
	if length == math.MaxInt {
		return ""
	}
	return s[start : start+length]
}

func checkInclusion(t string, s string) bool {
	need := make(map[byte]int)
	window := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	left, right := 0, 0
	valid := 0
	for right < len(s) {
		current := s[right]
		right++
		if _, ok := need[current]; ok {
			window[current]++
			if window[current] == need[current] {
				valid++
			}
		}
		for valid == len(need) {
			if right-left == len(t) {
				return true
			}
			toDelete := s[left]
			left++
			if _, ok := need[toDelete]; ok {
				if window[toDelete] == need[toDelete] {
					valid--
				}
				window[toDelete]--
			}
		}
	}
	return false
}

// https://leetcode.cn/problems/find-all-anagrams-in-a-string/description/
// 找到所有字符串中异位词
func findAnagrams(s string, t string) []int {
	need := make(map[byte]int)
	window := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	left, right := 0, 0
	valid := 0
	result := make([]int, 0)
	for right < len(s) {
		current := s[right]
		right++
		if _, ok := need[current]; ok {
			window[current]++
			if window[current] == need[current] {
				valid++
			}
		}
		for valid == len(need) {
			if right-left == len(t) {
				result = append(result, left)
			}
			toDelete := s[left]
			left++
			if _, ok := need[toDelete]; ok {
				if window[toDelete] == need[toDelete] {
					valid--
				}
				window[toDelete]--
			}
		}
	}
	return result
}
