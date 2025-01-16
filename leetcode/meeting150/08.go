package meeting150

import "slices"

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	res := dummy
	add := 0
	for l1 != nil || l2 != nil {
		sum := add
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		if sum >= 10 {
			add = 1
		} else {
			add = 0
		}
		dummy.Next = &ListNode{Val: sum % 10}
		dummy = dummy.Next
	}
	if add == 1 {
		dummy.Next = &ListNode{Val: 1}
	}
	return res.Next
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	p := dummy
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			p.Next = &ListNode{
				Val: list1.Val,
			}
			list1 = list1.Next
		} else {
			p.Next = &ListNode{
				Val:  list2.Val,
				Next: nil,
			}
			list2 = list2.Next
		}
		p = p.Next
	}
	if list1 != nil {
		p.Next = list1
	}
	if list2 != nil {
		p.Next = list2
	}
	return dummy.Next
}

func copyRandomList(head *Node) *Node {
	hashMap := make(map[*Node]*Node)
	for p := head; p != nil; p = p.Next {
		hashMap[p] = &Node{Val: p.Val}
	}
	for p := head; p != nil; p = p.Next {
		if p.Next != nil {
			hashMap[p].Next = hashMap[p.Next]
		}
		if p.Random != nil {
			hashMap[p].Random = hashMap[p.Random]
		}
	}
	return hashMap[head]
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == 1 {
		return reverseN(head, right)
	}
	p := head
	// 找到需要翻转的前驱节点
	for i := 1; i < left-1; i++ {
		p = p.Next
	}
	p.Next = reverseN(p.Next, right-left+1)
	return head
}

// 返回翻转后的Node的头节点
func reverseN(head *ListNode, n int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var pre, cur, next *ListNode
	cur = head
	for n > 0 {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
		n--
	}
	head.Next = cur
	return pre
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}

	a, b := head, head
	for i := k; i > 0; i-- {
		if b == nil {
			return head
		}
		b = b.Next
	}
	last := reverseN(a, k)
	a.Next = reverseKGroup(b, k)
	return last
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	fast, slow := dummy, dummy
	for n > 0 {
		fast = fast.Next
		n--
	}
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}

func rotateRight(head *ListNode, k int) *ListNode {
	list := []int{}

	for head != nil {
		list = append(list, head.Val)
		head = head.Next
	}
	k = k % len(list)
	slices.Reverse(list)
	slices.Reverse(list[:k])
	slices.Reverse(list[k:])
	res := &ListNode{Val: list[0]}
	p := res
	for i := 1; i < len(list); i++ {
		p.Next = &ListNode{Val: list[i]}
		p = p.Next
	}
	return res

}

func partition(head *ListNode, x int) *ListNode {
	dummy1 := &ListNode{}
	dummy2 := &ListNode{}
	p1 := dummy1
	p2 := dummy2
	for head != nil {
		if head.Val < x {
			p1.Next = &ListNode{Val: head.Val}
			p1 = p1.Next
		} else {
			p2.Next = &ListNode{Val: head.Val}
			p2 = p2.Next
		}
		head = head.Next
	}
	p1.Next = dummy2.Next
	return dummy1.Next
}
