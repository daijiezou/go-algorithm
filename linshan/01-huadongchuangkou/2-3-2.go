package _1_huadongchuangkou

/*
一般要写 ans += right - left + 1 .
滑动窗口的内层循环结束时，右端点固定在 right，左端点在 left, left + 1，•••，right的所有子数组
（子串）都是合法的，这一共有 right—left+1个。
*/

// 包含所有三种字符的子字符串数目
// 给你一个字符串 s ，它只包含三种字符 a, b 和 c 。
// 请你返回 a，b 和 c 都 至少 出现过一次的子字符串数目。
// https://leetcode.cn/problems/number-of-substrings-containing-all-three-characters/description/
func numberOfSubstrings(s string) int {
	res := 0
	window := make(map[byte]int, 0)
	left := 0
	for i := 0; i < len(s); i++ {
		window[s[i]]++
		for len(window) == 3 {
			// 该子数组满足，后面的包含该子数组的都满足
			res += len(s) - i
			window[s[left]]--
			left++
			if window[s[left]] == 0 {
				delete(window, s[left])
			}
		}
	}
	return res
}

// https://leetcode.cn/problems/count-substrings-that-satisfy-k-constraint-i/
/*
给你一个 二进制字符串 s 和一个整数 K。
如果一个 二进制字符串 满足以下任一条件，则认为该字符串满足k约束：
•字符串中0的数量最多为K。
•字符串中1的数量最多为 K。
返回一个整数，表示 s的所有满足k约束 的子字符串的数量。
*/
func countKConstraintSubstrings(s string, k int) int {
	cnt0 := 0
	cnt1 := 0
	left := 0
	res := 0
	for right := 0; right < len(s); right++ {
		if s[right] == '0' {
			cnt0++
		} else {
			cnt1++
		}
		for cnt0 > k && cnt1 > k {
			if s[left] == '0' {
				cnt0--
			} else {
				cnt1--
			}
			left++
		}
		res += right - left + 1
	}
	return res
}

/*
一个数组的 分数 定义为数组之和 乘以 数组的长度。
•比方说，［1，2，3，4，5］） 的分数为（1+2+3+4+5）*5=75。
给你一个正整数数组 nums 和一个整数 K），请你返回 nums 中分数 严格小于K的非空整数子数组数目。
子数组 是数组中的一个连续元素序列。
*/

// https://leetcode.cn/problems/count-subarrays-with-score-less-than-k/

func countSubarrays2(nums []int, k int64) int64 {
	res := 0
	left, right := 0, 0
	sum := 0
	for ; right < len(nums); right++ {
		sum += nums[right]
		for int64(sum*(right-left+1)) >= k {
			sum -= nums[left]
			left++
		}
		res += right - left + 1
	}
	return int64(res)
}

// https://leetcode.cn/problems/continuous-subarrays/
func continuousSubarrays(nums []int) int64 {
	minQ := make([]int, 0)
	maxQ := make([]int, 0)
	left := 0
	ans := int64(0)
	for right, x := range nums {
		for len(minQ) > 0 && x <= nums[minQ[len(minQ)-1]] {
			minQ = minQ[:len(minQ)-1]
		}
		minQ = append(minQ, right)
		for len(maxQ) > 0 && x >= nums[maxQ[len(maxQ)-1]] {
			maxQ = maxQ[:len(maxQ)-1]
		}
		maxQ = append(maxQ, right)

		// 检查最大值减最小值是否大于2

		/*
			条件是判断是否窗口内最大值减最小值是否大于2，
			所以说当窗口中只有一个值的时候肯定是满足条件的，所以说单调队列并不会为空，
		*/
		for nums[maxQ[0]]-nums[minQ[0]] > 2 {

			// 需要缩减窗口
			left++
			// 检测单调队列里的值是否已经出队
			if minQ[0] < left {
				minQ = minQ[1:]
			}
			if maxQ[0] < left {
				maxQ = maxQ[1:]
			}
		}
		ans += int64(right - left + 1)
	}
	return ans
}

func beautifulBouquet(flowers []int, cnt int) int {
	const mod = 1e9 + 7

	window := make(map[int]int)
	left := 0
	res := 0
	for right, x := range flowers {
		window[x]++
		for window[x] > cnt {
			window[flowers[left]]--
			left++
		}
		res += (right - left + 1) % mod
	}
	return res
}

/*
2799. 统计完全子数组的数目
给你一个由 正 整数组成的数组 nums 。
如果数组中的某个子数组满足下述条件，则称之为 完全子数组 ：
子数组中 不同 元素的数目等于整个数组不同元素的数目。
请返回 nums 中完全子数组的数目。
*/
func countCompleteSubarrays(nums []int) int {
	set := make(map[int]struct{})

	for i := 0; i < len(nums); i++ {
		set[nums[i]] = struct{}{}
	}
	n := len(set)
	left, right := 0, 0
	window := make(map[int]int)
	sum := 0
	for ; right < len(nums); right++ {
		window[nums[right]]++
		for len(window) == n {
			window[nums[left]]--
			if window[nums[left]] == 0 {
				delete(window, nums[left])
			}
			left++
		}
		sum += left
	}
	return sum
}
