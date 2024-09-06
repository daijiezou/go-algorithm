package _1_huadongchuangkou

import (
	"math"
	"sort"
	"strconv"
)

// https://leetcode.cn/problems/maximum-number-of-vowels-in-a-substring-of-given-length/
func maxVowels(s string, k int) int {
	yuanyin := map[byte]struct{}{
		'a': {},
		'e': {},
		'i': {},
		'o': {},
		'u': {},
	}

	n := len(s)
	end := k
	if k > n {
		end = n
	}
	ans := 0
	for i := 0; i < end; i++ {
		if _, ok := yuanyin[s[i]]; ok {
			ans++
		}
	}
	res := ans
	left := 0
	right := k
	for right < n {
		if res == k {
			break
		}
		if _, ok := yuanyin[s[right]]; ok {
			ans++
		}
		if _, ok := yuanyin[s[left]]; ok {
			ans--
		}
		res = max(res, ans)
		right++
		left++
	}
	return res
}
func maxVowels1(s string, k int) int {
	ans := 0
	res := 0
	for i, in := range s {
		if in == 'a' || in == 'e' || in == 'i' || in == 'o' || in == 'u' {
			ans++
		}
		if i < k-1 { // 窗口大小不足 k
			continue
		}
		res = max(res, ans)
		out := s[i-k+1]
		if out == 'a' || out == 'e' || out == 'i' || out == 'o' || out == 'u' {
			ans--
		}
	}
	return res
}

// https://leetcode.cn/problems/find-the-k-beauty-of-a-number/
func divisorSubstrings(num int, k int) int {
	res := 0
	numStr := strconv.Itoa(num)
	var temp []byte
	for i := 0; i < len(numStr); i++ {
		temp = append(temp, numStr[i])
		if i < k-1 { // 窗口大小不足 k
			continue
		}
		tempStr := string(temp)
		tempInt, _ := strconv.Atoi(tempStr)

		if tempInt != 0 && num%tempInt == 0 {
			res++
		}
		temp = temp[1:]
	}
	return res
}

func divisorSubstrings2(num int, k int) int {
	res := 0
	numStr := strconv.Itoa(num)
	for i := k; i < len(numStr); i++ {
		// 枚举所有长度为k的子串
		tempInt, _ := strconv.Atoi(numStr[i-k : i])
		if tempInt != 0 && num%tempInt == 0 {
			res++
		}
	}
	return res
}

// https://leetcode.cn/problems/minimum-difference-between-highest-and-lowest-of-k-scores/
func minimumDifference(nums []int, k int) int {
	if k == 1 {
		return 0
	}
	sort.Ints(nums)
	ans := math.MaxInt32
	for i := k; i <= len(nums); i++ {
		ans = min(ans, nums[i-1]-nums[i-k])
	}
	return ans
}

func findMaxAverage(nums []int, k int) float64 {

	ans := float64(0)
	curTotal := 0
	for i := 0; i < k; i++ {
		curTotal += nums[i]
	}
	ans = float64(curTotal) / float64(k)
	for i := k; i < len(nums); i++ {
		curTotal += nums[i]
		curTotal -= nums[i-k]
		ans = math.Max(ans, float64(curTotal)/float64(k))
	}
	return ans
}

// 不用每次都计算平均数，算出最大的sum，最后再返回平均数即可
func findMaxAverage2(nums []int, k int) float64 {
	maxTotal := 0
	curTotal := 0
	for i := 0; i < k; i++ {
		curTotal += nums[i]
	}
	maxTotal = curTotal
	for i := k; i < len(nums); i++ {
		curTotal += nums[i]
		curTotal -= nums[i-k]
		maxTotal = max(maxTotal, curTotal)
	}
	return float64(maxTotal) / float64(k)
}

// https://leetcode.cn/problems/number-of-sub-arrays-of-size-k-and-average-greater-than-or-equal-to-threshold/
func numOfSubarrays(arr []int, k int, threshold int) int {
	total := k * threshold
	windowSum := 0
	res := 0
	for i := 0; i < len(arr); i++ {
		windowSum += arr[i]
		if i < k-1 {
			continue
		}
		if windowSum >= total {
			res++
		}
		windowSum -= arr[i-k+1]
	}
	return res
}

// https://leetcode.cn/problems/k-radius-subarray-averages/
func getAverages(nums []int, k int) []int {
	n := len(nums)
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = -1
	}
	length := k*2 + 1
	if length > n {
		return res
	}
	windowSum := 0
	for i := 0; i < n; i++ {
		windowSum += nums[i]
		if i < length-1 {
			continue
		}
		avg := windowSum / length
		res[i-k] = avg
		windowSum -= nums[i-length+1]
	}
	return res
}

func minimumRecolors(blocks string, k int) int {
	res := len(blocks)
	n := len(blocks)
	cur := 0
	for i := 0; i < n; i++ {
		if blocks[i] == 'W' {
			cur++
		}
		if i < k-1 {
			continue
		}
		res = min(res, cur)
		if blocks[i-k+1] == 'W' {
			cur--
		}
	}
	return res
}

