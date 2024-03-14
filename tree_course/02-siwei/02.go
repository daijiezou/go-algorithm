package _2_siwei

import "Golang-algorithm/tree_course"

func invertTree(root *tree_course.TreeNode) *tree_course.TreeNode {
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

func invertTree2(root *tree_course.TreeNode) *tree_course.TreeNode {
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
