package leetcode

import "math"

func minChanges(n int, k int) int {
	if n == k {
		return 0
	}
	res := 0
	for k > 0 || n > 0 {
		modK := k % 2
		modN := n % 2
		if modK != modN {
			if modN == 0 {
				return -1
			}
			res++
		}
		k = k / 2
		n = n / 2
	}
	return res
}

// https://leetcode.cn/problems/shopping-offers/description/
func shoppingOffers(price []int, special [][]int, needs []int) int {
	// 去除那些大礼包还不如单买更便宜的情况
	specials := filterSpecials(price, special)
	return shoppingOffersBackTrack(price, specials, needs, make([]int, len(price)))
}

func filterSpecials(price []int, specials [][]int) [][]int {
	newSpecials := [][]int{}
	for _, special := range specials {
		cost := 0
		for j := 0; j < len(special)-1; j++ {
			cost += special[j] * price[j]
		}
		if cost > special[len(special)-1] {
			newSpecials = append(newSpecials, special)
		}
	}
	return newSpecials
}

func shoppingOffersBackTrack(price []int, special [][]int, needs []int, current []int) int {
	flag := false
	for i := 0; i < len(current); i++ {
		if current[i] != needs[i] {
			flag = true
			break
		}
	}
	// 所有的都已经满足
	if !flag {
		return 0
	}
	res := math.MaxInt
	selectSepecial := false
loop1:
	for i := 0; i < len(special); i++ {
		for j := 0; j < len(special[i])-1; j++ {
			// 超过所需不能选
			if current[j]+special[i][j] > needs[j] {
				continue loop1
			}
		}
		selectSepecial = true
		for j := 0; j < len(special[i])-1; j++ {
			current[j] += special[i][j]
		}
		res = min(shoppingOffersBackTrack(price, special, needs, current)+special[i][len(special[i])-1], res)
		for j := 0; j < len(special[i])-1; j++ {
			current[j] -= special[i][j]
		}
	}
	if selectSepecial {
		return res
	}
	sum := 0
	for i := 0; i < len(needs); i++ {
		sum += (needs[i] - current[i]) * price[i]
	}
	return sum
}

// https://leetcode.cn/problems/sum-of-square-numbers/description/
// 直接使用2分查找解决
func judgeSquareSum(c int) bool {
	// 下界为0
	left := 0
	// 上界为sqrt(c)
	right := int(math.Sqrt(float64(c)))
	for left <= right {
		sum := left*left + right*right
		if sum == c {
			return true
		} else if sum > c {
			right--
		} else {
			left++
		}
	}
	return false
}

func losingPlayer(x int, y int) string {
	res := 1
	for x > 1 && y >= 4 {
		res ^= 1
		x -= 1
		y -= 4
	}
	if res == 0 {
		return "Alice"
	}
	return "Bob"
}
