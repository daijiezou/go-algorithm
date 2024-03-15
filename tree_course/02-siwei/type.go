package _2_siwei

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Node struct {
	Val               int
	Left, Right, Next *Node
}
