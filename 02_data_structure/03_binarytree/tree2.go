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

func buildTree(preorder []int, inorder []int) *TreeNode {
	return build(preorder, 0, len(preorder)-1, inorder, 0, len(inorder)-1)
}

func build(preorder []int, preStart, preEnd int, inorder []int, inorderStart, inorderEnd int) *TreeNode {
	if preStart > preEnd {
		return nil
	}

	root := &TreeNode{Val: preorder[preStart]}
	index := inorder[inorderStart]
	for i := inorderStart; i <= inorderEnd; i++ {
		if inorder[i] == root.Val {
			index = i
			break
		}
	}
	//左子树的节点个数
	leftLength := index - inorderStart

	root.Left = build(preorder, preStart+1, preStart+leftLength, inorder, inorderStart, inorderStart+leftLength-1)
	root.Right = build(preorder, preStart+1+leftLength, preEnd, inorder, index+1, inorderEnd)
	return root
}
