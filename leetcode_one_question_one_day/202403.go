package leetcode_one_question_one_day

import "strings"

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
