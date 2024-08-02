package leetcode

import (
	"math"
	"sort"
)

// 20040801
// https://leetcode.cn/problems/uOAnQW/
func maxmiumScore3(cards []int, cnt int) int {
	sort.Slice(cards, func(i, j int) bool {
		return cards[i] > cards[j]
	})
	ans := 0
	sum := 0
	minODD := math.MaxInt32
	minEven := math.MaxInt32
	for i := 0; i < cnt; i++ {
		sum += cards[i]
		if cards[i]%2 == 0 {
			minEven = min(minEven, cards[i])
		} else {
			minODD = min(minODD, cards[i])
		}
	}
	if sum%2 == 0 {
		return sum
	}
	nextOdd, nextEven := -1, -1
	for i := cnt; i < len(cards); i++ {
		if (nextOdd != -1) && (nextEven != -1) {
			break
		}
		if cards[i]%2 == 0 {
			if nextEven == -1 {
				nextEven = cards[i]
			}

		} else {
			if nextOdd == -1 {
				nextOdd = cards[i]
			}

		}
	}
	if minEven != math.MinInt32 && nextOdd != -1 {
		ans = max(ans, sum-minEven+nextOdd)
	}
	if minODD != math.MinInt32 && nextEven != -1 {
		ans = max(ans, sum-minODD+nextEven)
	}
	return ans
}

func numberOfRightTriangles(grid [][]int) int64 {
	m := len(grid)
	n := len(grid[0])
	hangMap := make(map[int]int, m)
	lieMap := make(map[int]int, n)
	oneList := make([][]int, 0)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				hangMap[i] += 1
				lieMap[j] += 1
				oneList = append(oneList, []int{i, j})
			}
		}
	}
	ans := 0
	for _, v := range oneList {
		i, j := v[0], v[1]
		hangCount := hangMap[i]
		lieCount := lieMap[j]
		if hangCount < 2 {
			continue
		}
		if lieCount < 2 {
			continue
		}
		ans += (hangCount - 1) * (lieCount - 1)
	}
	return int64(ans)
}
