package _4_datastruct

/*
给你一个仅由 大写 英文字符组成的字符串 s 。
你可以对此字符串执行一些操作，在每一步操作中，你可以从 s 中删除 任一个 "AB" 或 "CD" 子字符串。
通过执行操作，删除所有 "AB" 和 "CD" 子串，返回可获得的最终字符串的 最小 可能长度。
注意，删除子串后，重新连接出的字符串可能会产生新的 "AB" 或 "CD" 子串
*/
func minLength(s string) int {
	stack := make([]rune, 0)
	for _, x := range s {
		switch x {
		case 'B':
			if len(stack) > 0 && stack[len(stack)-1] == 'A' {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, x)
			}
		case 'D':
			if len(stack) > 0 && stack[len(stack)-1] == 'C' {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, x)
			}
		default:
			stack = append(stack, x)
		}
	}
	return len(stack)
}

func removeDuplicates(s string) string {
	stack := make([]rune, 0)
	for _, x := range s {
		if len(stack) > 0 && stack[len(stack)-1] == x {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, x)
		}
	}
	return string(stack)
}

func makeGood(s string) string {
	offset := 'A' - 'a'
	stack := make([]rune, 0)
	for _, x := range s {
		if len(stack) > 0 {
			top := stack[len(stack)-1]
			if x <= 'z' && top-x == offset {
				stack = stack[:len(stack)-1]
			} else if x-top == offset {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, x)
			}
		} else {
			stack = append(stack, x)
		}
	}
	return string(stack)
}

// 3561. 移除相邻字符
func resultingString(s string) string {
	stack := make([]rune, 0)
	for _, x := range s {
		if len(stack) > 0 {
			top := stack[len(stack)-1]
			switch x {
			case 'a':
				if top == 'b' || top == 'z' {
					stack = stack[:len(stack)-1]
				} else {
					stack = append(stack, x)
				}
			case 'z':
				if top == 'y' || top == 'a' {
					stack = stack[:len(stack)-1]
				} else {
					stack = append(stack, x)
				}
			default:
				if top == x-1 || top == x+1 {
					stack = stack[:len(stack)-1]
				} else {
					stack = append(stack, x)
				}
			}
		} else {
			stack = append(stack, x)
		}
	}
	return string(stack)
}

// 1003. 检查替换后的
func isValid(s string) bool {
	stack := make([]rune, 0) //栈
	for _, x := range s {
		if len(stack) >= 2 && x == 'c' {
			top := stack[len(stack)-1]
			top2 := stack[len(stack)-2]
			if top == 'b' && top2 == 'a' {
				stack = stack[:len(stack)-2]
			} else {
				stack = append(stack, x)
			}
		} else {
			stack = append(stack, x)
		}
	}
	return len(stack) == 0
}

func isValid2(s string) bool {
	stack := make([]rune, 0) //栈
	for _, x := range s {
		if x > 'a' {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if top != x-1 {
				return false
			}
		}
		if x < 'c' {
			stack = append(stack, x)
		}
	}
	return len(stack) == 0
}

// 2216. 美化数组
func minDeletion(nums []int) int {
	stack := make([]int, 0)
	delCnt := 0
	for _, x := range nums {
		if len(stack) > 0 {
			top := stack[len(stack)-1]
			if x == top {
				delCnt++
				continue
			} else {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, x)
		}

	}
	return delCnt + len(stack)
}

// 1209 暴力做法
func removeDuplicates_1(s string, k int) string {
	stack := make([]rune, 0)
	for _, x := range s {
		if len(stack) < k-1 {
			stack = append(stack, x)
		} else {
			isDup := true
			for i := len(stack) - k + 1; i < len(stack); i++ {
				if stack[i] != x {
					isDup = false
					break
				}
			}
			if isDup {
				stack = stack[:len(stack)-k+1]
			} else {
				stack = append(stack, x)
			}
		}
	}
	return string(stack)
}

// 1209
func removeDuplicates1209(s string, k int) string {
	type pair struct {
		ch   byte
		freq int
	}

	stack := make([]pair, 0, len(s))
	for i := 0; i < len(s); i++ {
		ch := s[i]
		n := len(stack)
		if n > 0 && stack[n-1].ch == ch {
			stack[n-1].freq++
			if stack[n-1].freq == k {
				stack = stack[:n-1]
			}
		} else {
			stack = append(stack, pair{ch: ch, freq: 1})
		}
	}

	res := make([]byte, 0, len(s))
	for _, p := range stack {
		for i := 0; i < p.freq; i++ {
			res = append(res, p.ch)
		}
	}
	return string(res)
}

// 3703.移除K-平衡子字符串
func removeSubstring(s string, k int) string {
	type pair struct {
		ch   byte
		freq int
	}
	st := make([]pair, 0, len(s))
	for i := 0; i < len(s); i++ {
		b := s[i]
		if len(st) > 0 && st[len(st)-1].ch == b {
			st[len(st)-1].freq++ // 连续相同括号个数 +1
		} else {
			st = append(st, pair{b, 1}) // 新的括号
		}
		top := st[len(st)-1]
		if len(st) > 1 && top.ch == ')' && top.freq == k && st[len(st)-2].freq >= k {
			st = st[:len(st)-1]
			st[len(st)-1].freq -= k
			if st[len(st)-1].freq == 0 {
				st = st[:len(st)-1]
			}
		}
	}

	res := make([]byte, 0, len(s))
	for _, p := range st {
		for i := 0; i < p.freq; i++ {
			res = append(res, p.ch)
		}
	}
	return string(res)
}
