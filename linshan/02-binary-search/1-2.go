package _2_binary_search

import (
	"math"
	"sort"
)

/*
	求最小
*/

func smallestDivisor(nums []int, threshold int) int {
	left := 1
	right := 1000000
	for left <= right {
		mid := left + (right-left)>>1
		if GetDivisor(nums, mid) > threshold {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

func GetDivisor(nums []int, divisor int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		if nums[i]%divisor != 0 {
			sum++
		}
		sum += nums[i] / divisor
	}

	return sum
}

// 875. 爱吃香蕉的珂珂
// https://leetcode.cn/problems/koko-eating-bananas/
func minEatingSpeed(piles []int, h int) int {
	left := 1
	right := 1000000000
	for left <= right {
		mid := left + (right-left)>>1
		if getHours(piles, mid) > h {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

func getHours(piles []int, k int) int {
	sum := 0
	for i := 0; i < len(piles); i++ {
		sum += piles[i] / k
		if piles[i]%k != 0 {
			sum++
		}
	}
	return sum
}

func findRadius2(houses, heaters []int) (ans int) {
	sort.Ints(heaters)
	for _, house := range houses {
		// 则 j 是满足 heaters[j]>house 的最小下标
		j := sort.SearchInts(heaters, house+1)
		minDis := math.MaxInt32
		if j < len(heaters) {
			minDis = heaters[j] - house
		}
		// 则 i 是满足 heaters[i]<=house 的最大下标
		i := j - 1
		if i >= 0 && house-heaters[i] < minDis {
			minDis = house - heaters[i]
		}
		if minDis > ans {
			ans = minDis
		}
	}
	return
}

func findRadius(houses []int, heaters []int) int {
	sort.Ints(houses)
	sort.Ints(heaters)
	r := 0
	for i := 0; i < len(houses); i++ {
		if houses[i] > r {
			r = houses[i]
		}
	}
	for i := 0; i < len(heaters); i++ {
		if heaters[i] > r {
			r = heaters[i]
		}
	}
	l := 0
	for l <= r {
		mid := l + (r-l)>>1
		if check(houses, heaters, mid) {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l
}

func check(houses []int, heaters []int, r int) bool {
	j := 0
	for i := 0; i < len(houses); i++ {
		for j < len(heaters) && houses[i] > heaters[j]+r {
			if j < len(heaters) && heaters[j]-r <= houses[i] && houses[i] <= heaters[j]+r {
				continue
			}
			return false
		}
	}
	return true
}

func repairCars(ranks []int, cars int) int64 {
	right := cars*cars*100 + 1
	left := 1
	for left <= right {
		mid := left + (right-left)>>1
		if getRepairCars(ranks, mid) >= cars {
			right = mid - 1

		} else {
			left = mid + 1
		}
	}
	return int64(left)
}

func getRepairCars(ranks []int, k int) int {
	sum := 0
	for i := 0; i < len(ranks); i++ {
		sum += int(math.Sqrt(float64(k) / float64(ranks[i])))
	}
	return sum
}

/*
给你一个整数数组 bloomDay，以及两个整数 m 和 k 。
现需要制作 m 束花。制作花束时，需要使用花园中 相邻的 k 朵花 。
花园中有 n 朵花，第 i 朵花会在 bloomDay[i] 时盛开，恰好 可以用于 一束 花中。
请你返回从花园中摘 m 束花需要等待的最少的天数。如果不能摘到 m 束花则返回 -1 。
*/
func minDays(bloomDay []int, m int, k int) int {
	n := len(bloomDay)
	if m*k > n {
		return -1
	}
	right := 1000000000
	left := 1
	for left <= right {
		mid := left + (right-left)/2
		if checkBloom(bloomDay, m, k, mid) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

func checkBloom(bloomDay []int, m int, k int, days int) bool {
	n := len(bloomDay)
	total := 0
	for i := 0; i < n; i++ {
		// 成熟了
		if bloomDay[i] <= days {
			canMakeBloom := true
			if n-i < k {
				break
			}
			j := i
			for ; j < i+k; j++ {
				if bloomDay[j] > days {
					canMakeBloom = false
					i = j
					break
				}
			}
			if canMakeBloom {
				i = j - 1
				total++
			}
		}
	}
	return total >= m
}
