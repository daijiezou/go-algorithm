package leetcode

import (
	"math"
	"slices"
	"strings"
)

func partition(s string) [][]string {
	n := len(s)
	res := make([][]string, 0)
	curent := make([]string, 0)
	var backtrack func(i int)
	backtrack = func(i int) {
		// 到达终点
		if i == n {
			temp := make([]string, len(curent))
			copy(temp, curent)
			res = append(res, temp)
			return
		}
		for j := i; j < n; j++ {
			if isPalindrome(s, i, j) {
				curent = append(curent, s[i:j+1])
				backtrack(j + 1)
				curent = curent[:len(curent)-1]
			}
		}
	}
	backtrack(0)
	return res
}

func isPalindrome(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// https://leetcode.cn/problems/palindrome-partitioning-ii/
func minCut(s string) int {
	n := len(s)
	palMemo := make([][]int, n)
	for i := 0; i < n; i++ {
		palMemo[i] = make([]int, n)
		for j := 0; j < n; j++ {
			// 表示没有计算过
			palMemo[i][j] = -1
		}
	}

	var isPal func(left, right int) bool
	isPal = func(left, right int) bool {
		if left >= right {
			palMemo[left][right] = 1
		}
		p := &palMemo[left][right]
		if *p != -1 {
			return *p == 1
		}
		res := s[left] == s[right] && isPal(left+1, right-1)
		if res {
			*p = 1
		} else {
			*p = 0
		}
		return res
	}
	memo := make([]int, n)
	for i := 0; i < n; i++ {
		// 表示没有计算过
		memo[i] = -1
	}
	// 把 s[:r+1] 切 i 刀，分成 i+1 个子串，每个子串改成回文串的最小总修改次数
	var dfs func(i int) int
	dfs = func(i int) int {
		if isPal(0, i) {
			return 0
		}
		if memo[i] != -1 {
			return memo[i]
		}
		res := math.MaxInt
		for left := 1; left <= i; left++ {
			if isPal(left, i) {
				res = min(res, dfs(left-1)+1)
			}
		}
		memo[i] = res
		return res
	}
	return dfs(n - 1)
}

func palindromePartition(s string, k int) int {
	n := len(s)
	memoChange := make([][]int, n)
	for i := 0; i < n; i++ {
		memoChange[i] = make([]int, n)
		for j := 0; j < n; j++ {
			memoChange[i][j] = -1 // 表示没有计算过
		}
	}
	var minChange func(i, j int) int //表示s[i:j+1]修改为回文串的最小修改次数
	minChange = func(i, j int) int {
		if i >= j {
			return 0
		}
		if memoChange[i][j] != -1 {
			return memoChange[i][j]
		}
		res := minChange(i+1, j-1)
		if s[i] != s[j] {
			res++
		}
		memoChange[i][j] = res
		return res
	}
	memoDfs := make([][]int, k)
	for i := range memoDfs {
		memoDfs[i] = make([]int, n)
		for j := range memoDfs[i] {
			memoDfs[i][j] = -1 // -1 表示没有计算过
		}
	}
	// i表示还需要切i刀
	// j表示剩余字符串的右端点
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i == 0 {
			return minChange(0, j)
		}
		if memoDfs[i][j] != -1 {
			return memoDfs[i][j]
		}
		res := math.MaxInt
		// 由于不能有空串，所以右端点的初始位置必须>=i
		for l := i; l <= j; l++ {
			res = min(res, dfs(i-1, l-1)+minChange(l, j))
		}
		memoDfs[i][j] = res
		return res
	}
	return dfs(k-1, n-1)
}

