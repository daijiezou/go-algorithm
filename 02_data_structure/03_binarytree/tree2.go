package _3_binarytree

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := invertTree(root.Left)
	right := invertTree(root.Right)
	root.Left, root.Right = left, right
	return root
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	connectTraverse(root.Left, root.Right)
	return root
}

func connectTraverse(node1 *Node, node2 *Node) {
	if node1 == nil || node2 == nil {
		return
	}
	node1.Next = node2
	connectTraverse(node1.Left, node1.Right)
	connectTraverse(node2.Left, node2.Right)
	connectTraverse(node2.Right, node2.Left)
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
	for p != nil {
		p = p.Right
	}
	p.Right = right
}
