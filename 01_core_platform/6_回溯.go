package _1_core_platform

// Subsets : The function mimics the method subsets in Java class
func subsets2(nums []int) [][]int {
	var res [][]int
	track := []int{}
	// global function
	var backtrack func(nums []int, i int)
	backtrack = func(nums []int, i int) {
		if i == len(nums) {
			dst := make([]int, len(track))
			copy(dst, track)
			res = append(res, dst)
			return
		}
		// 做第一种选择，元素在子集中
		track = append(track, nums[i])
		backtrack(nums, i+1)
		// 撤销选择
		track = track[0 : len(track)-1]
		// 做第二种选择，元素不在子集中
		backtrack(nums, i+1)
	}
	backtrack(nums, 0)
	return res
}
