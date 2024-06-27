package _7game

import "math"

// https://leetcode.cn/problems/super-egg-drop/
func superEggDrop(k int, n int) int {
	memo := make([][]int, k)
	for i := 0; i < k+1; i++ {
		memo[i] = make([]int, n+1)
		for j := 0; j < n+1; j++ {
			memo[i][j] = -1
		}
	}
	return superEggDropDP(k, n, memo)
}

// 在有k个鸡蛋，n层楼的情况下，最少需要试几次
func superEggDropDP(k int, n int, memo [][]int) int {
	if n == 0 {
		return 0
	}
	// 只有一个鸡蛋需要试n次
	if k == 1 {
		return n
	}
	if memo[k][n] != -1 {
		return memo[k][n]
	}

	/*	// 进行状态转移
		res := math.MaxInt32
		for i := 1; i <= n; i++ {
			// 在第 i 层扔，有两种可能：碎或不碎
			// 根据这两种可能得到两个子问题：
			// 1. 在前 i-1 层楼中扔鸡蛋，鸡蛋没碎，此时鸡蛋个数不会发生变化，问题转化成
			//    了 K 个鸡蛋、n-i 层楼，即 dp(k, n-i)。
			// 2. 在第 i 层楼中扔鸡蛋，鸡蛋碎了，此时应该将鸡蛋个数减一，问题转化成
			//    了 K-1 个鸡蛋、n-1 层楼，即 dp(k-1, N-1)。
			// 因为是求至少需要多少次，所以需要在两者中取较大值
			// 两者取最大值，再加上在第 i 层扔的一次，就是这个状态下的最少实验次数。
			res = min(
				res,
				max(superEggDropDP(k, n-i, memo), superEggDropDP(k-1, i-1, memo))+1,
			)
		}
	*/

	lo := 1
	high := n
	res := math.MaxInt
	for lo <= high {
		// superEggDropDP是一个随着n增加单调递增函数，n越大，需要重试的次数就越多
		mid := lo + (high-lo)/2
		//鸡蛋碎了，楼层-1，鸡蛋-1
		broken := superEggDropDP(k-1, mid-1, memo)
		//鸡蛋没碎，表示下面的楼层都不会碎了，楼层-i
		noBroken := superEggDropDP(k, n-mid, memo)
		if broken > noBroken {
			high = mid - 1
			res = min(res, broken+1)
		} else {
			lo = mid + 1
			res = min(res, noBroken+1)
		}
	}
	memo[k][n] = res
	return res
}
