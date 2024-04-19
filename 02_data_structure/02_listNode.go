package _2_data_structure

// 把存在的重复元素全部去除
func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{Val: -1}
	p := dummy
	for head != nil {
		if head.Next != nil && head.Val == head.Next.Val {
			for head.Next != nil && head.Val == head.Next.Val {
				head = head.Next
			}
			// 把自己也去除掉
			head = head.Next
			if head == nil {
				p.Next = nil
			}

		} else {
			p.Next = head
			p = p.Next
			head = head.Next
		}
	}
	return dummy.Next
}
