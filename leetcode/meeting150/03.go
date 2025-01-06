package meeting150

import "math"

/*
滑动窗口
*/

func minSubArrayLen(target int, nums []int) int {
	windowSum := 0
	left, right := 0, 0
	n := len(nums)
	res := math.MaxInt
	for right < n {
		windowSum += nums[right]
		for windowSum >= target {
			res = min(res, right-left+1)
			windowSum -= nums[left]
			left++
		}
		right++
	}
	if res == math.MaxInt {
		return 0
	}
	return res
}

// 无重复子串
func lengthOfLongestSubstring(s string) int {
	window := make(map[byte]int)
	left, right := 0, 0
	n := len(s)
	res := 0
	for right < n {
		cur := s[right]
		window[cur]++
		for window[cur] > 1 {
			leave := s[left]
			window[leave]--
			if window[leave] == 0 {
				delete(window, leave)
			}
			left++
		}
		res = max(res, right-left+1)
		right++
	}
	return res
}

// 最小覆盖子串
func minWindow(s string, t string) string {
	need := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	needCnt := len(need)
	window := make(map[byte]int)
	left, right := 0, 0
	n := len(s)
	valid := 0
	start, end := 0, n+1
	for ; right < n; right++ {
		cur := s[right]
		window[cur]++
		if window[cur] == need[cur] {
			valid++
		}
		for valid == needCnt {
			if right-left < end-start {
				start = left
				end = right
			}
			leave := s[left]
			left++
			if window[leave] == need[leave] {
				valid--
			}
			window[leave]--
		}

	}
	if end == n+1 {
		return ""
	}
	return s[start : end+1]
}

func findSubstring(s string, words []string) []int {
	wordLength := len(words[0])
	totalLength := wordLength * len(words)
	wordMap := make(map[string]int)
	for i := 0; i < len(words); i++ {
		wordMap[words[i]]++
	}
	res := []int{}
	for i := 0; i < wordLength; i++ {
		valid := 0
		curMap := make(map[string]int)
		left := i
		for right := i + wordLength; right <= len(s); right += wordLength {
			curWord := s[right-wordLength : right]
			if _, ok := wordMap[curWord]; ok {
				curMap[curWord]++
				if wordMap[curWord] == curMap[curWord] {
					valid++
				}
			}
			if right-left == totalLength {
				if valid == len(wordMap) {
					res = append(res, left)
				}
				// 移动左边界(先判是否在哈希表中,若在先更新哈希表)
				leave := s[left : left+wordLength]
				if _, ok := wordMap[leave]; ok {
					if curMap[leave] == wordMap[leave] {
						valid--
					}
					curMap[leave]--
				}
				left += wordLength
			}
		}

	}
	return res
}

func check(s string, start int, wordMap map[string]int, wordLength int, totalLength int) bool {
	valid := 0
	curMap := make(map[string]int)
	for i := start; i < start+totalLength; i += wordLength {
		curWord := s[i : i+wordLength]
		if _, ok := wordMap[curWord]; !ok {
			return false
		}
		curMap[curWord]++
		if wordMap[curWord] == curMap[curWord] {
			valid++
		}
	}
	return valid == len(wordMap)
}
