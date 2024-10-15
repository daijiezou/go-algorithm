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
