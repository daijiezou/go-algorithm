package tree_course

/*
公共祖先问题
*/

// https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/description/
// 因为题目说了 p 和 q 一定存在于二叉树中(这点很重要），
// 所以即便我们遇到 q 就直接返回，根本没遍历到 p，也依然可以断定 p 在 q 底下，q 就是 LCA 节点。
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// 前序位置，看看 root 是不是目标值
	if root.Val == q.Val || root.Val == p.Val {
		// 如果遇到目标值，直接返回
		return root
	}
	// 去左右子树寻找
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		// 当前节点是 LCA 节点
		return root
	}
	if left != nil {
		return left
	} else {
		return right
	}
}

func lowestCommonAncestorHelper(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val < p.Val {
		// 当前节点太小去右子树找
		return lowestCommonAncestor(root.Right, p, q)
	}
	if root.Val > q.Val {
		// 当前节点太大去左子树找
		return lowestCommonAncestor(root.Left, p, q)
	}
	// val1 <= root.val <= val2
	// 则当前节点就是最近公共祖先
	return root

}
