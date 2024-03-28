package _1_listNodeAndArray

// 反转链表
// https://leetcode.cn/problems/reverse-linked-list/
func reverseList(head *ListNode) *ListNode {
	dummy := &ListNode{Val: -1}
	reveseListNode := dummy
	nodeVal := []int{}
	for head.Next != nil {
		head = head.Next
		nodeVal = append(nodeVal, head.Val)
	}
	for i := len(nodeVal) - 1; i >= 0; i-- {
		reveseListNode.Next = &ListNode{Val: nodeVal[i]}
		reveseListNode = reveseListNode.Next
	}
	return dummy.Next
}
