package meeting150

import (
	"strconv"
	"strings"
)

/*
stack
*/

func isValid(s string) bool {
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			stack = append(stack, s[i])
		} else {
			if len(stack) < 1 {
				return false
			}
			pop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			switch s[i] {
			case ')':
				if pop != '(' {
					return false
				}
			case '}':
				if pop != '{' {
					return false
				}
			case ']':
				if pop != '[' {
					return false
				}
			}
		}
	}
	return len(stack) == 0
}

func simplifyPath(path string) string {
	pathSlice := strings.Split(path, "/")
	vaild := []string{}
	for i := 0; i < len(pathSlice); i++ {
		switch pathSlice[i] {
		case ".", "":
			continue
		case "..":
			if len(vaild) > 0 {
				vaild = vaild[:len(vaild)-1]
			}
		default:
			vaild = append(vaild, pathSlice[i])
		}
	}
	res := strings.Join(vaild, "/")
	return "/" + res
}

func evalRPN(tokens []string) int {
	stack := []int{}
	for i := 0; i < len(tokens); i++ {
		switch tokens[i] {
		case "+", "-", "*", "/":
			a := stack[len(stack)-1]
			b := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			if tokens[i] == "+" {
				stack = append(stack, a+b)
			} else if tokens[i] == "-" {
				stack = append(stack, a-b)
			} else if tokens[i] == "*" {
				stack = append(stack, a*b)
			} else {
				stack = append(stack, b/a)
			}
		default:
			num, _ := strconv.Atoi(tokens[i])
			stack = append(stack, num)
		}
	}
	return stack[0]
}
