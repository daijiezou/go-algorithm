package __graph

func isBipartite(graph [][]int) bool {

	colors := make([]bool, len(graph))
	visited := make([]bool, len(graph))
	flag := true

	// 因为所有的图并不一定是全部相邻的，所以得每个都作为起点来遍历一遍
	for i := 0; i < len(graph); i++ {
		isBipartiteTransve(graph, colors, visited, &flag, i)
	}
	return flag
}

func isBipartiteTransve(graph [][]int, color []bool, visted []bool, flag *bool, x int) {
	if visted[x] {
		return
	}
	visted[x] = true
	for i := 0; i < len(graph[x]); i++ {
		neighbor := graph[x][i]
		if visted[neighbor] {
			if color[x] == color[graph[x][i]] {
				*flag = false
				return
			}
		} else {
			// 相邻节点 graph[x][i] 没有被访问过
			color[neighbor] = !color[x]
			isBipartiteTransve(graph, color, visted, flag, neighbor)
		}
	}
}
