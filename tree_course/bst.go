package tree_course

// https://leetcode.cn/problems/search-in-a-binary-search-tree/
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	if root.Val < val {
		return searchBST(root.Right, val)
	}
	if root.Val > val {
		return searchBST(root.Left, val)
	}
	return nil
}

// https://leetcode.cn/problems/kth-smallest-element-in-a-bst/submissions/494158554/
func kthSmallest(root *TreeNode, k int) int {
	var orderIndex int
	var kthSmallestRes int
	var kthSmallestHelper func(root *TreeNode, k int)
	kthSmallestHelper = func(root *TreeNode, k int) {
		if root == nil {
			return
		}
		kthSmallestHelper(root.Left, k)
		orderIndex++
		if orderIndex == k {
			kthSmallestRes = root.Val
			return
		}
		kthSmallestHelper(root.Right, k)
		return
	}
	kthSmallestHelper(root, k)
	return kthSmallestRes
}

// bstToGst
func bstToGst(root *TreeNode) *TreeNode {
	var bstToGstSum int
	var bstToGstHelp func(root *TreeNode)
	bstToGstHelp = func(root *TreeNode) {
		if root == nil {
			return
		}
		bstToGstHelp(root.Right)
		bstToGstSum += root.Val
		root.Val = bstToGstSum
		bstToGstHelp(root.Left)
	}
	bstToGstHelp(root)
	return root
}

// 从BST里删除一个节点
// https://leetcode.cn/problems/delete-node-in-a-bst/
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == key {
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		minNode := getMinNode(root.Right)
		root.Right = deleteNode(root.Right, minNode.Val)
		minNode.Left = root.Left
		minNode.Right = root.Right
		root = minNode
	}
	if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	}
	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	}
	return root
}

func getMinNode(node *TreeNode) *TreeNode {
	for node.Left != nil {
		node = node.Left
	}
	return node
}
