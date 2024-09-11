package leetcode

import "sort"

func busyStudent(startTime []int, endTime []int, queryTime int) int {
	n := len(startTime)
	res := 0
	for i := 0; i < n; i++ {
		if startTime[i] <= queryTime && queryTime <= endTime[i] {
			res++
		}
	}
	return res
}

func maxConsecutiveAnswers(answerKey string, k int) int {
	n := len(answerKey)
	left, right := 0, 0
	maxConsecutive := 0
	countT := 0
	countF := 0
	for right < n {
		if answerKey[right] == 'T' {
			countT++
		} else {
			countF++
		}
		right++
		for countT > k && countF > k {
			if answerKey[left] == 'T' {
				countT--
			} else {
				countF--
			}
			left++
		}
		maxConsecutive = max(maxConsecutive, right-left)
	}
	return maxConsecutive
}

func maxStrength(nums []int) int64 {
	if len(nums) == 1 {
		return int64(nums[0])
	}
	minIndex := 0
	minNum := -10
	count := 0
	zeroCnt := 0
	res := int64(1)
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			zeroCnt++
			continue
		}
		res *= int64(nums[i])
		if nums[i] < 0 {
			count++
			if nums[i] > minNum {
				minNum = nums[i]
				minIndex = i
			}
		}
	}

	// 当数组不包含正数，且负数元素小于等于 1 个时，最大积为 0。
	if zeroCnt+1 == len(nums) && count == 1 || zeroCnt == len(nums) {
		return 0
	}

	if count%2 == 0 {
		return res
	} else {
		return res / int64(nums[minIndex])
	}
}

/*
如果能够满足下述两个条件之一，则认为第 i 位学生将会保持开心：
这位学生被选中，并且被选中的学生人数 严格大于 nums[i] 。
这位学生没有被选中，并且被选中的学生人数 严格小于 nums[i] 。
*/
func countWays(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	res := 0
	// 可以都不选
	if nums[0] > 0 {
		res++
	}
	for i := 1; i < n; i++ {
		// i代表被选中的人数

		if nums[i-1] < i && // 被选中的学生人数 严格大于 nums[i]
			i < nums[i] { // 被选中的学生人数 严格小于 nums[i]
			res++
		}
	}

	// 0 <= nums[i] < nums.length
	// 一定可以都选
	return res + 1
}

/*
https://leetcode.cn/problems/clear-digits/description/
删除 第一个数字字符 以及它左边 最近 的 非数字 字符。
返回删除所有数字字符以后剩下的字符串。
*/
func clearDigits(s string) string {
	res := []byte{}
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			if len(res) > 0 {
				res = res[:len(res)-1]
			}
		} else {
			res = append(res, s[i])
		}
	}
	return string(res)
}

func maximumLength(nums []int, k int) int {
	// dp[i][j] 来表示以 nums[i] 结尾组成的最长合法序列中有 j 个数字与其在序列中的后一个数字不相等。
	//其中 i 的取值为 nums 的长度，j 不超过 k。初始时，有 dp[i][0]=1。
	n := len(nums)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, k+1)
		for j := 0; j <= k; j++ {
			dp[i][j] = -1
		}
	}
	ans := 0
	for i := 0; i < n; i++ {
		dp[i][0] = 1
		for l := 0; l <= k; l++ {
			for j := 0; j < i; j++ {
				add := 0
				if nums[i] != nums[j] {
					add = 1
				}
				if l-add >= 0 && dp[j][l-add] != -1 {
					dp[i][l] = max(dp[i][l], dp[j][l-add]+1)
				}
			}
			ans = max(ans, dp[i][l])
		}
	}
	return ans
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.cn/problems/merge-nodes-in-between-zeros/description/
func mergeNodes(head *ListNode) *ListNode {
	dummy := &ListNode{
		Val:  -1,
		Next: nil,
	}
	p := dummy
	curNum := 0
	head = head.Next
	for head != nil {
		curNum += head.Val
		if head.Val == 0 {
			p.Next = &ListNode{
				Val:  curNum,
				Next: nil,
			}
			p = p.Next
			curNum = 0
		}
		head = head.Next

	}
	return dummy.Next
}

func mergeNodes2(head *ListNode) *ListNode {
	tail := head
	cur := head.Next
	for cur.Next != nil {
		if cur.Val != 0 {
			tail.Val += cur.Val
		} else {
			tail = tail.Next
			tail.Val = 0
		}
		cur = cur.Next

	}
	tail.Next = nil
	return head
}

func maximizeWin(prizePositions []int, k int) int {
	n := len(prizePositions)
	mx := make([]int, n+1)
	ans, left := 0, 0
	for right, p := range prizePositions {
		for p-prizePositions[left] > k {
			left++
		}
		ans = max(ans, right-left+1+mx[left])
		mx[right+1] = max(mx[right], right-left+1)
	}
	return ans
}
