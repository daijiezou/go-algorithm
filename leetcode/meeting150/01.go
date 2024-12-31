package meeting150

func merge(nums1 []int, m int, nums2 []int, n int) {
	mi, ni, p := m-1, n-1, len(nums1)-1
	for mi >= 0 && ni >= 0 {
		if nums1[mi] > nums2[ni] {
			nums1[p] = nums1[mi]
			p--
			mi--
		} else {
			nums1[p] = nums2[ni]
			p--
			ni--
		}
	}
	for mi >= 0 {
		nums1[p] = nums1[mi]
		p--
		mi--
	}
	for ni >= 0 {
		nums1[p] = nums2[ni]
		p--
		ni--
	}
}

func removeElement(nums []int, val int) int {
	slow := 0
	n := len(nums)
	res := 0
	for i := 0; i < n; i++ {
		if nums[i] != val {
			nums[slow] = nums[i]
			slow++
			res++
		}
	}
	return res
}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
		fast++
	}
	return slow + 1
}

func removeDuplicates2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	slow, fast := 0, 0
	cnt := 1
	for fast < len(nums) {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		} else if slow < fast && cnt < 3 {
			slow++
			nums[slow] = nums[fast]
		}
		fast++
		cnt++
		if fast < len(nums) && nums[fast] != nums[fast-1] {
			cnt = 1
		}
	}
	return slow + 1
}

func majorityElement(nums []int) int {
	//if len(nums) == 1 {
	//	return nums[0]
	//}
	//n := len(nums)
	//slices.Sort(nums)
	//for i := 1; i < n; i++ {
	//	cnt := 1
	//	for i < n && nums[i] == nums[i-1] {
	//		cnt++
	//		if cnt > n/2 {
	//			return nums[i]
	//		}
	//		i++
	//	}
	//}
	//return -1
	numCnt := make(map[int]int)
	n := len(nums)
	targetCnt := n / 2
	for i := 0; i < n; i++ {
		numCnt[nums[i]]++
		if numCnt[nums[i]] > targetCnt {
			return nums[i]
		}
	}
	return -1
}

func maxProfit(prices []int) int {
	s := make([]int, 0)
	res := 0
	revenuePrice := make(map[int]int)
	for i := 0; i < len(prices); i++ {
		revenue := 0
		for len(s) > 0 && prices[i] > prices[s[len(s)-1]] {
			pop := s[len(s)-1]
			s = s[:len(s)-1]
			revenue = max(prices[i]-prices[pop]+revenuePrice[pop], revenue)
			revenuePrice[i] = revenue
			res = max(res, revenue)
		}
		s = append(s, i)
	}
	return res
}

func maxProfit2(prices []int) int {
	minPrice := prices[0]
	res := 0
	for i := 0; i < len(prices); i++ {
		minPrice = min(prices[i], minPrice)
		res = max(prices[i]-minPrice, res)
	}
	return res
}

func maxProfit_k1(prices []int) int {
	dp := make([][2]int, 0)
	n := len(prices)
	for i := 0; i < n; i++ {
		if i == 0 {
			dp[i][0] = 0
			dp[i][1] = -prices[i]
			continue
		}
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], -prices[i])
	}
	return dp[n-1][0]
}

func maxProfit_kk(prices []int) int {

	n := len(prices)
	dp := make([][2]int, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			dp[i][0] = 0
			dp[i][1] = -prices[i]
			continue
		}
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i]) // 能继承之前的，表示能够操作多次

	}
	return dp[n-1][0]
}

func canJump(nums []int) bool {
	nextRight := 0
	curRight := 0
	n := len(nums)
	for i := 0; i < n-1; i++ {
		nextRight = max(nextRight, nums[i]+i)
		if curRight == i {
			if nextRight == i { // 无法到达i+1
				return false
			}
			curRight = nextRight
		}
	}
	return true
}

func jump(nums []int) int {
	n := len(nums)
	memo := make(map[int]int)
	for i := 0; i < n; i++ {
		memo[i] = n
	}

	return jumpDp(nums, 0, memo)
}

func jumpDp(nums []int, start int, memo map[int]int) int {
	if start >= len(nums)-1 {
		return 0
	}
	if memo[start] != len(nums) {
		return memo[start]
	}
	steps := nums[start]
	for i := 1; i <= steps; i++ {
		// 穷举每一个选择
		// 计算每一个子问题的结果
		step := jumpDp(nums, start+i, memo) + 1
		memo[start] = min(memo[start], step)
	}
	return memo[start]
}

func jump2(nums []int) int {
	n := len(nums)
	curEnd := 0  // 当前终点
	nextEnd := 0 //下一个终点
	step := 0
	for i := 0; i < n; i++ {
		nextEnd = max(nextEnd, nums[i]+i)
		if curEnd == i { //到达当前终点
			curEnd = nextEnd //再造一个新桥
			step++
		}
	}
	return step
}

func minTaps(n int, ranges []int) int {
	rightMost := make([]int, n+1)
	for i, r := range ranges {
		left := max(i-r, 0)
		rightMost[left] = max(rightMost[left], i+r)
	}
	curEnd := 0  // 当前终点
	nextEnd := 0 //下一个终点
	cnt := 0
	for i := 0; i < n; i++ {
		nextEnd = max(nextEnd, rightMost[i]+i)
		if curEnd == i { //到达当前终点
			if nextEnd == i { // 无法到达i+1
				return -1
			}
			curEnd = nextEnd //再造一个新桥
			cnt++
		}
	}
	return cnt
}
