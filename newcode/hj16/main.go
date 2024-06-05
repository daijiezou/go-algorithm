package main

import (
	"fmt"
)

func main() {
	var totalMoney int
	var count int
	fmt.Scan(&totalMoney, &count)
	prices := make([]int, count+1)
	values := make([]int, count+1)
	owners := make([]int, count+1)
	//scanner := bufio.NewScanner(os.Stdin)
	for i := 1; i <= count; i++ {
		fmt.Scan(&prices[i]) //价格
		fmt.Scan(&values[i]) //价值
		values[i] = values[i] * prices[i]
		fmt.Scan(&owners[i]) //是否主件
	}
	//fmt.Println(goods)
	// dp[i][j] 表示选择在前i件物品里，总钱数为j的最大价值
	dp := make([][]int, count+1)
	for i := range dp {
		dp[i] = make([]int, totalMoney+1)
	}
	for i := 2; i <= count; i++ {
		for j := 2; j <= totalMoney; j++ {
			if owners[i-1] == 0 { //主键
				if prices[i-1] <= j {
					dp[i][j] = max(dp[i-1][j], dp[i-1][j-prices[i-1]]+values[i-1])
				}
			} else { //附件
				if prices[i-1]+prices[owners[i-1]] <= j {
					//附件的话 加上主件一起算
					dp[i][j] = max(dp[i-1][j], dp[i-1][j-prices[i-1]]+values[i-1])
				}
			}

		}
	}
	fmt.Println(dp[count][totalMoney])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
