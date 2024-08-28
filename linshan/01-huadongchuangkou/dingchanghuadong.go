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
