package aim_offer

// https://leetcode.cn/problems/xuan-zhuan-shu-zu-de-zui-xiao-shu-zi-lcof/?envType=problem-list-v2&envId=8LSpuXqD
func inventoryManagement(stock []int) int {
	n := len(stock)
	left := -1
	right := n - 1
	for left+1 < right { //left+1=right 退出循环，此时区间内不包括任何元素了
		mid := left + (right-left)/2
		if stock[mid] < stock[right] {
			right = mid
		} else if stock[mid] == stock[right] {
			right = right - 1
		} else {
			left = mid
		}
	}
	// 此时返回right是目标值
	return stock[right]
}
