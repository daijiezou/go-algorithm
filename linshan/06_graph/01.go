package _6_graph

type union struct {
	parent []int
	Cnt    int
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
