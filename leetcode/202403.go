package leetcode

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
