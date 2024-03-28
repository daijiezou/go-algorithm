package _1_listNodeAndArray

/*
	快慢指针
*/

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	fast := head
	slow := head
	for fast != nil {
		if fast.Val != slow.Val {
			slow = slow.Next
			slow.Val = fast.Val

		}
		fast = fast.Next
	}
	// 断开后面的节点
	slow.Next = nil
	return head
}

// https://leetcode.cn/problems/remove-element/description/
func removeElement(nums []int, val int) int {
	fast, slow := 0, 0
	length := len(nums)
	for fast < length {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}

// https://leetcode.cn/problems/move-zeroes/submissions/503416318/
func moveZeroes(nums []int) {
	res := removeElement(nums, 0)
	for i := res; i < len(nums); i++ {
		nums[i] = 0
	}
}

// https://leetcode.cn/problems/two-sum-ii-input-array-is-sorted/description/
func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		}
		if sum < target {
			left++
		}
		if sum > target {
			right--
		}
	}
	return []int{}
}

// https://leetcode.cn/problems/longest-palindromic-substring/
func longestPalindrome(s string) string {
	sr := []rune(s)
	length := len(sr)
	var res []rune
	for i := 0; i < length; i++ {
		s1 := palindrome(sr, i, i)
		s2 := palindrome(sr, i, i+1)
		if len(s1) > len(res) {
			res = s1
		}
		if len(s2) > len(res) {
			res = s2
		}
	}
	return string(res)
}

func palindrome(s []rune, left int, right int) []rune {
	length := len(s)
	for left >= 0 && right < length && s[left] == s[right] {
		left--
		right++
	}
	// left多减了一次需要加回来，right因为是包前不包后的所以不需要再减回来
	return s[left+1 : right]
}
