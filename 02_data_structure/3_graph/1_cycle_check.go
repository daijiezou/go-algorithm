package __graph

/*
你这个学期必须选修 numCourses 门课程，记为 0 到 numCourses - 1 。
在选修某些课程之前需要一些先修课程。 先修课程按数组 prerequisites 给出，其中 prerequisites[i] = [ai, bi] ，
表示如果要学习课程 ai 则 必须 先学习课程  bi 。
例如，先修课程对 [0, 1] 表示：想要学习课程 0 ，你需要先完成课程 1 。
请你判断是否可能完成所有课程的学习？如果可以，返回 true ；否则，返回 false 。
*/

func buildGraph(n int, prerequisites [][]int) [][]int {
	newGraph := make([][]int, n)
	for i := 0; i < n; i++ {
		newGraph[i] = make([]int, 0, 2)
	}
	for i := 0; i < len(prerequisites); i++ {
		// [1,0] 表示要想学习课程1，必须先学习课程0
		// 用图来表示，就是有有一条从0到1的路径
		edge := prerequisites[i]
		from, to := edge[1], edge[0]
		newGraph[from] = append(newGraph[from], to)
	}
	return newGraph
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := buildGraph(numCourses, prerequisites)
	visited := make([]bool, numCourses)
	path := make([]bool, numCourses)
	hasCycle := false
	// note: 图中并不是所有节点都相连，
	// 所以要用一个 for 循环将所有节点都作为起点调用一次 DFS 搜索算法。
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
	// 将不需要依赖的课程先加入到队列中
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
/*
现在你总共有 numCourses 门课需要选，记为 0 到 numCourses - 1。
给你一个数组 prerequisites ，其中 prerequisites[i] = [ai, bi] ，表示在选修课程 ai 前 必须 先选修 bi 。
例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示：[0,1] 。
返回你为了学完所有课程所安排的学习顺序。可能会有多个正确的顺序，你只要返回 任意一种 就可以了。如果不可能完成所有课程，返回 一个空数组 。
*/

// 使用拓扑排序
func findOrder1(numCourses int, prerequisites [][]int) []int {
	graph := buildGraph(numCourses, prerequisites)
	visted := make([]bool, numCourses)
	path := make([]bool, numCourses)
	hasCycle := false
	post := make([]int, 0)
	for i := 0; i < numCourses; i++ {
		findOrder1Tranrvse(graph, visted, path, &hasCycle, &post, i)
	}
	if hasCycle {
		return []int{}
	}
	reverse(post)
	return post
}

// 数组翻转函数
func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func findOrder1Tranrvse(graph [][]int, visted []bool, path []bool, hasCycle *bool, post *[]int, s int) {
	// 要遍历的节点已经在路径中，表示有环
	if path[s] {
		*hasCycle = true
		return
	}

	if visted[s] {
		return
	}
	visted[s] = true
	path[s] = true
	for _, edge := range graph[s] {
		findOrder1Tranrvse(graph, visted, path, hasCycle, post, edge)
	}
	*post = append(*post, s)
	// 撤销选择
	path[s] = false

}

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
