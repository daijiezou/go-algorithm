package math

// https://leetcode.cn/problems/nim-game/submissions/542743035/
/*
你和你的朋友，两个人一起玩 Nim 游戏：
桌子上有一堆石头。
你们轮流进行自己的回合， 你作为先手 。
每一回合，轮到的人拿掉 1 - 3 块石头。
拿掉最后一块石头的人就是获胜者。
假设你们每一步都是最优解。请编写一个函数，来判断你是否可以在给定石头数量为 n 的情况下赢得游戏。如果可以赢，返回 true；否则，返回 false 。
*/
func canWinNim(n int) bool {
	return n%4 == 0
}

// https://leetcode.cn/problems/stone-game/
func stoneGame(piles []int) bool {
	length := len(piles)
	dp := make([][]int, 0)
	/*
		定义二维数组 dp，其行数和列数都等于数组的长度，dp[i][j] 表示当数组剩下的部分为下标 i 到下标 j 时，即在下标范围 [i,j] 中，
		当前玩家与另一个玩家的分数之差的最大值，注意当前玩家不一定是先手。
		只有当 i≤j 时，数组剩下的部分才有意义，因此当 i>j 时，dp[i][j]=0。
		当 i=j 时，只剩一个数字，当前玩家只能拿取这个数字，因此对于所有 0≤i<nums.length，都有 dp[i][i]=nums[i]。
	*/
	for i := 0; i < length; i++ {
		dp[i] = make([]int, length)
		dp[i][i] = piles[i]
	}
	for i := length - 1; i >= 0; i-- {
		for j := i + 1; j < length; j++ {
			dp[i][j] = max(piles[i]-dp[i+1][j], piles[j]-dp[i][j-1])
		}
	}
	return dp[0][length-1] >= 0
}
