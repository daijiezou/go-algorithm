package _4_datastruct

import "strings"

/*
3.1 栈
*/

func buildArray(target []int, n int) []string {
	mx := target[len(target)-1]
	res := []string{}
	index := 0
	for i := 1; i <= mx; i++ {
		res = append(res, "push")
		if i == target[index] {
			index++
		} else {
			res = append(res, "pop")
		}
	}
	return res
}

func backspaceCompare(s string, t string) bool {
	s1 := []byte{}
	s2 := []byte{}
	for _, x := range s {
		if x != '#' {
			s1 = append(s1, byte(x))
		} else {
			if len(s1) > 0 {
				s1 = s1[:len(s1)-1]
			}
		}
	}
	for _, x := range t {
		if x != '#' {
			s2 = append(s2, byte(x))
		} else {
			if len(s2) > 0 {
				s2 = s2[:len(s2)-1]
			}
		}
	}
	return string(s1) == string(s2)
}

// O(1) 空间复杂度的优化版本
func backspaceCompare2(s string, t string) bool {
	// 辅助函数：找到下一个有效字符的位置
	nextValidChar := func(str string, idx int) int {
		skip := 0
		for idx >= 0 {
			if str[idx] == '#' {
				skip++
				idx--
			} else if skip > 0 {
				skip--
				idx--
			} else {
				break
			}
		}
		return idx
	}

	i, j := len(s)-1, len(t)-1

	for i >= 0 || j >= 0 {
		// 找到 s 中下一个有效字符
		i = nextValidChar(s, i)
		// 找到 t 中下一个有效字符
		j = nextValidChar(t, j)

		// 如果一个到达开头，另一个没有，返回 false
		if (i >= 0) != (j >= 0) {
			return false
		}

		// 如果都还有字符，比较是否相同
		if i >= 0 && j >= 0 {
			if s[i] != t[j] {
				return false
			}
		}

		i--
		j--
	}

	return true
}

func removeStars(s string) string {
	res := []byte{}
	for i, _ := range s {
		if s[i] != '*' {
			res = append(res, s[i])
		} else {
			if len(res) > 0 {
				res = res[:len(res)-1]
			}
		}
	}
	return string(res)
}

func validateStackSequences(pushed []int, popped []int) bool {
	stask := make([]int, 0, len(popped))
	j := 0
	for _, x := range pushed {
		stask = append(stask, x)
		for len(stask) > 0 && stask[len(stask)-1] == popped[j] {
			stask = stask[:len(stask)-1]
			j++
		}
	}
	return j == len(popped)
}

// 3412. 计算字符串的镜像分数
func calculateScore(s string) int64 {
	sum := 0
	stacks := make(map[byte][]int)
	for i := 0; i < len(s); i++ {
		target := 'z' - (s[i] - 'a')
		if v, ok := stacks[target]; ok && len(v) > 0 {
			j := v[len(v)-1]
			stacks[target] = v[:len(v)-1]
			sum += i - j
		} else {
			stacks[s[i]] = append(stacks[s[i]], i)
		}
	}
	return int64(sum)
}

func calculateScore2(s string) int64 {
	sum := 0
	stacks := make(map[byte][]int)
	for i := 0; i < len(s); i++ {
		target := 'a' + 'z' - s[i]
		if v, ok := stacks[target]; ok && len(v) > 0 {
			// 找到镜像对，计算分数
			j := v[len(v)-1]
			sum += i - j
			// 弹出栈顶
			stacks[target] = v[:len(v)-1]
		} else {
			// 没有镜像对，将当前字符入栈
			stacks[s[i]] = append(stacks[s[i]], i)
		}
	}
	return int64(sum)
}

// 71.简化路径
// https://leetcode.cn/problems/find-mirror-score-of-a-string/description/
func simplifyPath(path string) string {
	stack := []string{}
	for _, x := range strings.Split(path, "/") {
		if x == "." || x == "/" {
			continue
		}
		if x != ".." {
			stack = append(stack, x)
		} else if len(stack) > 0 {
			stack = stack[:len(stack)-1]
		}
	}
	return "/" + strings.Join(stack, "/")
}
