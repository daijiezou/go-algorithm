package main

import (
	"fmt"
	"math"
	"slices"
)

func main() {

	fs := create()
	for i := 0; i < len(fs); i++ {
		fs[i]()
	}
}

func create() (fs [2]func()) {
	for i := 0; i < 2; i++ {
		fs[i] = func() {
			fmt.Println(i)
		}
	}
	return
}

func lengthOfLIS(nums []int) int {
	length := len(nums)
	dp := make([]int, length)
	for i := 0; i < length; i++ {
		// 初始化最小的长度为1
		dp[i] = 1
	}
	// dp数组定义为以dp[i],为结尾的最长递增子数组的值
	for i := 1; i < length; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	maxRes := 0
	for i := 0; i < length; i++ {
		if dp[i] > maxRes {
			maxRes = dp[i]
		}
	}
	return maxRes
}

func merge(nums1 []int, nums2 []int) []int {
	res := make([]int, 0, len(nums1)+len(nums2))
	for len(nums1) > 0 && len(nums2) > 0 {
		if nums1[0] <= nums2[0] {
			res = append(res, nums1[0])
			nums1 = nums1[1:]
		} else {
			res = append(res, nums2[0])
			nums2 = nums2[1:]
		}
	}
	for len(nums1) > 0 {
		res = append(res, nums1[0])
		nums1 = nums1[1:]
	}
	for len(nums2) > 0 {
		res = append(res, nums2[0])
		nums2 = nums2[1:]
	}
	return res
}

func NextGreater(nums []int) []int {
	length := len(nums)
	smallStack := make([]int, 0)
	res := make([]int, 0)
	for i := length - 1; i > 0; i-- {
		for len(smallStack) > 0 && nums[i] >= smallStack[len(smallStack)-1] {
			smallStack = smallStack[:len(smallStack)-1]
		}
		if len(smallStack) == 0 {
			res = append(res, 0)
		} else {
			res = append(res, smallStack[len(smallStack)-1])
		}
		smallStack = append(smallStack, nums[i])
	}
	return res
}

func minimumFuelCost(roads [][]int, seats int) int64 {
	for _, road := range roads {
		for _, v := range road {
			if v == 0 {

			}
		}
	}
	return 0
}

func FindNeighbour(roads [][]int, root int, preRoot, index int, resultMap map[int]int) {
	var neighbourList []int
	for _, road := range roads {
		for i := 0; i < len(road); i++ {
			if road[i] == root && road[(i+1)%2] != preRoot {
				neighbourList = append(neighbourList, road[(i+1)%2])
			}
		}
	}

	if len(neighbourList) == 0 {
		return
	}
	fmt.Println(neighbourList)
	fmt.Println(index)
	resultMap[index] = len(neighbourList)
	index++
	for i := 0; i < len(neighbourList); i++ {
		FindNeighbour(roads, neighbourList[i], root, index, resultMap)
	}
}

// 走楼梯
func Zoulouti(num int) int {
	// 还剩1阶只有1种走法
	if num == 1 {
		return 1
	}
	// 还剩2阶只有2种走法
	if num == 2 {
		return 2
	}
	return Zoulouti(num-1) + Zoulouti(num-2)
}

func minWindow(s string, t string) string {
	sByte := []byte(s)
	tByte := []byte(t)
	need := make(map[byte]int)
	window := make(map[byte]int)
	for _, b := range tByte {
		need[b]++
	}
	left, right := 0, 0 // 滑动窗口
	var start int
	var valid int
	length := math.MaxInt32
	for right < len(sByte) {
		c := sByte[right]
		right++
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
			fmt.Printf("window: [%d, %d)\n", left, right)
		}
		for valid == len(need) {
			if right-left < length {
				length = right - left
				start = left
			}
			d := sByte[left]
			left++
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}

	if length == math.MaxInt32 { // 如果最小子串长度没有更新，则返回空格
		return ""
	}
	return string(sByte[start : start+length+1])
}

func checkInclusion(s string, t string) bool {
	sByte := []byte(s)
	tByte := []byte(t)
	need := make(map[byte]int)
	window := make(map[byte]int)
	for _, b := range tByte {
		need[b]++
	}
	left, right := 0, 0
	valid := 0
	for right < len(sByte) {
		c := sByte[right]
		right++
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}
		for right-left >= len(tByte) {
			if valid == len(need) {
				return true
			}
			c := sByte[left]
			if _, ok := need[c]; ok {
				if window[c] == need[c] {
					valid--
				}
				window[c]--
			}
			left++
		}
	}

	return false
}

