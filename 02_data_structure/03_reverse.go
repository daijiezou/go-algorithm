package _2_data_structure

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

var nextNode *ListNode

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	defer func() {
		nextNode = nil
	}()

	/*	if left == 1 {
			return reverseN(head, right)
		}
		// 前进到反转的起点触发 base case
		head.Next = reverseBetween(head.Next, left-1, right-1)
		return head*/

	n := right - left + 1
	temp := head
	var pre *ListNode
	if left == 1 {
		pre = head
	}
	for ; left > 1; left-- {
		if left == 2 {
			pre = temp
		}
		temp = temp.Next
	}

	last := reverseN(temp, n)
	pre.Next = last
	return head

}
