package leetcode

import (
	"strings"
)

// https://leetcode.cn/problems/maximum-odd-binary-number/?envType=daily-question&envId=2024-03-13
func maximumOddBinaryNumber(s string) string {
	//sByte := []byte(s)
	//res := ""
	//firstOne := true
	//for _, value := range sByte {
	//	if string(value) == "1" {
	//		if firstOne {
	//			firstOne = false
	//			continue
	//
	//		} else {
	//			res = "1" + res
	//		}
	//	} else {
	//		res = res + "0"
	//	}
	//}
	//return res + "1"
	oneCnt := strings.Count(s, "1")
	return strings.Repeat("1", oneCnt-1) + strings.Repeat("0", len(s)-oneCnt) + "1"
}

func maximumBinaryString(binary string) string {
	if len(binary) < 2 {
		return binary
	}

	binaryBytes := []byte(binary)
	var preOne int
	var ans []byte
	for i := 0; i < len(binaryBytes); i++ {
		if binaryBytes[i] == '1' {
			preOne++
			ans = append(ans, '1')
		} else {
			break
		}
	}
	if preOne == len(binary) {
		return string(ans)
	}
	var suf int
	for i := preOne; i < len(binaryBytes); i++ {
		if binaryBytes[i] == '1' {
			suf++
		}
	}
	for i := 0; i < len(binary)-preOne-suf-1; i++ {
		ans = append(ans, '1')
	}
	ans = append(ans, '0')
	for i := 0; i < suf; i++ {
		ans = append(ans, '1')
	}
	return string(ans)
}

// leetcode 004
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	//res := merge(nums1, nums2)
	//if len(res)%2 == 0 {
	//	return float64(res[len(res)/2]+res[len(res)/2-1]) / 2
	//} else {
	//	return float64(res[len(res)/2])
	//}
	n := len(nums1)
	m := len(nums2)
	k1 := (n + m + 1) / 2
	k2 := (n + m + 2) / 2

	return (getKth(nums1, 0, n-1, nums2, 0, m-1, k1) + getKth(nums1, 0, n-1, nums2, 0, m-1, k2)) / 2
}

func getKth(nums1 []int, start1, end1 int, nums2 []int, start2, end2 int, k int) float64 {
	len1 := end1 - start1 + 1
	len2 := end2 - start2 + 1
	if len1 == 0 {
		return float64(nums2[start2+k-1])
	}
	if len2 == 0 {
		return float64(nums1[start1+k-1])
	}
	if k == 1 {
		return float64(Mymin(nums1[start1+k-1], nums2[start2+k-1]))
	}
	i := start1 + Mymin(len1, k/2) - 1
	j := start2 + Mymin(len2, k/2) - 1
	if nums1[i] > nums2[j] {
		// k 要减去排除的数字的个数
		return getKth(nums1, start1, end1, nums2, j+1, end2, k-(j-start2+1))
	} else {
		return getKth(nums1, i+1, end1, nums2, start2, end2, k-(i-start1+1))
	}
}

func Mymin(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func merge(nums1 []int, nums2 []int) []int {
	res := make([]int, 0, len(nums1)+len(nums2))
	for len(nums1) > 0 && len(nums2) > 0 {
		if nums1[0] <= nums2[0] {
			res = append(res, nums1[0])
			nums1 = nums1[1:]
		} else {
			res = append(res, nums2[0])
			nums2 = nums2[1:]
		}
	}
	for len(nums1) > 0 {
		res = append(res, nums1[0])
		nums1 = nums1[1:]
	}
	for len(nums2) > 0 {
		res = append(res, nums2[0])
		nums2 = nums2[1:]
	}
	return res
}
