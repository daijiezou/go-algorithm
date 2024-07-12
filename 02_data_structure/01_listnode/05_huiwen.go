package _1_listnode

var isPalindromeleft *ListNode

func isPalindrome(head *ListNode) bool {
	//defer func() {
	//	isPalindromeleft = nil
	//}()
	//isPalindromeleft = head
	//return isPalindromeTraverse(head)
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	if fast != nil {
		slow = slow.Next
	}
	rev := myReverse(slow)
	for rev != nil && head != nil {
		if rev.Val != head.Val {
			return false
		}
		rev = rev.Next
		head = head.Next
	}

	return true
}

func myReverse(head *ListNode) *ListNode {
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

// 使用后续遍历的方式来判断是否是回文链表
func isPalindromeTraverse(head *ListNode) bool {
	left := head
	res := true
	var traverse func(head *ListNode)
	traverse = func(head *ListNode) {
		if head == nil {
			return
		}
		traverse(head.Next)
		// 后序
		if left.Val != head.Val {
			res = false
			return
		}
		left = left.Next
	}
	traverse(head)
	return res
}
