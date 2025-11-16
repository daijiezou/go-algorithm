package _4_datastruct

import "math"

func numPairsDivisibleBy60(time []int) int {
	numsCnt := make(map[int]int)
	res := 0
	for _, v := range time {
		mod := v % 60
		need := (60 - mod) % 60
		res += numsCnt[need]
		numsCnt[mod]++
	}
	return res
}

func similarPairs(words []string) int {
	cntsMap := make(map[[26]int8]int)
	res := 0
	for i := 0; i < len(words); i++ {
		key := [26]int8{}
		for _, v := range words[i] {
			key[v-'a'] = 1
		}
		res += cntsMap[key]
		cntsMap[key]++
	}
	return res
}

func maximumTripletValue(nums []int) int64 {
	mx := math.MinInt
	mn := math.MaxInt
	dmx := math.MinInt
	dmn := math.MaxInt
	res := 0
	for i := 0; i < len(nums); i++ {
		x := nums[i]

		if i >= 2 {
			res = max(res, x*dmn, x*dmx)
		}
		if i >= 1 {
			dmx = max(mx-x, dmx)
			dmn = min(mn-x, dmn)
		}

		mx = max(x, mx)
		mn = min(x, mn)
	}
	return int64(res)
}

// 两个无重叠数组的最大和

/*
给你一个整数数组 nums 和两个整数 firstLen 和 secondLen，
请你找出并返回两个无重叠 子数组 中元素的最大和，长度分别为 firstLen 和 secondLen 。
长度为 firstLen 的子数组可以出现在长为 secondLen 的子数组之前或之后，但二者必须是无重叠。
子数组是数组的一个 连续 部分。
*/
func maxSumTwoNoOverlap(nums []int, firstLen int, secondLen int) int {
	n := len(nums)
	preSum := make([]int, n+1)
	for i := 0; i < len(nums); i++ {
		preSum[i+1] = preSum[i] + nums[i]
	}
	sumA := 0
	res := 0
	// 在枚举第二段的过程中，检查第一段的最大值，顺便计算总的最大值
	for i := firstLen + secondLen; i <= n; i++ {
		sumA = max(sumA, preSum[i-secondLen]-preSum[i-secondLen-firstLen])
		res = max(res, sumA+preSum[i]-preSum[i-secondLen])
	}
	sumA = 0
	for i := firstLen + secondLen; i <= n; i++ {
		sumA = max(sumA, preSum[i-firstLen]-preSum[i-secondLen-firstLen])
		res = max(res, sumA+preSum[i]-preSum[i-firstLen])
	}
	return res
}

/*
在 X轴 上有一些奖品。给你一个整数数组 prizePositions ，
它按照 非递减 顺序排列，其中 prizePositions[i] 是第 i 件奖品的位置。数轴上一个位置可能会有多件奖品。再给你一个整数 k 。
你可以同时选择两个端点为整数的线段。每个线段的长度都必须是 k 。
你可以获得位置在任一线段上的所有奖品（包括线段的两个端点）。注意，两个线段可能会有相交。

比方说 k = 2 ，你可以选择线段 [1, 3] 和 [2, 4] ，你可以获得满足 1 <= prizePositions[i] <= 3 或者 2 <= prizePositions[i] <= 4 的所有奖品 i 。
请你返回在选择两个最优线段的前提下，可以获得的 最多 奖品数目。
*/
func maximizeWin(prizePositions []int, k int) int {
	if 2*k+1 >= len(prizePositions) {
		return len(prizePositions)
	}
	ans := 0
	mx := make([]int, len(prizePositions)+1)
	left := 0
	for i := 0; i < len(prizePositions); i++ {
		index := prizePositions[i]
		for index-prizePositions[left] > k {
			left++
		}
		ans = max(ans, mx[left]+i-left+1)
		mx[i+1] = max(mx[i], i-left+1)
	}
	return ans
}

/*
给你一个 下标从 0 开始 的整数数组 nums ，返回满足下述条件的 不同 四元组 (a, b, c, d) 的 数目 ：

nums[a] + nums[b] + nums[c] == nums[d] ，且
a < b < c < d

*/

func countQuadruplets(nums []int) int {
	ans := 0
	cntMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			ans += cntMap[nums[j]-nums[i]]
		}
		// 为下一轮做准备
		for z := 0; z < i; z++ {
			cntMap[nums[z]+nums[i]]++
		}
	}
	return ans
}
