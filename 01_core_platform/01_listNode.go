package _1_core_platform

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.cn/problems/merge-two-sorted-lists/submissions/517454095/
// 合并两个有序链表
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{Val: -1, Next: nil}
	p := dummy
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			p.Next = l1
			l1 = l1.Next
		} else {
			p.Next = l2
			l2 = l2.Next
		}
		p = p.Next
	}
	for l1 != nil {
		p.Next = l1
		p = p.Next
		l1 = l1.Next
	}
	for l2 != nil {
		p.Next = l2
		p = p.Next
		l2 = l2.Next
	}
	return dummy.Next
}
