package linshan

import (
	"math"
	"sort"
)

// https://leetcode.cn/problems/two-sum-ii-input-array-is-sorted/
func twoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1
	for left < right {
		if numbers[left]+numbers[right] == target {
			return []int{left + 1, right + 1}
		} else if numbers[left]+numbers[right] < target {
			left++
		} else if numbers[left]+numbers[right] > target {
			right--
		}
	}
	return []int{-1, -1}
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	var res [][]int
	for i := 0; i < len(nums)-2; i++ {
		// 避免产生重复
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j := i + 1
		k := n - 1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[j], nums[k]})

				// 避免产生重复
				j++
				for j < k && nums[j] == nums[j-1] {
					j++
				}

				// 避免产生重复
				k--
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			} else if sum > 0 {
				k--
			} else {
				j++
			}
		}

	}
	return res
}

func maxArea(height []int) int {
	n := len(height)
	left := 0
	right := n - 1
	res := math.MinInt
	for left < right {
		area := (right - left) * min(height[left], height[right])
		res = max(res, area)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return res
}

// 接雨水
func trap(height []int) int {
	n := len(height)
	leftMax := make([]int, n)
	leftMax[0] = height[0]
	for i := 1; i < n; i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}
	rightMax := make([]int, n)
	rightMax[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}
	res := 0
	for i := 0; i < n; i++ {
		res += min(rightMax[i], leftMax[i]) - height[i]
	}
	return res
}

func trap2(height []int) int {
	n := len(height)
	leftMax := 0
	rightMax := 0
	left := 0
	right := n - 1
	res := 0
	for left < right {
		leftMax = max(leftMax, height[left])
		rightMax = max(rightMax, height[right])
		if leftMax < rightMax {
			res += leftMax - height[left]
			left++
		} else {
			res += rightMax - height[right]
			right--
		}
	}
	return res
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{
		Val:  0,
		Next: head,
	}
	fast := dummy
	slow := dummy
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	// 此时slow指向倒数第n+1的节点
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}

// 给定一个已排序的链表的头 head ， 删除所有重复的元素，使每个元素只出现一次 。返回 已排序的链表 。
func deleteDuplicates1(head *ListNode) *ListNode {
	//if head == nil {
	//	return head
	//}
	//p := head
	//for head.Next != nil {
	//	if head.Val == head.Next.Val {
	//		// 移除
	//		head.Next = head.Next.Next
	//	} else {
	//		head = head.Next
	//	}
	//}
	//return p
	cur := head
	for cur != nil {
		for cur.Next != nil && cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		}
		cur = cur.Next
	}
	return head
}

// 给定一个已排序的链表的头 head ， 删除所有重复的元素，剩下来的元素保证只出现过一次。返回 已排序的链表 。
func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	cur := dummy
	for cur.Next != nil && cur.Next.Next != nil {
		val := cur.Next.Val
		if cur.Next.Next.Val == val {
			for cur.Next != nil && cur.Next.Val == val {
				cur.Next = cur.Next.Next
			}
		} else {
			cur = cur.Next
		}
	}
	return dummy.Next
}
