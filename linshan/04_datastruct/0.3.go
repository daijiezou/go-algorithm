package _4_datastruct

import (
	"sort"
)

/*
给你一个大小为 n x n 的整数方阵 grid。返回一个经过如下调整的矩阵：

左下角三角形（包括中间对角线）的对角线按 非递增顺序 排序。
右上角三角形 的对角线按 非递减顺序 排序。
*/
func sortMatrix(grid [][]int) [][]int {
	n := len(grid)
	for i := 0; i < n; i++ {
		tmp := make([]int, 0)
		for j := 0; i+j < n; j++ {
			tmp = append(tmp, grid[i+j][j])
		}
		sort.Sort(sort.Reverse(sort.IntSlice(tmp)))
		for j := 0; i+j < n; j++ {
			grid[i+j][j] = tmp[j]
		}
	}

	for j := 1; j < n; j++ {
		tmp := []int{}
		for i := 0; j+i < n; i++ {
			tmp = append(tmp, grid[i][j+i])
		}
		sort.Ints(tmp)
		for i := 0; j+i < n; i++ {
			grid[i][j+i] = tmp[i]
		}
	}
	return grid
}
