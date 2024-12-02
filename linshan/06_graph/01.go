package _6_graph

type union struct {
	parent []int
	Cnt    int
}

func (u *union) find(x int) (parent int) {
	root := x
	// 找到自己的父节点
	for u.parent[root] != root {
		root = u.parent[root]
	}
	oldParent := u.parent[x]
	// 然后把 x 到根节点之间的所有节点直接接到根节点下面
	for x != root {
		u.parent[x] = root
		x = oldParent
		oldParent = u.parent[x]
	}

	return root
}

// 将节点 p 和节点 q 连通
func (u *union) union(p, q int) {
	rootP := u.find(p)
	rootQ := u.find(q)

	if rootP == rootQ {
		return
	}

	u.parent[rootQ] = rootP
	// 两个连通分量合并成一个连通分量
	u.Cnt--
}
func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	uni := union{parent: parent, Cnt: n}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if isConnected[i][j] == 1 {
				uni.union(i, j)
			}
		}
	}
	return uni.Cnt
}

func validPath(n int, edges [][]int, source int, destination int) bool {
	graph := make([][]int, n)
	for _, e := range edges {
		graph[e[0]] = append(graph[e[0]], e[1])
		graph[e[1]] = append(graph[e[1]], e[0])
	}
	visited := make([]bool, n)
	queue := make([]int, 0, n)
	queue = append(queue, source)
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		visited[cur] = true
		if cur == destination {
			return true
		}
		for _, w := range graph[cur] {
			if !visited[w] {
				queue = append(queue, w)
			}
		}
	}
	return false
}

// https://leetcode.cn/problems/all-paths-from-source-to-target/
func allPathsSourceTarget(graph [][]int) [][]int {
	res := make([][]int, 0)
	path := []int{0}
	var btk = func(start int, path []int) {}
	btk = func(start int, path []int) {
		if start == len(graph)-1 {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}
		for _, w := range graph[start] {
			// 做选择
			path = append(path, w)

			btk(w, path)
			// 取消选择
			path = path[:len(path)-1]
		}
	}
	btk(0, path)
	return res
}

// https://leetcode.cn/problems/keys-and-rooms/
func canVisitAllRooms(rooms [][]int) bool {
	n := len(rooms)
	visited := make([]bool, n)
	visited[0] = true
	cnt := 1
	queue := []int{0}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		for _, w := range rooms[cur] {
			if !visited[w] {
				visited[w] = true
				cnt++
				queue = append(queue, w)
			}
		}
	}
	return cnt == n
}

// https://leetcode.cn/problems/count-unreachable-pairs-of-nodes-in-an-undirected-graph/
func countPairs(n int, edges [][]int) int64 {
	res := 0
	graph := make([][]int, n)
	for _, e := range edges {
		graph[e[0]] = append(graph[e[0]], e[1])
		graph[e[1]] = append(graph[e[1]], e[0])
	}
	vis := make([]bool, n)
	var dfs func(x int) int
	dfs = func(x int) int {
		vis[x] = true
		size := 1
		for _, w := range graph[x] {
			if !vis[w] {
				size += dfs(w)
			}
		}
		return size
	}
	total := 0
	for i, b := range vis {
		// 说明是一个新的连通块
		if !b {
			size := dfs(i)
			res += size * total
			total += size
		}
	}
	return int64(res)

	//for i := 0; i < n; i++ {
	//	notReach := n - 1 - canReachCnt(i, graph, n)
	//	res += notReach
	//}
	//return int64(res) / 2
}

func canReachCnt(start int, edges [][]int, n int) int {
	queue := []int{start}
	visited := make([]bool, n)
	visited[start] = true
	if start > n-1 {
		return 0
	}
	cnt := 0
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if cur >= len(edges) {
			return cnt
		}
		for _, w := range edges[cur] {
			if !visited[w] {
				cnt++
				visited[w] = true
				queue = append(queue, w)
			}
		}
	}
	return cnt
}
