package _8_binarytree_listNode

type ListNode struct {
	Val  int
	Next *ListNode
}

// 206.反转链表
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var pre, cur, next *ListNode = nil, head, nil
	for cur != nil {
		// 保存原先变量的next
		next = cur.Next
		// 将当前节点的next的换成上一个节点
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

// 92.反转链表2
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := new(ListNode)
	dummy.Next = head
	// 翻转区间的上一个节点
	p0 := dummy
	for i := 0; i < left-1; i++ {
		p0 = p0.Next
	}
	var pre, cur, next *ListNode = nil, p0.Next, nil
	for i := left; i <= right; i++ {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	// 现在的p0.Next指向left位置的节点，所以要把p0.Next.Next指向right+1位置的节点
	p0.Next.Next = cur

	// 将p0节点的下一个节点指向right
	p0.Next = pre

	// dummy的next就是现在的头结点
	return dummy.Next
}
