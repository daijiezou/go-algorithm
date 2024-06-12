package main

// https://www.nowcoder.com/discuss/564543169094799360?sourceSSR=users

func findLastValidIndex(S, L string) int {
	sLen, lLen := len(S), len(L)
	i, j := 0, 0
	lastIndex := -1

	for i < sLen && j < lLen {
		if S[i] == L[j] {
			lastIndex = j
			i++
		}
		j++
	}

	if i == sLen {
		return lastIndex
	}

	return -1
}

// https://www.nowcoder.com/discuss/583953228107108352?sourceSSR=users
func countPeaks(heights []int) int {
	n := len(heights)
	if n == 0 {
		return 0
	}

	peakCount := 0

	for i := 0; i < n; i++ {
		if i == 0 {
			// 第一个元素
			if n > 1 && heights[i] > heights[i+1] {
				peakCount++
			}
		} else if i == n-1 {
			// 最后一个元素
			if heights[i] > heights[i-1] {
				peakCount++
			}
		} else {
			// 中间元素
			if heights[i] > heights[i-1] && heights[i] > heights[i+1] {
				peakCount++
			}
		}
	}

	return peakCount
}
