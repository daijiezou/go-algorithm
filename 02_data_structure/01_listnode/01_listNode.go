package _1_listnode

type ListNode struct {
	Val  int
	Next *ListNode
}

// 单链表的分解

func partition(head *ListNode, x int) *ListNode {
	dummy1 := &ListNode{Val: -1}
	dummy2 := &ListNode{Val: -1}

	// p1, p2 指针负责生成结果链表
	p1 := dummy1
	p2 := dummy2
	for head != nil {
		if head.Val < x {
			p1.Next = head
			p1 = p1.Next
		} else {
			p2.Next = head
			p2 = p2.Next
		}
		temp := head.Next
		head.Next = nil
		head = temp
	}
	p1.Next = dummy2.Next
	return dummy1.Next
}

// https://leetcode.cn/problems/remove-nth-node-from-end-of-list/description/
// 删除链表的倒数第N个节点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	length := getLength(head)
	dummy := &ListNode{0, head}
	cur := dummy
	for i := 0; i < length-n; i++ {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	return dummy.Next
}

func getLength(head *ListNode) (length int) {
	for ; head != nil; head = head.Next {
		length++
	}
	return
}

// https://leetcode.cn/problems/intersection-of-two-linked-lists/description/
// 相交链表
func getIntersectionNode(headA *ListNode, headB *ListNode) *ListNode {
	p1 := headA
	p2 := headB
	for p1 != p2 {
		if p1 == nil {
			p1 = headB
		} else {
			p1 = p1.Next
		}

		if p2 == nil {
			p2 = headA
		} else {
			p2 = p2.Next
		}
	}
	return p1
}

// https://leetcode.cn/problems/sort-list/description/
func sortList(head *ListNode) *ListNode {
	return sort(head, nil)
}

func sort(head *ListNode, tail *ListNode) *ListNode {
	if head == nil {
		return head
	}
	// 说明只有一个节点了
	if head.Next == tail {
		return head
	}

	// 寻找到中间节点
	fast := head
	slow := head
	for fast != tail && fast.Next != tail {
		fast = fast.Next.Next
		slow = slow.Next
	}
	mid := slow
	return mergeTwoLists(sort(head, mid), sort(mid, tail))
}

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

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	pq := NewMinPQ()
	for _, listNode := range lists {
		if listNode != nil {
			pq.insert(listNode)
		}
	}
	dummy := &ListNode{Val: -1}
	cur := dummy
	for pq.size > 0 {
		node := pq.pop()
		cur.Next = node
		cur = cur.Next
		if node.Next != nil {
			pq.insert(node.Next)
		}
	}
	return dummy.Next
}
