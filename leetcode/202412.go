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
