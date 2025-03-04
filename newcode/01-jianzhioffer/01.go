package _1_jianzhioffer

func duplicate(numbers []int) int {
	// write code here
	cnt := make(map[int]struct{})
	for i := 0; i < len(numbers); i++ {
		if _, ok := cnt[numbers[i]]; ok {
			return numbers[i]
		}
		cnt[numbers[i]] = struct{}{}
	}
	return -1
}

func duplicate2(numbers []int) int {
	// write code here
	for i := 0; i < len(numbers); i++ {
		for numbers[i] != i {
			if numbers[i] == numbers[numbers[i]] {
				return numbers[i]
			}
			numbers[i], numbers[numbers[i]] = numbers[numbers[i]], numbers[i]
		}
	}
	return -1
}

/*
在一个长度为n+1的数组里的所有数字都在1～n的范围内，所以数组
中至少有一个数字是重复的。请找出数组中任意一个重复的数字，但不能
修改输入的数组。例如，如果输入长度为8的数组 2,3,5,4,3,2,6,7那
么对应的输出是重复的数字2或者3。
*/

func getDup(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	left := 1
	right := len(nums) - 1
	for right >= left {
		mid := left + (right-left)/2
		cnt := countRange(left, mid, nums)
		if left == right {
			if cnt > 1 {
				return left
			} else {
				return -1
			}
		}
		// 说明重复数字在left到mid之间
		if cnt > mid-left+1 {
			// 这里right不能等于mid-1，因为有可能mid是重复的数字
			right = mid
		} else {
			// 说明重复数字在mid+1到right之间
			left = mid + 1
		}
	}
	return -1
}

func countRange(left, right int, nums []int) int {
	cnt := 0
	for i := 0; i < len(nums); i++ {
		if left <= nums[i] && nums[i] <= right {
			cnt++
		}
	}
	return cnt
}

/*
给定一个二维数组，其每一行从左到右递增排序，从上到下也是递增排序。给定一个数，判断这个数是否在该二维数组中。

Consider the following matrix:
[

	[1,   4,  7, 11, 15],
	[2,   5,  8, 12, 19],
	[3,   6,  9, 16, 22],
	[10, 13, 14, 17, 24],
	[18, 21, 23, 26, 30]

]
*/

// https://leetcode.cn/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof/
func findTargetIn2DPlants(plants [][]int, target int) bool {
	m := len(plants)
	if m == 0 {
		return false
	}
	n := len(plants[0])
	row := 0
	col := n - 1
	for row < m && col >= 0 {
		if plants[row][col] == target {
			return true
		} else if plants[row][col] > target {
			col--
		} else if plants[row][col] < target {
			row++
		}
	}
	return false
}

func findTargetIn2DPlants2(plants [][]int, target int) bool {
	m := len(plants)
	if m == 0 {
		return false
	}
	n := len(plants[0])
	row := m - 1
	col := 0
	for row >= 0 && col < n {
		if plants[row][col] == target {
			return true
		} else if plants[row][col] > target {
			row--
		} else if plants[row][col] < target {
			col++
		}
	}
	return false
}

// ReplaceSpace 将字符串中的空格替换为"%20"
// 例如: "hello world" -> "hello%20world"
// 使用原地替换的方式，从后向前遍历，避免多次移动字符
func ReplaceSpace(input string) string {
	// 将字符串转换为字节数组以便修改
	strBytes := []byte(input)
	originalLen := len(input) - 1
	spaceCount := 0

	// 统计空格数量
	for i := 0; i <= originalLen; i++ {
		if strBytes[i] == ' ' {
			spaceCount += 2 // 每个空格需要额外2个字节(从1个字符变成3个字符)
		}
	}

	// 扩展原数组，为替换预留空间
	extraSpace := make([]byte, spaceCount)
	strBytes = append(strBytes, extraSpace...)
	newLen := len(strBytes) - 1

	// 从后向前遍历，替换空格
	for originalLen >= 0 && newLen > originalLen {
		char := strBytes[originalLen]
		originalLen--

		if char == ' ' {
			// 替换空格为"%20"
			strBytes[newLen] = '0'
			newLen--
			strBytes[newLen] = '2'
			newLen--
			strBytes[newLen] = '%'
			newLen--
		} else {
			// 非空格字符直接移动
			strBytes[newLen] = char
			newLen--
		}
	}

	return string(strBytes)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func printListFromTailToHead(head *ListNode) []int {
	// write code here
	res := []int{}
	if head == nil {
		return nil
	}
	res = printListFromTailToHead(head.Next)
	res = append(res, head.Val)
	return res
}

// 重建二叉树

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func reConstructBinaryTree(preOrder []int, vinOrder []int) *TreeNode {
	// write code here
	if len(preOrder) == 0 || len(vinOrder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preOrder[0]}
	leftIndex := 0
	for i := 0; i < len(vinOrder); i++ {
		if vinOrder[i] == root.Val {
			leftIndex = i
			break
		}
	}
	root.Left = reConstructBinaryTree(preOrder[1:leftIndex+1], vinOrder[:leftIndex])
	root.Right = reConstructBinaryTree(preOrder[leftIndex+1:], vinOrder[leftIndex+1:])
	return root
}

type TreeLinkNode struct {
	Val   int
	Left  *TreeLinkNode
	Right *TreeLinkNode
	Next  *TreeLinkNode
}

// GetNext 获取二叉树中序遍历的下一个节点
// 1. 如果当前节点有右子树，则下一个节点是右子树中最左边的节点
// 2. 如果当前节点没有右子树，则需要向上找第一个当前节点是其父节点的左子树的节点
func GetNext(pNode *TreeLinkNode) *TreeLinkNode {
	if pNode == nil {
		return nil
	}
	var pNext *TreeLinkNode
	// 如果有右子树,则下一个节点是右子树中最左边的节点
	if pNode.Right != nil {
		pNext = pNode.Right
		for pNext.Left != nil {
			pNext = pNext.Left
		}
		return pNext
	}

	// 如果没有右子树,
	// 则找第一个当前节点是其父节点的左子树的节点
	for pNode.Next != nil {
		root := pNode.Next
		if root.Left == pNode {
			return root
		}
		pNode = pNode.Next
	}
	return pNext
}
