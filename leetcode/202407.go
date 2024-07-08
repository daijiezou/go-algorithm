package leetcode

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
