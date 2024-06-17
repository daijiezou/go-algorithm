package _1base

import (
	"container/list"
	"strings"
)

// https://leetcode.cn/problems/word-break/submissions/536513055/
func wordBreakHuiShuo(s string, wordDict []string) bool {
	wordDictMap := make(map[string]bool)
	for _, word := range wordDict {
		wordDictMap[word] = true
	}
	var found bool
	n := len(s)
	var wordBreakHuiShuoBacktrack func(start int) bool
	wordBreakHuiShuoBacktrack = func(start int) bool {
		if found {
			return true
		}
		if start == n {
			found = true
			return true
		}
		for i := start; i <= len(s); i++ {
			if wordDictMap[s[start:i]] {
				wordBreakHuiShuoBacktrack(i)
			}
		}
		return false
	}
	wordBreakHuiShuoBacktrack(0)
	return found
}

// https://leetcode.cn/problems/word-break/submissions/536513055/
func wordBreak(s string, wordDict []string) bool {
	wordDictMap := make(map[string]bool)
	for _, word := range wordDict {
		wordDictMap[word] = true
	}
	memo := make(map[int]int)
	for i := 0; i < len(s); i++ {
		memo[i] = -1
	}
	return wordBreakDp(s, 0, wordDictMap, memo)
}

func wordBreakDp(s string, start int, wordDict map[string]bool, memo map[int]int) bool {
	if start == len(s) {
		return true
	}
	if memo[start] != -1 {
		return memo[start] == 1
	}
	for length := 1; start+length <= len(s); length++ {
		if wordDict[s[start:start+length]] {
			// 因为s[0:len] 已经可以被拼出
			// 只要 s[start+len..] 可以被拼出，s[start..] 就能被拼出
			if wordBreakDp(s, start+length, wordDict, memo) == true {
				memo[start] = 1
				return true
			}
		}
	}
	memo[start] = 0
	return false
}

// https://leetcode.cn/problems/word-break-ii/submissions/536524271/
func wordBreak2(s string, wordDict []string) []string {
	res := make([]string, 0)
	backTrack(s, wordDict, 0, &res, []string{})
	return res
}

func backTrack(s string, wordDict []string, start int, res *[]string, track []string) {
	if start == len(s) {
		*res = append(*res, strings.Join(track, " "))
		return
	}
	for _, word := range wordDict {
		lenWord := len(word)
		if start+lenWord <= len(s) && s[start:start+lenWord] == word {
			// 选择
			track = append(track, word)
			backTrack(s, wordDict, start+lenWord, res, track)
			// 撤销选择
			track = track[:len(track)-1]
		}
	}
}

// 注意：go 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码不保证正确性，仅供参考。如有疑惑，可以参照我写的 java 代码对比查看。

func wordBreak2Dp(s string, wordDict []string) []string {
	// 集合快速查询
	wordSet := make(map[string]bool)
	for _, word := range wordDict {
		wordSet[word] = true
	}
	// 备忘录
	memo := make([]*list.List, len(s))

	// 定义一个下标从 i 开始的子串可以被拆分成符合条件的字符串的集合
	var dp func(i int) *list.List
	dp = func(i int) *list.List {
		// 如果在最后一个下标后面，就返回一个空字符串
		if i == len(s) {
			res := list.New()
			res.PushBack("")
			return res
		}
		// 如果这个下标的状态已经被计算过了，则直接返回
		if memo[i] != nil {
			return memo[i]
		}
		res := list.New()
		for length := 1; i+length <= len(s); length++ {
			// 取到这个下标开始的子串
			sub := s[i : i+length]
			// 如果这个子串在字典里，则递归调用找到字符串集合中剩下的组合
			if wordSet[sub] {
				subProblem := dp(i + length)
				for e := subProblem.Front(); e != nil; e = e.Next() {
					// 在前面加入这个子串
					str := e.Value.(string)
					if str == "" {
						res.PushBack(sub)
					} else {
						res.PushBack(sub + " " + str)
					}
				}
			}
		}
		// 将结果存入 memo，返回这个下标的字符串集合
		memo[i] = res
		return res
	}

	// 最终返回从 0 开始的字符串集合
	var res []string
	for e := dp(0).Front(); e != nil; e = e.Next() {
		res = append(res, e.Value.(string))
	}
	return res
}
