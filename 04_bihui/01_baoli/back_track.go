package _1_baoli

import "strings"

func solveNQueens(n int) [][]string {
	res := [][]string{}
	board := make([]string, n)
	for i := 0; i < n; i++ {
		board[i] = strings.Repeat(".", n)
	}
	backtrack(board, 0, &res, n)
	return res
}

// 路径：board 中小于 row 的那些行都已经成功放置了皇后
// 选择列表：第 row 行的所有列都是放置皇后的选择
// 结束条件：row 超过 board 的最后一行
func backtrack(board []string, row int, res *[][]string, n int) {
	// 触发结束条件
	if row == n {
		newRow := make([]string, len(board))
		copy(newRow, board)
		*res = append(*res, newRow)
		return
	}
	for col := 0; col < n; col++ {
		// 排除不合法选择
		if !isValid(board, row, col) {
			continue
		}
		// 做选择
		newLine := []byte(board[row])
		newLine[col] = 'Q'
		board[row] = string(newLine)
		// 进入下一行决策
		backtrack(board, row+1, res, n)
		// 撤销选择
		newLine[col] = '.'
		board[row] = string(newLine)
	}
}

func isValid(board []string, row, col int) bool {

	// 因为皇后是一行一行从上往下放的，所以左下方，右下方和正下方不用检查（还没放皇后）；
	// 因为一行只会放一个皇后，所以每行不用检查。也就是最后只用检查上面，左上，右上三个方向。
	n := len(board)
	// 检查列是否有皇后冲突
	for i := 0; i < n; i++ {
		if board[i][col] == 'Q' {
			return false
		}
	}
	// 检查右上方是否有皇后冲突
	for i, j := row-1, col+1; i >= 0 && j < n; {
		if board[i][j] == 'Q' {
			return false
		}
		i--
		j++
	}
	// 检查左上方是否有皇后冲突
	for i, j := row-1, col-1; i >= 0 && j >= 0; {
		if board[i][j] == 'Q' {
			return false
		}
		i--
		j--
	}
	return true
}

// https://leetcode.cn/problems/sudoku-solver/description/
func solveSudoku(board [][]byte) {
	solveSudokuBackTrack(board, 0, 0, 9)
}

func solveSudokuBackTrack(
	board [][]byte, // 棋盘
	row int, // 行数
	col int,
	n int, // n
) bool {

	// 列超过9，跳到下一行
	if col == n {
		return solveSudokuBackTrack(board, row+1, 0, n)
	}

	// 回溯结束
	if row == n {
		return true
	}

	// 本身就是数字，跳到下一列
	if board[row][col] != '.' {
		return solveSudokuBackTrack(board, row, col+1, n)
	}

	for ch := '1'; ch <= '9'; ch++ {
		// 如果遇到不合法的数字，就跳过
		if !check(board, row, col, byte(ch)) {
			continue
		}

		board[row][col] = byte(ch)
		// 如果找到一个可行解，立即结束
		if solveSudokuBackTrack(board, row, col+1, n) {
			return true
		}
		board[row][col] = '.'
	}
	// 穷举完 1~9，依然没有找到可行解，此路不通
	return false

}

func check(board [][]byte, row int, col int, n byte) bool {
	for i := 0; i < 9; i++ {
		// 判断行是否存在重复
		if board[row][i] == n {
			return false
		}
		// 判断列是否存在重复
		if board[i][col] == n {
			return false
		}

		// 判断 3 x 3 方框是否存在重复
		/*
			(r/3)*3 & (c/3)*3 求的是行和列的偏移量，
			对于r和c来说，是3的几倍就偏移几倍的3格 (注意整型除法会先向下取整）
			因为 i=0 -> 8， 所以 i / 3 在每一次循环的变化是 0, 0, 0, 1, 1, 1, 2, 2, 2
			因为 i=0 -> 8， 所以 i % 3 在每一次循环的变化是 0, 1, 2, 0, 1, 2, 0, 1, 2
		*/
		if board[(row/3)*3+i/3][(col/3)*3+i%3] == n {
			return false
		}
	}
	return true
}
