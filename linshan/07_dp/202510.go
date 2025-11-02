package _7_dp

func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i <= target; i++ {
		sum := 0
		for j := 0; j < len(nums); j++ {

			if i-nums[j] >= 0 {
				sum += dp[i-nums[j]]
			}
			dp[i] = sum
		}
	}
	return dp[target]
}

// https://leetcode.cn/problems/count-ways-to-build-good-strings/description/
func countGoodStrings3(low int, high int, zero int, one int) int {
	mod := 1_000_000_007
	dp := make([]int, high+1)
	dp[0] = 1
	for i := 1; i <= high; i++ {
		sum := 0
		if i-zero >= 0 {
			sum += dp[i-zero] % mod
		}
		if i-one >= 0 {
			sum += dp[i-one] % mod
		}
		dp[i] = sum % mod
	}
	res := 0
	for i := low; i <= high; i++ {
		res += dp[i] % mod
	}
	return res % mod
}

// https://leetcode.cn/problems/count-number-of-texts/
func countTexts2(pressedKeys string) int {
	mod := 1_000_000_007
	n := len(pressedKeys)
	dp := make([]int, n+1)
	dp[0] = 1 // 基准情况：空字符串有一种解码方式（即什么都不做）

	for i := 1; i <= n; i++ {
		// 状态转移：dp[i] 表示解码前 i 个字符的方法数

		// 1. 单独解码第 i 个字符 (按键1次)
		dp[i] = dp[i-1]

		// 2. 结合前一个字符解码 (按键2次)
		if i >= 2 && pressedKeys[i-1] == pressedKeys[i-2] {
			dp[i] = (dp[i] + dp[i-2]) % mod
		} else {
			continue // 如果连续2个都不同，就不可能连续3个或4个
		}

		// 3. 结合前两个字符解码 (按键3次)
		if i >= 3 && pressedKeys[i-1] == pressedKeys[i-3] {
			dp[i] = (dp[i] + dp[i-3]) % mod
		} else {
			continue
		}

		// 4. 结合前三个字符解码 (按键4次, 仅限'7'和'9')
		if (pressedKeys[i-1] == '7' || pressedKeys[i-1] == '9') && i >= 4 && pressedKeys[i-1] == pressedKeys[i-4] {
			dp[i] = (dp[i] + dp[i-4]) % mod
		}
	}

	return dp[n]
}
