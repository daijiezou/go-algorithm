package tree_course

import "fmt"

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
