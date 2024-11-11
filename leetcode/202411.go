package leetcode

import (
	"fmt"
	"math"
	"sort"
)

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

/*
https://leetcode.cn/problems/find-the-power-of-k-size-subarrays-ii/description/
给你一个长度为 n 的整数数组 nums 和一个正整数 k 。
一个数组的 能量值 定义为：
如果 所有 元素都是依次 连续 且 上升 的，那么能量值为 最大 的元素。
否则为 -1 。
你需要求出 nums 中所有长度为 k 的
子数组的能量值。
请你返回一个长度为 n - k + 1 的整数数组 results ，其中 results[i] 是子数组 nums[i..(i + k - 1)] 的能量值。
*/
func resultsArray(nums []int, k int) []int {
	n := len(nums)
	res := make([]int, n-k+1)
	for i := range res {
		res[i] = -1
	}
	cnt := 0
	for i := 0; i < n-1; i++ {
		if i == 0 || nums[i] == nums[i-1]+1 {
			cnt++
		} else {
			cnt = 1
		}
		if cnt >= k {
			res[i-k+1] = nums[i]
		}
	}
	return res
}

type NeighborSum struct {
	grid    [][]int
	n       int
	Address map[int][2]int
}

func Constructor2(grid [][]int) NeighborSum {
	addr := make(map[int][2]int)
	for rowIndex, row := range grid {
		for colIndex, cell := range row {
			addr[cell] = [2]int{rowIndex, colIndex}
		}
	}
	return NeighborSum{grid: grid, Address: addr, n: len(grid)}
}

func (this *NeighborSum) AdjacentSum(value int) int {
	sum := 0
	addr := this.Address[value]
	row := addr[0]
	col := addr[1]
	if row > 0 {
		sum += this.grid[row-1][col]
	}
	if row < this.n-1 {
		sum += this.grid[row+1][col]
	}
	if col > 0 {
		sum += this.grid[row][col-1]
	}
	if col < this.n-1 {
		sum += this.grid[row][col+1]
	}
	return sum
}

func (this *NeighborSum) DiagonalSum(value int) int {
	sum := 0
	addr := this.Address[value]
	row := addr[0]
	col := addr[1]
	if row > 0 && col > 0 {
		sum += this.grid[row-1][col-1]
	}
	if row > 0 && col < this.n-1 {
		sum += this.grid[row-1][col+1]
	}
	if row < this.n-1 && col > 0 {
		sum += this.grid[row+1][col-1]
	}
	if row < this.n-1 && col < this.n-1 {
		sum += this.grid[row+1][col+1]
	}
	return sum
}

/*
有序数组中的单一元素
https://leetcode.cn/problems/single-element-in-a-sorted-array/description/
*/
func singleNonDuplicate(nums []int) int {
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		res ^= nums[i]
	}
	return res
}

/*
https://leetcode.cn/problems/minimum-cost-to-cut-a-stick/
*/
func minCost2(n int, cuts []int) int {
	cuts = append(cuts, 0, n)
	sort.Ints(cuts)
	fmt.Println(n)
	res := getCost(n, 0, len(cuts)-1, cuts)
	return res
}

func getCost(n, left, right int, cuts []int) int {
	if right-left <= 1 {
		return 0
	}
	res := math.MaxInt
	for i := left + 1; i < right; i++ {
		res = min(res, getCost(cuts[i], left, i, cuts)+getCost(n-cuts[i], i, right, cuts))
	}
	fmt.Println("length", cuts[right]-cuts[left])
	fmt.Println("=====")
	// 切割之前的木棍长度
	return res + cuts[right] - cuts[left]
}
