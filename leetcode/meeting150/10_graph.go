package meeting150

import "slices"

// 岛屿数量
func numIslands(grid [][]byte) int {
	cnt := 0
	m := len(grid)
	n := len(grid[0])
	var dfs func(grid [][]byte, i, j int)
	dfs = func(grid [][]byte, i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n {
			return
		}
		if grid[i][j] == '0' {
			return
		}
		grid[i][j] = '0'
		dfs(grid, i, j+1)
		dfs(grid, i, j-1)
		dfs(grid, i-1, j)
		dfs(grid, i+1, j)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				cnt++
				dfs(grid, i, j)
			}
		}
	}
	return cnt
}

func solve(board [][]byte) {
	m := len(board)
	n := len(board[0])
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n || board[i][j] != 'O' {
			return
		}
		board[i][j] = 'Q'
		dfs(i, j+1)
		dfs(i, j-1)
		dfs(i-1, j)
		dfs(i+1, j)
	}
	for i := 0; i < m; i++ {
		dfs(i, 0)
		dfs(i, n-1)
	}
	for j := 0; j < n; j++ {
		dfs(0, j)
		dfs(m-1, j)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
			if board[i][j] == 'Q' {
				board[i][j] = 'O'
			}
		}
	}
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses)

	for i := 0; i < len(prerequisites); i++ {
		from, to := prerequisites[i][1], prerequisites[i][0]
		graph[from] = append(graph[from], to)
	}

	var dfs func(i int, hasCycle *bool, visited *map[int]bool, onPath *[]bool)
	dfs = func(i int, hasCycle *bool, visited *map[int]bool, onPath *[]bool) {
		if *hasCycle {
			return
		}

		if (*onPath)[i] {
			*hasCycle = true
			return
		}
		if (*visited)[i] {
			return
		}
		(*visited)[i] = true
		(*onPath)[i] = true
		for _, to := range graph[i] {
			dfs(to, hasCycle, visited, onPath)
		}
		(*onPath)[i] = false
	}
	visited := make(map[int]bool)
	// 记录递归堆栈中的节点
	onPath := make([]bool, numCourses)
	hasCycle := false
	for i := 0; i < numCourses; i++ {
		dfs(i, &hasCycle, &visited, &onPath)
	}
	return !hasCycle
}

func findOrder(numCourses int, prerequisites [][]int) []int {
	graph := make([][]int, numCourses)

	for i := 0; i < len(prerequisites); i++ {
		from, to := prerequisites[i][1], prerequisites[i][0]
		graph[from] = append(graph[from], to)
	}
	postorder := make([]int, 0)
	visited := make(map[int]bool)
	// 记录递归堆栈中的节点
	onPath := make([]bool, numCourses)
	hasCycle := false
	var dfs func(i int)
	dfs = func(i int) {
		if hasCycle {
			return
		}

		if (onPath)[i] {
			hasCycle = true
			return
		}
		if (visited)[i] {
			return
		}
		(visited)[i] = true
		(onPath)[i] = true
		for _, to := range graph[i] {
			dfs(to)
		}
		postorder = append(postorder, i)
		(onPath)[i] = false
	}

	for i := 0; i < numCourses; i++ {
		dfs(i)
	}
	if hasCycle {
		return []int{}
	}
	slices.Reverse(postorder)
	return postorder
}
