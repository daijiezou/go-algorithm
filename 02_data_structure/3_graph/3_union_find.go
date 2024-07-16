package __graph

// 并查集
type UFI interface {
	// 将 p 和 q 连接
	union(p int, q int)

	// 判断 p 和 q 是否连通
	connected(p int, q int) bool

	// 返回图中有多少个连通分量
	count() int
}

type UF struct {
	// 记录连通分量
	Count int
	// 节点 x 的父节点是 parent[x]
	parent []int
}

// NewUF /* 构造函数，n 为图的节点总数 */
func NewUF(n int) *UF {
	// 一开始互不连通
	uf := &UF{Count: n, parent: make([]int, n)}
	// 父节点指针初始指向自己
	for i := 0; i < n; i++ {
		uf.parent[i] = i
	}
	return uf
}

/* 返回某个节点 x 的根节点 */
func (uf *UF) find(x int) int {
	// 根节点的 parent[x] == x
	for uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x])
	}
	return x
}

/* 返回当前的连通分量个数 */
func (uf *UF) count() int {
	return uf.Count
}

func (uf *UF) union(p int, q int) {
	rootP := uf.find(p)
	rootQ := uf.find(q)
	if rootP == rootQ {
		return
	}
	uf.parent[rootQ] = rootP
	uf.Count--
}

func (uf *UF) connected(p int, q int) bool {
	rootP := uf.find(p)
	rootQ := uf.find(q)
	return rootP == rootQ
}
