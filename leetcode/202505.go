package leetcode

import "math"

func pushDominoes(dominoes string) string {
	bytes := []byte("L" + dominoes + "R")
	pre := 0
	for i := 0; i < len(bytes); i++ {
		if bytes[i] == '.' {
			continue
		}
		if bytes[i] == bytes[pre] {
			fill(bytes[pre:i], bytes[i])
		} else if bytes[i] == 'L' {
			cnt := i - pre - 1
			if cnt%2 == 0 {
				fill(bytes[pre+1:pre+cnt/2+1], 'R') // 前一半变 R
				fill(bytes[pre+cnt/2+1:i], 'L')     // 后一半变 L
			} else {
				fill(bytes[pre+1:pre+cnt/2+1], 'R') // 前一半变 R
				fill(bytes[pre+cnt/2+2:i], 'L')     // 后一半变 L
			}
		}
		pre = i
	}
	return string(bytes[1 : len(bytes)-1])
}

func fill(bytes []byte, s byte) {
	for i := 0; i < len(bytes); i++ {
		bytes[i] = s
	}
}

// https://leetcode.cn/problems/minimum-domino-rotations-for-equal-row/solutions/3042326/du-bian-cheng-tops0-huo-zhe-bottoms0pyth-zvnj/?envType=daily-question&envId=2025-05-03
func minDominoRotations(tops []int, bottoms []int) int {
	n := len(tops)
	topCnt := make(map[int]int)
	botCnt := make(map[int]int)
	dup := make(map[int]int)
	for i := 0; i < n; i++ {
		topCnt[tops[i]]++
		botCnt[bottoms[i]]++
		if tops[i] == bottoms[i] {
			dup[tops[i]]++
		}
	}
	res := -1
	for i := 1; i <= 6; i++ {
		if topCnt[i]+botCnt[i]-dup[i] >= n {
			res = min(n-topCnt[i], n-botCnt[i])
		}
	}
	return res
}

func minDominoRotations1(tops []int, bottoms []int) int {
	var minCnts func(target int) int
	n := len(tops)
	minCnts = func(target int) int {
		topCnt := 0
		botCnt := 0
		for i := 0; i < n; i++ {
			if bottoms[i] != target && tops[i] != target {
				return math.MaxInt
			}
			if bottoms[i] != target {
				botCnt++
			} else if tops[i] != target {
				topCnt++
			}

		}
		return min(topCnt, botCnt)
	}
	// 想要每一行都一样，则必须跟第一列的bot或者top的数据是一致的
	ans := min(minCnts(tops[0]), minCnts(bottoms[0]))
	if ans == math.MaxInt {
		return -1
	}
	return ans
}
