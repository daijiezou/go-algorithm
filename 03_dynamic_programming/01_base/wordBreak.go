package _1_base

import (
	"container/list"
)

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

func wordBreakDp(s string, i int, wordDict map[string]bool, memo map[int]int) bool {
	if i == len(s) {
		return true
	}
	if memo[i] != -1 {
		return memo[i] == 1
	}
	for length := 1; i+length <= len(s); length++ {
		if wordDict[s[i:i+length]] {
			// 只要 s[i+len..] 可以被拼出，s[i..] 就能被拼出
			if wordBreakDp(s, i+length, wordDict, memo) == true {
				memo[i] = 1
				return true
			}
		}
	}
	memo[i] = 0
	return false
}

// https://leetcode.cn/problems/word-break-ii/submissions/536524271/
func wordBreak2(s string, wordDict []string) []string {
	wordDictMap := make(map[string]bool)
	for _, word := range wordDict {
		wordDictMap[word] = true
	}
	memo := make([]*list.List, len(s))
	var res []string
	for e := wordBreak2Dp(s, 0, wordDictMap, memo).Front(); e != nil; e = e.Next() {
		res = append(res, e.Value.(string))
	}
	return res
}

func wordBreak2Dp(s string, i int, wordDict map[string]bool, memo []*list.List) *list.List {
	if i == len(s) {
		res := list.New()
		res.PushBack("")
		return res
	}
	if memo[i] != nil {
		return memo[i]
	}
	res := list.New()
	for length := 1; i+length <= len(s); length++ {
		sub := s[i : i+length]
		if wordDict[sub] {
			tempRes := wordBreak2Dp(s, i+length, wordDict, memo)
			for e := tempRes.Front(); e != nil; e = e.Next() {
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
	memo[i] = res
	return memo[i]
}
