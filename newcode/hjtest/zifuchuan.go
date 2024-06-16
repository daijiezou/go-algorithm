package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

// https://www.nowcoder.com/discuss/564543169094799360?sourceSSR=users

func findLastValidIndex(S, L string) int {
	sLen, lLen := len(S), len(L)
	i, j := 0, 0
	lastIndex := -1

	for i < sLen && j < lLen {
		if S[i] == L[j] {
			lastIndex = j
			i++
		}
		j++
	}

	if i == sLen {
		return lastIndex
	}

	return -1
}

// https://www.nowcoder.com/discuss/583953228107108352?sourceSSR=users
func countPeaks(heights []int) int {
	n := len(heights)
	if n == 0 {
		return 0
	}

	peakCount := 0

	for i := 0; i < n; i++ {
		if i == 0 {
			// 第一个元素
			if n > 1 && heights[i] > heights[i+1] {
				peakCount++
			}
		} else if i == n-1 {
			// 最后一个元素
			if heights[i] > heights[i-1] {
				peakCount++
			}
		} else {
			// 中间元素
			if heights[i] > heights[i-1] && heights[i] > heights[i+1] {
				peakCount++
			}
		}
	}

	return peakCount
}

// https://www.nowcoder.com/discuss/597000574252400640?sourceSSR=users
/*
使用递归和回溯的思想来生成不同的字符串。具体的逻辑如下：
*/
func countValidStrings(characters []byte, N int) int {
	if N == 0 {
		return 1
	}
	if len(characters) == 0 {
		return 0
	}

	results := make(map[string]struct{})

	var backtrack func(current []byte, used []bool)
	backtrack = func(current []byte, used []bool) {
		if len(current) == N {
			results[string(current)] = struct{}{}
			return
		}

		for i := 0; i < len(characters); i++ {
			if used[i] || (len(current) > 0 && characters[i] == current[len(current)-1]) {
				continue
			}
			used[i] = true
			current = append(current, characters[i])
			backtrack(current, used)
			current = current[:len(current)-1]
			used[i] = false
		}
	}

	backtrack([]byte{}, make([]bool, len(characters)))
	return len(results)
}

// https://www.nowcoder.com/discuss/597839627168378880?sourceSSR=users
func CalOne() {
	// 读取初始信息
	var w, h, x, y, sx, sy, t int
	fmt.Scan(&w, &h, &x, &y, &sx, &sy, &t)

	// 读取矩阵信息
	matrix := make([][]int, h)
	for i := 0; i < h; i++ {
		matrix[i] = make([]int, w)
		var row string
		fmt.Scan(&row)
		for j := 0; j < w; j++ {
			matrix[i][j] = int(row[j] - '0')
		}
	}

	// 初始化计数器
	count := 0

	// 模拟运动
	for i := 0; i <= t; i++ {
		// 统计当前点是否是1
		if matrix[y][x] == 1 {
			count++
		}

		// 更新位置
		x += sx
		y += sy

		// 检查边界反射
		if x < 0 {
			x = 0
			sx = -sx
		}
		if x >= w {
			x = w - 1
			sx = -sx
		}
		if y < 0 {
			y = 0
			sy = -sy
		}
		if y >= h {
			y = h - 1
			sy = -sy
		}
	}
}

// https://www.nowcoder.com/discuss/597841686588366848?sourceSSR=users
func countPaths(grid [][]int, rows, cols int) int {
	// 检查起点和终点是否可达
	if grid[0][0] == 1 || grid[rows-1][cols-1] == 1 {
		return 0
	}

	dp := make([][]int, rows)
	for i := range dp {
		dp[i] = make([]int, cols)
	}

	// 初始化起点
	dp[0][0] = 1

	// 初始化第一列
	for i := 1; i < rows; i++ {
		if grid[i][0] == 0 {
			dp[i][0] = dp[i-1][0]
		}
	}

	// 初始化第一行
	for j := 1; j < cols; j++ {
		if grid[0][j] == 0 {
			dp[0][j] = dp[0][j-1]
		}
	}

	// 填充dp数组
	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			if grid[i][j] == 0 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}

	return dp[rows-1][cols-1]
}

type Task struct {
	start, end int
}

// https://www.nowcoder.com/discuss/599560893734690816?sourceSSR=users
func maxTask2(tasks []Task) int {
	// 按照任务的结束时间进行排序
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].end < tasks[j].end
	})

	// Initialize a map to keep track of the days
	days := make(map[int]bool)

	// Try to fit each task into the earliest possible day
	maxTaskCnt := 0
	for _, task := range tasks {
		for day := task.start; day <= task.end; day++ {
			if !days[day] {
				days[day] = true
				maxTaskCnt++
				break
			}
		}
	}
	return maxTaskCnt
}

func maxBananas(numbers []int, n int) int {
	length := len(numbers)
	var maxSum int
	for i := 0; i < n; i++ {
		sum := 0
		for left := 0; left < i; left++ {
			sum += numbers[left]
		}
		for right := 0; right < n-1; right++ {
			sum += numbers[length-1-right]
		}
		if sum > maxSum {
			maxSum = sum
		}
	}
	return maxSum
}

func summarize(s string) string {
	// 去除字符串中的非字母符号
	var cleaned strings.Builder
	for _, char := range s {
		if unicode.IsLetter(char) {
			cleaned.WriteRune(char)
		}
	}
	s = cleaned.String()

	s = strings.ToLower(s) // 将输入字符串转换为小写
	runes := []rune(s)     // 将字符串转换为 Unicode 字符数组
	var summary []string   // 创建空的摘要字符串数组

	i := 0
	for i < len(runes) {
		count := 1
		j := i + 1
		for j < len(runes) && runes[i] == runes[j] {
			count++
			j++
		}

		if count > 1 { // 如果字符连续出现超过1次
			summary = append(summary, fmt.Sprintf("%c%d", runes[i], count)) // 添加字符及其连续出现次数到摘要数组
			i = j
		} else { // 如果字符不连续出现
			nextCount := strings.Count(s[j:], string(runes[i]))                 // 计算该字符在后续字符串中出现的次数
			summary = append(summary, fmt.Sprintf("%c%d", runes[i], nextCount)) // 添加字符及其后续出现次数到摘要数组
			i++
		}
	}

	sort.Slice(summary, func(i, j int) bool { // 对摘要数组进行排序
		if len(summary[i]) == len(summary[j]) {
			return summary[i] < summary[j]
		}
		return len(summary[i]) > len(summary[j])
	})

	return strings.Join(summary, "") // 将排序后的摘要数组连接成摘要字符串并返回
}

func countLuckyNumbers(k, n, m int) int {
	// 将十进制的 k 转换为 m 进制的字符串
	kStr := strconv.FormatInt(int64(k), m)

	// 计算幸运数字在 kStr 中的出现次数
	count := 0
	for _, digit := range kStr {
		if int(digit-'0') == n {
			count++
		}
	}

	return count
}

type Task2 struct {
	SLA, Value int
}

func getMaxValue(tasks []Task2, T int) int {
	// 根据任务的价值（积分）进行排序，降序
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].Value > tasks[j].Value
	})

	timeUsed := 0
	totalValue := 0
	for _, task := range tasks {
		if timeUsed < T && timeUsed < task.SLA {
			totalValue += task.Value
			timeUsed++
		}
	}

	return totalValue
}
