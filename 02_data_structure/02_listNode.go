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

// 从未排序的链表中移除重复元素
func deleteDuplicatesUnsorted(head *ListNode) *ListNode {
	valCount := make(map[int]int)
	p := head
	for p != nil {
		if _, ok := valCount[p.Val]; !ok {
			valCount[p.Val] = 1
		} else {
			valCount[p.Val]++
		}
		p = p.Next
	}
	dummy := &ListNode{Val: -1}
	p = dummy
	for head != nil {
		for valCount[head.Val] > 1 {
			head = head.Next
		}
		p.Next = head
		p = p.Next
		head = head.Next
	}
	return dummy.Next
}

// https://leetcode.cn/problems/ugly-number-ii/description/
func nthUglyNumber(n int) int {
	p2, p3, p5 := 1, 1, 1
	product2, product3, product5 := 1, 1, 1
	// 可以理解为最终合并的有序链表（结果链表）
	ugly := make([]int, n+1)
	// 可以理解为结果链表上的指针
	for i := 1; i <= n; i++ {
		res := Mymin(product2, product3, product5)
		ugly[i] = res
		if res == product2 {
			product2 = ugly[p2] * 2
			p2++
		}
		if res == product3 {
			product3 = ugly[p3] * 3
			p3++
		}
		if res == product5 {
			product5 = ugly[p5] * 5
			p5++
		}
	}
	return ugly[n]
}

// 取三个数的最小值
func Mymin(i, j, k int) int {
	if i < j {
		if i < k {
			return i
		} else {
			return k
		}
	} else {
		if j < k {
			return j
		} else {
			return k
		}
	}
}
