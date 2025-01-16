package mytest

// 二叉树的最大深度
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)
	return max(leftDepth, rightDepth) + 1
}

var depth int
var res int

func maxDepthTraverse(root *TreeNode) {
	if root == nil {
		return
	}
	// 前序位置
	depth++
	if root.Left == nil && root.Right == nil {
		res = max(res, depth)
	}
	maxDepthTraverse(root.Left)
	maxDepthTraverse(root.Right)
	// 后续位置
	depth--
}

func maxDepth2(root *TreeNode) int {
	ans := 0
	var f func(root *TreeNode, cnt int)
	f = func(root *TreeNode, cnt int) {
		if root == nil {
			return
		}
		cnt++
		ans = max(cnt, ans)
		f(root.Left,cnt)
		f(root.Right,cnt)
	}
	f(root, 0)
	return ans
}
