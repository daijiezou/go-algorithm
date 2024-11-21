package _4_grid_chart

import (
	"sort"
)

/*
网格图 DFS
*/

func numIslands(grid [][]byte) int {
	m := len(grid)
	n := len(grid[0])
	cnt := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				landDfs2(grid, i, j)
				cnt++
			}
		}
	}
	return cnt
}

func landBfs(grid [][]byte, i, j int) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] == '0' {
		return
	}
	grid[i][j] = '0'
	landBfs(grid, i+1, j)
	landBfs(grid, i-1, j)
	landBfs(grid, i, j+1)
	landBfs(grid, i, j-1)
}

func landDfs2(grid [][]byte, i, j int) {
	grid[i][j] = '0'
	temp := [2]int{-1, 1}
	for _, iBias := range temp {
		if i+iBias >= 0 && i+iBias < len(grid) && grid[i+iBias][j] == '1' {
			landDfs2(grid, i+iBias, j)
		}

	}
	for _, jBias := range temp {
		if j+jBias >= 0 && j+jBias < len(grid[0]) && grid[i][j+jBias] == '1' {
			landDfs2(grid, i, j+jBias)
		}
	}
}

func maxAreaOfIsland(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				ans = max(ans, maxAreaOfIslandBfs(grid, i, j))
			}
		}
	}
	return ans
}

var dirs2 = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func maxAreaOfIslandBfs(grid [][]int, i, j int) int {
	grid[i][j] = 0
	ans := 1
	for _, bias := range dirs2 {
		iBias, jBias := i+bias[0], j+bias[1]
		if iBias >= 0 && iBias < len(grid) && jBias >= 0 && jBias < len(grid[0]) && grid[iBias][jBias] == 1 {
			ans += maxAreaOfIslandBfs(grid, iBias, iBias)
		}
	}
	return ans
}

func pondSizes(land [][]int) []int {
	m := len(land)
	n := len(land[0])
	res := []int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if land[i][j] == 0 {
				res = append(res, pondSizesDfs(land, i, j))
			}
		}
	}
	sort.Ints(res)
	return res
}

var dirs = [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

func pondSizesDfs(grid [][]int, i, j int) int {
	ans := 1
	grid[i][j] = 1
	for _, d := range dirs {
		x0, y0 := i+d[0], j+d[1]
		if x0 >= 0 && x0 < len(grid) && y0 >= 0 && y0 < len(grid[0]) && grid[x0][y0] == 0 {
			ans += pondSizesDfs(grid, x0, y0)
		}
	}
	return ans
}

func largestArea(grid []string) int {
	m := len(grid)
	n := len(grid[0])
	newGrid := make([][]int, m+2)

	for i := 0; i < m+2; i++ {
		newGrid[i] = make([]int, n+2)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			newGrid[i+1][j+1] = int(grid[i][j] - '0')
		}
	}
	ans := 0
	for i := 0; i < m+2; i++ {
		for j := 0; j < n+2; j++ {
			if 0 < newGrid[i][j] && newGrid[i][j] < 6 {
				ans = max(ans, largestAreaDfs(newGrid, i, j, newGrid[i][j]))
			}
		}
	}
	return ans
}

func largestAreaDfs(grid [][]int, i, j int, topic int) int {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] == 0 {
		return -250000
	}
	if grid[i][j] != topic {
		return 0
	}

	grid[i][j] = 6 // 表示已经遍历过
	ans := 1
	for _, d := range [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
		iBias, jBias := i+d[0], j+d[1]
		ans += largestAreaDfs(grid, iBias, jBias, topic)
	}
	return ans
}

/*
https://leetcode.cn/problems/island-perimeter/
岛屿的周长
*/
func islandPerimeter(grid [][]int) int {

	m := len(grid)
	n := len(grid[0])
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				continue
			}
			ans += 4
			// 查看这个岛屿四周是不是陆地
			for _, d := range [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
				iBias, jBias := i+d[0], j+d[1]
				if iBias >= 0 && iBias < m && jBias >= 0 && jBias < n && grid[iBias][jBias] == 1 {
					ans -= 1
				}
			}
		}
	}
	return ans
}

// https://leetcode.cn/problems/maximum-number-of-fish-in-a-grid/
func findMaxFish(grid [][]int) int {
	res := 0
	m := len(grid)
	n := len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] > 0 {
				res = max(res, findMaxFishDFS(grid, i, j))
			}
		}
	}
	return res
}

func findMaxFishDFS(grid [][]int, i, j int) int {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] == 0 {
		return 0
	}
	res := grid[i][j]
	grid[i][j] = 0
	res += findMaxFishDFS(grid, i+1, j)
	res += findMaxFishDFS(grid, i-1, j)
	res += findMaxFishDFS(grid, i, j+1)
	res += findMaxFishDFS(grid, i, j-1)
	return res
}

func numEnclaves(grid [][]int) int {
	res := 0
	m := len(grid)
	n := len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] > 0 {
				cnt := numEnclavesDFS(grid, i, j)
				if cnt > 0 {
					res += cnt
				}
			}
		}
	}
	return res
}

func numEnclavesDFS(grid [][]int, i, j int) int {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) {
		return -1000000
	}
	if grid[i][j] == 0 {
		return 0
	}
	res := 1
	grid[i][j] = 0
	res += numEnclavesDFS(grid, i+1, j)
	res += numEnclavesDFS(grid, i-1, j)
	res += numEnclavesDFS(grid, i, j+1)
	res += numEnclavesDFS(grid, i, j-1)
	return res
}
