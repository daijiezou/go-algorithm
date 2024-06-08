package newcode

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

func brotherword() {
	var count int
	var k int
	var x string
	fmt.Scan(&count)
	words := make([]string, count)
	for i := 0; i < count; i++ {
		fmt.Scan(&words[i])
	}
	fmt.Scan(&x)
	fmt.Scan(&k)

	xMap := make(map[byte]int)
	for i := 0; i < len(x); i++ {
		xMap[x[i]]++
	}
	brotherList := make([]string, 0)
lool1:
	for i := 0; i < count; i++ {
		word := words[i]
		if len(word) != len(x) {
			continue
		}
		if word == x {
			continue
		}
		for key, val := range xMap {
			if strings.Count(word, string(key)) != val {
				continue lool1
			}
		}
		brotherList = append(brotherList, word)
	}
	sort.Strings(brotherList)
	fmt.Println(len(brotherList))
	if k <= len(brotherList) {
		fmt.Println(brotherList[k-1])
	}
}

func CoordinateShift() {
	origin := []int{0, 0}
	var s1 string
	fmt.Scan(&s1)
	orderList := strings.Split(s1, ";")
	for _, s := range orderList {
		if s == "" {
			continue
		}
		if len(s) == 1 {
			continue
		}
		fangxiang := s[0]
		distance := s[1:]
		distanceInt, err := strconv.Atoi(distance)
		if err != nil {
			continue
		}
		switch fangxiang {
		case 'A':
			origin[0] -= distanceInt
		case 'S':
			origin[1] -= distanceInt
		case 'W':
			origin[1] += distanceInt
		case 'D':
			origin[0] += distanceInt
		default:
			continue
		}
	}
	x := strconv.Itoa(origin[0])
	y := strconv.Itoa(origin[1])
	fmt.Println(x + "," + y)
}

// 华为机试雨花石问题
func Yuhuashi(nums []int) int {
	totalSum := 0
	for i := 0; i < len(nums); i++ {
		totalSum += nums[i]
	}
	if totalSum%2 != 0 {
		return -1
	}
	targetSum := totalSum / 2
	n := len(nums)

	//dp数组含义
	// dp[i][j] 表示达到每一个重量用的最多的雨花石数量
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, targetSum+1)
		dp[i][0] = 0
	}
	for i := 1; i <= n; i++ {
		currentWeight := nums[i-1]
		for j := 1; j <= targetSum; j++ {
			// 剩余的重量小于当前石头重量，不能拿
			if j < currentWeight {
				//只能选择不拿
				dp[i][j] = dp[i-1][j]
			} else {
				// 选择不拿
				option1 := dp[i-1][j]
				// 选择拿
				option2 := dp[i-1][j-currentWeight] + 1
				dp[i][j] = max(option1, option2)
			}
		}
	}
	if dp[n][targetSum] == n {
		return -1
	}
	return min(dp[n][targetSum], n-dp[n][targetSum])
}

/*
提取字符串中的最长合法简单数学表达式，字符串长度最长的，并计算表达式的值。如果没有，则返回 0 。
简单数学表达式只能包含以下内容：

输入：1-2abcd
输出-1
*/
func LongestExpression(s string) int64 {
	n := len(s)
	var maxLength int
	var currentStart int
	var valid bool
	startIndex := -1
	for i := 0; i < n; i++ {
		if unicode.IsDigit(rune(s[i])) {
			// 合法的表达式必须是以数字开头
			if currentStart == -1 {
				currentStart = i
			}
			valid = true
		} else if isOperator(rune(s[i])) {
			// 检查是不是第一个操作符，
			// 检查前一个是不是数字
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
	if startIndex == -1 {
		return 0
	}
	fmt.Println(startIndex, maxLength)
	return evaluateExpression(s[startIndex : startIndex+maxLength])
}

func isOperator(c rune) bool {
	return c == '+' || c == '-' || c == '*'
}

func evaluateExpression(expr string) int64 {
	var result, current int64
	var operator = '+'
	for i := 0; i < len(expr); {
		if unicode.IsDigit(rune(expr[i])) {
			start := i
			for i < len(expr) && unicode.IsDigit(rune(expr[i])) {
				i++
			}
			num, _ := strconv.ParseInt(expr[start:i], 10, 64)
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
			return 0
		}
	}
	result += current
	return result
}

// 解密犯罪时间
func CrimeTime(s1 string) string {
	maps := make(map[byte]struct{})
	maps[s1[0]] = struct{}{}
	maps[s1[1]] = struct{}{}
	maps[s1[2]] = struct{}{}
	maps[s1[3]] = struct{}{}
	maps[s1[4]] = struct{}{}
	H1 := s1[0:2]
	M1 := s1[3:5]
	h1, _ := strconv.Atoi(H1)
	m1, _ := strconv.Atoi(M1)
	bytes := make([]byte, 5)
	bytes[2] = ':'
loop1:
	for {
		m1++
		if m1 == 60 {
			m1 = 0
			h1++
			bytes[3] = '0'
			bytes[4] = '0'
		} else {
			m1Str := strconv.Itoa(m1)
			if m1 > 10 {
				bytes[3] = m1Str[0]
				bytes[4] = m1Str[1]
			} else {
				bytes[3] = '0'
				bytes[4] = m1Str[0]
			}

		}
		if h1 == 24 {
			h1 = 0
			bytes[0] = '0'
			bytes[1] = '0'
		} else {
			h1Str := strconv.Itoa(h1)
			if h1 > 10 {
				bytes[0] = h1Str[0]
				bytes[1] = h1Str[1]
			} else {
				bytes[0] = '0'
				bytes[1] = h1Str[0]
			}

		}
		for i := 0; i < 5; i++ {
			if _, ok := maps[bytes[i]]; !ok {
				continue loop1
			}
		}
		return string(bytes)
	}
}
