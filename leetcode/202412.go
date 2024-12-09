package leetcode

func minMovesToCaptureTheQueen(a int, b int, c int, d int, e int, f int) int {

	// 先判断车和后是否在一条直线
	if a == e {
		if c != a {
			return 1
		} else {
			if !inBetween(b, d, f) {
				return 1
			}
		}

	}
	if b == f {
		if d != f {
			return 1
		} else {
			if !inBetween(a, c, e) {
				return 1
			}
		}

	}

	// 判断皇后和象是否在一条斜线
	if (c - e) == (d - f) {
		if (c - a) != (d - b) {
			return 1
		} else {
			if !inBetween(c, a, e) {
				return 1
			}
		}
	}

	if (c + d) == (e + f) {
		if (a + b) != (e + f) {
			return 1
		} else {
			if !inBetween(c, a, e) {
				return 1
			}
		}
	}
	return 2
}

func inBetween(l, m, r int) bool {
	return min(l, r) < m && m < max(l, r)
}

func numRookCaptures(board [][]byte) int {
	cnt := 0
	rRow := 0
	rCol := 0
loop1:
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 'R' {
				rRow = i
				rCol = j
				break loop1
			}
		}
	}

	for i := rRow + 1; i < 8; i++ {
		if board[i][rCol] == 'p' {
			cnt++
			break
		}
		if board[i][rCol] == 'B' {
			break
		}
	}
	for i := rRow - 1; i >= 0; i-- {
		if board[i][rCol] == 'p' {
			cnt++
			break
		}
		if board[i][rCol] == 'B' {
			break
		}
	}

	for j := rCol - 1; j >= 0; j-- {
		if board[rRow][j] == 'p' {
			cnt++
			break
		}
		if board[rRow][j] == 'B' {
			break
		}
	}
	for j := rCol + 1; j < 8; j++ {
		if board[rRow][j] == 'p' {
			cnt++
			break
		}
		if board[rRow][j] == 'B' {
			break
		}
	}
	return cnt
}

// https://leetcode.cn/problems/determine-color-of-a-chessboard-square/
func squareIsWhite(coordinates string) bool {
	bytes := []byte(coordinates)
	row := bytes[0] - 'a'
	col := bytes[1] - '0'
	if row%2 == 0 {
		if col%2 == 0 {
			return true
		} else {
			return false
		}
	} else {
		if col%2 == 0 {
			return false
		} else {
			return true
		}
	}
}
