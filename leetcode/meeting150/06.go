package meeting150

import (
	"fmt"
	"sort"
	"strconv"
)

/*
区间
*/

func summaryRanges(nums []int) []string {
	res := []string{}
	start := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1]+1 {
			if i-start > 1 {
				res = append(res, fmt.Sprintf("%d->%d", nums[start], nums[i-1]))
			} else {
				res = append(res, strconv.Itoa(nums[start]))
			}
			start = i
		}
	}
	if len(nums)-start > 1 {
		res = append(res, fmt.Sprintf("%d->%d", nums[start], nums[len(nums)-1]))
	} else if len(nums)-start == 1 {
		res = append(res, strconv.Itoa(nums[start]))
	}
	return res
}

func summaryRanges2(nums []int) []string {
	res := []string{}
	i, n := 0, len(nums)
	for i < n {
		start := i
		for i < n-1 && nums[i]+1 == nums[i+1] {
			i++
		}
		s := strconv.Itoa(nums[start])
		if start < i {
			s += "->" + strconv.Itoa(nums[i])
		}
		res = append(res, s)
		i++
	}
	return res
}

func merge2(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		curr := intervals[i]
		last := res[len(res)-1]
		if curr[0] > last[1] {
			res = append(res, curr)
		} else {
			res[len(res)-1][1] = max(curr[1], res[len(res)-1][1])
		}
	}

	return res
}

func insert(intervals [][]int, newInterval []int) [][]int {
	n := len(intervals)
	i := 0
	res := [][]int{}
	for i < n && intervals[i][1] < newInterval[0] {
		res = append(res, intervals[i])
		i++
	}
	for i < n && newInterval[1] >= intervals[i][0] {
		newInterval[0] = min(newInterval[0], intervals[i][0])
		newInterval[1] = max(newInterval[1], intervals[i][1])
		i++
	}
	res = append(res, newInterval)
	for i < n {
		res = append(res, intervals[i])
		i++
	}
	return res
}

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})
	cnt := 1
	right := points[0][1]
	for i := 1; i < len(points); i++ {
		cur := points[i]
		if cur[0] > right {
			cnt++
			right = cur[1]
		} else {
			right = min(right, cur[1])
		}
	}
	return cnt
}

func findMinArrowShots1(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	cnt := 1
	right := points[0][1]
	for i := 1; i < len(points); i++ {
		cur := points[i]
		if cur[0] > right {
			cnt++
			right = cur[1]
		}
	}
	return cnt
}
