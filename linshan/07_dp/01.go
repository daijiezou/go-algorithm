package _7_dp

func climbStairs(n int) int {
	memo := make([]int, n+1)
	return climbStairsDp(memo, n)
}

func climbStairsDp(memo []int, n int) int {
	if n == 2 {
		return 2
	}
	if n == 1 {
		return 1
	}
	if memo[n] != 0 {
		return memo[n]
	}
	memo[n] = climbStairsDp(memo, n-1) + climbStairsDp(memo, n-2)
	return memo[n]
}

func climbStairs2(n int) int {
	dp := make([]int, n+1)
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// https://leetcode.cn/problems/count-ways-to-build-good-strings/
func countGoodStrings(low int, high int, zero int, one int) int {
	memo := make([]int, high+1)
	for i := 0; i <= high; i++ {
		memo[i] = -1
	}
	mod1 := 1000000007
	var dfs func(i int) int
	dfs = func(i int) int {
		if i < 0 {
			return 0
		}
		if i == 0 {
			return 1
		}
		if memo[i] != -1 {
			return memo[i]
		}
		memo[i] = (dfs(i-one) + dfs(i-zero)) % mod1
		return memo[i]
	}
	res := 0
	for i := low; i <= high; i++ {
		res = (res + dfs(i)) % mod1
	}
	return res
}

func countGoodStrings2(low int, high int, zero int, one int) int {
	memo := make([]int, high+1)
	for i := 0; i <= high; i++ {
		memo[i] = -1
	}
	mod1 := 1000000007
	dp := make([]int, high+1)
	res := 0
	dp[0] = 1
	for i := 1; i <= high; i++ {
		if i >= one {
			dp[i] = dp[i-one] % mod1
		}
		if i >= zero {
			dp[i] = dp[i-one] + dp[i-zero]%mod1
		}
		if i >= low {
			res += dp[i] % mod1
		}
	}

	return res
}

const mx = 100_001
const mod = 1_000_000_007

var dp3 = [mx]int{1, 1, 2, 4}
var dp4 = dp3

func init() {
	for i := 4; i < mx; i++ {
		dp3[i] = (dp3[i-1] + dp3[i-2] + dp3[i-3]) % mod
		dp4[i] = (dp4[i-1] + dp4[i-2] + dp4[i-3] + dp4[i-4]) % mod
	}
}

// https://leetcode.cn/problems/count-number-of-texts/
func countTexts(pressedKeys string) int {
	ans, cnt := 1, 0
	for i, c := range pressedKeys {
		cnt++
		if i == len(pressedKeys)-1 || pressedKeys[i+1] != byte(c) { // 找到一个完整的组
			if c != '7' && c != '9' {
				ans = ans * dp3[cnt] % mod
			} else {
				ans = ans * dp4[cnt] % mod
			}
			cnt = 0
		}
	}
	return ans
}
