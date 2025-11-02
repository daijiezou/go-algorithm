package _5_grid_chart

// https://leetcode.cn/problems/nearest-exit-from-entrance-in-maze/
func nearestExit(maze [][]byte, entrance []int) int {
	m := len(maze)
	n := len(maze[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	step := 0
	queue := [][]int{entrance}
	visited[entrance[0]][entrance[1]] = true
	for len(queue) > 0 {
		sz := len(queue)
		step++
		for i := 0; i < sz; i++ {
			cur := queue[0]
			queue = queue[1:]
			for _, d := range [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
				x := cur[0] + d[0]
				y := cur[1] + d[1]
				if x < 0 || x >= m || y < 0 || y >= n || visited[x][y] || maze[x][y] == '+' {
					continue
				}
				if x == 0 || x == m-1 || y == 0 || y == n-1 {
					// 走到边界（出口）
					return step
				}
				visited[x][y] = true
				queue = append(queue, []int{x, y})
			}
		}
	}
	return -1
}

func shortestPathBinaryMatrix(grid [][]int) int {
	if grid[0][0] == 1 {
		return -1
	}
	m := len(grid)
	n := len(grid[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	step := 1
	queue := [][]int{{0, 0}}
	for len(queue) > 0 {
		sz := len(queue)
		for i := 0; i < sz; i++ {
			cur := queue[0]
			queue = queue[1:]
			if cur[0] == m-1 && cur[1] == n-1 {
				return step
			}
			for _, d := range [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}} {
				x := cur[0] + d[0]
				y := cur[1] + d[1]
				if x < 0 || x >= m || y < 0 || y >= n || visited[x][y] || grid[x][y] == 1 {
					continue
				}
				visited[x][y] = true
				queue = append(queue, []int{x, y})
			}
		}
		step++
	}
	return -1
}
