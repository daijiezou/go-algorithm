package _1_core_platform

// 去除数组中重复元素
func removeDuplicates(nums []int) int {
	length := len(nums)
	slow := 0
	fast := 0
	for ; fast < length; fast++ {
		if nums[slow] != nums[fast] {
			slow++
			nums[slow] = nums[fast]
		}
	}
	return slow + 1
}

func removeElement(nums []int, val int) int {
	length := len(nums)
	slow := 0
	fast := 0
	for ; fast < length; fast++ {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}

func moveZeroes(nums []int) {
	length := len(nums)
	slow := 0
	fast := 0
	for ; fast < length; fast++ {
		if nums[fast] != 0 {
			nums[slow] = nums[fast]
			slow++
		}
	}
	for ; slow < length; slow++ {
		nums[slow] = 0
	}
}

func twoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left, right}
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

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow, fast := head, head
	for fast != nil {
		if fast.Val != slow.Val {
			slow.Next = fast
			slow = slow.Next
		}
		fast = fast.Next
	}
	slow.Next = nil
	return head
}

// 反转字符串
func reverseString(s []byte) {
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

// 回文串
func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// 最长回文子串
func longestPalindrome(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		// 以 s[i] 为中心的最长回文子串
		s1 := getlongestPalindrome(s, i, i)
		//奇数情况
		s2 := getlongestPalindrome(s, i, i+1)
		if len(s1) > len(res) {
			res = s1
		}
		if len(s2) > len(res) {
			res = s2
		}
	}
	return res
}

func getlongestPalindrome(s string, left int, right int) string {
	for left >= 0 && right < len(s) {
		if s[left] == s[right] {
			left--
			right++
		} else {
			break
		}
	}
	return s[left+1 : right]
}
