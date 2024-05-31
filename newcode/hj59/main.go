package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewScanner(os.Stdin)
	for reader.Scan() {
		s := reader.Text()
		checkPwd(s)
	}
}

func checkPwd(str1 string) {
	if len(str1) < 8 {
		fmt.Println("NG")
		return
	}
	classCnt := make(map[int]struct{})
	checkRepeat := make(map[string]struct{})
	for i := 0; i < len(str1); i++ {
		if i < len(str1)-3 {
			if _, ok := checkRepeat[str1[i:i+3]]; ok {
				fmt.Println("NG")
				return
			} else {
				checkRepeat[str1[i:i+2]] = struct{}{}
			}
		}
		if '0' <= str1[i] && str1[i] <= '9' {
			classCnt[1] = struct{}{}
			continue
		}
		if 'a' <= str1[i] && str1[i] <= 'z' {
			classCnt[2] = struct{}{}
			continue
		}
		if 'A' <= str1[i] && str1[i] <= 'Z' {
			classCnt[3] = struct{}{}
			continue
		}
		classCnt[4] = struct{}{}
	}

	if len(classCnt) >= 3 {
		fmt.Println("OK")
	} else {
		fmt.Println("NG")
	}
}

func GetOne(str1 string) {
	myMap := make(map[uint8]int)
	for i := 0; i < len(str1); i++ {
		if _, ok := myMap[str1[i]]; ok {
			myMap[str1[i]]++
		} else {
			myMap[str1[i]] = 1
		}
	}

	for i := 0; i < len(str1); i++ {
		if myMap[str1[i]] == i {
			fmt.Println(string(str1[i]))
			return
		}
	}
	fmt.Println(-1)

}
