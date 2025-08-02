package _1_core_platform

var totalSum int

// https://leetcode.cn/problems/3Etpl5/
// 从根节点到叶子节点之合
func sumNumbers(root *TreeNode) int {
	defer func() {
		totalSum = 0
	}()
	if root == nil {
		return 0
	}
	// valStr := strconv.Itoa(root.Val)
	sumNumbersHelper(root, 0)
	return totalSum
}

func sumNumbersHelper(root *TreeNode, currentVal int) {
	if root == nil {
		return
	}
	currentVal = currentVal*10 + root.Val
	if root.Left == nil && root.Right == nil {
		// val,_ := strconv.Atoi(currentVal)
		totalSum += currentVal
	}
	sumNumbersHelper(root.Left, currentVal)
	sumNumbersHelper(root.Right, currentVal)
}

// 二叉树的最大深度

// 使用分解方法
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := maxDepth(root.Left) + 1
	rightDepth := maxDepth(root.Right) + 1
	return max(leftDepth, rightDepth)
}

var (
	currentDepth int
	maxDep       int
)

func maxDepth2(root *TreeNode) int {
	defer func() {
		currentDepth = 0
		maxDep = 0
	}()
	maxDepth2Traverse(root)
	return maxDep
}

func maxDepth2Traverse(root *TreeNode) {
	if root == nil {
		return
	}
	currentDepth++
	// 到达叶子节点
	if root.Left == nil && root.Right == nil {
		if currentDepth > maxDep {
			maxDep = currentDepth
		}
	}
	maxDepth2Traverse(root.Left)
	maxDepth2Traverse(root.Right)
	currentDepth--
}

func preorderTraverse(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}

	// 前序遍历的结果，root.val 在第一个
	res = append(res, root.Val)
	// 利用函数定义，后面接着左子树的前序遍历结果
	res = append(res, preorderTraverse(root.Left)...)
	// 利用函数定义，最后接着右子树的前序遍历结果
	res = append(res, preorderTraverse(root.Right)...)

	return res
}

// 层序遍历
func levelTraverse(root *TreeNode) {
	if root == nil {
		return
	}
	treeList := []*TreeNode{root}
	res := make([]int, 0)

	for len(treeList) > 0 {
		length := len(treeList)
		for length > 0 {
			current := treeList[0]
			treeList = treeList[1:]
			if current.Left != nil {
				treeList = append(treeList, current.Left)
			}
			if current.Right != nil {
				treeList = append(treeList, current.Right)
			}
			res = append(res, current.Val)
			length--
		}
	}
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	treeList := []*TreeNode{root}
	res := make([][]int, 0)

	for len(treeList) > 0 {
		length := len(treeList)
		levelRes := make([]int, 0)
		for length > 0 {
			current := treeList[0]
			treeList = treeList[1:]
			if current.Left != nil {
				treeList = append(treeList, current.Left)
			}
			if current.Right != nil {
				treeList = append(treeList, current.Right)
			}
			levelRes = append(levelRes, current.Val)
			length--
		}
		res = append(res, levelRes)
	}
	return res
}

func rob(root *TreeNode) int {
	var dfs func(root *TreeNode) (int, int)
	dfs = func(root *TreeNode) (int, int) {
		if root == nil {
			return 0, 0
		}
		robleft, notRobleft := dfs(root.Left)
		robright, notRobright := dfs(root.Right)
		return notRobleft + notRobright + root.Val, max(robleft, notRobleft) + max(robright, notRobright)
	}
	return max(dfs(root))
}