func findAnagrams(s string, t string) []int {
	sByte := []byte(s)
	tByte := []byte(t)
	need := make(map[byte]int)
	window := make(map[byte]int)
	for _, b := range tByte {
		need[b]++
	}
	var left, right int
	var valid int
	result := make([]int, 0)
	for right < len(sByte) {
		c := sByte[right]
		right++
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}

		}
		for right-left >= len(tByte) {
			if valid == len(need) {
				result = append(result, left)
			}
			d := sByte[left]
			left++
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	return result
}

func lengthOfLongestSubstring(s string) int {
	window := make(map[byte]int)
	var left, right int
	res := 0
	for right < len(s) {
		c := s[right]
		right++
		window[c]++
		for window[c] > 1 {
			d := s[left]
			left++
			window[d]--
		}
		res = max(res, right-left)
	}
	return res
}

// 给你一个字符串数组 words ，找出并返回 length(words[i]) * length(words[j]) 的最大值，
// 并且这两个单词不含有公共字母。如果不存在这样的两个单词，返回 0 。
func maxProduct(words []string) int {
	var result int
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if !checkIsDup(words[i], words[j]) {
				ji := len(words[i]) * len(words[j])
				if ji > result {
					result = ji
				}
			}
		}
	}
	return result
}

func checkIsDup(a, b string) bool {
	for _, i := range a {
		for _, k := range b {
			if i == k {
				return true
			}
		}
	}
	return false
}

func removeDuplicates(nums []int) int {
	var slow, fast int
	for fast < len(nums) {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
		fast++
	}
	return slow + 1
}

func removeElement(nums []int, val int) int {
	var slow, fast int
	for fast < len(nums) {
		if nums[fast] != val {
			slow++
			nums[slow] = nums[fast]
		}
		fast++
	}
	return slow + 1
}

func uniqueLetterString(s string) int {
	byteSlice := []byte(s)
	if len(byteSlice) == 1 {
		return 1
	}
	length := len(byteSlice)
	var result int
	countUniqueCharsMap := make(map[string]int)
	// 从长度为2的子串开始计算，一直到length-1
	for i := 2; i < length; i++ {
		for j := 0; j <= length-i; j++ {
			if count, ok := countUniqueCharsMap[string(byteSlice[j:j+i])]; ok {
				result += count
			} else {
				result += countUniqueChars(byteSlice[j : j+i])
			}
		}
	}
	return result + countUniqueChars(byteSlice) + length
}

func countUniqueChars(req []byte) int {
	var result int
	tempMap := make(map[byte]int)
	for _, k := range req {
		if _, ok := tempMap[k]; ok {
			tempMap[k]++
		} else {
			tempMap[k] = 1
		}
	}
	for _, v := range tempMap {
		if v == 1 {
			result++
		}
	}
	return result
}

func uniqueLetterString2(s string) int {
	byteSlice := []byte(s)
	if len(byteSlice) == 1 {
		return 1
	}
	numIndex := make(map[byte][]int)
	for index, value := range byteSlice {
		numIndex[value] = append(numIndex[value], index)
	}
	var result int
	for _, v := range numIndex {
		vIndex := append(append([]int{-1}, v...), len(s))
		for i := 1; i < len(vIndex)-1; i++ {
			result += (vIndex[i] - vIndex[i-1]) * (vIndex[i+1] - vIndex[i])
		}
	}
	return result
}

var myresult int

func sumSubarrayMins(arr []int) int {
	defer func() {
		myresult = 0
	}()
	BuildsumSubarrayMins(arr)
	return myresult % (1e9 + 7)
}

func BuildsumSubarrayMins(arr []int) {
	length := len(arr)
	if length == 1 {
		myresult += arr[0]
		return
	}
	if length == 0 {
		return
	}
	var minIndex int
	minNum := math.MaxInt32
	for i := 0; i < length; i++ {
		if arr[i] < minNum {
			minNum = arr[i]
			minIndex = i
		}
	}
	myresult += (minIndex + 1) * (length - minIndex) * minNum
	BuildsumSubarrayMins(arr[0:minIndex])
	BuildsumSubarrayMins(arr[minIndex+1:])
}

func sumSubarrayMins2(arr []int) int {
	var ans int
	arr = append(arr, -1)
	st := []int{-1} // 哨兵
	for r, x := range arr {
		for len(st) > 1 && arr[st[len(st)-1]] >= x {
			i := st[len(st)-1]
			st = st[:len(st)-1]
			ans += arr[i] * (i - st[len(st)-1]) * (r - i) // 累加贡献
		}
		st = append(st, r)
	}
	return ans % (1e9 + 7)
}

