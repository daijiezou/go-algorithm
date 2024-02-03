package _00questions

import (
	"Golang-algorithm/tree"
	"strconv"
)

var binaryTreePathsRes []string
var binaryTreePath []int

func binaryTreePaths(root *tree.TreeNode) []string {
	binaryTreePathsRes = []string{}
	binaryTreePath = []int{}
	binaryTreePathsHelper(root)
	return binaryTreePathsRes
}
func binaryTreePathsHelper(node *tree.TreeNode) {
	if node == nil {
		return
	}
	binaryTreePath = append(binaryTreePath, node.Val)
	if node.Left == nil && node.Right == nil {
		tempStr := strconv.Itoa(binaryTreePath[0])
		for i := 1; i < len(binaryTreePath); i++ {
			tempStr += "->"
			tempStr += strconv.Itoa(binaryTreePath[i])
		}
		binaryTreePathsRes = append(binaryTreePathsRes, tempStr)
	}
	binaryTreePathsHelper(node.Left)
	binaryTreePathsHelper(node.Right)
	binaryTreePath = binaryTreePath[:len(binaryTreePath)-1]
}

// https://leetcode.cn/problems/sum-root-to-leaf-numbers/
// 求根节点到叶节点数字之和
var sumNumbersRes int
var sumNumbersPath string

func sumNumbers(root *tree.TreeNode) int {
	sumNumbersRes = 0
	sumNumbersPath = ""
	sumNumbersHelper(root)
	return sumNumbersRes
}

func sumNumbersHelper(root *tree.TreeNode) {
	if root == nil {
		return
	}
	// 到达叶子节点
	if root.Left == nil && root.Right == nil {
		pathsum, _ := strconv.Atoi(sumNumbersPath)
		sumNumbersRes += pathsum
		return
	}
	sumNumbersPath += strconv.Itoa(root.Val)
	sumNumbersHelper(root.Left)
	sumNumbersHelper(root.Right)
	sumNumbersPath = sumNumbersPath[:len(sumNumbersPath)-1]
}

// 二叉树的右视图
// https://leetcode.cn/problems/binary-tree-right-side-view/
func rightSideView(root *tree.TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	var trees []*tree.TreeNode
	trees = append(trees, root)
	for len(trees) > 0 {
		length := len(trees)
		for i := 0; i < length; i++ {
			cur := trees[0]
			trees = trees[1:]
			if cur.Left != nil {
				trees = append(trees, cur.Left)
			}
			if cur.Right != nil {
				trees = append(trees, cur.Right)
			}
			// 这一层的最后的一个节点就是右侧能看到的
			if i == length-1 {
				res = append(res, cur.Val)
			}
		}
	}
	return res
}

var smallestFromLeafRes string
var smallestFromLeafPath []int

// 988. 从叶结点开始的最小字符串
// https://leetcode.cn/problems/smallest-string-starting-from-leaf/description/
func smallestFromLeaf(root *TreeNode) string {
	smallestFromLeafRes = ""
	smallestFromLeafPath = []int{}
	smallestFromLeafHelper(root)
	return smallestFromLeafRes
}

func smallestFromLeafHelper(root *TreeNode) {
	if root == nil {
		return
	}

	smallestFromLeafPath = append(smallestFromLeafPath, root.Val)
	if root.Left == nil && root.Right == nil {
		tempRes := ""
		for i := len(smallestFromLeafPath) - 1; i >= 0; i-- {
			tempRes += string([]byte{byte(smallestFromLeafPath[i] + 'a')})
		}
		if smallestFromLeafRes == "" || tempRes < smallestFromLeafRes {
			smallestFromLeafRes = tempRes
		}
	}
	smallestFromLeafHelper(root.Left)
	smallestFromLeafHelper(root.Right)
	smallestFromLeafPath = smallestFromLeafPath[:len(smallestFromLeafPath)-1]
}

// https://leetcode.cn/problems/sum-of-root-to-leaf-binary-numbers/
// 从根到叶的二进制数之和
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

// https://leetcode.cn/problems/flip-binary-tree-to-match-preorder-traversal/
var flipMatchVoyageRes []int
var voyageFlag bool
var flipMatchVoyageIndex int

func flipMatchVoyage(root *tree.TreeNode, voyage []int) []int {
	flipMatchVoyageRes = []int{}
	voyageFlag = true
	flipMatchVoyageIndex = 0
	flipMatchVoyageHelper(root, voyage)
	if !voyageFlag {
		return []int{-1}
	}
	return flipMatchVoyageRes
}

func flipMatchVoyageHelper(node *tree.TreeNode, voyage []int) {
	if node == nil || !voyageFlag {
		return
	}

	if node.Val != voyage[flipMatchVoyageIndex] {
		voyageFlag = false
		return
	}
	flipMatchVoyageIndex++
	if node.Left != nil && node.Left.Val != voyage[flipMatchVoyageIndex] {
		temp := node.Left
		node.Left = node.Right
		node.Right = temp
		flipMatchVoyageRes = append(flipMatchVoyageRes, node.Val)
	}
	flipMatchVoyageHelper(node.Left, voyage)
	flipMatchVoyageHelper(node.Right, voyage)
}
