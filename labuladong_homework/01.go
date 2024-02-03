package labuladong_homework

import "fmt"

var path []rune

func generateBinaryNumber(n int) {
	if n == 0 {
		fmt.Println(path)
		return
	}
	for i := 0; i < 2; i++ {
		path = append(path, rune(i))
		generateBinaryNumber(n - 1)
		path = path[0 : len(path)-1]
	}
}
