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
	return mergeSort(head)
}

func mergeSort(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 找到链表中点进行切分
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	mid := slow
	// 断开链表
	midNext := mid.Next
	mid.Next = nil
	left := mergeSort(head)
	right := mergeSort(midNext)
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

func sortListDJ(head *ListNode) *ListNode {
	return sortDj(head)
}

func sortDj(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	// 只有一个节点
	if head.Next == nil {
		return head
	}
	fast := head
	slow := head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	mid := slow.Next
	slow.Next = nil
	return mergeTwoList(sortDj(head), sortDj(mid))
}

func mergeTwoList(A, B *ListNode) *ListNode {
	dummy := &ListNode{Val: -1, Next: nil}
	p := dummy
	for A != nil && B != nil {
		if A.Val < B.Val {
			p.Next = A
			A = A.Next
		} else {
			p.Next = B
			B = B.Next
		}
		p = p.Next
	}
	if A != nil {
		p.Next = A
	} else {
		p.Next = B
	}
	return dummy.Next
}

func deleteDuplicatesDJ(head *ListNode) *ListNode {
	dummy := &ListNode{Val: -1, Next: nil}
	p := dummy
	for head != nil {
		if head.Next != nil && head.Val == head.Next.Val {
			for head.Next != nil && head.Next.Val == head.Val {
				head = head.Next
			}
			// 把自己也去除掉
			head = head.Next
			if head == nil {
				p.Next = nil
			}
		} else {
			p.Next = head
			head = head.Next
			p = p.Next
		}
	}
	return dummy.Next
}

func deleteDuplicatesUnsortedDJ(head *ListNode) *ListNode {
	count := make(map[int]int)
	// 先遍历一遍链表，记录每个值出现的次数
	p := head
	for p != nil {
		count[p.Val]++
		p = p.Next
	}
	// 虚拟头结点（哨兵节点），存放结果链表
	dummy := &ListNode{Val: -1, Next: head}
	// 再遍历一遍节点，把重复出现的节点剔除
	p = dummy
	for head != nil {
		for count[head.Val] > 1 {
			// 跳过重复节点，直到找到不重复的节点
			head = head.Next
		}
		// 接入不重复的节点或尾部空指针
		p.Next = head
		// p 前进，继续寻找不重复节点
		head = head.Next
	}
	return dummy.Next
}
