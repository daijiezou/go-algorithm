package _2_queueAndStack

import (
	"fmt"
	"strconv"
	"strings"
)

// https://leetcode.cn/problems/reorder-list/description/
// 重排链表
func reorderList(head *ListNode) {
	nodeStack := make([]*ListNode, 0)
	p := head
	for p != nil {
		nodeStack = append(nodeStack, p)
		p = p.Next
	}
	p = head
	for p != nil {
		lastNode := nodeStack[len(nodeStack)-1]
		nodeStack = nodeStack[:len(nodeStack)-1]
		next := p.Next

		// 偶数节点，lastNode已经被接到链表的最后
		if lastNode == next {
			lastNode.Next = nil
			break
		}

		if lastNode.Next == next {
			lastNode.Next = nil
			break
		}

		p.Next = lastNode
		lastNode.Next = next
		p = next
	}
}

// https://leetcode.cn/problems/valid-parentheses/
func isValid(s string) bool {
	sByte := []byte(s)
	stack := []string{}
	for _, v := range sByte {
		switch string(v) {
		case "(", "{", "[":
			stack = append(stack, string(v))
		case ")":
			if len(stack) < 1 {
				return false
			}
			res := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if res != "(" {
				return false
			}
		case "}":
			if len(stack) < 1 {
				return false
			}
			res := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if res != "{" {
				return false
			}
		case "]":
			if len(stack) < 1 {
				return false
			}
			res := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if res != "[" {
				return false
			}
		}
	}
	return len(stack) == 0
}

// https://leetcode.cn/problems/evaluate-reverse-polish-notation/
func evalRPN(tokens []string) int {
	stack := []int{}
	for _, v := range tokens {
		switch v {
		case "+", "-", "*", "/":
			// 取出两个数字
			first := stack[len(stack)-1]
			second := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			switch v {
			case "+":
				second += first
				stack = append(stack, second)
			case "-":
				second -= first
				stack = append(stack, second)
			case "*":
				second *= first
				stack = append(stack, second)
			case "/":
				second /= first
				stack = append(stack, second)
			}
		default:
			// 数字
			num, _ := strconv.Atoi(v)
			stack = append(stack, num)
		}
	}
	return stack[0]
}

// https://leetcode.cn/problems/longest-absolute-file-path/description/
func lengthLongestPath(input string) int {
	var stack []string
	maxLen := 0
	for _, part := range strings.Split(input, "\n") {
		level := strings.LastIndex(part, "\t") + 1
		fmt.Println(part, level)
		// 让栈中只保留当前目录的父路径
		for level < len(stack) {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, part[level:])
		// 如果是文件，就计算路径长度
		if strings.Contains(part, ".") {
			sum := 0
			for _, s := range stack {
				sum += len(s)
			}
			// 加上父路径的分隔符
			sum += len(stack) - 1
			maxLen = max(maxLen, sum)
		}
	}
	return maxLen
}
