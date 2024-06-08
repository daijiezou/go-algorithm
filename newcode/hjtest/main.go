package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func isOperator(c rune) bool {
	return c == '+' || c == '-' || c == '*'
}

func evaluateExpression(expr string) (int64, error) {
	var result, current int64
	var operator rune = '+'
	for i := 0; i < len(expr); {
		if unicode.IsDigit(rune(expr[i])) {
			start := i
			for i < len(expr) && unicode.IsDigit(rune(expr[i])) {
				i++
			}
			num, err := strconv.ParseInt(expr[start:i], 10, 64)
			if err != nil {
				return 0, err
			}
			switch operator {
			case '+':
				result += current
				current = num
			case '-':
				result += current
				current = -num
			case '*':
				current *= num
			}
		} else if isOperator(rune(expr[i])) {
			operator = rune(expr[i])
			i++
		} else {
			return 0, fmt.Errorf("invalid character in expression")
		}
	}
	result += current
	return result, nil
}

func longestValidExpression(s string) int64 {
	n := len(s)
	if n == 0 {
		return 0
	}

	maxLength := 0
	startIndex := -1
	currentStart := -1
	valid := true

	for i := 0; i < n; i++ {
		if unicode.IsDigit(rune(s[i])) {
			if currentStart == -1 {
				currentStart = i
			}
			valid = true
		} else if isOperator(rune(s[i])) {
			if i == 0 || !unicode.IsDigit(rune(s[i-1])) {
				valid = false
				currentStart = -1
			}
		} else {
			valid = false
			currentStart = -1
		}

		if valid && currentStart != -1 && (i-currentStart+1 > maxLength) {
			maxLength = i - currentStart + 1
			startIndex = currentStart
		}
	}

	if maxLength == 0 {
		return 0
	}

	longestExpr := s[startIndex : startIndex+maxLength]
	result, err := evaluateExpression(longestExpr)
	if err != nil {
		return 0
	}
	return result
}

func main() {
	s := "a+1--23*45-+67*89b"
	fmt.Println(longestValidExpression(s)) // 示例输出
}
