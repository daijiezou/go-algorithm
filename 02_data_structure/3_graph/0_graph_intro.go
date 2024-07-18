package __graph

// graph[x] 存储 x 的所有邻居节点
var graph [][]int

// 邻接矩阵
// matrix[x][y] 记录 x 是否有一条指向 y 的边
var matrix [][]bool

/*
图论中特有的度（degree）的概念，在无向图中，「度」就是每个节点相连的边的条数。
由于有向图的边有方向，所以有向图中每个节点「度」被细分为入度（indegree）和出度（outdegree）
*/

// 图遍历框架
/*var visited map[int]bool

func traverse(graph *Graph, v int) {
	// 防止走回头路进入死循环
	if visited[v] {
		return
	}
	// 前序遍历位置，标记节点 v 已访问
	visited[v] = true
	for _, neighbor := range graph.Neighbors(v) {
		traverse(graph, neighbor)
	}
}*/

// https://leetcode.cn/problems/all-paths-from-source-to-target/
// 纪录从起点到终点的所有路径
func allPathsSourceTarget(graph [][]int) [][]int {
	// 记录所有路径
	res := [][]int{}
	// 维护递归过程中经过的路径
	path := []int{}
	traverse(graph, 0, &path, &res)
	return res
}

func traverse(graph [][]int, start int, path *[]int, res *[][]int) {
	*path = append(*path, start)
	// 到达终点
	if start == len(graph)-1 {
		temp := make([]int, len(*path))
		copy(temp, *path)
		*res = append(*res, temp)
		// note:从路径移出节点 s
		*path = (*path)[:len(*path)-1]
		return
	}

	for i := 0; i < len(graph[start]); i++ {
		traverse(graph, graph[start][i], path, res)
	}
	// 从路径移出节点 s
	*path = (*path)[:len(*path)-1]
}
