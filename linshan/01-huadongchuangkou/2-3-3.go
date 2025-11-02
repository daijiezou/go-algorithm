package _1_huadongchuangkou

/*
§2.3.3 恰好型滑动窗口
例如，要计算有多少个元素和恰好等于 k 的子数组，可以把问题变成：

计算有多少个元素和 ≥k 的子数组。
计算有多少个元素和 >k，也就是 ≥k+1 的子数组。
答案就是元素和 ≥k 的子数组个数，减去元素和 ≥k+1 的子数组个数。这里把 > 转换成 ≥，从而可以把滑窗逻辑封装成一个函数 f，然后用 f(k) - f(k + 1) 计算，无需编写两份滑窗代码。

总结：「恰好」可以拆分成两个「至少」，也就是两个「越长越合法」的滑窗问题。

注：也可以把问题变成 ≤k 减去 ≤k−1（两个至多）。可根据题目选择合适的变形方式。

注：也可以把两个滑动窗口合并起来，维护同一个右端点 right 和两个左端点 left
*/
func numSubarraysWithSum(nums []int, goal int) int {
	goalCnt := 0
	goal1Cnt := 0
	left := 0
	sum := 0
	sum1 := 0
	length := len(nums)
	left2 := 0
	for right, x := range nums {
		sum += x

		for sum >= goal && left <= right {
			goalCnt += length - right
			sum -= nums[left]
			left++
		}
		sum1 += x
		for sum1 >= goal+1 && left2 <= right {
			goal1Cnt += length - right
			sum1 -= nums[left2]
			left2++
		}
	}
	return goalCnt - goal1Cnt
}

/*
https://leetcode.cn/problems/count-number-of-nice-subarrays/description/
*/
func numberOfSubarrays(nums []int, k int) int {
	left1, left2 := 0, 0
	oddCnt1 := 0
	oddCnt2 := 0
	sum1 := 0
	sum2 := 0
	length := len(nums)
	for right, x := range nums {
		if x%2 == 1 {
			oddCnt1++
			oddCnt2++
		}
		for oddCnt1 >= k && left1 <= right {
			sum1 += length - right
			if nums[left1]%2 == 1 {
				oddCnt1--
			}
			left1++
		}

		for oddCnt2 >= k+1 && left2 <= right {
			sum2 += length - right
			if nums[left2]%2 == 1 {
				oddCnt2--
			}
			left2++
		}
	}
	return sum1 - sum2
}
