package _2_queueAndStack

/*
	单调栈经典习题
*/

// https://leetcode.cn/problems/next-greater-node-in-linked-list/description/
// 链表中下一个更大的节点
func nextLargerNodes(head *ListNode) []int {
	//先把转为数组
	nums := make([]int, 0)
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}
	return NextGreaterElement(nums)
}

// https://leetcode.cn/problems/number-of-visible-people-in-a-queue/description/
// 队列中可以看到的人数
func canSeePersonsCount(heights []int) []int {
	length := len(heights)
	res := make([]int, length)
	p := make([]int, 0)
	for i := length - 1; i >= 0; i-- {
		count := 0
		for len(p) > 0 && heights[i] > p[len(p)-1] {
			p = p[:len(p)-1]
			count++
		}
		if len(p) == 0 {
			res[i] = count
		} else {
			res[i] = count + 1
		}
		p = append(p, heights[i])
	}
	return res
}

// https://leetcode.cn/problems/final-prices-with-a-special-discount-in-a-shop/

func finalPrices(prices []int) []int {
	length := len(prices)
	res := make([]int, length)
	p := make([]int, 0)
	for i := length - 1; i >= 0; i-- {
		for len(p) > 0 && prices[i] < p[len(p)-1] {
			p = p[:len(p)-1]
		}
		if len(p) == 0 {
			res[i] = 0
		} else {
			res[i] = p[len(p)-1]
		}
		p = append(p, prices[i])
	}
	trueRes := make([]int, length)
	for i := 0; i < length; i++ {
		trueRes[i] = prices[i] - res[i]
	}
	return trueRes
}
