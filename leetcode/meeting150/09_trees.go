package meeting150

import (
	"math"
	"slices"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	return max(left, right) + 1
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	return root
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return Symmetric(root.Left, root.Right)
}

func Symmetric(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}
	return Symmetric(p.Left, q.Right) && Symmetric(p.Right, q.Left)
}

//func buildTree(preorder []int, inorder []int) *TreeNode {
//	if len(preorder) == 0 {
//		return nil
//	}
//	root := &TreeNode{Val: preorder[0]}
//	if len(preorder) == 1 {
//		return root
//	}
//	leftIndex := 0
//	for i := 0; i < len(inorder); i++ {
//		if inorder[i] == root.Val {
//			leftIndex = i
//			break
//		}
//	}
//	root.Left = buildTree(preorder[1:leftIndex+1], inorder[0:leftIndex+1])
//	root.Right = buildTree(preorder[leftIndex+1:], inorder[leftIndex+1:])
//
//	return root
//}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: postorder[len(postorder)-1]}
	if len(postorder) == 1 {
		return root
	}
	leftIndex := 0
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == root.Val {
			leftIndex = i
			break
		}
	}
	root.Left = buildTree(inorder[0:leftIndex+1], postorder[:leftIndex])
	root.Right = buildTree(inorder[leftIndex+1:], postorder[leftIndex:len(postorder)-1])

	return root
}

type Node struct {
	Val    int
	Left   *Node
	Right  *Node
	Next   *Node
	Random *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	queue := []*Node{root}
	for len(queue) > 0 {
		size := len(queue)
		var pre *Node
		for size > 0 {
			cur := queue[0]
			queue = queue[1:]
			if pre != nil {
				pre.Next = cur
			}
			pre = cur
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
			size--
		}

	}
	return root
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		if root.Val == targetSum {
			return true
		} else {
			return false
		}

	} else {
		return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
	}

}

func sumNumbers(root *TreeNode) int {
	//if root == nil {
	//	return 0
	//}
	return num(root, 0)
}

func num(node *TreeNode, preSum int) int {
	if node == nil {
		return 0
	}
	preSum = preSum*10 + node.Val
	if node.Left == nil && node.Right == nil {
		return preSum
	}
	return num(node.Left, preSum) + (num(node.Right, preSum))
}

func maxPathSum(root *TreeNode) int {
	res := math.MinInt32
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		leftMax := dfs(root.Left)
		rightMax := dfs(root.Right)
		if root.Val+leftMax+rightMax > res {
			res = root.Val + leftMax + rightMax
		}
		return max(0, max(leftMax, rightMax, 0)+root.Val)
	}
	dfs(root)
	return res
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l, r := root, root
	lh := 0
	rh := 0
	for l != nil {
		l = l.Left
		lh++
	}
	for r != nil {
		r = r.Right
		rh++
	}
	if rh == lh {
		return int(math.Pow(2, float64(lh))) - 1
	}
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}

// 二叉树的右视图
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	queue := []*TreeNode{root}
	res := make([]int, 0)
	for len(queue) > 0 {
		size := len(queue)
		for size > 0 {
			cur := queue[0]
			queue = queue[1:]
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
			if size == 1 {
				res = append(res, cur.Val)
			}
			size--
		}
	}
	return res
}

func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return []float64{}
	}
	queue := []*TreeNode{root}
	res := make([]float64, 0)
	for len(queue) > 0 {
		size := len(queue)
		sum := float64(0)
		curSize := size
		for size > 0 {
			cur := queue[0]
			queue = queue[1:]
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
			size--
			sum += float64(cur.Val)
		}
		res = append(res, sum/float64(curSize))
	}
	return res
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := []*TreeNode{root}
	res := make([][]int, 0)
	flag := true
	for len(queue) > 0 {
		size := len(queue)
		curVals := []int{}
		for size > 0 {
			cur := queue[0]
			queue = queue[1:]
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
			curVals = append(curVals, cur.Val)
			size--
		}
		if flag {
			res = append(res, curVals)
		} else {
			slices.Reverse(curVals)
			res = append(res, curVals)
		}
		flag = !flag
	}
	return res
}

func getMinimumDifference(root *TreeNode) int {
	ans := math.MaxInt
	var prev *TreeNode
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if prev != nil {
			ans = min(node.Val-prev.Val, ans)
		}
		prev = node
		if ans == 1 {
			return
		}
		dfs(node.Right)
	}
	dfs(root)
	return ans
}

func kthSmallest(root *TreeNode, k int) int {
	ans := 0

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		k--
		if k == 0 {
			ans = node.Val
			return
		}
		dfs(node.Right)
	}
	dfs(root)
	return ans
}

func isValidBST(root *TreeNode) bool {
	var dfs func(node *TreeNode)
	var prev *TreeNode
	valid := true
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if prev != nil && node.Val <= prev.Val {
			valid = false
			return
		}
		prev = node
		dfs(node.Right)
	}
	dfs(root)
	return valid
}
