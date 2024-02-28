package _1_listNodeAndArray

import (
	"fmt"
	"strings"
)

// https://leetcode.cn/problems/remove-duplicates-from-sorted-list-ii/description/
func deleteDuplicates2(head *ListNode) *ListNode {
	dummy := &ListNode{-1, nil}
	p, q := dummy, head
	for q != nil {
		// 发现重复元素
		if q.Next != nil && q.Val == q.Next.Val {
			// 一直检查看是否有多个重复元素
			for q.Next != nil && q.Val == q.Next.Val {
				q = q.Next
			}
			// 跳过这段重复元素
			q = q.Next
			// 如果一直到最后元素都是重复的，则将p的后面的链表全部删除
			if q == nil {
				p.Next = nil
			}
		} else {
			// 没有重复把q接到p的后面
			p.Next = q
			q = q.Next
			p = p.Next
		}
	}

	return dummy.Next
}

// 使用递归的方法解决上面的问题
func deleteDuplicates3(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	if head.Val != head.Next.Val {
		head.Next = deleteDuplicates(head.Next)
		return head
	}
	for head.Next != nil && head.Val == head.Next.Val {
		head = head.Next
	}
	return deleteDuplicates(head.Next)
}

// https://leetcode.cn/problems/remove-duplicates-from-an-unsorted-linked-list
// 给定一个链表的第一个节点 head，找到链表中所有出现多于一次的元素，
// 并删除这些元素所在的节点，返回删除后的链表。
func deleteDuplicatesUnsorted(head *ListNode) *ListNode {
	p := head
	valCountMap := make(map[int]int)
	for p != nil {
		if _, ok := valCountMap[p.Val]; ok {
			valCountMap[p.Val]++
		} else {
			valCountMap[p.Val] = 1
		}
		p = p.Next
	}
	dummy := &ListNode{-1, head}
	p = dummy
	for p != nil {
		// unique负责寻找不重复的节点
		unique := p.Next
		// 跳过重复的节点
		for unique != nil && valCountMap[p.Val] > 1 {
			unique = unique.Next
		}
		// 接入不重复的节点或尾部空指针
		p.Next = unique
		// p前进寻找下一个不重复的节点
		p = p.Next
	}
	return dummy.Next
}

// 丑数1
func isUgly(n int) bool {
	if n <= 0 {
		return false
	}
	for n%2 == 0 {
		n = n / 2
	}
	for n%3 == 0 {
		n = n / 3
	}
	for n%5 == 0 {
		n = n / 5
	}
	return n == 1
}

// 丑数2
// https://leetcode.cn/problems/ugly-number-ii/description/
func nthUglyNumber(n int) int {
	// 三个有序链表头节点的指针
	p2, p3, p5 := 1, 1, 1
	product2, product3, product5 := 1, 1, 1
	ugly := make([]int, n+1)
	p := 1
	for p <= n {
		// 取最小值加到结果链上，顺便得到了去重的效果
		min := myMin(product2, product3, product5)
		ugly[p] = min
		p++
		// 查看现在在哪条链上
		if min == product2 {
			product2 = 2 * ugly[p2]
			p2++
		}
		if min == product3 {
			product3 = 3 * ugly[p3]
			p3++
		}
		if min == product5 {
			product5 = 5 * ugly[p5]
			p5++
		}
	}
	return ugly[n]
}

func myMin(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
	} else {
		if b < c {
			return b
		}
	}
	return c
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	lists := make([][][]int, 0)
	for _, numsV := range nums1 {
		nums1List := [][]int{}
		for _, nums2V := range nums2 {
			shudui := []int{numsV, nums2V}
			nums1List = append(nums1List, shudui)
		}
		lists = append(lists, nums1List)
	}
	return kSmallestPairsMerge(0, len(lists)-1, lists)[:k]
}

