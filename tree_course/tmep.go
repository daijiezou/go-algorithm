package tree_course

import (
	"fmt"
	"strconv"
)

// OrderListNode 将链表的最后一个元素置为空
func OrderListNode(node *ListNode) *ListNode {
	if node.Next == nil {
		return nil
	}
	node.Next = OrderListNode(node)
	return node
}

// SizeListNode 递归的方法计算链表的长度
func SizeListNode(node *ListNode) int {
	return sizeListNode(node, 1)
}

func sizeListNode(node *ListNode, size int) int {
	if node.Next == nil {
		return size
	}
	return sizeListNode(node.Next, size+1)
}

// OrderPrintArray 递归方式倒序打印数组元素
func OrderPrintArray(nums []int) {
	orderPrintArray(nums, len(nums)-1)
}
func orderPrintArray(nums []int, index int) {
	if index < 0 {
		return
	}
	fmt.Println(nums[index])
	orderPrintArray(nums, index-1)
}

func OrderPrintNode(node *ListNode) {
	if node == nil {
		return
	}
	OrderPrintNode(node.Next)
	fmt.Println(node.Val)
}

// ParseTree 解析字符串表示的二叉树
// 格式: (1(2(7,8),3)) 表示节点1，左子树2(左7右8)，右子树3
func ParseTree(s string) *TreeNode {
	if s == "" {
		return nil
	}

	pos := 0
	return parseTreeHelper(s, &pos)
}

func parseTreeHelper(s string, pos *int) *TreeNode {
	if *pos >= len(s) || s[*pos] == ')' || s[*pos] == ',' {
		return nil
	}

	// 解析节点值
	start := *pos
	for *pos < len(s) && s[*pos] != '(' && s[*pos] != ')' && s[*pos] != ',' {
		(*pos)++
	}

	valStr := s[start:*pos]
	if valStr == "" {
		return nil
	}

	val, err := strconv.Atoi(valStr)
	if err != nil {
		return nil
	}

	node := &TreeNode{Val: val}

	// 检查是否有子节点
	if *pos < len(s) && s[*pos] == '(' {
		(*pos)++ // 跳过 '('

		// 解析左子树
		node.Left = parseTreeHelper(s, pos)

		// 跳过逗号
		if *pos < len(s) && s[*pos] == ',' {
			(*pos)++
			// 解析右子树
			node.Right = parseTreeHelper(s, pos)
		}

		// 跳过 ')'
		if *pos < len(s) && s[*pos] == ')' {
			(*pos)++
		}
	}

	return node
}

// RightSideView 返回二叉树的右视图（层序遍历每层最右边的节点）
func RightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)

		// 遍历当前层的所有节点
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			// 如果是当前层的最后一个节点，加入结果
			if i == levelSize-1 {
				result = append(result, node.Val)
			}

			// 将子节点加入队列
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return result
}

// LevelOrder 层序遍历（用于验证树结构）
func LevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := [][]int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		level := []int{}

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		result = append(result, level)
	}

	return result
}

// RightSideViewDirect 直接从字符串解析右视图，不构建二叉树
// 格式: (1(2(7,8),3)) 表示节点1，左子树2(左7右8)，右子树3
func RightSideViewDirect(s string) []int {
	if s == "" {
		return []int{}
	}

	result := []int{}
	pos := 0
	depth := 0

	// 记录每层最右边的值（用map存储，key是深度，value是该深度最右边的值）
	rightmost := make(map[int]int)
	maxDepth := -1

	parseRightView(s, &pos, depth, rightmost, &maxDepth)

	// 按深度顺序提取结果
	for i := 0; i <= maxDepth; i++ {
		result = append(result, rightmost[i])
	}

	return result
}

func parseRightView(s string, pos *int, depth int, rightmost map[int]int, maxDepth *int) {
	if *pos >= len(s) {
		return
	}

	// 跳过左括号
	if s[*pos] == '(' {
		(*pos)++
	}

	// 解析节点值
	if *pos >= len(s) || s[*pos] == ')' || s[*pos] == ',' {
		return
	}

	start := *pos
	for *pos < len(s) && s[*pos] != '(' && s[*pos] != ')' && s[*pos] != ',' {
		(*pos)++
	}

	valStr := s[start:*pos]
	if valStr == "" {
		return
	}

	val, err := strconv.Atoi(valStr)
	if err != nil {
		return
	}

	// 更新当前深度的最右值（后面的会覆盖前面的）
	rightmost[depth] = val
	if depth > *maxDepth {
		*maxDepth = depth
	}

	// 检查是否有子节点
	if *pos < len(s) && s[*pos] == '(' {
		// 解析左子树
		parseRightView(s, pos, depth+1, rightmost, maxDepth)

		// 跳过逗号
		if *pos < len(s) && s[*pos] == ',' {
			(*pos)++
			// 解析右子树
			parseRightView(s, pos, depth+1, rightmost, maxDepth)
		}

		// 跳过右括号
		if *pos < len(s) && s[*pos] == ')' {
			(*pos)++
		}
	}

	// 跳过外层右括号
	if *pos < len(s) && s[*pos] == ')' {
		(*pos)++
	}
}
