package meeting150

import "math"

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
