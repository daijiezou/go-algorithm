package leetcode

import (
	"slices"
	"strconv"
)

func mostPoints(questions [][]int) int64 {
	n := len(questions)
	memo := make([]int, len(questions))
	var dfs func(i int) int
	dfs = func(i int) int {
		if i >= n {
			return 0
		}
		if memo[i] != 0 {
			return memo[i]
		}
		q := questions[i]
		memo[i] = max(q[0]+dfs(i+q[1]+1), dfs(i+1))
		return memo[i]
	}
	return int64(dfs(0))
}

func mostPoints2(questions [][]int) int64 {
	n := len(questions)
	dp := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		q := questions[i]
		j := min(n, i+q[1]+1)
		dp[i] = max(dp[i+1], dp[j]+q[0])
	}
	return int64(dp[0])
}

func maximumTripletValue(nums []int) int64 {
	n := len(nums)
	res := 0
	//preMax := make([]int, n+1)
	sufMax := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		sufMax[i] = max(sufMax[i+1], nums[i])
	}

	preMax := 0
	for i := 0; i < n; i++ {
		res = max(res, (preMax-nums[i])*sufMax[i+1])
		preMax = max(preMax, nums[i])

	}
	return int64(res)
}

func subsetXORSum(nums []int) int {
	n := len(nums)
	x := 0
	for i := 0; i < n; i++ {
		x |= nums[i]
	}
	return x << (n - 1)
}

func subsetXORSum2(nums []int) int {
	n := len(nums)
	var dfs func(res, i int) int
	dfs = func(res int, i int) int {
		if i == n {
			return res
		}
		return dfs(res^nums[i], i+1) + dfs(res, i+1)
	}
	return dfs(0, 0)
}

func largestDivisibleSubset(nums []int) []int {
	slices.Sort(nums)
	n := len(nums)
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	maxSize := 1
	maxVal := 1
	for i := 1; i < n; i++ {
		for j, v := range nums[:i] {
			if nums[i]%v == 0 && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
			}
		}
		if dp[i] > maxSize {
			maxSize, maxVal = dp[i], nums[i]
		}
	}
	if maxSize == 1 {
		return []int{nums[0]}
	}
	res := make([]int, 0)
	for i := n - 1; i >= 0 && maxSize > 0; i-- {
		if dp[i] == maxSize && maxVal%nums[i] == 0 {
			res = append(res, nums[i])
			maxSize--
			maxVal = nums[i]
		}
	}
	return res
}

func canPartition(nums []int) bool {
	sum := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		sum += nums[i]
	}
	if sum%2 != 0 {
		return false
	}
	memo := make([][]int, n)
	for i := 0; i < n; i++ {
		memo[i] = make([]int, sum)
	}
	target := sum / 2
	var dfs func(start int, preSum int) bool
	dfs = func(start int, preSum int) (res bool) {
		if start == n {
			return false
		}
		if memo[start][preSum] != 0 {
			return memo[start][preSum] == 1
		}
		p := &memo[start][preSum]
		defer func() {
			if res {
				*p = 1
			} else {
				*p = 2
			}
		}()
		x := nums[start]
		if preSum+x == target {
			return true
		}
		if preSum+x > target {
			return dfs(start+1, preSum)
		}
		return dfs(start+1, preSum+x) || dfs(start+1, preSum)
	}
	return dfs(0, 0)
}

func minimumOperations(nums []int) int {
	lastRepeatIndex := 0
	dup := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := dup[nums[i]]; ok {
			lastRepeatIndex = max(dup[nums[i]], lastRepeatIndex)
		}
		dup[nums[i]] = i + 1
	}
	if lastRepeatIndex == 0 {
		return 0
	}
	add := 0
	if lastRepeatIndex%3 != 0 {
		add = 1
	}
	return (lastRepeatIndex)/3 + add
}

func minimumOperations2(nums []int) int {
	dup := make(map[int]int)
	n := len(nums)
	for i := n - 1; i >= 0; i-- {
		if _, ok := dup[nums[i]]; ok {
			return i/3 + 1
		}
		dup[nums[i]] = i
	}
	return 0
}

func minOperations(nums []int, k int) int {

	n := len(nums)
	tempMap := make(map[int]struct{})
	for i := 0; i < n; i++ {
		if nums[i] < k {
			return -1
		}
		tempMap[nums[i]] = struct{}{}
	}
	delete(tempMap, k)
	return len(tempMap)
}

func countSymmetricIntegers(low int, high int) int {
	res := 0
	for i := low; i <= high; i++ {
		strNum := strconv.Itoa(i)
		if len(strNum)%2 != 0 {
			continue
		}
		n := len(strNum)
		pre := 0
		suf := 0
		for j := 0; j < n/2; j++ {
			pre += int(strNum[j] - '0')
			suf += int(strNum[n-j-1] - '0')
		}
		if pre == suf {
			res++
		}
	}
	return res
}
