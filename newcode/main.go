package newcode

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func brotherword() {
	var count int
	var k int
	var x string
	fmt.Scan(&count)
	words := make([]string, count)
	for i := 0; i < count; i++ {
		fmt.Scan(&words[i])
	}
	fmt.Scan(&x)
	fmt.Scan(&k)

	xMap := make(map[byte]int)
	for i := 0; i < len(x); i++ {
		xMap[x[i]]++
	}
	brotherList := make([]string, 0)
lool1:
	for i := 0; i < count; i++ {
		word := words[i]
		if len(word) != len(x) {
			continue
		}
		if word == x {
			continue
		}
		for key, val := range xMap {
			if strings.Count(word, string(key)) != val {
				continue lool1
			}
		}
		brotherList = append(brotherList, word)
	}
	sort.Strings(brotherList)
	fmt.Println(len(brotherList))
	if k <= len(brotherList) {
		fmt.Println(brotherList[k-1])
	}
}

func CoordinateShift() {
	origin := []int{0, 0}
	var s1 string
	fmt.Scan(&s1)
	orderList := strings.Split(s1, ";")
	for _, s := range orderList {
		if s == "" {
			continue
		}
		if len(s) == 1 {
			continue
		}
		fangxiang := s[0]
		distance := s[1:]
		distanceInt, err := strconv.Atoi(distance)
		if err != nil {
			continue
		}
		switch fangxiang {
		case 'A':
			origin[0] -= distanceInt
		case 'S':
			origin[1] -= distanceInt
		case 'W':
			origin[1] += distanceInt
		case 'D':
			origin[0] += distanceInt
		default:
			continue
		}
	}
	x := strconv.Itoa(origin[0])
	y := strconv.Itoa(origin[1])
	fmt.Println(x + "," + y)
}
