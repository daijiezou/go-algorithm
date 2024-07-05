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

// https://leetcode.cn/problems/all-paths-from-source-to-target/
// 纪录从七点到终点的所有路径
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
	newGraph := make([][]int, n)
	for i := 0; i < n; i++ {
		newGraph[i] = make([]int, 0, 2)
	}
	for i := 0; i < len(prerequisites); i++ {
		edge := prerequisites[i]
		from, to := edge[1], edge[0]
		newGraph[from] = append(newGraph[from], to)
	}
	return newGraph
}

// [{1,0}]
// 要学习课程1，先得学习课程0
func canFinishBFS(numCourses int, prerequisites [][]int) bool {
	// graph记录的是被依赖的情况
	graph := buildGraph(numCourses, prerequisites)
	inDegree := make([]int, numCourses)
	for _, edge := range prerequisites {
		_, to := edge[1], edge[0] //
		// 纪录需要先修课程的数量
		inDegree[to]++
	}
	// 根据入度初始化队列中的节点，和环检测算法相同
	queue := make([]int, 0)
	for i, val := range inDegree {
		if val == 0 {
			queue = append(queue, i)
		}
	}
	count := 0
	for len(queue) > 0 {
		count++
		cur := queue[0]
		queue = queue[1:]
		for _, next := range graph[cur] {

			//next是依赖cur的
			inDegree[next]--

			// 如果next的所有依赖都以完成，则加入队列
			if inDegree[next] == 0 {
				queue = append(queue, next)
			}
		}
	}
	// 如果所有节点都被遍历过，说明不成环
	return count == numCourses
}

// https://leetcode.cn/problems/course-schedule-ii/description/
func findOrder(numCourses int, prerequisites [][]int) []int {
	// graph记录的是被依赖的情况
	graph := buildGraph(numCourses, prerequisites)
	inDegree := make([]int, numCourses)
	for _, edge := range prerequisites {
		_, to := edge[1], edge[0] //
		// 纪录需要先修课程的数量
		inDegree[to]++
	}
	// 根据入度初始化队列中的节点，和环检测算法相同
	queue := make([]int, 0)
	for i, val := range inDegree {
		if val == 0 {
			queue = append(queue, i)
		}
	}
	count := 0
	res := make([]int, 0)
	for len(queue) > 0 {
		count++
		cur := queue[0]
		res = append(res, cur)
		queue = queue[1:]
		for _, next := range graph[cur] {

			//next是依赖cur的
			inDegree[next]--

			// 如果next的所有依赖都以完成，则加入队列
			if inDegree[next] == 0 {
				queue = append(queue, next)
			}
		}
	}
	// 如果所有节点都被遍历过，说明不成环
	if len(res) != numCourses {
		return []int{}
	}
	return res
}
