package meeting150

import "fmt"

// 矩阵

func isValidSudoku(board [][]byte) bool {
	var rows, columns [9][9]int
	var subboxes [3][3][9]int
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			index := board[i][j] - '1'
			rows[i][index]++
			columns[j][index]++
			subboxes[i/3][j/3][index]++
			if rows[i][index] > 1 || columns[j][index] > 1 || subboxes[i/3][j/3][index] > 1 {
				return false
			}
		}
	}

	return true
}

func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	n := len(matrix[0])
	res := make([]int, 0)
	upper := 0
	lower := m - 1
	leftBound := 0
	rightBound := n - 1
	for len(res) < m*n {
		if upper <= lower {
			for j := leftBound; j <= rightBound; j++ {
				res = append(res, matrix[upper][j])
			}
			upper++
		}
		if leftBound <= rightBound {
			for j := upper; j <= lower; j++ {
				res = append(res, matrix[j][rightBound])
			}
			rightBound--
		}
		if upper <= lower {
			for j := rightBound; j >= leftBound; j-- {
				res = append(res, matrix[lower][j])
			}
			lower--
		}
		if leftBound <= rightBound {
			for j := lower; j >= upper; j-- {
				res = append(res, matrix[j][leftBound])
			}
			leftBound++
		}
	}
	return res
}

// 旋转图像
func rotate(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])
	for i := 0; i < m; i++ {
		//note: 这里是j=i
		for j := i; j < n; j++ {
			tmep := matrix[i][j]
			matrix[i][j] = matrix[j][i]
			matrix[j][i] = tmep
			//matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	fmt.Println(matrix)
	for i := 0; i < m; i++ {
		for j := 0; j < n/2; j++ {
			matrix[i][j], matrix[i][n-j-1] = matrix[i][n-j-1], matrix[i][j]
		}
	}

}

func setZeroes(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])

	zeroCol := make([]bool, m)
	zeroRow := make([]bool, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				zeroCol[i] = true
				zeroRow[j] = true
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if zeroCol[i] || zeroRow[j] {
				matrix[i][j] = 0
			}
		}
	}
}

func gameOfLife(board [][]int) {
	poss := [][2]int{{-1, 0}, {-1, -1}, {0, -1}, {1, -1}, {1, 0}, {1, 1}, {0, 1}, {-1, 1}}
	m := len(board)
	n := len(board[0])
	life := [][2]int{}
	die := [][2]int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			lifeCnt := 0
			for _, pos := range poss {
				dx := i + pos[0]
				dy := j + pos[1]
				if 0 <= dx && dx < m && dy >= 0 && dy < n {
					if board[dx][dy] == 1 {
						lifeCnt++
					}
				}
			}
			if board[i][j] == 1 {
				if lifeCnt < 2 || lifeCnt > 3 {
					die = append(die, [2]int{i, j})
				}
			} else {
				if lifeCnt == 3 {
					life = append(life, [2]int{i, j})
				}
			}
		}
	}
	for _, pos := range life {
		board[pos[0]][pos[1]] = 1
	}
	for _, pos := range die {
		board[pos[0]][pos[1]] = 0
	}
}