// https://leetcode.cn/problems/most-beautiful-item-for-each-query/?envType=daily-question&envId=2025-03-09
func maximumBeauty(items [][]int, queries []int) []int {
	ans := make([]int, len(queries))
	slices.SortFunc(items, func(a, b []int) int { return a[0] - b[0] })
	for i := 1; i < len(items); i++ {
		items[i][1] = max(items[i][1], items[i-1][1])
	}
	for i, price := range queries {
		left, right := 0, len(items)-1
		for left <= right {
			mid := left + (right-left)/2
			if items[mid][0] >= price+1 {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
		if left == 0 || items[left-1][0] > price {
			ans[i] = 0
			continue
		}
		ans[i] = items[left-1][1]
	}
	return ans
}

// https://leetcode.cn/problems/sum-of-beauty-in-the-array/description/?envType=daily-question&envId=2025-03-11
func sumOfBeauties(nums []int) int {
	n := len(nums)
	sufMin := make([]int, n)
	sufMin[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		sufMin[i] = min(sufMin[i+1], nums[i])
	}

	preMax := nums[0]
	res := 0
	for i := 1; i < n-1; i++ {
		x := nums[i]
		if preMax < x && x < sufMin[i+1] {
			res += 2
		} else if x > nums[i-1] && x < nums[i+1] {
			res += 1
		}
		preMax = max(preMax, x)
	}
	return res
}

/*
根据这个思路，本题等价于如下两个问题：

每个元音字母至少出现一次，并且至少包含 k 个辅音字母的子串个数。记作 f k

	。

每个元音字母至少出现一次，并且至少包含 k+1 个辅音字母的子串个数。记作 fk+1

	。

作者：灵茶山艾府
链接：https://leetcode.cn/problems/count-of-substrings-containing-every-vowel-and-k-consonants-ii/solutions/2934309/liang-ci-hua-chuang-pythonjavacgo-by-end-2lpz/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func countOfSubstrings(word string, k int) int {
	return f(word, k) - f(word, k+1)
}

func f(word string, k int) int {
	yuan := make(map[byte]struct{})
	yuan[byte('a')] = struct{}{}
	yuan[byte('e')] = struct{}{}
	yuan[byte('i')] = struct{}{}
	yuan[byte('o')] = struct{}{}
	yuan[byte('u')] = struct{}{}

	cnt1 := make(map[byte]int)
	cnt2 := 0
	left := 0
	res := 0
	for i := 0; i < len(word); i++ {
		x := word[i]
		if _, ok := yuan[x]; ok {
			cnt1[x]++
		} else {
			cnt2++
		}
		for len(cnt1) == 5 && cnt2 >= k {
			leave := word[left]
			if _, ok := yuan[leave]; ok {
				cnt1[leave]--
				if cnt1[leave] == 0 {
					delete(cnt1, leave)
				}
			} else {
				cnt2--
			}
			left++
		}
		// 至少有，前面的都符合
		res += left
	}
	return res
}

func f2(word string, k int) (ans int64) {
	// 这里用哈希表实现，替换成数组会更快
	cnt1 := map[byte]int{} // 每种元音的个数
	cnt2 := 0              // 辅音个数
	left := 0
	for _, b := range word {
		if strings.ContainsRune("aeiou", b) {
			cnt1[byte(b)]++
		} else {
			cnt2++
		}
		for len(cnt1) == 5 && cnt2 >= k {
			out := word[left]
			if strings.ContainsRune("aeiou", rune(out)) {
				cnt1[out]--
				if cnt1[out] == 0 {
					delete(cnt1, out)
				}
			} else {
				cnt2--
			}
			left++
		}
		ans += int64(left)
	}
	return
}

func isBalanced(num string) bool {
	bytes := []byte(num)
	sum1 := 0
	sum2 := 0
	for i := 0; i < len(bytes); i++ {
		if i%2 == 0 {
			sum1 += int(bytes[i] - '0')
		} else {
			sum2 += int(bytes[i] - '0')
		}
	}
	return sum1 == sum2
}

func scoreOfString(s string) int {
	sum := 0
	for i := 0; i < len(s)-1; i++ {
		sum += Myabs(int(s[i]), int(s[i+1]))
	}
	return sum
}

func Myabs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	res := dp[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1], 0) + nums[i]
		res = max(res, dp[i])
	}
	return res
}

// https://leetcode.cn/problems/substring-with-largest-variance/?envType=daily-question&envId=2025-03-16
func largestVariance(s string) int {
	res := 0
	for a := 'a'; a <= 'z'; a++ {
		for b := 'a'; b <= 'z'; b++ {
			if a == b {
				continue
			}
			f0, f1 := 0, math.MinInt
			for _, x := range s {
				if x == a {
					f0 = max(f0, 0) + 1
					f1 = f1 + 1
				} else {
					f0 = max(f0, 0) - 1
					f1 = f0
				}
			}
			res = max(f1)
		}
	}
	return res
}

func minSwaps(s string) int {
	sByte := []byte(s)
	res := 0
	stack := 0
	j := len(s) - 1
	for _, x := range sByte {
		if x == '[' {
			stack++
		} else {
			if stack > 0 {
				stack--
			} else {
				// 和最右边的左括号交换
				for s[j] != '[' {
					j--
				}
				sByte[j] = ']'
				stack++
				res++
			}
		}
	}
	return res
}

func diagonalPrime(nums [][]int) int {
	res := 0
	n := len(nums)
	for i := 0; i < len(nums); i++ {
		if checkIsZhi(nums[i][i]) {
			res = max(nums[i][i], res)
		}
		if checkIsZhi(nums[i][n-i-1]) {
			res = max(nums[i][n-i-1], res)
		}
	}
	return res

}

func checkIsZhi(num int) bool {
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return num >= 2
}

func findMatrix(nums []int) [][]int {
	numsMpa := make(map[int]int)
	res := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		numsMpa[nums[i]]++
	}
	for len(numsMpa) > 0 {
		size := len(numsMpa)
		temp := make([]int, 0)
		for k, v := range numsMpa {
			temp = append(temp, k)
			numsMpa[k]--
			if v == 1 {
				delete(numsMpa, k)
			}
			size--
			if size == 0 {
				break
			}
		}
		res = append(res, temp)
	}
	return res
}

/*
每次操作可以将某个整数乘 2，因此要想使得所有整数的或值更大，我们应该尽量选择二进制表示中具有更高位为 1 的数字。
而这样的整数有很多个，我们需要一一遍历找到令答案最大的那个，
一旦将其乘 2 之后，接下来的 k−1 次操作都需要基于该整数进行。
*/
func maximumOr(nums []int, k int) int64 {
	n := len(nums)
	suf := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		suf[i] = nums[i] | suf[i+1]
	}
	pre := 0
	res := 0
	for i := 0; i < n; i++ {
		res = max(res, pre|nums[i]<<k|suf[i+1])
		pre |= nums[i]
	}
	return int64(res)
}

func canBeValid(s string, locked string) bool {
	stack := 0
	n := len(s)
	memo := make([][]int, n)
	for i := 0; i < len(s); i++ {
		memo[i] = make([]int, n)
		for j := 0; j < n; j++ {
			memo[i][j] = -1
		}
	}
	return f3(s, locked, stack, 0, memo)
}

func f3(s string, locked string, stack, start int, memo [][]int) (res bool) {
	if start == len(s) {
		return stack == 0
	}
	p := &memo[start][stack]
	if *p != -1 {
		return *p == 0
	}
	defer func() {
		if res {
			*p = 0
		} else {
			*p = 1
		}
	}()
	if locked[start] == '0' {
		if stack > 0 {
			return f3(s, locked, stack-1, start+1, memo) || f3(s, locked, stack+1, start+1, memo)
		} else {
			return f3(s, locked, stack+1, start+1, memo)
		}
	} else {
		if s[start] == '(' {
			return f3(s, locked, stack+1, start+1, memo)
		} else {
			if stack > 0 {
				return f3(s, locked, stack-1, start+1, memo)
			} else {
				*p = 1
				return false
			}
		}
	}
}

func countPrefixes(words []string, s string) int {
	res := 0
	s0 := []byte(s)
	sLen := len(s)
	for i := 0; i < len(words); i++ {
		if len(words[i]) > sLen {
			continue
		}
		length := min(len(words[i]), sLen)
		target := string(s0[:length])
		if strings.HasPrefix(words[i], target) {
			res++
		}
	}
	return res
}

func minimumSum(n int, k int) int {
	res := 0

	forbid := make(map[int]struct{})
	for i := 1; i < (k-1)/2+1; i++ {
		forbid[k-i] = struct{}{}
	}
	j := 0
	for i := 1; i < n+k; i++ {
		if _, ok := forbid[i]; !ok {
			res += i
			j++
		}
		if j == n {
			return res
		}
	}
	return 0
}

//func minimumCost(s string) int64 {
//
//}

func minimizedStringLength(s string) int {
	letterCnt := make(map[byte]struct{})
	for i := 0; i < len(s); i++ {
		letterCnt[s[i]] = struct{}{}
	}
	return len(letterCnt)
}
