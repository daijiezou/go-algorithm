package _1_listNodeAndArray

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.cn/problems/partition-list/description/
// 分隔两个链表
func partition(head *ListNode, x int) *ListNode {
	dummy1 := &ListNode{Val: -1, Next: nil}
	dummy2 := &ListNode{Val: -1, Next: nil}
	p1, p2 := dummy1, dummy2
	p := head
	for p != nil {
		if p.Val >= x {
			p2.Next = p
			p2 = p2.Next
		} else {
			p1.Next = p
			p1 = p1.Next
		}
		// 不能直接让 p 指针前进，
		// p = p.Next
		// 断开原链表中的每个节点的 Next 指针
		temp := p.Next
		p.Next = nil
		p = temp

	}
	p1.Next = dummy2.Next
	return dummy1.Next
}

// https://leetcode.cn/problems/merge-two-sorted-lists/
// 合并两个有序链表
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{Val: -1, Next: nil}
	p := dummy
	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			p.Next = list2
			list2 = list2.Next
		} else {
			p.Next = list1
			list1 = list1.Next
		}
		p = p.Next
	}
	for list1 != nil {
		p.Next = list1
		list1 = list1.Next
		p = p.Next
	}

	for list2 != nil {
		p.Next = list2
		list2 = list2.Next
		p = p.Next
	}
	return dummy.Next
}

// 合并K个升序链表
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	length := len(lists)
	preList := mergeTwoLists(lists[0], lists[1])
	for i := 2; i < length; i++ {
		preList = mergeTwoLists(preList, lists[i])
	}
	return preList
}

// 使用分治合并K个升序链表
func mergeKLists2(lists []*ListNode) *ListNode {
	return merge(lists, 0, len(lists)-1)
}

func merge(lists []*ListNode, left int, right int) *ListNode {
	if left == right {
		return lists[left]
	}
	if left > right {
		return nil
	}
	mid := (left + right) >> 1
	return mergeTwoLists(merge(lists, left, mid), merge(lists, mid+1, right))
}

// 链表的中间节点
func middleNode(head *ListNode) *ListNode {
	fast, slow := head, head
	// 快指针走到末尾时停止
	for fast != nil && fast.Next != nil {
		// 快指针走两步，慢指针走一步
		// 等到快指针走到末尾时，慢指针正好走到中间节点
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

// 判断列表是否包含环
func hasCycle(head *ListNode) bool {
	// 快慢指针初始化指向 head
	slow, fast := head, head
	// 快指针走到末尾时停止
	for fast != nil && fast.Next != nil {
		// 慢指针走一步，快指针走两步
		slow = slow.Next
		fast = fast.Next.Next
		// 快慢指针相遇，说明含有环
		if slow == fast {
			return true
		}
	}
	// 不包含环
	return false
}

// 让 p1 遍历完链表 A 之后开始遍历链表 B，让 p2 遍历完链表 B 之后开始遍历链表 A，这样相当于「逻辑上」两条链表接在了一起。
// 如果这样进行拼接，就可以让 p1 和 p2 同时进入公共部分，也就是同时到达相交节点 c1：
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	p1 := headA
	p2 := headB
	for p1 != p2 {
		if p1 != nil {
			p1 = p1.Next
		} else {
			p1 = headB
		}

		if p2 != nil {
			p2 = p2.Next
		} else {
			p2 = headA
		}
	}

	return p1
}
