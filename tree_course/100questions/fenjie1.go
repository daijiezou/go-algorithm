package _00questions

import (
	"fmt"
	"math"
	"strings"
)

func buildTree(preorder []int, inorder []int) *TreeNode {
	//获取中序的所有节点值的index
	inorderIndex := make(map[int]int)
	for k, v := range inorder {
		inorderIndex[v] = k
	}

	return build(preorder, 0, len(preorder)-1,
		inorder, 0, len(inorder)-1, inorderIndex)
}

func build(pre []int, preStart int, preEnd int, in []int, inStart int, inEnd int, indexMap map[int]int) *TreeNode {
	if preStart > preEnd {
		return nil
	}
	rootVal := pre[preStart]
	index := indexMap[rootVal]
	root := &TreeNode{Val: rootVal}
	leftSize := index - inStart
	root.Left = build(pre, preStart+1, preStart+leftSize, in, inStart, index-1, indexMap)
	root.Right = build(pre, preStart+leftSize+1, preEnd, in, index+1, inEnd, indexMap)
	return root
}

// https://leetcode.cn/problems/construct-binary-tree-from-inorder-and-postorder-traversal/
func buildTree2(inorder []int, postorder []int) *TreeNode {
	//获取中序的所有节点值的index
	inorderIndex := make(map[int]int)
	for k, v := range inorder {
		inorderIndex[v] = k
	}
	return build2(inorder, 0, len(inorder)-1, postorder, 0, len(postorder)-1, inorderIndex)
}

func build2(inorder []int, inStart int, inEnd int, postorder []int, postStart int, postEnd int, indexMap map[int]int) *TreeNode {
	if inStart > inEnd {
		return nil
	}
	// root 节点对应的值就是后序遍历数组的最后一个元素
	rootVal := postorder[postEnd]
	// rootVal 在中序遍历数组中的索引
	index := indexMap[rootVal]
	// 左子树的节点个数
	leftSize := index - inStart
	root := &TreeNode{Val: rootVal}
	root.Left = build2(inorder, inStart, index-1, postorder, postStart, postStart+leftSize-1, indexMap)
	root.Right = build2(inorder, index+1, inEnd, postorder, postStart+leftSize, postEnd-1, indexMap)
	return root
}

// 将字符串转化成列表
func isValidSerialization2(preorder string) bool {
	// 将字符串转化成列表
	nodes := strings.Split(preorder, ",")
	return deserialize(&nodes) && len(nodes) == 0
}

// 改造后的前序遍历反序列化函数
// 详细解析：https://mp.weixin.qq.com/s/DVX2A1ha4xSecEXLxW_UsA
func deserialize(nodes *[]string) bool {
	if len(*nodes) == 0 {
		return false
	}

	/****** 前序遍历位置 ******/
	// 列表最左侧就是根节点
	first := (*nodes)[0]
	*nodes = (*nodes)[1:]
	if first == "#" {
		return true
	}
	/**************************/
	left := deserialize(nodes)
	right := deserialize(nodes)
	return left && right
}

// https://leetcode.cn/problems/all-possible-full-binary-trees/description/
// 所有可能的真二叉树
func allPossibleFBT(n int) []*TreeNode {
	if n%2 == 0 {
		return []*TreeNode{}
	}
	memo := make(map[int][]*TreeNode)
	return allPossibleFBTCount(n, memo)
}

func allPossibleFBTCount(n int, memo map[int][]*TreeNode) []*TreeNode {
	var res []*TreeNode
	if n == 1 {
		res = append(res, &TreeNode{Val: 0})
		return res
	}
	if _, ok := memo[n]; ok {
		return memo[n]
	}
	for i := 1; i < n; i += 2 {
		j := n - i - 1
		leftSubTrees := allPossibleFBTCount(i, memo)
		RightSubTrees := allPossibleFBTCount(j, memo)
		// 左右子树的不同排列也能构成不同的二叉树
		for _, left := range leftSubTrees {
			for _, right := range RightSubTrees {
				// 生成根节点
				root := &TreeNode{Val: 0}
				// 组装出一种可能的二叉树形状
				root.Left = left
				root.Right = right
				// 加入结果列表
				res = append(res, root)
			}
		}
	}
	memo[n] = res
	return res
}

