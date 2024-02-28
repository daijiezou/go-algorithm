package tree_course

import (
	"math"
)

// https://leetcode.cn/problems/add-one-row-to-tree/
func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if depth == 1 {
		newNode := &TreeNode{Val: val, Left: root}
		return newNode
	}
	return addOneRowHelper(root, val, depth, 2)
}

func addOneRowHelper(node *TreeNode, val int, depth, curDepth int) *TreeNode {
	if node == nil {
		return nil
	}
	if depth == curDepth {
		newLeft := &TreeNode{Val: val, Left: node.Left}
		newRight := &TreeNode{Val: val, Right: node.Right}
		node.Left = newLeft
		node.Right = newRight
		return node
	}
	node.Left = addOneRowHelper(node.Left, val, depth, curDepth+1)
	node.Right = addOneRowHelper(node.Right, val, depth, curDepth+1)
	return node
}

//func levelOrder(root *TreeNode) [][]int {
//	var res [][]int
//	if root == nil {
//		return res
//	}
//	layerQueue := []*TreeNode{root}
//	for len(layerQueue) > 0 {
//		length := len(layerQueue)
//		var layer []int
//		// 从左到右遍历每一层的每个节点
//		for i := 0; i < length; i++ {
//			cur := layerQueue[0]
//			// 将下一层节点放入队列
//			if cur.Left != nil {
//				layerQueue = append(layerQueue, cur.Left)
//			}
//			if cur.Right != nil {
//				layerQueue = append(layerQueue, cur.Right)
//			}
//			layer = append(layer, layerQueue[0].Val)
//			// 将遍历完的节点踢出去
//			layerQueue = layerQueue[1:]
//		}
//		res = append(res, layer)
//	}
//	return res
//}

// invertTree:翻转二叉树
// https://leetcode.cn/problems/invert-binary-tree/
func invertTree(root *TreeNode) *TreeNode {
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

// flatten：将以 root 为根的树拉平为链表
// https://leetcode.cn/problems/flatten-binary-tree-to-linked-list/
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Right)
	flatten(root.Left)

	right := root.Right
	root.Right = root.Left
	root.Left = nil
	p := root
	for p.Right != nil {
		p = p.Right
	}
	p.Right = right
}

var sum int

func convertBST(root *TreeNode) *TreeNode {
	// 反向中序遍历
	reversInorder(root)
	return root
}

// reversInorder 反向中序遍历
func reversInorder(root *TreeNode) {
	if root == nil {
		return
	}
	reversInorder(root.Right)
	sum += root.Val
	root.Val = sum
	reversInorder(root.Left)
}

// 记录遍历到的节点的深度
func reverseOddLevels(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	q := make([]*TreeNode, 0)
	q = append(q, root)
	var depth int
	// 从上到下遍历二叉树的每一层
	for len(q) > 0 {
		sz := len(q)
		if depth%2 == 1 {
			n := len(q)
			for i := 0; i < n/2; i++ {
				nodex, nodey := q[i], q[n-1-i]
				nodex.Val, nodey.Val = nodey.Val, nodex.Val
			}
		}
		// 从左到右遍历每一层的每个节点
		for i := 0; i < sz; i++ {
			cur := q[0]
			q = q[1:]
			// 将下一层节点放入队列
			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
		}
		depth++
	}
	return root
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)
	return max(leftDepth, rightDepth) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

var res int
var depth int

func minDepth(root *TreeNode) int {
	res = math.MaxInt32
	traverse(root)
	return res
}

func traverse(root *TreeNode) {
	if root == nil {
		return
	}
	depth++
	if root.Left == nil && root.Right == nil {
		res = min(depth, res)
	}
	traverse(root.Left)
	traverse(root.Right)
	depth--
}
func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

var res2 []int

func preorderTraversal(root *TreeNode) []int {
	res2 = make([]int, 0)
	if root == nil {
		return res2
	}
	traverse2(root)
	return res2
}

