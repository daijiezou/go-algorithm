package tree

func buildTree(preorder []int, inorder []int) *TreeNode {
	// 根据函数定义，用 preorder 和 reversInorder 构造二叉树
	return build(preorder, 0, len(preorder)-1,
		inorder, 0, len(inorder)-1)
}

// 构造篇
func build(pre []int, preStart int, preEnd int, in []int, inStart int, inEnd int) *TreeNode {
	if preStart > preEnd {
		return nil
	}
	rootVal := pre[preStart]
	var index int
	for i := inStart; i <= inEnd; i++ {
		if in[i] == rootVal {
			index = i
			break
		}
	}
	root := &TreeNode{Val: rootVal}
	leftSize := index - inStart
	root.Left = build(pre, preStart+1, preStart+leftSize, in, inStart, index-1)
	root.Right = build(pre, preStart+leftSize+1, preEnd, in, index+1, inEnd)
	return root
}

func build2(inorder []int, inStart int, inEnd int, postorder []int, postStart int, postEnd int) *TreeNode {
	if inStart > inEnd {
		return nil
	}
	rootVal := postorder[postEnd]
	var index int
	for i := inStart; i <= inEnd; i++ {
		if inorder[i] == rootVal {
			index = i
			break
		}
	}
	root := &TreeNode{Val: rootVal}
	// 左子树的节点数
	leftSize := index - inStart
	leftPostEnd := postStart + leftSize - 1 // 下标等于数组长度减1
	root.Left = build2(inorder, inStart, index-1, postorder, postStart, leftPostEnd)
	root.Right = build2(inorder, index+1, inEnd, postorder, postStart+leftSize, postEnd-1)
	return root
}

func build3(pre []int, preStart int, preEnd int, postorder []int, postStart int, postEnd int) *TreeNode {
	if preStart > preEnd {
		return nil
	}
	if preStart == preEnd {
		return &TreeNode{Val: pre[preStart]}
	}
	root := &TreeNode{Val: pre[preStart]}
	var index int
	for i := postStart; i <= postEnd; i++ {
		if postorder[i] == pre[preStart+1] {
			index = i
			break
		}
	}
	leftSize := index - postStart + 1
	root.Left = build3(pre, preStart+1, preStart+leftSize, postorder, postStart, index)
	root.Right = build3(pre, preStart+1+leftSize, preEnd, postorder, index+1, postEnd-1)
	return root
}

// removeLeafNodes
// https://leetcode.cn/problems/delete-leaves-with-a-given-value/
func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left = removeLeafNodes(root.Left, target)
	root.Right = removeLeafNodes(root.Right, target)
	if root.Left == nil && root.Right == nil && root.Val == target {
		return nil
	}
	return root
}