// https://leetcode.cn/problems/defuse-the-bomb/
/*

你有一个炸弹需要拆除，时间紧迫！你的情报员会给你一个长度为 n 的 循环 数组 code 以及一个密钥 k 。
为了获得正确的密码，你需要替换掉每一个数字。所有数字会 同时 被替换。
如果 k > 0 ，将第 i 个数字用 接下来 k 个数字之和替换。
如果 k < 0 ，将第 i 个数字用 之前 k 个数字之和替换。
如果 k == 0 ，将第 i 个数字用 0 替换。
*/
func decrypt(code []int, k int) []int {
	n := len(code)
	res := make([]int, n)
	if k == 0 {

		return res
	}
	code = append(code, code...)

	l, r := 1, k+1
	if k < 0 {
		l, r = n+k, n
	}
	// 算出第一个数的解码后的值
	sum := 0
	for _, v := range code[l:r] {
		sum += v
	}
	for i := 0; i < n; i++ {
		res[i] = sum
		sum += code[r]
		sum -= code[l]
		r++
		l++
	}
	return res
}

// https://leetcode.cn/problems/grumpy-bookstore-owner/
func maxSatisfied(customers []int, grumpy []int, minutes int) int {
	maxCrease := 0
	sum := 0
	left, right := 0, 0
	crease := 0
	for ; right < len(customers); right++ {
		if grumpy[right] == 0 {
			sum += customers[right]
		}
		if grumpy[right] == 1 {
			crease += customers[right]
		}
		if right < minutes-1 { // 窗口长度不足 minutes
			continue
		}
		maxCrease = max(maxCrease, crease)
		if grumpy[left] == 1 {
			crease -= customers[left]
		}
		left++
	}
	return maxCrease + sum
}

func maxSum(nums []int, m int, k int) int64 {
	left := 0
	var sum int
	count := map[int]int{}
	res := 0
	for i := 0; i < len(nums); i++ {
		count[nums[i]]++
		sum += nums[i]
		if i < k-1 {
			continue
		}
		if len(count) >= m {
			res = max(res, sum)
		}
		sum -= nums[left]
		count[nums[left]]--
		if count[nums[left]] == 0 {
			delete(count, nums[left])
		}
		left++
	}
	return int64(res)
}

func maximumSubarraySum(nums []int, k int) int64 {
	left := 0
	var sum int
	count := map[int]int{}
	res := 0
	for i := 0; i < len(nums); i++ {
		count[nums[i]]++
		sum += nums[i]
		if i < k-1 {
			continue
		}
		if len(count) == k {
			res = max(res, sum)
		}
		sum -= nums[left]
		count[nums[left]]--
		if count[nums[left]] == 0 {
			delete(count, nums[left])
		}
		left++
	}
	return int64(res)
}

func maxScore(cardPoints []int, k int) int {
	count := len(cardPoints) - k
	leftSum := 0
	sum := 0
	minLeftSum := math.MaxInt64
	left := 0
	for i := 0; i < len(cardPoints); i++ {
		sum += cardPoints[i]
		leftSum += cardPoints[i]
		if i < count-1 {
			continue
		}
		minLeftSum = min(minLeftSum, leftSum)
		leftSum -= cardPoints[left]
		left++
	}
	if count == 0 {
		minLeftSum = 0
	}
	return sum - minLeftSum
}

func maxFreq(s string, maxLetters int, minSize int, maxSize int) int {
	left := 0
	count := map[byte]int{}
	res := 0
	ansCount := map[string]int{}
	for i := 0; i < len(s); i++ {
		count[s[i]]++
		if i < minSize-1 {
			continue
		}
		for i-left+1 > maxSize {
			count[s[left]]--
			if count[s[left]] == 0 {
				delete(count, s[left])
			}
			left++
		}
		for len(count) > maxLetters && i-left+1 > minSize {
			count[s[left]]--
			if count[s[left]] == 0 {
				delete(count, s[left])
			}
			left++
		}
		if len(count) <= maxLetters {
			ansCount[s[left:i+1]]++
			for i-left+1 > minSize {
				count[s[left]]--
				if count[s[left]] == 0 {
					delete(count, s[left])
				}
				left++
				ansCount[s[left:i+1]]++
			}
		}
	}
	for _, v := range ansCount {
		if v > res {
			res = v
		}
	}
	return res
}

/*
优化
假设字符串 T 在给定的字符串 S 中出现的次数为 k，那么 T 的任意一个子串出现的次数至少也为 k，即 T 的任意一个子串在 S 中出现的次数不会少于 T 本身。这样我们就可以断定，在所有满足条件且出现次数最多的的字符串中，一定有一个的长度恰好为 minSize。
*/

func maxFreq2(s string, maxLetters int, minSize int, maxSize int) int {
	left := 0
	count := map[byte]int{}
	res := 0
	ansCount := map[string]int{}
	for i := 0; i < len(s); i++ {
		count[s[i]]++
		if i < minSize-1 {
			continue
		}
		if len(count) <= maxLetters {
			ansCount[s[left:i+1]]++
		}
		count[s[left]]--
		if count[s[left]] == 0 {
			delete(count, s[left])
		}
		left++
	}
	for _, v := range ansCount {
		if v > res {
			res = v
		}
	}
	return res
}