func merge2SmallestPairs(shudui1 [][]int, shudui2 [][]int) [][]int {
	result := [][]int{}
	for len(shudui1) > 0 && len(shudui2) > 0 {
		if ListSum(shudui1[0]) > ListSum(shudui2[0]) {
			result = append(result, shudui2[0])
			shudui2 = shudui2[1:]
		} else {
			result = append(result, shudui1[0])
			shudui1 = shudui1[1:]
		}
	}
	if len(shudui2) > 0 {
		result = append(result, shudui2...)
	}
	if len(shudui1) > 0 {
		result = append(result, shudui1...)
	}
	return result

}

func ListSum(list []int) int {
	sum := 0
	for _, i2 := range list {
		sum += i2
	}
	return sum
}

func kSmallestPairsMerge(left, right int, lists [][][]int) [][]int {
	if right == left {
		return lists[right]
	}
	if left > right {
		return [][]int{}
	}
	mid := (left + right) >> 1
	return merge2SmallestPairs(kSmallestPairsMerge(left, mid, lists), kSmallestPairsMerge(mid+1, right, lists))
}

// 将双指针初始化在数组的尾部，然后从后向前进行合并，这样即便覆盖了 nums1 中的元素，这些元素也必然早就被用过了，不会影响答案的正确性。

func mergeNum(nums1 []int, m int, nums2 []int, n int) {
	i, j, p := m-1, n-1, len(nums1)-1 // 将指针放到末尾
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[p] = nums1[i]
			i--
		} else {
			nums1[p] = nums2[j]
			j--
		}
		p--
	}

	for j >= 0 {
		nums1[p] = nums2[j]
		j--
		p--
	}
	fmt.Println(nums1)
}

// https://leetcode.cn/problems/squares-of-a-sorted-array/
// 有序数组的平方
func sortedSquares(nums []int) []int {
	left, right := 0, len(nums)-1
	result := make([]int, len(nums))
	index := len(nums) - 1
	for index >= 0 {
		if nums[left]*nums[left] > nums[right]*nums[right] {
			result[index] = nums[left] * nums[left]
			left++
		} else {
			result[index] = nums[right] * nums[right]
			right--
		}
		index--
	}
	return result
}

// 给你一个已经排好序的整数数组 nums 和整数 a, b, c。
// 对于数组中的每一个元素 nums[i]，计算函数值 f(x) = ax2 + bx + c，请按升序返回结果数组。
func sortTransformedArray(nums []int, a int, b int, c int) []int {
	f := func(x int) int {
		return a*x*x + b*x + c
	}
	result := make([]int, len(nums))
	left, right := 0, len(nums)-1
	if a > 0 {
		index := len(nums) - 1
		for index >= 0 {
			if f(nums[left]) > f(nums[right]) {
				result[index] = f(nums[left])
				left++
			} else {
				result[index] = f(nums[right])
				right--
			}
			index--
		}
	} else {
		index := 0
		for index < len(nums) {
			if f(nums[left]) < f(nums[right]) {
				result[index] = f(nums[left])
				left++
			} else {
				result[index] = f(nums[right])
				right--
			}
			index++
		}
	}
	return result
}

// https://leetcode.cn/problems/reverse-words-in-a-string/description/
func reverseWords(s string) string {
	for strings.Index(s, "  ") >= 0 {
		s = strings.Replace(s, "  ", " ", -1)
	}
	fmt.Println(s)
	srune := []rune(s)
	if string(srune[0]) == " " {
		srune = srune[1:]
	}
	if string(srune[len(srune)-1]) == " " {
		srune = srune[:len(srune)-1]
	}
	s = string(srune)
	res := strings.Split(s, " ")

	filterRes := []string{}
	for _, v := range res {
		if v != " " {
			filterRes = append(filterRes, v)
		}
	}
	for i := 0; i < len(filterRes)/2; i++ {
		temp := filterRes[i]
		filterRes[i] = filterRes[len(filterRes)-1-i]
		filterRes[len(filterRes)-1-i] = temp

	}

	return strings.Join(filterRes, " ")
}
