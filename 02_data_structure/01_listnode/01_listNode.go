package _1_listnode

import "container/heap"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 单链表的分解
// https://leetcode.cn/problems/partition-list/description/
/*
给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。

你应当 保留 两个分区中每个节点的初始相对位置。
*/
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
		/*
			果我们需要把原链表的节点接到新链表上，
			而不是 new 新节点来组成新链表的话，那么断开节点和原链表之间的链接可能是必要的。
		*/
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

	pq := ListNodeHeap{}

	heap.Init(pq)
	// 将 k 个链表的头结点加入最小堆
	for _, head := range lists {
		if head != nil {
			heap.Push(&pq, head)
		}
	}

	dummy := &ListNode{Val: -1}
	cur := dummy
	for pq.Len() > 0 {
		node := heap.Pop(&pq).(*ListNode)
		cur.Next = node
		cur = cur.Next
		if node.Next != nil {
			heap.Push(&pq, node.Next)
		}
	}
	return dummy.Next
}

// 链表排序，给
