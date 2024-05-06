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

// 使用后续遍历的方式来判断是否是回文链表
func isPalindromeTraverse(right *ListNode) bool {
	if right == nil {
		return true
	}
	res := isPalindromeTraverse(right.Next)
	res = res && isPalindromeleft.Val == right.Val
	isPalindromeleft = isPalindromeleft.Next
	return res
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
