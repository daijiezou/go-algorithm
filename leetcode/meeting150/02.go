package meeting150

import "strings"

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	left := 0
	right := len(s) - 1
	for left < right {
		for left < right && !(s[left] >= 'a' && s[left] <= 'z' || s[left] >= '0' && s[left] <= '9') {
			left++
		}
		for left < right && !(s[right] >= 'a' && s[right] <= 'z' || s[right] >= '0' && s[right] <= '9') {
			right--
		}
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

//	给定字符串 s 和 t ，判断 s 是否为 t 的子序列。
//	字符串的一个子序列是原始字符串删除一些（也可以不删除）字符而不改变剩余字符相对位置形成的新字符串。
//
// （例如，"ace"是"abcde"的一个子序列，而"aec"不是）。
func isSubsequence(s string, t string) bool {
	index := 0
	cnt := 0
loop1:
	for i := 0; i < len(s); i++ {
		for j := index; j < len(t); j++ {
			if s[i] == t[j] {
				index = j
				cnt++
				continue loop1
			}
		}
	}
	return cnt == len(s)
}
