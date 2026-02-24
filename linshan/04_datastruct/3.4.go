package _4_datastruct

// 921. 使括号有效的最少添加
/*
只有满足下面几点之一，括号字符串才是有效的：

它是一个空字符串，或者
它可以被写成 AB （A 与 B 连接）, 其中 A 和 B 都是有效字符串，或者
它可以被写作 (A)，其中 A 是有效字符串。
给定一个括号字符串 s ，在每一次操作中，你都可以在字符串的任何位置插入一个括号

例如，如果 s = "()))" ，你可以插入一个开始括号为 "(()))" 或结束括号为 "())))" 。
返回 为使结果字符串 s 有效而必须添加的最少括号数。
*/
func minAddToMakeValid(s string) int {
	stack := make([]rune, 0)
	for _, x := range s {
		if x == ')' && len(stack) > 0 && stack[len(stack)-1] == '(' {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, x)
		}
	}
	return len(stack)
}

func minAddToMakeValid2(s string) int {
	leftNeed := 0
	balance := 0 // 左括号-右括号的差
	for _, x := range s {
		if x == '(' {
			balance++
		} else {
			balance--
			if balance < 0 {
				leftNeed++
				balance = 0
			}
		}
	}
	return balance+leftNeed
}
