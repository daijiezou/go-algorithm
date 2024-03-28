package __核心框架

func removeDuplicates(nums []int) int {
	length := len(nums)
	slow := 0
	fast := 0
	for ; fast < length; fast++ {
		if nums[slow] != nums[fast] {
			slow++
			nums[slow] = nums[fast]
		}
	}
	return slow + 1
}

func removeElement(nums []int, val int) int {
	length := len(nums)
	slow := 0
	fast := 0
	for ; fast < length; fast++ {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}

func moveZeroes(nums []int) {
	length := len(nums)
	slow := 0
	fast := 0
	for ; fast < length; fast++ {
		if nums[fast] != 0 {
			nums[slow] = nums[fast]
			slow++
		}
	}
	for ; slow < length; slow++ {
		nums[slow] = 0
	}
}

func twoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left, right}
		}
		if sum < target {
			left++
		}
		if sum > target {
			right--
		}
	}
	return []int{}
}
