package tree_course

import (
	"math"
)

/*
	层序遍历
*/

// 层序遍历的标准模板
func LevelTraverse(root *TreeNode) {
	if root == nil {
		return
	}
	q := make([]*TreeNode, 0)
	q = append(q, root)
	for len(q) > 0 {
		sz := len(q)
		for i := 0; i < sz; i++ {
			cur := q[0]
			// 在这里可以做些事情
			q = q[1:]
			// 将下一层节点放入队列
			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
		}
	}
}

// https://leetcode.cn/problems/binary-tree-level-order-traversal/description/
func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	var trees []*TreeNode
	trees = append(trees, root)
	for len(trees) > 0 {
		length := len(trees)
		var levelTress []int
		for i := 0; i < length; i++ {
			cur := trees[0]
			trees = trees[1:]
			levelTress = append(levelTress, cur.Val)
			if cur.Left != nil {
				trees = append(trees, cur.Left)
			}
			if cur.Right != nil {
				trees = append(trees, cur.Right)
			}
		}
		res = append(res, levelTress)
	}
	return res
}

var levelOrder2Res [][]int

// 使用递归的方式遍历
func levelOrder2(root *TreeNode) [][]int {
	if root == nil {
		return levelOrder2Res
	}
	levelOrder2Traverse(root, 0)
	return levelOrder2Res
}

func levelOrder2Traverse(root *TreeNode, depth int) {
	if root == nil {
		return
	}
	if len(levelOrder2Res) == depth {
		levelOrder2Res = append(levelOrder2Res, []int{})
	}
	levelOrder2Res[depth] = append(levelOrder2Res[depth], root.Val)
	levelOrder2Traverse(root.Left, depth+1)
	levelOrder2Traverse(root.Right, depth+1)
}

// https://leetcode.cn/problems/binary-tree-level-order-traversal-ii/
func levelOrderBottom(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	var trees []*TreeNode
	trees = append(trees, root)
	for len(trees) > 0 {
		lenth := len(trees)
		var levelTress []int
		for i := 0; i < lenth; i++ {
			cur := trees[0]
			trees = trees[1:]
			levelTress = append(levelTress, cur.Val)
			if cur.Left != nil {
				trees = append(trees, cur.Left)
			}
			if cur.Right != nil {
				trees = append(trees, cur.Right)
			}
		}
		// 把每一层添加到头部，就是自底向上的层序遍历。
		res = append([][]int{levelTress}, res...)
	}
	return res
}

func largestValues(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	var trees []*TreeNode
	trees = append(trees, root)
	for len(trees) > 0 {
		levelMax := -math.MaxInt64
		lenth := len(trees)
		for i := 0; i < lenth; i++ {
			cur := trees[0]
			trees = trees[1:]
			if cur.Val >= levelMax {
				levelMax = cur.Val
			}
			if cur.Left != nil {
				trees = append(trees, cur.Left)
			}
			if cur.Right != nil {
				trees = append(trees, cur.Right)
			}
		}
		res = append(res, levelMax)
	}
	return res
}

type Node struct {
	Val      int
	Children []*Node
}

func levelNOrder(root *Node) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	var trees []*Node
	trees = append(trees, root)
	for len(trees) > 0 {
		lenth := len(trees)
		var levelTress []int
		for i := 0; i < lenth; i++ {
			cur := trees[0]
			trees = trees[1:]
			levelTress = append(levelTress, cur.Val)
			for _, child := range cur.Children {
				trees = append(trees, child)
			}
		}
		res = append(res, levelTress)
	}
	return res
}

func findBottomLeftValue(root *TreeNode) int {
	var trees []*TreeNode
	trees = append(trees, root)
	res := root
	for len(trees) > 0 {
		res = trees[0]
		trees = trees[1:]
		if res.Right != nil {
			trees = append(trees, res.Right)
		}
		if res.Left != nil {
			trees = append(trees, res.Left)
		}
	}
	return res.Val
}
