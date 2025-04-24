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

// https://leetcode.cn/problems/count-good-numbers/?envType=daily-question&envId=2025-04-13
func countGoodNumbers(n int64) int {
	mod := 1000000007
	x := 5
	x2 := 4
	x1Cnt := n/2 + n%2
	x2Cnt := n / 2

	return (pow_mod(x, int(x1Cnt), mod) % mod) * (pow_mod(x2, int(x2Cnt), mod) % mod) % mod
}

func pow_mod(x, y, mod int) int {
	res := 1
	for y > 0 {
		if (y & 1) == 1 {
			res = res * x % mod
		}
		x = x * x % mod
		y >>= 1
	}
	return res
}

func countGoodTriplets(arr []int, a int, b int, c int) int {
	n := len(arr)
	res := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if Myabs(arr[i], arr[j]) > a {
				continue
			}
			for k := j + 1; k < n; k++ {
				if Myabs(arr[i], arr[k]) <= b && Myabs(arr[j], arr[k]) <= c {
					res++
				}
			}
		}
	}
	return res
}

// https://leetcode.cn/problems/count-the-number-of-good-subarrays/?envType=daily-question&envId=2025-04-16
func countGood(nums []int, k int) int64 {
	n := len(nums)
	res := int64(0)
	left := 0
	numCnt := make(map[int]int)
	curCnt := 0
	for right := 0; right < n; right++ {
		x := nums[right]
		curCnt += numCnt[x]
		numCnt[x]++

		for curCnt >= k {

			leftx := nums[left]
			numCnt[leftx]--
			curCnt -= numCnt[leftx]
			left++
		}
		res += int64(left)
	}
	return res
}

func countPairs(nums []int, k int) int {
	res := 0
	n := len(nums)
	equalValueNumbs := make(map[int][]int)
	numbCntMods := make(map[int]int)

	for i := 0; i < n; i++ {
		x := nums[i]
		if i%k == 0 {
			res += len(equalValueNumbs[x])
			numbCntMods[x]++
		} else {
			for _, index := range equalValueNumbs[x] {
				if (i*index)%k == 0 {
					res++
				}
			}
		}
		equalValueNumbs[x] = append(equalValueNumbs[x], i)
	}

	return res
}

func numIdenticalPairs(nums []int) int {
	numsCnt := make(map[int]int)
	res := 0
	for _, x := range nums {
		res += numsCnt[x]
		numsCnt[x]++
	}
	return res
}

func countBadPairs(nums []int) int64 {
	numsCnt := make(map[int]int)
	n := len(nums)
	res := n * (n - 1) / 2
	for i, x := range nums {
		res -= numsCnt[x-i]
		numsCnt[x-i]++
	}
	return int64(res)
}

func countFairPairs(nums []int, lower int, upper int) int64 {
	slices.Sort(nums)

	res := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		x := nums[i]
		start, ok := slices.BinarySearch(nums, lower-x)
		if !ok {
			continue
		}
		end, _ := slices.BinarySearch(nums, upper-x+1)
		end = end - 1
		start = max(start, i+1)

		cnt := end - start + 1
		if cnt >= 0 {
			res += cnt
		}
	}
	return int64(res)
}

func numRabbits(answers []int) int {
	numCnt := make(map[int]int)
	res := 0
	for _, v := range answers {
		numCnt[v+1]++
	}

	for k, v := range numCnt {
		add := 0
		if v%(k) != 0 {
			add = 1
		}
		res += (v/(k) + add) * k
	}
	return res
}

func numberOfArrays(differences []int, lower int, upper int) int {
	origin := 0
	add := 0
	sub := 0
	for i := 0; i < len(differences); i++ {
		origin += differences[i]
		add = max(add, origin)
		sub = min(sub, origin)
	}
	return max(0, upper-lower-(add-sub)+1)
}

func countLargestGroup(n int) int {
	cnt := make(map[int]int)
	res := 0
	maxCnt := 0
	for i := 1; i <= n; i++ {
		ds := 0
		for j := i; j != 0; j /= 10 {
			ds += j % 10
		}
		cnt[ds]++
		if cnt[ds] > maxCnt {
			maxCnt = cnt[ds]
			res = 1
		} else if cnt[ds] == maxCnt {
			res++
		}
	}
	return res
}

func countCompleteSubarrays(nums []int) int {
	set := make(map[int]struct{})
	n := len(nums)
	for i := 0; i < n; i++ {
		set[nums[i]] = struct{}{}
	}
	target := len(set)
	window := make(map[int]int)
	left := 0
	res := 0
	for i := 0; i < n; i++ {
		window[nums[i]]++
		for len(window) == target {
			x := nums[left]
			left++
			window[x]--
			if window[x] == 0 {
				delete(window, x)
			}
		}
		res += left
	}
	return res
}
