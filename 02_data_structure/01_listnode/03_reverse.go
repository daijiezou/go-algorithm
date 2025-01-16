package _1_listnode

// https://leetcode.cn/problems/reverse-linked-list/
// 反转链表
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	last := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last
}

func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var pre, cur, next *ListNode
	pre, cur, next = nil, head, head.Next
	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

// 反转链表的前N个节点
func reverseN(head *ListNode, n int) *ListNode {
	if n == 1 {
		nextNode = head.Next
		return head
	}
	last := reverseN(head.Next, n-1)
	head.Next.Next = head
	head.Next = nextNode
	return last
}

func reverseN2(head *ListNode, n int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var pre, cur, next *ListNode
	pre, cur, next = nil, head, head.Next
	for n > 0 {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
		n--
	}
	head.Next = cur
	return pre
}

func reverseBetween2(head *ListNode, m int, n int) *ListNode {
	if m == 1 {
		return reverseN2(head, n)
	}
	pre := head
	for i := 1; i < m-1; i++ { // 此处比较重要，从1开始到m-1
		pre = pre.Next
	}
	pre.Next = reverseN2(pre.Next, n-m+1)
	return head
}

var nextNode *ListNode

// 反转一个区间里的链表
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	defer func() {
		nextNode = nil
	}()

	if left == 1 {
		return reverseN(head, right)
	}
	// 前进到反转的起点触发 base case
	head.Next = reverseBetween(head.Next, left-1, right-1)
	return head
}

func ReverseN(head *ListNode, n int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var pre, cur, nxt *ListNode
	pre, cur, nxt = nil, head, head.Next
	for n > 0 {
		cur.Next = pre
		pre = cur
		cur = nxt
		if nxt != nil {
			nxt = nxt.Next
		}
		n--
	}
	head.Next = cur
	return pre
}

func DJreverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == 1 {
		return reverseN(head, right)
	}
	pre := head
	// 找到需翻转节点的前驱
	for i := 1; i < left-1; i++ {
		pre = pre.Next
	}
	last := ReverseN(pre.Next, right-left+1)
	pre.Next = last
	return head
}

func re(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var pre, cur, nxt *ListNode
	pre, cur, nxt = nil, head.Next, head
	for cur != nil {
		cur.Next = pre
		pre = cur
		cur = nxt
		if nxt != nil {
			nxt = nxt.Next
		}
	}
	return pre
}