func closeStrings(word1 string, word2 string) bool {
	byte1 := []byte(word1)
	byte2 := []byte(word2)
	map1 := make(map[byte]int)
	for i := 0; i < len(byte1); i++ {
		if _, ok := map1[byte1[i]]; ok {
			map1[byte1[i]]++
		} else {
			map1[byte1[i]] = 1
		}
	}
	map2 := make(map[byte]int)
	for i := 0; i < len(byte2); i++ {
		if _, ok := map2[byte2[i]]; ok {
			map2[byte2[i]]++
		} else {
			map2[byte2[i]] = 1
		}
	}
	// 判断是否有没有的字符
	for value := range map1 {
		if _, ok := map2[value]; !ok {
			return false
		}
	}
	var list1 []int
	for _, count := range map1 {
		list1 = append(list1, count)
	}
	var list2 []int
	for _, count := range map2 {
		list2 = append(list2, count)
	}
	slices.Sort(list1)
	slices.Sort(list2)
	// 将两个map的字符数量排序
	return slices.Equal(list1, list2)
}

func firstCompleteIndex(arr []int, mat [][]int) int {
	rowLength, colLength := len(mat), len(mat[0])
	mp := make(map[int][2]int)
	for row := 0; row < rowLength; row++ {
		for col := 0; col < colLength; col++ {
			num := mat[row][col]
			mp[num] = [2]int{row, col}
		}
	}
	rowCnt, colCnt := make([]int, rowLength), make([]int, colLength)
	for i := 0; i < rowLength; i++ {
		rowCnt[i] = 0
	}
	for j := 0; j < colLength; j++ {
		colCnt[j] = 0
	}

	for i := 0; i < len(arr); i++ {
		v := mp[arr[i]]
		//这一行被涂色的元素+1
		rowCnt[v[0]]++
		// 如果这一行被涂色的元素等于列长度，说明这一行被涂满，返回数组下标
		if rowCnt[v[0]] == colLength {
			return i
		}
		colCnt[v[1]]++
		if colCnt[v[1]] == rowLength {
			return i
		}
	}
	return -1

}

func lengthOfLongestSubstring1(s string) int {
	sByte := []byte(s)
	window := make(map[byte]int)
	var left, right = 0, 0
	length := len(sByte)
	resLength := 0
	for right < length {
		c := sByte[right]
		right++
		window[c]++
		for window[c] > 1 {
			d := s[left]
			window[d]--
			left++
		}
		length := right - left
		if length > resLength {
			resLength = length
		}
	}
	return resLength
}

func minCostClimbingStairs(cost []int) int {
	length := len(cost)
	dp := make([]int, length+1)
	for i := 2; i <= length; i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}
	return dp[length]
}

func climbStairs(n int) int {
	if n == 2 {
		return 2
	}
	if n == 1 {
		return 1
	}
	return climbStairs(n-1) + climbStairs(n-2)
}

// 计算质数
// https://leetcode.cn/problems/count-primes/description/
func countPrimes(n int) int {
	count := 0
	for i := 2; i < n; i++ {
		if isPrime(i) {
			count++
		}
	}
	return count
}

func isPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func countPrimes2(n int) int {
	isPrime := make([]bool, n)
	// 将数组都初始化为 true
	for i := range isPrime {
		isPrime[i] = true
	}
	for i := 2; i < n; i++ {
		if isPrime[i] {
			// i 的倍数不可能是素数了
			for j := i * i; j < n; j += i {
				isPrime[j] = false
			}
		}
	}
	count := 0
	for i := 2; i < n; i++ {
		if isPrime[i] {
			count++
		}
	}
	return count
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{Val: -1}
	p := dummy
	addOne := 0
	for l1 != nil || l2 != nil || addOne > 0 {
		val := addOne
		if l1 != nil {
			val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val += l2.Val
			l2 = l2.Next
		}
		if val > 9 {
			val = val - 10
			addOne = 1
		} else {
			addOne = 0
		}
		node := &ListNode{Val: val}
		p.Next = node
		p = p.Next

	}
	return dummy.Next
}

func nSumTarget(nums []int, n int, start int, target int64) [][]int {
	res := make([][]int, 0)
	if len(nums) < n {
		return res
	}
	length := len(nums)
	if n == 2 {
		left := start
		right := length - 1
		for left < right {
			leftVal := nums[left]
			rightVal := nums[right]
			sum := nums[right] + nums[left]
			if int64(sum) > target {
				right--
			} else if int64(sum) < target {
				left++
			} else {
				res = append(res, []int{left, right})
				// 跳过所有重复的元素
				for left < right && nums[left] == leftVal {
					left++
				}
				for left < right && nums[right] == rightVal {
					right--
				}
			}
		}

	} else {
		for i := start; i < len(nums); i++ {
			subs := nSumTarget(nums, n-1, i+1, target-int64(nums[i]))
			for _, sub := range subs {
				sub = append(sub, nums[i])
				res = append(res, sub)
			}
			for i < length-1 && nums[i] == nums[i+1] {
				i++
			}
		}
	}
	return res

}
