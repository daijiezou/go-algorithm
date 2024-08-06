package __graph

func isBipartite(graph [][]int) bool {

	colors := make([]bool, len(graph))
	visited := make([]bool, len(graph))
	flag := true

	// 因为所有的图并不一定是全部相邻的，所以得每个都作为起点来遍历一遍
	for i := 0; i < len(graph); i++ {
		isBipartiteDFS(graph, colors, visited, &flag, i)
	}
	return flag
}

func isBipartiteDFS(graph [][]int, color []bool, visted []bool, flag *bool, x int) {
	if visted[x] {
		return
	}
	visted[x] = true
	for i := 0; i < len(graph[x]); i++ {
		neighbor := graph[x][i]
		// 相邻节点已经被访问过，查看颜色是否冲突
		if visted[neighbor] {
			if color[x] == color[neighbor] {
				*flag = false
				return
			}
		} else {
			// 相邻节点 neighbor 没有被访问过
			color[neighbor] = !color[x]
			isBipartiteDFS(graph, color, visted, flag, neighbor)
		}
	}
}

func isBipartiteBFS(graph [][]int, color []bool, visted []bool, flag *bool, x int) {
	if visted[x] {
		return
	}
	visted[x] = true
	queue := []int{x}
	for len(queue) > 0 && *flag {
		cur := queue[0]
		queue = queue[1:]
		for i := 0; i < len(graph[cur]); i++ {
			neighbor := graph[cur][i]
			// 相邻节点已经被访问过
			if visted[neighbor] {
				if color[cur] == color[neighbor] {
					*flag = false
					return
				}
			} else {
				// 相邻节点 graph[x][i] 没有被访问过
				// 给它标记颜色
				color[neighbor] = !color[cur]

				// 标记为已访问
				visted[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}
}

// https://leetcode.cn/problems/possible-bipartition/
func possibleBipartition(n int, dislikes [][]int) bool {
	colors := make([]bool, n+1)
	visited := make([]bool, n+1)
	flag := true
	graph := buildGraph2(dislikes, n)
	for i := 1; i <= n; i++ {
		isBipartiteBFS(graph, colors, visited, &flag, i)
		if !flag {
			return flag
		}
	}
	return flag
}

func buildGraph2(dislikes [][]int, n int) [][]int {
	graph := make([][]int, n+1)
	for i := 0; i < n; i++ {
		graph[i] = make([]int, 0)
	}
	for i := 0; i < len(dislikes); i++ {
		x1, x2 := dislikes[i][0], dislikes[i][1]
		graph[x1] = append(graph[x1], x2)
		graph[x2] = append(graph[x2], x1)
	}
	return graph
}
