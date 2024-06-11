package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
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
	var avg int
	fmt.Scan(&avg)
	scanner := bufio.NewScanner(os.Stdin)
	nums := make([]string, 0)
	for scanner.Scan() {
		nums = strings.Split(scanner.Text(), ",")

	}
	numsInt := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		numsInt[i], _ = strconv.Atoi(nums[i])
	}
	return
}

/*
// https://www.nowcoder.com/discuss/626731698809479168?sourceSSR=users
查找接口成功率最优时间段
找出平均值小于等于minAverageLost的最长时间段，输出数组下标对，格式{beginIndex}-{endIndx}(下标从0开始)，
如果同时存在多个最长时间段，则输出多个下标对且下标对之间使用空格(” “)拼接，多个下标对按下标从小到大排序。
*/
func Get(nums []int, avg float64) []string {
	left := 0
	right := 0
	windowSum := 0
	resultMap := make(map[int]string)
	for right < len(nums) {
		current := nums[right]
		right++
		windowSum += current
		windowAvg := float64(windowSum) / float64(right-left)
		if windowAvg <= avg && (right-left > 1) {
			resultMap[left] = strconv.Itoa(left) + "-" + strconv.Itoa(right-1)
		}
		for windowAvg > avg && left < right {
			windowSum -= nums[left]
			left++
		}

	}
	res := make([]string, 0, len(resultMap))
	resKey := make([]int, 0, len(resultMap))
	for k, _ := range resultMap {
		resKey = append(resKey, k)
	}
	sort.Ints(resKey)
	for _, k := range resKey {
		res = append(res, resultMap[k])
	}
	return res
}

type MyMap struct {
	key   int
	count int
	next  *MyMap
}

func FromXiaoqu(nums []int) int {
	maps := make(map[int]*MyMap)
	for _, num := range nums {
		if mymap, ok := maps[num]; ok {
			for mymap.next != nil {
				mymap = mymap.next
			}
			if mymap.count >= num {
				mymap.next = &MyMap{
					key:   num,
					count: 1,
					next:  nil,
				}
			} else {
				mymap.count++
			}
		} else {
			maps[num] = &MyMap{
				key:   num,
				count: 1,
				next:  nil,
			}
		}
	}
	totalCount := 0
	for _, mymap := range maps {
		totalCount += mymap.key + 1
		for mymap.next != nil {
			mymap = mymap.next
			totalCount += mymap.key + 1
		}
	}
	return totalCount
}
