package main

import (
	"fmt"
	"sort"
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
/*
任务排序：我们按照任务的结束时间对任务进行排序。这样可以尽量早地完成任务，从而空出更多的时间给后续任务。
遍历任务：对于每一个任务，如果它的开始时间比当前天数大，则可以处理该任务，并更新当前天数为任务的结束时间。
调整条件：如果任务的开始时间小于或等于当前天数，但结束时间大于当前天数，则可以处理该任务，并将当前天数增加1。

*/
func maxTasks(tasks []Task) int {
	// 按照任务的结束时间进行排序
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].end < tasks[j].end
	})

	count := 0
	currentDay := 0

	// 遍历任务
	for _, task := range tasks {
		// 如果当前任务可以在 currentDay 处理
		if task.start > currentDay {
			count++
			currentDay = task.end
		} else if task.start <= currentDay && task.end > currentDay {
			count++
			currentDay++
		}
	}

	return count
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
