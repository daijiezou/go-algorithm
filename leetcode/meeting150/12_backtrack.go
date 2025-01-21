package meeting150

import (
	"strconv"
)

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	maps := map[string][]string{
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"},
		"9": {"w", "x", "y", "z"},
	}
	res := []string{}
	n := len(digits)
	var back func(start int, pre string)
	back = func(start int, pre string) {
		if start == n {
			res = append(res, pre)
			return
		}
		for _, v := range maps[strconv.Itoa(int(digits[start]-'0'))] {
			pre += v
			back(start+1, pre)
			pre = pre[:len(pre)-1]
		}
	}
	back(0, "")
	return res
}

// 组合
func combine(n int, k int) [][]int {
	res := [][]int{}
	var backtrack func(pre []int, start int)
	backtrack = func(pre []int, start int) {
		if len(pre) == k {
			temp := make([]int, k)
			copy(temp, pre)
			res = append(res, temp)
			return
		}
		for i := start; i <= n; i++ {
			pre = append(pre, i)
			backtrack(pre, i+1)
			pre = pre[:len(pre)-1]
		}
	}
	backtrack([]int{}, 1)
	return res
}

// 全排列
func permute(nums []int) [][]int {
	res := [][]int{}
	n := len(nums)
	selectd := make(map[int]bool)
	var backtrack func(pre []int)
	backtrack = func(pre []int) {
		if len(pre) == n {
			temp := make([]int, n)
			copy(temp, pre)
			res = append(res, temp)
			return
		}
		for i := 0; i < n; i++ {
			if selectd[nums[i]] {
				continue
			}
			selectd[nums[i]] = true
			pre = append(pre, nums[i])
			backtrack(pre)
			selectd[nums[i]] = false
			pre = pre[:len(pre)-1]
		}
	}
	backtrack([]int{})
	return res
}

func combinationSum(candidates []int, target int) [][]int {
	n := len(candidates)
	var backtrack func(pre []int, preSum int, start int)
	res := [][]int{}
	//reDup := make(map[[41]int]struct{})
	backtrack = func(pre []int, preSum int, start int) {
		if preSum == target {
			//key := genKey(pre)
			//if _, ok := reDup[genKey(pre)]; ok {
			//	return
			//}
			//reDup[key] = struct{}{}
			temp := make([]int, len(pre))
			copy(temp, pre)
			res = append(res, temp)
			return
		}
		if preSum > target {
			return
		}
		for i := start; i < n; i++ {
			preSum += candidates[i]
			pre = append(pre, candidates[i])
			backtrack(pre, preSum, i)
			pre = pre[:len(pre)-1]
			preSum -= candidates[i]
		}
	}
	backtrack([]int{}, 0, 0)
	return res
}

func genKey(nums []int) [41]int {
	keyNums := [41]int{}
	for i := 0; i < len(nums); i++ {
		keyNums[nums[i]]++
	}
	return keyNums
}

// n皇后问题
func totalNQueens(n int) int {
	board := make([][]int, n)
	for i := 0; i < n; i++ {
		board[i] = make([]int, n)
	}
	var backtrack func(depth int)
	res := 0
	backtrack = func(depth int) {
		if depth >= n {
			res++
			return
		}
		for i := 0; i < n; i++ {
			board[depth][i] = 1
			if !checkNQueue(depth, board, i, n) {
				board[depth][i] = 0
				continue
			}
			backtrack(depth + 1)
			board[depth][i] = 0
		}
	}
	backtrack(0)
	return res
}

func checkNQueue(depth int, board [][]int, row int, n int) bool {
	for i := 0; i < depth; i++ {
		if board[i][row] == 1 {
			return false
		}
	}

	curDep := depth - 1
	curRow := row - 1
	for curDep >= 0 && curDep < n && curRow >= 0 && curRow < n {
		if board[curDep][curRow] == 1 {
			return false
		}
		curDep--
		curRow--
	}
	curDep = depth - 1
	curRow = row + 1
	for curDep >= 0 && curDep < n && curRow >= 0 && curRow < n {
		if board[curDep][curRow] == 1 {
			return false
		}
		curDep--
		curRow++
	}
	return true
}

// 括号生成
func generateParenthesis(n int) []string {
	res := []string{}
	var backtrack func(pre string, left, right int)
	backtrack = func(pre string, left, right int) {
		if n-left > n-right {
			return
		}
		if left > n || right > n {
			return
		}
		if left == n && right == n {
			res = append(res, pre)
			return
		}
		pre += "("
		backtrack(pre, left+1, right)
		pre = pre[:len(pre)-1]
		pre += ")"
		backtrack(pre, left, right+1)
		pre = pre[:len(pre)-1]
	}
	backtrack("", 0, 0)
	return res
}

func checkParenthesis(s string) bool {
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, s[i])
		} else {
			if len(stack) < 1 {
				return false
			}
			pop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if pop != '(' {
				return false
			}
		}
	}
	return len(stack) == 0
}

func exist(board [][]byte, word string) bool {
	res := false
	m := len(board)
	n := len(board[0])
	poss := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	used := make(map[[2]int]bool)
	var backtrack func(i, j int, start int)
	backtrack = func(i, j int, start int) {
		if res {
			return
		}
		if start == len(word) {
			res = true
			return
		}
		for _, pos := range poss {
			newX := i + pos[0]
			newY := j + pos[1]
			if used[[2]int{newX, newY}] {
				continue
			}
			if 0 <= newX && newX < m && newY >= 0 && newY < n && board[newX][newY] == word[start] {
				used[[2]int{newX, newY}] = true
				backtrack(newX, newY, start+1)
				used[[2]int{newX, newY}] = false
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == word[0] {
				used[[2]int{i, j}] = true
				backtrack(i, j, 1)
				used[[2]int{i, j}] = false
				if res {
					return true
				}
			}
		}
	}
	return res
}
