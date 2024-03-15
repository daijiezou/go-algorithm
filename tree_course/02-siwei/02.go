package _2_siwei

// https://leetcode.cn/problems/invert-binary-tree/
// 遍历的方式
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	tmp := root.Left
	root.Left = root.Right
	root.Right = tmp
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

// 分解的方式
func invertTree2(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	// 利用函数定义，先翻转左右子树
	left := invertTree(root.Left)
	right := invertTree(root.Right)

	// 再交换左右节点
	root.Right = left
	root.Left = right
	return root
}

// https://leetcode.cn/problems/populating-next-right-pointers-in-each-node/
func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	// 遍历「三叉树」，连接相邻节点
	connecttraverse(root.Left, root.Right)
	return root
}

func connecttraverse(node1, node2 *Node) {
	if node1 == nil || node2 == nil {
		return
	}
	node1.Next = node2
	connecttraverse(node1.Left, node1.Right)
	connecttraverse(node2.Left, node2.Right)
	connecttraverse(node1.Right, node2.Left)
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Left)
	flatten(root.Right)
	left := root.Left
	right := root.Right

	root.Left = nil
	root.Right = left
	p := root
	for p.Right != nil {
		p = p.Right
	}
	p.Right = right
	return
}
