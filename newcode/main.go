package newcode

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

/*
定义一个单词的“兄弟单词”为：交换该单词字母顺序（注：可以交换任意次），而不添加、删除、修改原有的字母就能生成的单词。

兄弟单词要求和原来的单词不同。例如：ab和ba是兄弟单词。ab和ab则不是兄弟单词。

现在给定你n个单词，另外再给你一个单词str，让你寻找str的兄弟单词里，按字典序排列后的第k个单词是什么？

注意：字典中可能有重复单词。本题含有多组输入数据。

输入描述：先输入单词的个数n，再输入n个单词。 再输入一个单词，为待查找的单词x 最后输入数字k

输出描述：输出查找到x的兄弟单词的个数m 然后输出查找到的按照字典顺序排序后的第k个兄弟单词，没有符合第k个的话则不用输出。
*/
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

// 雨花石问题
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
	fmt.Println(s[startIndex : startIndex+maxLength])
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

// https://www.nowcoder.com/discuss/626732760618590208?sourceSSR=users
/*
在一个机房中，服务器的位置标识在 n*m 的整数矩阵网格中，1 表示单元格上有服务器，0 表示没有。如果两台服务器位于同一行或者同一列中紧邻的位置，则认为它们之间可以组成一个局域网。

请你统计机房中最大的局域网包含的服务器个数。
*/
func Wangluofuwuqi(nums [][]int) int {
	visited := make(map[[2]int]struct{})
	res := make(map[[2]int]int)
	m := len(nums)
	n := len(nums[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if nums[i][i] == 1 {
				visited[[2]int{i, j}] = struct{}{}
				res[[2]int{i, j}] = GetLan(nums, i, j, visited)
			}
		}
	}
	maxRes := 0
	for _, v := range res {
		if v > maxRes {
			maxRes = v
		}
	}
	return maxRes
}

func GetLan(nums [][]int, i, j int, visted map[[2]int]struct{}) int {
	if nums[i][j] == 0 {
		return 0
	}
	res := 1
	if i > 0 {
		if _, ok := visted[[2]int{i - 1, j}]; !ok {
			visted[[2]int{i - 1, j}] = struct{}{}
			res += GetLan(nums, i-1, j, visted)
		}
	}
	if i < len(nums[0])-1 {
		if _, ok := visted[[2]int{i + 1, j}]; !ok {
			visted[[2]int{i + 1, j}] = struct{}{}
			res += GetLan(nums, i+1, j, visted)
		}
	}
	if j > 0 {
		if _, ok := visted[[2]int{i, j - 1}]; !ok {
			visted[[2]int{i, j - 1}] = struct{}{}
			res += GetLan(nums, i, j-1, visted)
		}
	}
	if j < len(nums)-1 {
		if _, ok := visted[[2]int{i, j + 1}]; !ok {
			visted[[2]int{i, j + 1}] = struct{}{}
			res += GetLan(nums, i, j+1, visted)
		}
	}
	return res
}

func largestServerNetwork(grid [][]int) int {
	// 定义方向：上、下、左、右
	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	// 检查是否在网格范围内
	inBounds := func(x, y int) bool {
		return x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0])
	}
	visited := make(map[[2]int]bool)
	var bfs func(x, y int) int
	bfs = func(x, y int) int {
		queue := [][2]int{{x, y}}
		visited[[2]int{x, y}] = true
		count := 0
		for len(queue) > 0 {
			current := queue[0]
			cx, cy := current[0], current[1]
			queue = queue[1:]
			count++
			for _, direction := range directions {
				nx, ny := cx+direction[0], cy+direction[1]
				if inBounds(nx, ny) && !visited[[2]int{nx, ny}] && grid[nx][ny] == 1 {
					queue = append(queue, [2]int{nx, ny})
					visited[[2]int{nx, ny}] = true
				}
			}
		}
		return count
	}
	m := len(grid)
	n := len(grid[0])
	// 统计最大的局域网
	maxNetworkSize := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 && !visited[[2]int{i, j}] {
				maxNetworkSize = max(maxNetworkSize, bfs(i, j))
			}
		}
	}

	return maxNetworkSize

}
