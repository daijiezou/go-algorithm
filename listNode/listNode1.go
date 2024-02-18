package listNode

type ListNode struct {
	Val  int
	Next *ListNode
}

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

// 使用分治
func mergeKLists2(lists []*ListNode) *ListNode {
	return merge(lists, 0, len(lists)-1)
}
