package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

func main() {

}

func isOperator(c rune) bool {
	return c == '+' || c == '-' || c == '*'
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

// https://www.nowcoder.com/discuss/628594841575809024?sourceSSR=users
// 最多几个直角三角形
func countRightTriangles(lengths []int) int {
	n := len(lengths)
	count := 0
	triangleSet := make(map[[3]int]struct{})
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				a, b, c := lengths[i], lengths[j], lengths[k]
				if isRightTriangle(a, b, c) {
					triangle := [3]int{a, b, c}
					if _, ok := triangleSet[triangle]; !ok {
						triangleSet[triangle] = struct{}{}
						count++
					}
				}
			}
		}
	}

	return count
}

func isRightTriangle(a, b, c int) bool {
	// Check for a^2 + b^2 = c^2
	return a*a+b*b == c*c
}

type ExcelCell struct {
	Origin     string
	startIndex int
	referIndex int
	output     string
}

// https://www.nowcoder.com/discuss/628538017472344064?sourceSSR=users
func HandleExcel(cells []string) {
	excelCells := make([]*ExcelCell, len(cells))
	for i, cell := range cells {
		excelCell := new(ExcelCell)
		refIdx, startIndex := getIndex(cell)
		excelCell.Origin = cell
		excelCell.referIndex = refIdx
		excelCell.startIndex = startIndex
		if refIdx < 0 {
			excelCell.output = cell
		}
		excelCells[i] = excelCell
	}
	res := ""
	for i := 0; i < len(excelCells); i++ {
		out := getOut(excelCells, i)
		res += out
		if i < len(excelCells)-1 {
			res += ","
		}
	}
	fmt.Println(res)
}

func getIndex(cell string) (int, int) {
	idx := strings.Index(cell, "<")
	// shuo
	if idx >= 0 {
		u := cell[idx+1]
		return int(u - 'A'), idx
	} else {
		return -1, idx
	}
}

func getOut(cells []*ExcelCell, i int) (output string) {
	if cells[i].referIndex != -1 {
		res := getOut(cells, cells[i].referIndex)
		cells[i].output = strings.Replace(cells[i].Origin, cells[i].Origin[cells[i].startIndex:cells[i].startIndex+3], res, 1)
		cells[i].referIndex = -1
		return cells[i].output
	} else {
		return cells[i].output
	}
}

/*
前缀和数组：构建 prefixSum 数组，用于快速计算任意子数组的和。在处理环形数组时，数组长度扩展为两倍，以便处理跨越数组末尾和起点的子数组。
动态规划数组：dp[i][j] 表示从 i 到 j 子数组中，吃货能获得的最大披萨块总和。
状态转移：
对于长度为 1 的子数组，直接等于对应的披萨块大小。
对于长度大于 1 的子数组，使用状态转移方程 dp[l][r] = totalSum - min(dp[l+1][r], dp[l][r-1])，其中 totalSum 是子数组的总和。这个方程考虑了吃货的最佳选择策略，以及馋嘴的最优选择策略。
寻找最佳起点：由于披萨块是环形的，我们需要从每个可能的起点计算最大值，因此遍历所有起点 i 并计算 dp[i][i+n-1]，找到其中的最大值。
*/
func maxPizzaSum(pizzaSizes []int) int {
	n := len(pizzaSizes)
	if n == 1 {
		return pizzaSizes[0]
	}

	// 构建前缀和数组
	prefixSum := make([]int, 2*n+1)
	for i := 1; i <= 2*n; i++ {
		prefixSum[i] = prefixSum[i-1] + pizzaSizes[(i-1)%n]
	}

	// dp[i][j]表示从i到j这段内吃货能拿到的最大披萨块总和
	dp := make([][]int, 2*n)
	for i := range dp {
		dp[i] = make([]int, 2*n)
	}

	// 枚举长度
	for length := 1; length <= n; length++ {
		// 枚举左端点
		for l := 0; l < 2*n; l++ {
			r := l + length - 1
			if r >= 2*n {
				continue
			}
			totalSum := prefixSum[r+1] - prefixSum[l]
			if length == 1 {
				// 对于长度为 1 的子数组，直接等于对应的披萨块大小。
				dp[l][r] = pizzaSizes[l%n]
			} else {
				// 使用状态转移方程 dp[l][r] = totalSum - min(dp[l+1][r], dp[l][r-1])，
				// 其中 totalSum 是子数组的总和。这个方程考虑了吃货的最佳选择策略，以及馋嘴的最优选择策略。
				dp[l][r] = totalSum - min(dp[l+1][r], dp[l][r-1])
			}
		}
	}

	// 寻找最大值
	maxSum := 0
	for i := 0; i < n; i++ {
		maxSum = max(maxSum, dp[i][i+n-1])
	}

	return maxSum
}
