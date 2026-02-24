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

/*
给出由小写字母组成的字符串 s，重复项删除操作会选择两个相邻且相同的字母，并删除它们。

在 s 上反复执行重复项删除操作，直到无法继续删除。

在完成所有重复项删除操作后返回最终的字符串。答案保证唯一。
*/
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

/*
给你一个由大小写英文字母组成的字符串 s 。

一个整理好的字符串中，两个相邻字符 s[i] 和 s[i+1]，其中 0<= i <= s.length-2 ，要满足如下条件:

若 s[i] 是小写字符，则 s[i+1] 不可以是相同的大写字符。
若 s[i] 是大写字符，则 s[i+1] 不可以是相同的小写字符。
请你将字符串整理好，每次你都可以从字符串中选出满足上述条件的 两个相邻 字符并删除，直到字符串整理好为止。

请返回整理好的 字符串 。题目保证在给出的约束条件下，测试样例对应的答案是唯一的。

注意：空字符串也属于整理好的字符串，尽管其中没有任何字符。
*/
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
/*
给你一个下标从 0 开始的整数数组 nums ，如果满足下述条件，则认为数组 nums 是一个 美丽数组 ：

nums.length 为偶数
对所有满足 i % 2 == 0 的下标 i ，nums[i] != nums[i + 1] 均成立
注意，空数组同样认为是美丽数组。

你可以从 nums 中删除任意数量的元素。当你删除一个元素时，被删除元素右侧的所有元素将会向左移动一个单位以填补空缺，而左侧的元素将会保持 不变 。

返回使 nums 变为美丽数组所需删除的 最少 元素数目。
*/
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

		// 检查最上方是否为 ) 个数是否足够，然后再检查前一个格式是否足够
		if len(st) > 1 && top.ch == ')' && top.freq == k && st[len(st)-2].freq >= k {
			// 如果足够则移除最上方
			st = st[:len(st)-1]
			st[len(st)-1].freq -= k
			// 查看是否已经用完了所有的数量，如已用完则删除次上方的字符
			if st[len(st)-1].freq == 0 {
				st = st[:len(st)-1]
			}
		}
	}

	// 拼接成最后的答案
	res := make([]byte, 0, len(s))
	for _, p := range st {
		for i := 0; i < p.freq; i++ {
			res = append(res, p.ch)
		}
	}
	return string(res)
}

// 1717
/*
给你一个字符串 s 和两个整数 x 和 y 。你可以执行下面两种操作任意次。

删除子字符串 "ab" 并得到 x 分。
比方说，从 "cabxbae" 删除 ab ，得到 "cxbae" 。
删除子字符串"ba" 并得到 y 分。
比方说，从 "cabxbae" 删除 ba ，得到 "cabxe" 。
请返回对 s 字符串执行上面操作若干次能得到的最大得分。
*/
func maximumGain(s string, x int, y int) int {
	// 贪心：优先删除分数高的对
	// 如果 x > y，优先删除 "ab"；否则优先删除 "ba"
	first, second := byte('a'), byte('b')
	firstScore, secondScore := x, y
	if y > x {
		first, second = 'b', 'a'
		firstScore, secondScore = y, x
	}

	res := 0

	// 第一遍：删除高分对
	stack := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if len(stack) > 0 && stack[len(stack)-1] == first && ch == second {
			// 匹配到高分对，删除
			stack = stack[:len(stack)-1]
			res += firstScore
		} else {
			stack = append(stack, ch)
		}
	}

	// 第二遍：删除低分对（first 和 second 互换）
	stack2 := make([]byte, 0, len(stack))
	for _, ch := range stack {
		if len(stack2) > 0 && stack2[len(stack2)-1] == second && ch == first {
			// 匹配到低分对，删除
			stack2 = stack2[:len(stack2)-1]
			res += secondScore
		} else {
			stack2 = append(stack2, ch)
		}
	}

	return res
}
