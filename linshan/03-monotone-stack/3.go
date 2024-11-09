package _3_monotone_stack

/*
字典序
*/
func removeKdigits(num string, k int) string {
	if len(num) == k {
		return "0"
	}
	s := []byte{}
	n := len(num)
	cnt := 0
	for i := 0; i < n; i++ {
		for len(s) > 0 && s[len(s)-1] > num[i] && cnt < k {
			cnt++
			s = s[:len(s)-1]
		}
		s = append(s, num[i])
	}
	for i := cnt; i < k && len(s) > 1; i++ {
		s = s[:len(s)-1]
	}
	for len(s) > 1 && s[0] == '0' {
		s = s[1:]
	}
	return string(s)
}

func mostCompetitive(nums []int, k int) []int {
	n := len(nums)
	s := make([]int, 0)
	for i := 0; i < n; i++ {
		for len(s) > 0 && s[len(s)-1] > nums[i] && len(s)+n-i > k {
			s = s[:len(s)-1]
		}
		s = append(s, nums[i])
	}
	if len(s) > k {
		s = s[:k]
	}
	return s
}

// https://leetcode.cn/problems/remove-duplicate-letters/
func removeDuplicateLetters(s string) string {
	cnt := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		cnt[s[i]]++
	}

	st := []byte{}
	inStack := [256]bool{}
	for i := 0; i < len(s); i++ {
		cnt[s[i]]--
		// 前面已经有这个字符了
		if inStack[s[i]] {
			continue
		}
		for len(st) > 0 && st[len(st)-1] > s[i] {
			// 后面没这个字符了，不能弹出这个字符
			if cnt[st[len(st)-1]] <= 0 {
				break
			}
			x := st[len(st)-1]
			inStack[x] = false
			st = st[:len(st)-1]
		}
		st = append(st, s[i])
		inStack[s[i]] = true
		if len(st) >= 26 {
			break
		}
	}
	return string(st)
}
