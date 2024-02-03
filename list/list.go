package list

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var slow, fast int
	length := len(nums)
	for ; fast < length; fast++ {
		if nums[slow] != nums[fast] {
			slow++
			nums[slow] = nums[fast]
		}
	}
	return slow + 1
}
