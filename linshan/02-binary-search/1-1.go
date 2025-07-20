package _2_binary_search

import "sort"

/*
>=x 直接使用lowerBound
< 相当于(>=x)-1 lowerBound(x)-1

>  相当于>=x+1 lowerBound(x+1)
<= 相当于(>x)-1 lowerBound(x+1)-1


x>=9 相当于 x
*/

// 适用于>=
func lowerBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] >= target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

// 34. 在排序数组中查找元素的第一个和最后一个位置
func searchRange(nums []int, target int) []int {
	start := lowerBound(nums, target)
	if start == len(nums) || nums[start] != target {
		return []int{-1, -1}
	}
	end := lowerBound(nums, target+1) - 1
	return []int{start, end}
}

// firstMax
func FirstMax(nums []int, target int) int {
	//第一个大于target的数
	return lowerBound(nums, target+1)
}

func nextGreatestLetter(letters []byte, target byte) byte {
	n := len(letters)
	left, right := 0, n-1
	for left <= right {
		mid := left + (right-left)>>1
		// >x 等效于 >=x+1
		if letters[mid] >= target+1 {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return letters[left%(n)]
}

func maximumCount(nums []int) int {
	// 找到第一个大于0的数
	// 找到第一个小于0的数 <0
	bigger := lowerBound(nums, 1)      // 找到>=1的第一个数
	smaller := lowerBound(nums, 0) - 1 // 找到>=0的第一个数-1
	n := len(nums)

	bigCnt := n - bigger
	smallCnt := smaller + 1
	return max(bigCnt, smallCnt)

}

// 给你一个下标从 0 开始、长度为 n 的整数数组 nums ，和两个整数 lower 和 upper ，返回 公平数对的数目 。
func countFairPairs(nums []int, lower int, upper int) int64 {
	sort.Ints(nums)
	ans := 0
	for i, num := range nums {
		l := lowerBound(nums[:i], lower-num)       // j >= lower - num
		r := lowerBound(nums[:i], upper-num+1) - 1 // j <= upper - num 相当于
		ans += r - l + 1
	}
	return int64(ans)
}

func isPossibleToSplit(nums []int) bool {
	numCnt := make([]int, 101)
	for i := 0; i < len(nums); i++ {
		numCnt[nums[i]]++
		if numCnt[nums[i]] > 2 {
			return false
		}
	}
	return true
}

func findPeakElement(nums []int) int {
	left := -1
	right := len(nums)
	for left+1 < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[mid+1] {
			right = mid
		} else {
			left = mid
		}
	}
	return right
}

func search(nums []int, target int) int {
	left, right := -1, len(nums)
	for left+1 < right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid
		} else if nums[mid] < target {
			left = mid
		}
	}
	return -1
}
