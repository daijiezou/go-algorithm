package __graph

// graph[x] 存储 x 的所有邻居节点
var graph [][]int

// 邻接矩阵
// matrix[x][y] 记录 x 是否有一条指向 y 的边
var matrix [][]bool

/*
图论中特有的度（degree）的概念，在无向图中，「度」就是每个节点相连的边的条数。
由于有向图的边有方向，所以有向图中每个节点「度」被细分为入度（indegree）和出度（outdegree），比如下图：
*/

// https://leetcode.cn/problems/all-paths-from-source-to-target/
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
		// 从路径移出节点 s
		*path = (*path)[:len(*path)-1]
		return
	}

	for i := 0; i < len(graph[start]); i++ {
		traverse(graph, graph[start][i], path, res)
	}
	// 从路径移出节点 s
	*path = (*path)[:len(*path)-1]
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := buildGraph(numCourses, prerequisites)
	visited := make([]bool, numCourses)
	path := make([]bool, numCourses)
	hasCycle := false
	for i := 0; i < numCourses; i++ {
		// 遍历图中的所有节点
		canFinishTr(graph, i, visited, path, &hasCycle)
	}

	return !hasCycle
}

func canFinishTr(graph [][]int, s int, visited []bool, path []bool, hasCycle *bool) {
	if path[s] {
		// 发现环！！！
		*hasCycle = true
		return
	}
	if visited[s] || *hasCycle {
		return
	}
	// 将节点 s 标记为已遍历
	visited[s] = true
	// 开始遍历节点 s
	path[s] = true
	for _, t := range graph[s] {
		canFinishTr(graph, t, visited, path, hasCycle)
	}
	// 节点 s 遍历完成
	path[s] = false
}

func buildGraph(n int, prerequisites [][]int) [][]int {
	graph := make([][]int, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]int, 0, 2)
	}
	for i := 0; i < len(prerequisites); i++ {
		from, to := prerequisites[i][1], prerequisites[i][0]
		graph[from] = append(graph[from], to)
	}
	return graph
}
