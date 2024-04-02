package _1_core_platform

import (
	"fmt"
	"sort"
)

/*
https://leetcode.cn/problems/permutations/
全排列问题
*/
func permute(nums []int) [][]int {
	res := [][]int{}
	track := []int{}
	used := make([]bool, len(nums))

	backtrack(nums, track, used, &res)
	fmt.Println(res)
	return res
}

func backtrack(nums []int, track []int, used []bool, res *[][]int) {
	if len(track) == len(nums) {
		temp := make([]int, len(track))
		copy(temp, track)
		*res = append(*res, temp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}
		used[i] = true
		track = append(track, nums[i])
		backtrack(nums, track, used, res)
		used[i] = false
		track = track[:len(track)-1]
	}
}

/*
https://leetcode.cn/problems/subsets/submissions/518831435/
输入一个无重复元素的数组 nums，其中每个元素最多使用一次，请你返回 nums 的所有子集。
*/
func subsets(nums []int) [][]int {
	res := [][]int{}
	track := []int{}
	subsetsBacktrack(0, nums, track, &res)
	return res
}

func subsetsBacktrack(start int, nums []int, track []int, res *[][]int) {
	temp := make([]int, len(track))
	copy(temp, track)
	*res = append(*res, temp)
	for i := start; i < len(nums); i++ {
		// 做出选择
		track = append(track, nums[i])
		// 递归进入下一个状态
		subsetsBacktrack(i+1, nums, track, res)
		// 撤销选择
		track = track[:len(track)-1]
	}
}

/*
https://leetcode.cn/problems/combinations/
给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。
*/
func combine(n int, k int) [][]int {
	res := [][]int{}
	track := []int{}
	combineBacktrack(0, n, k, track, &res)
	return res
}

func combineBacktrack(start int, n int, k int, track []int, res *[][]int) {
	if len(track) == k {
		temp := make([]int, len(track))
		copy(temp, track)
		*res = append(*res, temp)
		return
	}
	for i := start; i < n; i++ {
		// 做出选择
		track = append(track, i+1)
		// 递归进入下一个状态
		combineBacktrack(i+1, n, k, track, res)
		// 撤销选择
		track = track[:len(track)-1]
	}
}

/*
一个整数数组 nums，其中可能包含重复元素，请你返回该数组所有可能的子集。
https://leetcode.cn/problems/subsets-ii/
*/
func subsetsWithDup(nums []int) [][]int {
	res := [][]int{}
	track := []int{}
	sort.Ints(nums)
	subsetsWithDupTrack(0, nums, track, &res)
	return res
}

func subsetsWithDupTrack(start int, nums []int, track []int, res *[][]int) {
	temp := make([]int, len(track))
	copy(temp, track)
	*res = append(*res, temp)
	for i := start; i < len(nums); i++ {
		// 剪枝逻辑，值相同的相邻树枝，只遍历第一条
		if i > start && nums[i] == nums[i-1] {
			continue
		}
		// 做出选择
		track = append(track, nums[i])
		// 递归进入下一个状态
		subsetsWithDupTrack(i+1, nums, track, res)
		// 撤销选择
		track = track[:len(track)-1]
	}
}

/*
https://leetcode.cn/problems/combination-sum-ii/
输入 candidates 和一个目标和 target，从 candidates 中找出中所有和为 target 的组合。
candidates 可能存在重复元素，且其中的每个数字最多只能使用一次。
*/
func combinationSum2(candidates []int, target int) [][]int {
	res := [][]int{}
	track := []int{}
	sort.Ints(candidates)
	combinationSum2Track(0, target, candidates, track, &res)
	return res
}

func combinationSum2Track(start int, target int, nums []int, track []int, res *[][]int) {
	if Sum(track) == target {
		temp := make([]int, len(track))
		copy(temp, track)
		*res = append(*res, temp)
		return
	}
	if Sum(track) > target {
		return
	}

	for i := start; i < len(nums); i++ {
		// 剪枝逻辑，值相同的相邻树枝，只遍历第一条
		if i > start && nums[i] == nums[i-1] {
			continue
		}
		// 做出选择
		track = append(track, nums[i])
		// 递归进入下一个状态
		combinationSum2Track(i+1, target, nums, track, res)
		// 撤销选择
		track = track[:len(track)-1]
	}
}

func Sum(nums []int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	return sum
}

/*
https://leetcode.cn/problems/permutations-ii/
输入一个可包含重复数字的序列 nums，请你写一个算法，返回所有可能的全排列
*/
func permuteUnique(nums []int) [][]int {
	res := [][]int{}
	track := []int{}
	used := make([]bool, len(nums))
	sort.Ints(nums)
	permuteUniqueBacktrack(nums, track, used, &res)
	fmt.Println(res)
	return res
}

func permuteUniqueBacktrack(nums []int, track []int, used []bool, res *[][]int) {
	if len(track) == len(nums) {
		temp := make([]int, len(track))
		copy(temp, track)
		*res = append(*res, temp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}

		/*
			当出现重复元素时，比如输入 nums = [1,2,2',2'']，
			2' 只有在 2 已经被使用的情况下才会被选择，
			同理，2'' 只有在 2' 已经被使用的情况下才会被选择，
			这就保证了相同元素在排列中的相对位置保证固定。
		*/

		// 如果前面的相邻相等元素没有用过，则跳过
		if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
			continue
		}
		used[i] = true
		track = append(track, nums[i])
		permuteUniqueBacktrack(nums, track, used, res)
		used[i] = false
		track = track[:len(track)-1]
	}
}

func permuteUniqueBacktrack2(nums []int, track []int, used []bool, res *[][]int) {
	if len(track) == len(nums) {
		temp := make([]int, len(track))
		copy(temp, track)
		*res = append(*res, temp)
		return
	}
	prevNum := -666
	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}
		if nums[i] == prevNum {
			continue
		}
		prevNum = nums[i]
		used[i] = true
		track = append(track, nums[i])
		permuteUniqueBacktrack(nums, track, used, res)
		used[i] = false
		track = track[:len(track)-1]
	}
}
