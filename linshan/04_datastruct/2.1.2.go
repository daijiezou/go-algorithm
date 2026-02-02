package _4_datastruct

/*
给你一个二维整数数组 squares ，其中 squares[i] = [xi, yi, li] 表示一个与 x 轴平行的正方形的左下角坐标和正方形的边长。

找到一个最小的 y 坐标，它对应一条水平线，该线需要满足它以上正方形的总面积 等于 该线以下正方形的总面积。

答案如果与实际答案的误差在 10-5 以内，将视为正确答案。

注意：正方形 可能会 重叠。重叠区域应该被 多次计数 。
*/
func separateSquares(squares [][]int) float64 {
	// 1. 计算总面积和最大 y 坐标
	totArea := 0
	maxY := 0
	for _, sq := range squares {
		l := sq[2]
		totArea += l * l
		maxY = max(maxY, sq[1]+l)
	}

	// 2. 定义检查函数：计算 y 下方的总面积
	check := func(y float64) bool {
		area := 0.0
		for _, sq := range squares {
			yi := float64(sq[1])
			if yi < y {
				l := float64(sq[2])
				// 正方形在 y 下方的高度
				h := min(y-yi, l)
				// 贡献面积 = 宽度 × 高度
				area += l * h
			}
		}
		return area >= float64(totArea)/2
	}

	// 3. 二分查找分界线
	left, right := 0.0, float64(maxY)
	// 二分精度：需要达到 10^-5，二分次数约为 log2(maxY * 10^5)
	for i := 0; i < 50; i++ { // 50 次足够达到精度要求
		mid := (left + right) / 2
		if check(mid) {
			right = mid
		} else {
			left = mid
		}
	}

	return (left + right) / 2
}
