package __graph

import "math"

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
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x])
	}
	return uf.parent[x]
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

// https://leetcode.cn/problems/satisfiability-of-equality-equations/
func equationsPossible(equations []string) bool {
	// 26个英文字母
	uf := NewUF(26)

	// 将等号两边的字母联通起来
	for _, x := range equations {
		if x[1] == '=' {
			uf.union(int(x[0]-'a'), int(x[3]-'a'))
		}

	}
	for _, x := range equations {
		if x[1] == '!' {
			if uf.connected(int(x[0]-'a'), int(x[3]-'a')) {
				return false
			}
		}
	}
	return true
}

func equationsPossible2(equations []string) bool {
	// 26个英文字母
	parent := make([]int, 26)
	for i := 0; i < 26; i++ {
		// 一开始每个人的父节点都是自己
		parent[i] = i
	}

	// 将等号两边的字母联通起来
	for _, x := range equations {
		if x[1] == '=' {
			union(parent, int(x[0]-'a'), int(x[3]-'a'))
		}

	}
	for _, x := range equations {
		if x[1] == '!' {
			pRoot := find(parent, int(x[0]-'a'))
			qRoot := find(parent, int(x[3]-'a'))
			if pRoot == qRoot {
				return false
			}
		}
	}
	return true
}

func union(parent []int, p int, q int) {
	pRoot := find(parent, p)
	qRoot := find(parent, q)
	parent[pRoot] = qRoot
}

func find(parent []int, q int) int {
	if parent[q] != q {
		parent[q] = find(parent, parent[q])
	}
	return parent[q]
}

func robotWithString(s string) string {
	tByte := make([]byte, 0)
	res := make([]byte, 0)
	n := len(s)
	sufMin := make([]byte, n+1)
	sufMin[n] = math.MaxUint8
	for i := n - 1; i >= 0; i-- {
		sufMin[i] = min(sufMin[i+1], s[i])
	}
	for i := 0; i < n; i++ {
		tByte = append(tByte, s[i])
		for len(tByte) > 0 && tByte[len(tByte)-1] <= sufMin[i+1] {
			left := tByte[len(tByte)-1]
			tByte = tByte[:len(tByte)-1]
			res = append(res, left)
		}
	}
	return string(res)
}
