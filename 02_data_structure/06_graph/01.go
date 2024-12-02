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