func traverse2(root *TreeNode) {
	if root == nil {
		return
	}
	res2 = append(res2, root.Val)
	traverse2(root.Left)
	traverse2(root.Right)
}

// https://leetcode.cn/problems/path-sum/
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == root.Right {
		if root.Val == targetSum {
			return true
		}
	}
	return hasPathSum(root.Right, targetSum-root.Val) || hasPathSum(root.Left, targetSum-root.Val)
}

var res3 [][]int
var cursum int
var path []int

func pathSum2(root *TreeNode, targetSum int) [][]int {
	res3 = make([][]int, 0)
	path = make([]int, 0)
	traverse3(root, targetSum)
	return res3
}

func traverse3(root *TreeNode, targetSum int) {
	if root == nil {
		return
	}
	path = append(path, root.Val)
	cursum += root.Val
	if root.Left == root.Right {
		if cursum == targetSum {
			cp := make([]int, len(path))
			copy(cp, path)
			res3 = append(res3, cp)
		}
	}
	traverse3(root.Left, targetSum)
	traverse3(root.Right, targetSum)
	path = path[:len(path)-1]
	cursum -= root.Val
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return subRoot == nil
	}
	if isSameTree(root, subRoot) {
		return true
	}
	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func isSameTree(p, q *TreeNode) bool {
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

type ListNode struct {
	Val  int
	Next *ListNode
}

func isSubPath(head *ListNode, root *TreeNode) bool {
	if head == nil {
		return true
	}
	if root == nil {
		return false
	}
	if head.Val == root.Val {
		if isSamePath(head, root) {
			return true
		}
	}
	return isSubPath(head, root.Left) || isSubPath(head, root.Right)
}

func isSamePath(head *ListNode, root *TreeNode) bool {
	if head == nil {
		return true
	}
	if root == nil {
		return false
	}
	if root.Val != head.Val {
		return false
	}
	return isSamePath(head.Next, root.Left) || isSamePath(head.Next, root.Right)
}

// https://leetcode.cn/problems/sum-of-root-to-leaf-binary-numbers/
func sumRootToLeaf(root *TreeNode) int {
	var sumRootToLeaPath int
	var sumRootToLeafRes int
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		sumRootToLeaPath = sumRootToLeaPath<<1 | root.Val
		if root.Left == nil && root.Right == nil {
			sumRootToLeafRes += sumRootToLeaPath
		}
		dfs(root.Left)
		dfs(root.Right)
		sumRootToLeaPath = sumRootToLeaPath >> 1
	}
	dfs(root)
	return sumRootToLeafRes
}

func pseudoPalindromicPaths(root *TreeNode) int {
	var res int
	count := [10]int{}
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		count[root.Val]++
		if root.Left == nil && root.Right == nil {
			var oddCount int
			for _, i2 := range count {
				if i2%2 == 1 {
					oddCount++
				}
			}
			if oddCount <= 1 {
				res++
			}
		}
		dfs(root.Left)
		dfs(root.Right)
		count[root.Val]--
	}
	dfs(root)
	return res
}

// https://leetcode.cn/problems/path-sum-iii/
var path3SumMap map[int]int
var path3Sum int
var path3Res int

func pathSum3(root *TreeNode, targetSum int) int {
	path3SumMap = make(map[int]int)
	path3SumMap[0] = 1
	path3Sum = 0
	path3Res = 0
	pathSum3Helper(root, targetSum)
	return path3Res
}

func pathSum3Helper(node *TreeNode, targetSum int) {
	if node == nil {
		return
	}
	path3Sum += node.Val
	if count, ok := path3SumMap[path3Sum-targetSum]; ok {
		path3Res += count
	}
	path3SumMap[path3Sum]++
	pathSum3Helper(node.Left, targetSum)
	pathSum3Helper(node.Right, targetSum)
	path3SumMap[path3Sum]--
	path3Sum -= node.Val
}
