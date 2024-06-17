package _6greedy_algorithm

import (
	"sort"
)

/*
在一条环路上有 n 个加油站，其中第 i 个加油站有汽油 gas[i] 升。

你有一辆油箱容量无限的的汽车，从第 i 个加油站开往第 i+1 个加油站需要消耗汽油 cost[i] 升。你从其中的一个加油站出发，开始时油箱为空。

给定两个整数数组 gas 和 cost ，如果你可以按顺序绕环路行驶一周，则返回出发时加油站的编号，否则返回 -1 。如果存在解，则 保证 它是 唯一 的。
*/

/*
1、从区间集合 intvs 中选择一个区间 x，这个 x 是在当前所有区间中结束最早的（end 最小）。
2、把所有与 x 区间相交的区间从区间集合 intvs 中删除。
3、重复步骤 1 和 2，直到 intvs 为空为止。之前选出的那些 x 就是最大不相交子集
*/

// https://leetcode.cn/problems/non-overlapping-intervals/submissions/539972480/
func eraseOverlapIntervals(intvs [][]int) int {
	n := len(intvs)
	if n < 1 {
		return n
	}
	sort.Slice(intvs, func(i, j int) bool {
		// 按照结束时间来排序
		return intvs[i][1] < intvs[j][1]
	})
	total := 1
	endTime := intvs[0][1]
	for i := 1; i < n; i++ {
		// 判断startTime是否大于等于前一个结束时间的endTime
		if intvs[i][0] >= endTime {
			endTime = intvs[i][1]
			total++
		}
	}
	return n - total
}

// https://leetcode.cn/problems/minimum-number-of-arrows-to-burst-balloons/submissions/539972694/
func findMinArrowShots(intvs [][]int) int {
	n := len(intvs)
	if n < 1 {
		return n
	}
	sort.Slice(intvs, func(i, j int) bool {
		// 按照结束时间来排序
		return intvs[i][1] < intvs[j][1]
	})
	total := 1
	endTime := intvs[0][1]
	for i := 1; i < n; i++ {
		// 判断startTime是否大于前一个结束时间的endTime
		// note 与上面问题的差别就在于这里是 > 而不是 >=
		if intvs[i][0] > endTime {
			endTime = intvs[i][1]
			total++
		}
	}
	return total
}

// https://labuladong.online/algo/frequency-interview/cut-video/#%E6%80%9D%E8%B7%AF%E5%88%86%E6%9E%90
func videoStitching(clips [][]int, T int) int {
	if T == 0 {
		return 0
	}
	sort.Slice(clips, func(i, j int) bool {
		// 如果起点相同则按照结束时间的倒序排
		//if clips[i][0] == clips[j][0] {
		//	return clips[i][1] > clips[j][1]
		//}
		return clips[i][0] < clips[j][0]
	})
	// 记录选择的短视频个数
	videoCnt := 0
	curEnd, nextEnd := 0, 0
	for i := 0; i < len(clips) && clips[i][0] <= curEnd; {
		// 在第 res 个视频的区间内贪心选择下一个视频
		// 找到结束时间最长的那个视频
		for ; i < len(clips) && clips[i][0] <= curEnd; i++ {
			nextEnd = max(nextEnd, clips[i][1])
		}
		// 找到下一个视频，更新 curEnd
		videoCnt++
		curEnd = nextEnd
		if curEnd >= T {
			// 已经可以拼出区间 [0, T]
			return videoCnt
		}
	}
	return -1
}
