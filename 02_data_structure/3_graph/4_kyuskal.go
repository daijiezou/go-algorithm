package __graph

import (
	"sort"
)

// https://leetcode.cn/problems/min-cost-to-connect-all-points/
/*
先对所有边按照权重从小到大排序，从权重最小的边开始，选择合适的边加入 mst 集合，这样挑出来的边组成的树就是权重和最小的。
在挑选边的时候是有讲究的，如果一条边的两个节点已经是连通的，则这条边会使 mst 集合中出现环；如果最后的连通分量总数大于 1，
则说明形成的是多棵树（森林）而不是一棵最小生成树。
所以，Kruskal 算法用到了
Union-Find 并查集算法，来保证挑选出来的这些边组成的一定是一棵「树」，而不会包含环或者形成一片「森林」。
*/
func minCostConnectPoints(points [][]int) int {
	edges := make([][]int, 0, len(points)+1)
	n := len(points)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			cost := abs(points[i][0], points[j][0]) + abs(points[i][1], points[j][1])
			edges = append(edges, []int{i, j, cost})
		}
	}
	sort.Slice(edges, func(i, j int) bool {
		return edges[i][2] < edges[j][2]
	})
	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	totalCost := 0
	for _, edge := range edges {
		if find(parent, edge[0]) == find(parent, edge[1]) {
			continue
		}
		totalCost += edge[2]
		union(parent, edge[0], edge[1])
	}
	return totalCost
}

func abs(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}
