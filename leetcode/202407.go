package leetcode

import (
	"math/bits"
	"sort"
	"strconv"
	"strings"
)

// https://leetcode.cn/problems/check-if-move-is-legal/
func checkMove(board [][]byte, rMove int, cMove int, color byte) bool {
	// 从y轴正方向开始遍历
	//上、右上、右、右下、下、左下、左、左上
	dxs := []int{0, 1, 1, 1, 0, -1, -1, -1}
	dys := []int{-1, -1, 0, 1, 1, 1, 0, -1}
	for i := 0; i < 8; i++ {
		// 检查8个方向
		if checkGood(board, rMove, cMove, color, dxs[i], dys[i]) {
			return true
		}
	}
	return false
}

func checkGood(board [][]byte, rMove int, cMove int, color byte, dx, dy int) bool {
	x := rMove + dx
	y := cMove + dy
	step := 1
	for x >= 0 && x < 8 && y >= 0 && y < 8 {
		//第一步必须是其他颜色
		if step == 1 {
			if board[x][y] == color || board[x][y] == '.' {
				return false
			}
		} else {
			//中间不能有空棋盘
			if board[x][y] == '.' {
				return false
			}
			// 遍历到了终点
			if board[x][y] == color {
				return true
			}
		}
		x += dx
		y += dy
		step++
	}
	return false
}

// https://leetcode.cn/problems/find-pivot-index/description/?envType=daily-question&envId=2024-07-08
func pivotIndex(nums []int) int {
	length := len(nums)
	presum := make([]int, length+1)

	for i := 1; i < length+1; i++ {
		presum[i] = presum[i-1] + nums[i-1]
	}
	for i := 1; i < length+1; i++ {
		// 计算 nums[i-1] 左侧和右侧的元素和
		left := presum[i-1]
		right := presum[length] - presum[i]
		if left == right {
			return i
		}
	}
	return -1
}

// https://leetcode.cn/problems/count-the-number-of-incremovable-subarrays-i/?envType=daily-question&envId=2024-07-10
// 找到有几个递增子数组
func incremovableSubarrayCount(nums []int) int {
	current := []int{}
	count := 0
	currentIndex := make([]bool, len(nums))
	incremovableSubarrayCountBacktack(nums, 0, current, currentIndex, &count)
	return count + 1
}

func incremovableSubarrayCountBacktack(nums []int, start int, zijihe []int, index []bool, count *int) {
	for i := start; i < len(nums); i++ {
		zijihe = append(zijihe, nums[i])
		index[i] = true
		// 判断是否为递增子数组
		if heckIncremovable(index, zijihe) {
			*count++
		}
		incremovableSubarrayCountBacktack(nums, i+1, zijihe, index, count)
		zijihe = zijihe[:len(zijihe)-1]
		index[i] = false
	}
}

func heckIncremovable(index []bool, nums []int) bool {
	if len(nums) == len(index) {
		return false
	}
	leftIndex := make([]int, 0)
	for i := 0; i < len(index); i++ {
		if !index[i] {
			leftIndex = append(leftIndex, i)
		}
	}

	// 判断index是否连续
	for i := 0; i < len(leftIndex)-1; i++ {
		if leftIndex[i]+1 != leftIndex[i+1] {
			return false
		}
	}
	if len(nums) == 1 {
		return true
	}
	length := len(nums)
	for i := 0; i < length-1; i++ {
		if nums[i] >= nums[i+1] {
			return false
		}
	}

	return true
}

func numberGame(nums []int) []int {
	sort.Ints(nums)
	for i := 0; i < len(nums); i += 2 {
		nums[i], nums[i+1] = nums[i+1], nums[i]
	}
	return nums
}

// https://leetcode.cn/problems/find-if-array-can-be-sorted/description/?envType=daily-question&envId=2024-07-13
func canSortArray(nums []int) bool {
	oneCount := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		res := strconv.FormatInt(int64(nums[i]), 2)
		oneCount[i] = strings.Count(res, "1")
	}
	temp := make([][2]int, 0)
	left, right := 0, 0
	for right < len(nums) {
		if oneCount[right] == oneCount[left] {
			right++
			if right < len(nums) && oneCount[right] != oneCount[left] {
				temp = append(temp, [2]int{left, right})
			}
		} else {
			left++
		}
	}
	temp = append(temp, [2]int{left, right})
	numMax, _ := getNumsMaxAndMin(nums, temp[0])
	for i := 1; i < len(temp); i++ {
		tempMax, tempMin := getNumsMaxAndMin(nums, temp[i])
		if tempMin < numMax {
			return false
		}
		numMax = tempMax
	}
	return true
}

func getNumsMaxAndMin(nums []int, index2 [2]int) (max, min int) {

	left := index2[0]
	right := index2[1]
	numMax := nums[left]
	numMin := nums[left]
	for i := left + 1; i < right; i++ {
		if nums[i] > numMax {
			numMax = nums[i]
		}
		if nums[i] < numMin {
			numMin = nums[i]
		}
	}
	return numMax, numMin
}

// https://leetcode.cn/problems/find-if-array-can-be-sorted/description/?envType=daily-question&envId=2024-07-13
func canSortArray2(nums []int) bool {
	//当前组的最大值
	currentGroupMax := 0

	// 当前1的个数
	latestOneCnt := 0

	// 上一组的最大值
	lastGroupMax := 0
	for i := 0; i < len(nums); i++ {
		if bits.OnesCount(uint(nums[i])) == latestOneCnt {
			currentGroupMax = max(currentGroupMax, nums[i])
		} else {
			// 更新最新的1的个数
			latestOneCnt = bits.OnesCount(uint(nums[i]))

			// 将当前组的最大值赋予上一组
			lastGroupMax = currentGroupMax

			// 更新当前组的最大值
			currentGroupMax = nums[i]
		}

		// 后面组的每个都必须大于上个组的最大值，否则无法排序
		if nums[i] < lastGroupMax {
			return false
		}
	}
	return true
}

// https://leetcode.cn/problems/max-increase-to-keep-city-skyline/?envType=daily-question&envId=2024-07-14
func maxIncreaseKeepingSkyline(grid [][]int) int {
	// 找到每一个点所在行列的最大值，
	// 这个最多比这个点所在的行列最大值要小
	m := len(grid)
	n := len(grid[0])
	hangMax := make([]int, m)
	lieMax := make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			hangMax[i] = max(hangMax[i], grid[i][j])
			lieMax[j] = max(lieMax[j], grid[i][j])
		}
	}
	totalCount := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			minHeight := min(hangMax[i], lieMax[j])
			if grid[i][j] < minHeight {
				totalCount += minHeight - grid[i][j]
			}
		}
	}
	return totalCount
}

// https://leetcode.cn/problems/accounts-merge/
func accountsMerge(accounts [][]string) [][]string {
	emailId := make(map[string]int)
	emailName := make(map[string]string)
	for i := 0; i < len(accounts); i++ {
		name := accounts[i][0]
		for j := 1; j < len(accounts[i]); j++ {
			if _, ok := emailId[accounts[i][j]]; !ok {
				emailId[accounts[i][j]] = i
				emailName[accounts[i][j]] = name
			}
		}
	}

}
