package _00questions

// 路径总和
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		if root.Val == targetSum {
			return true
		}
	}
	return hasPathSum(root.Right, targetSum-root.Val) || hasPathSum(root.Left, targetSum-root.Val)
}

var hasPathSum2PathSum int
var hasPathSumFlag bool

func hasPathSum2(root *TreeNode, targetSum int) bool {
	hasPathSumFlag = false
	hasPathSum2PathSum = 0
	hasPathSum2Helper(root, targetSum)
	return hasPathSumFlag
}

func hasPathSum2Helper(root *TreeNode, targetSum int) {
	if root == nil {
		return
	}
	hasPathSum2PathSum += root.Val
	if root.Left == nil && root.Right == nil {
		if hasPathSum2PathSum == targetSum {
			hasPathSumFlag = true
			return
		}
	}
	hasPathSum2Helper(root.Left, targetSum)
	hasPathSum2Helper(root.Right, targetSum)
	hasPathSum2PathSum -= root.Val
}

var pathSum2Res [][]int
var path2 []int

// 路径总和2
func pathSum(root *TreeNode, targetSum int) [][]int {
	var Answers [][]int
	if root == nil {
		return Answers
	}
	if root.Left == root.Right && root.Val == targetSum {
		path := []int{root.Val}
		Answers = append(Answers, path)
	}
	leftAnswers := pathSum(root.Left, targetSum-root.Val)
	rightAnswers := pathSum(root.Right, targetSum-root.Val)
	for i := 0; i < len(leftAnswers); i++ {
		path := []int{root.Val}
		path = append(path, leftAnswers[i]...)
		Answers = append(Answers, path)
	}

	for i := 0; i < len(rightAnswers); i++ {
		path := []int{root.Val}
		path = append(path, rightAnswers[i]...)
		Answers = append(Answers, path)
	}
	return Answers
}

var node *TreeNode

func increasingBST2(root *TreeNode) *TreeNode {
	node = new(TreeNode)
	if root == nil {
		return node
	}
	node1 := node
	increasingBSTHelper(root)
	return node1.Right
}

func increasingBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// 先把左右子树拉平
	left := increasingBST(root.Left)
	right := increasingBST(root.Right)
	root.Left = nil
	root.Right = right
	// 左子树为空的话就不需要处理了
	if left == nil {
		return root
	}
	// 把根节点和右子树接在左子树的后面
	p := left
	for p.Right != nil {
		p = p.Right
	}
	p.Right = root
	return left
}

func increasingBSTHelper(root *TreeNode) {
	if root == nil {
		return
	}
	increasingBSTHelper(root.Left)
	for node.Right != nil {
		node = node.Right
	}
	node.Right = &TreeNode{
		Val: root.Val,
	}
	increasingBSTHelper(root.Right)
}

// 二叉搜索树的范围和：分解方法
func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}

	if root.Val > high {
		return rangeSumBST(root.Left, low, high)
	}
	if root.Val < low {
		return rangeSumBST(root.Right, low, high)
	}
	// 以 root 为根的这棵 BST 落在 [low, high] 之间的元素之和，
	// 等于 root.val 加上左右子树落在区间的元素之和
	return root.Val + rangeSumBST(root.Left, low, high) + rangeSumBST(root.Right, low, high)
}

// 二叉搜索树的范围和：遍历方法
var rangeSumBST2Sum int

func rangeSumBST2(root *TreeNode, low int, high int) int {
	rangeSumBST2Sum = 0
	rangeSumBST2Helper(root, low, high)
	return rangeSumBST2Sum
}

func rangeSumBST2Helper(root *TreeNode, low int, high int) {
	if root == nil {
		return
	}
	if root.Val > high {
		rangeSumBST2Helper(root.Left, low, high)
		return
	}
	if root.Val < low {
		rangeSumBST2Helper(root.Right, low, high)
		return
	}
	rangeSumBST2Helper(root.Left, low, high)
	rangeSumBST2Helper(root.Right, low, high)
	rangeSumBST2Sum += root.Val
}
