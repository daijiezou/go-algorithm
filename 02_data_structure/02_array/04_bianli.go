package _2_array

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	s = strings.Trim(s, " ")
	for strings.Contains(s, "  ") {
		s = strings.ReplaceAll(s, "  ", " ")
	}
	sbyte := []byte(s)
	reverseString(&sbyte, 0, len(sbyte)-1)
	fmt.Println(string(sbyte))
	for i := 0; i < len(sbyte); i++ {
		start := i
		for start < len(sbyte) && sbyte[start] != ' ' {
			start++
		}
		fmt.Printf("%d:%d\n", i, start-1)
		reverseString(&sbyte, i, start-1)
		i = start
	}
	return string(sbyte)
}

func reverseString(sBytes *[]byte, left, right int) {
	for left < right {
		(*sBytes)[left], (*sBytes)[right] = (*sBytes)[right], (*sBytes)[left]
		left++
		right--
	}
}

// 将二维矩阵原地顺时针旋转 90 度
func rotate(matrix [][]int) {
	m := len(matrix)

	// 先将矩阵按照对角线翻转
	for i := 0; i < m; i++ {
		for j := i; j < m; j++ {
			tmp := matrix[i][j]
			matrix[i][j] = matrix[j][i]
			matrix[j][i] = tmp
		}
	}
	// 翻转每一行
	for i := 0; i < m; i++ {
		reverseList(matrix[i])
	}

}

func reverseList(req []int) {
	length := len(req)
	for i := 0; i < length/2; i++ {
		req[i], req[length-i-1] = req[length-i-1], req[i]
	}
}
