package _1_listnode

func reverseListNode(head *ListNode) *ListNode {
	var prev, cur, next *ListNode
	prev, cur, next = nil, head, head
	for cur != nil {
		next = cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}

func deleteDuplicatesTest(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	p := dummy
	for head != nil {
		// 值相同
		if head.Next != nil && head.Val == head.Next.Val {
			for head.Next != nil && head.Next.Val == head.Val {
				head = head.Next
			}
			// 把自己也去掉
			head = head.Next
		} else {
			p.Next = head
			head = head.Next
			p = p.Next
		}
	}
	p.Next = head
	return dummy.Next
}

func deleteDuplicatesUnsortedTest(head *ListNode) *ListNode {
	count := make(map[int]int)
	p := head
	for p != nil {
		if _, ok := count[p.Val]; ok {
			count[p.Val]++
		} else {
			count[p.Val] = 1
		}
	}
	dummy := &ListNode{Val: 0, Next: head}
	q := dummy
	for head != nil {
		for count[head.Val] > 1 {
			head = head.Next
		}
		q.Next = head
		q = q.Next
		head = head.Next
	}
	return dummy.Next
}

// 链表排序
func sortListTest(head *ListNode) *ListNode {
	return mergeSort(head, nil)
}

func mergeSort(head *ListNode, tail *ListNode) *ListNode {
	if head == nil {
		return head
	}
	// 只有一个节点了
	if head.Next == tail {
		head.Next = nil
		return head
	}
	// 找到链表中点进行切分
	slow, fast := head, head
	for fast != tail && fast.Next != tail {
		slow = slow.Next
		fast = fast.Next.Next
	}
	mid := slow
	left := mergeSort(head, mid)
	right := mergeSort(mid, tail)
	return merge(left, right)
}

func merge(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{Val: -1, Next: nil}
	p := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
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
		l1 = l1.Next
		p = p.Next
	}
	for l2 != nil {
		p.Next = l2
		l2 = l2.Next
		p = p.Next
	}
	return dummy.Next
}