// 最大二叉树
// https://leetcode.cn/problems/maximum-binary-tree/
func constructMaximumBinaryTree(nums []int) *TreeNode {
	fmt.Println(nums)
	length := len(nums)
	if length == 0 {
		return nil
	}
	maxVal := 0
	var maxIndex int
	for i := 0; i < length; i++ {
		if nums[i] > maxVal {
			maxVal = nums[i]
			maxIndex = i
		}
	}
	root := &TreeNode{Val: nums[maxIndex]}
	root.Left = constructMaximumBinaryTree(nums[:maxIndex])
	root.Right = constructMaximumBinaryTree(nums[maxIndex+1:])
	return root
}

// 最大二叉树2
// https://leetcode.cn/problems/maximum-binary-tree/
func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	// 说明最右边的val是整个数组中的最大值
	if root.Val < val {
		temp := &TreeNode{Val: val}
		temp.Left = root
		return temp
	}
	// 如果 val 不是最大的，那么就应该在右子树上，
	// 因为 val 节点是接在原始数组 a 的最后一个元素
	root.Right = insertIntoMaxTree(root.Right, val)
	return root
}

// https://leetcode.cn/problems/delete-nodes-and-return-forest/
var delNodesRes []*TreeNode

func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	delNodesRes = []*TreeNode{}
	delSet := make(map[int]bool)
	for _, d := range to_delete {
		delSet[d] = true
	}
	delNodesHelper(root, false, delSet)
	return delNodesRes
}

func delNodesHelper(root *TreeNode, hasParent bool, delSet map[int]bool) *TreeNode {
	if root == nil {
		return nil
	}
	deleted := delSet[root.Val]
	// 没有父节点且不需要被删除就是一个新树
	if !hasParent && !deleted {
		delNodesRes = append(delNodesRes, root)
	}
	root.Left = delNodesHelper(root.Left, !deleted, delSet)
	root.Right = delNodesHelper(root.Right, !deleted, delSet)
	// 如果自己被删除了的话就return nil让上层去接收
	if deleted {
		return nil
	}
	return root
}

// https://leetcode.cn/problems/same-tree/
// 判断两棵树是否相同
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
	return isSameTree(p.Right, q.Right) && isSameTree(p.Left, p.Left)
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return checkIsSymmetric(root.Left, root.Right)
}

func checkIsSymmetric(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return checkIsSymmetric(left.Left, right.Right) && checkIsSymmetric(left.Right, right.Left)
}

// https://leetcode.cn/problems/flip-equivalent-binary-trees/description/
// 翻转等价二叉树
func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil {
		return false
	}

	if root1.Val != root2.Val {
		return false
	}

	if flipEquiv(root1.Left, root2.Left) && flipEquiv(root1.Right, root2.Right) {
		return true
	}
	if flipEquiv(root1.Left, root2.Right) && flipEquiv(root1.Right, root2.Left) {
		return true
	}
	return false
}

// https://leetcode.cn/problems/binary-tree-maximum-path-sum/description/
// 二叉树的最大单边和
var maxPathSumMaxRes int

func maxPathSum(root *TreeNode) int {
	maxPathSumMaxRes = math.MinInt32
	OneSideMax(root)
	return maxPathSumMaxRes
}

func OneSideMax(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 如果小于0，就放弃这边边
	leftMax := MyMax(OneSideMax(root.Left), 0)
	rightMax := MyMax(OneSideMax(root.Right), 0)
	// 计算最大路径和
	// 路径和等于两个单边和加上根节点的值
	if root.Val+leftMax+rightMax > maxPathSumMaxRes {
		maxPathSumMaxRes = root.Val + leftMax + rightMax
	}
	return MyMax(leftMax, rightMax) + root.Val
}

func MyMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
