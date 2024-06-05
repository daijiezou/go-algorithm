package main

import (
	"fmt"
)

type Good struct {
	v    int  //价格
	p    int  //价值：价格*满意度
	main bool //是否为主件
	a1   int  // 附件1
	a2   int  // 附件2
}

func main() {
	var totalMoney int
	var count int
	fmt.Scan(&totalMoney, &count)
	goods := make([]Good, count)
	//scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < count; i++ {
		goods[i] = Good{
			a1: -1,
			a2: -1,
		}
	}
	for i := 0; i < count; i++ {
		var v, p, q int
		fmt.Scan(&v) //价格
		fmt.Scan(&p) //价值
		p = v * p
		fmt.Scan(&q) //是否主件
		goods[i].v = v
		goods[i].p = p
		if q == 0 {
			goods[i].main = true
		} else {
			if goods[q-1].a1 == -1 {
				goods[q-1].a1 = i
			} else {
				goods[q-1].a2 = i
			}
		}
	}
	//fmt.Println(goods)
	// dp[i][j] 表示选择在前i件物品里，总钱数为j的最大价值
	dp := make([][]int, count+1)
	for i := range dp {
		dp[i] = make([]int, totalMoney+1)
	}
	for i := 1; i <= count; i++ {
		for j := 0; j <= totalMoney; j++ {
			// 什么都不选
			dp[i][j] = dp[i-1][j]
			if !goods[i-1].main { //附件直接跳过
				continue
			}

			//情况二：只选择主件
			if j >= goods[i-1].v {
				dp[i][j] = max(dp[i][j], dp[i-1][j-goods[i-1].v]+goods[i-1].p)
			}
			//情况三：只选择主件和第一个附件
			if goods[i-1].a1 != -1 && j >= goods[i-1].v+goods[goods[i-1].a1].v {
				dp[i][j] = max(dp[i][j], dp[i-1][j-goods[i-1].v-goods[goods[i-1].a1].v]+goods[i-1].p+goods[goods[i-1].a1].p)
			}
			//情况四：只选择主件和第二个附件
			if goods[i-1].a2 != -1 && j >= goods[i-1].v+goods[goods[i-1].a2].v {
				dp[i][j] = max(dp[i][j], dp[i-1][j-goods[i-1].v-goods[goods[i-1].a2].v]+goods[i-1].p+goods[goods[i-1].a2].p)
			}
			//情况五：选择主件和两个附件
			if goods[i-1].a1 != -1 && goods[i-1].a2 != -1 && j >= goods[i-1].v+goods[goods[i-1].a1].v+goods[goods[i-1].a2].v {
				dp[i][j] = max(dp[i][j], dp[i-1][j-goods[i-1].v-goods[goods[i-1].a1].v-goods[goods[i-1].a2].v]+goods[i-1].p+goods[goods[i-1].a1].p+goods[goods[i-1].a2].p)
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
