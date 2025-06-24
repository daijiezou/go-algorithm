package aim_offer

import "sync"

var (
	x     string
	xOnce sync.Once
)

func singleton() string {
	if x == "" {
		xOnce.Do(func() {
			x = "1"
		})
	}
	return x
}

// 2.
/*
在一个长度为n的数组里的所有数字都在0~n-1的范围内。数组中某
些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了
几次。请找出数组中任意一个重复的数字。例如，如果输入长度为7的数
组｛2,3，1,0.2,5.3｝，那么对应的输出是重复的数字2或者3。
*/

func dup(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for nums[i] != i {
			if nums[nums[i]] == nums[i] {
				return nums[i]
			} else {
				nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
			}
		}
	}
	return -1
}

/*
在一个长度为n的数组里的所有数字都在0~n-1的范围内。数组中某
些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了
几次。请找出数组中任意一个重复的数字。例如，如果输入长度为7的数
组｛2,3，1,0.2,5.3｝，那么对应的输出是重复的数字2或者3。

不可以修改原数组
*/
func dup2(nums []int) int {
	start := 0
	end := len(nums) - 1
	for start <= end {
		mid := start + (end-start)/2
		cnt := getCnt(nums, start, mid)
		if start == end {
			if cnt > 1 {
				return start
			} else {
				break
			}
		}
		if cnt > mid-start+1 {
			end = mid
		} else {
			start = mid + 1
		}
	}
	return -1
}

func getCnt(nums []int, start, end int) int {
	cnt := 0
	for i := 0; i < len(nums); i++ {
		if start <= nums[i] && nums[i] <= end {
			cnt++
		}
	}
	return cnt
}

/*
题目：在一个二维数组中，每一行都按照从左到右递增的顺序排序，
每一列都按照从上到下递增的顺序排序。请完成一个函数，输入这样的一
个二维数组和一个整数，判断数组中是否含有该整数。
*/

func findTargetIn2DPlants(plants [][]int, target int) bool {
	row := len(plants)
	if row == 0 {
		return false
	}
	col := len(plants[0])
	if col == 0 {
		return false
	}
	x := 0
	y := col - 1
	for x < row && y >= 0 {
		if plants[x][y] == target {
			return true
		} else if plants[x][y] > target {
			y--
		} else if plants[x][y] < target {
			x++
		}
	}
	return false
}
