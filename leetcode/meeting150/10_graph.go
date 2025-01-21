package meeting150

import (
	"slices"
)

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

type EdgeNode struct {
	Node   string
	Weight float64
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	graph := make(map[string][]EdgeNode)
	for i := 0; i < len(equations); i++ {
		from := equations[i][0]
		to := equations[i][1]
		if _, ok := graph[from]; !ok {
			graph[from] = make([]EdgeNode, 0)
		}
		graph[from] = append(graph[from], EdgeNode{
			Node:   to,
			Weight: values[i],
		})
		if _, ok := graph[to]; !ok {
			graph[to] = make([]EdgeNode, 0)
		}
		graph[to] = append(graph[to], EdgeNode{
			Node:   from,
			Weight: 1.0 / values[i],
		})
	}
	res := make([]float64, len(queries))
	for i := 0; i < len(queries); i++ {
		res[i] = calDfs(graph, queries[i][0], queries[i][1])
	}
	return res
}

func calDfs(graph map[string][]EdgeNode, from, to string) float64 {
	if _, ok := graph[from]; !ok {
		return -1.0
	}
	if _, ok := graph[to]; !ok {
		return -1.0
	}
	if from == to {
		return 1.0
	}
	queue := []string{from}
	weight := make(map[string]float64)
	visited := make(map[string]bool)
	visited[from] = true
	weight[from] = 1.0
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		for _, neighbor := range graph[cur] {
			if visited[neighbor.Node] {
				continue
			}

			weight[neighbor.Node] = weight[cur] * neighbor.Weight
			if neighbor.Node == to {
				return weight[neighbor.Node]
			}
			visited[neighbor.Node] = true
			queue = append(queue, neighbor.Node)
		}
	}
	return -1
}
