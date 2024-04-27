package _2_data_structure

// 迭代的方式反转链表
func reverse(head *ListNode) *ListNode {
	var pre, cur, next *ListNode
	pre, cur, next = nil, head, head
	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	return pre
}

func reverseAb(a, b *ListNode) *ListNode {
	var pre, cur, next *ListNode
	pre, cur, next = new(ListNode), a, a
	for cur != b {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre

}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	a, b := head, head
	for i := 0; i < k; i++ {
		if b == nil {
			return head
		}
		b = b.Next
	}
	// 反转前k个的元素
	last := reverseAb(a, b)
	a.Next = reverseKGroup(b, k)
	return last
}
