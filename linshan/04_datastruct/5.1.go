package _4_datastruct

import (
	"container/heap"
	"math"
	"sort"
)

type myInts []int

func (m myInts) Len() int {
	return len(m)
}

func (m myInts) Less(i, j int) bool {
	return m[i] > m[j]
}

func (m myInts) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *myInts) Push(x any) {
	*m = append(*m, x.(int))
}

// pop应该弹出最后一个，而不是第一个
func (m *myInts) Pop() any {
	n := len(*m)
	x := (*m)[n-1]
	*m = (*m)[:n-1]
	return x
}

func lastStoneWeight(stones []int) int {
	newStones := myInts(stones)
	q := &newStones
	heap.Init(q)
	for q.Len() > 1 {
		x, y := heap.Pop(q).(int), heap.Pop(q).(int)
		if x > y {
			heap.Push(q, x-y)
		}
	}
	if q.Len() > 0 {
		return (*q)[0]
	}
	return 0
}

/*
给你一个整数数组 gifts ，表示各堆礼物的数量。每一秒，你需要执行以下操作：

选择礼物数量最多的那一堆。
如果不止一堆都符合礼物数量最多，从中选择任一堆即可。
将堆中的礼物数量减少到堆中原来礼物数量的平方根，向下取整。
返回在 k 秒后剩下的礼物数量。
https://leetcode.cn/problems/take-gifts-from-the-richest-pile/description/
*/

func pickGifts(gifts []int, k int) int64 {
	h := &hp{gifts}
	heap.Init(h)
	for i := 0; i < k && gifts[0] > 1; i++ {
		gifts[0] = int(math.Sqrt(float64(gifts[0]))) // 直接修改堆顶
		heap.Fix(h, 0)
	}
	sum := 0
	for i := 0; i < len(gifts); i++ {
		sum += gifts[i]
	}
	return int64(sum)
}

/*
给你一个下标从 0 开始的整数数组 nums 和一个整数 k 。你的 起始分数 为 0 。

在一步 操作 中：

选出一个满足 0 <= i < nums.length 的下标 i ，
将你的 分数 增加 nums[i] ，并且
将 nums[i] 替换为 ceil(nums[i] / 3) 。
返回在 恰好 执行 k 次操作后，你可能获得的最大分数。

向上取整函数 ceil(val) 的结果是大于或等于 val 的最小整数。
*/
func maxKelements(nums []int, k int) int64 {
	h := &hp{nums}
	heap.Init(h)
	sum := 0
	for ; k > 0; k-- {
		x := heap.Pop(h).(int)
		sum += x
		x = (x + 2) / 3
		heap.Push(h, x)
	}
	return int64(sum)
}

/*
给你一个整数数组 piles ，数组 下标从 0 开始 ，
其中 piles[i] 表示第 i 堆石子中的石子数量。另给你一个整数 k ，请你执行下述操作 恰好 k 次：

选出任一石子堆 piles[i] ，并从中 移除 floor(piles[i] / 2) 颗石子。
注意：你可以对 同一堆 石子多次执行此操作。

返回执行 k 次操作后，剩下石子的 最小 总数。

floor(x) 为 小于 或 等于 x 的 最大 整数。（即，对 x 向下取整）。
https://leetcode.cn/problems/remove-stones-to-minimize-the-total/description/
*/
func minStoneSum(piles []int, k int) int {
	h := &hp{piles}
	heap.Init(h)
	for ; k > 0; k-- {
		piles[0] -= piles[0] / 2
		heap.Fix(h, 0)
	}
	sum := 0
	for _, x := range piles {
		sum += x
	}
	return sum
}

/*
设计一个找到数据流中第 k 大元素的类（class）。注意是排序后的第 k 大元素，不是第 k 个不同的元素。

请实现 KthLargest 类：

KthLargest(int k, int[] nums) 使用整数 k 和整数流 nums 初始化对象。
int add(int val) 将 val 插入数据流 nums 后，返回当前数据流中第 k 大的元素。
https://leetcode.cn/problems/kth-largest-element-in-a-stream/description/
*/
type hp struct{ sort.IntSlice }

func (h hp) Less(i, j int) bool {
	return h.IntSlice[i] > h.IntSlice[j]
} // 最大堆

func (h *hp) Push(x any) {
	h.IntSlice = append(h.IntSlice, x.(int))
}

func (h *hp) Pop() any {
	n := len(h.IntSlice)
	x := h.IntSlice[n-1]
	h.IntSlice = h.IntSlice[:n-1]
	return x
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] } // 小顶堆的Less规则
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type KthLargest struct {
	heap IntHeap
	k    int
}

func Constructor(k int, nums []int) KthLargest {
	h := &IntHeap{}
	heap.Init(h)
	kth := KthLargest{heap: *h, k: k}
	for _, num := range nums {
		kth.Add(num)
	}
	return kth
}

func (this *KthLargest) Add(val int) int {
	heap.Push(&this.heap, val)
	// 如果堆大小超过k，弹出最小的元素（堆顶），保持堆大小为k
	if this.heap.Len() > this.k {
		heap.Pop(&this.heap)
	}
	// 堆顶就是第k大元素
	return this.heap[0]
}

/*


给你一个正整数数组 nums 。每一次操作中，你可以从 nums 中选择 任意 一个数并将它减小到 恰好 一半。
（注意，在后续操作中你可以对减半过的数继续执行操作）

请你返回将 nums 数组和 至少 减少一半的 最少 操作数。
*/

type FloatHeap []float64

func (h FloatHeap) Len() int           { return len(h) }
func (h FloatHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h FloatHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *FloatHeap) Push(x interface{}) {
	*h = append(*h, x.(float64))
}

func (h *FloatHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func halveArray(nums []int) int {
	total := float64(0)
	h := &FloatHeap{}
	for _, num := range nums {
		val := float64(num)
		total += val
		heap.Push(h, val)
	}

	target := total / 2
	reduced := float64(0)
	ops := 0

	for reduced < target {
		maxVal := heap.Pop(h).(float64)
		reduced += maxVal / 2
		heap.Push(h, maxVal/2)
		ops++
	}

	return ops
}
func halveArray2(nums []int) int {
	total := 0
	for i := 0; i < len(nums); i++ {
		nums[i] = nums[i] << 20
		total += nums[i]
	}
	q := &hp{nums}
	heap.Init(q)
	sub := 0
	ops := 0
	for sub < total/2 {
		sub += q.IntSlice[0] / 2
		q.IntSlice[0] /= 2
		ops++
		heap.Fix(q, 0)
	}
	return ops
}

/*
给你一个非负整数数组 nums 和一个整数 k 。每次操作，你可以选择 nums 中 任一 元素并将它 增加 1 。

请你返回 至多 k 次操作后，能得到的 nums的 最大乘积 。由于答案可能很大，请你将答案对 109 + 7 取余后返回。
*/
type minHeap struct {
	sort.IntSlice
}

func (m *minHeap) Push(x any) {
	m.IntSlice = append(m.IntSlice, x.(int))
}

func (m *minHeap) Pop() any {
	n := len(m.IntSlice)
	x := m.IntSlice[n-1] // ✅ 取最后一个元素
	m.IntSlice = m.IntSlice[:n-1]
	return x
}

func maximumProduct_2233(nums []int, k int) int {
	// 每次都选择最小的数字加1
	h := &minHeap{nums}
	heap.Init(h)
	for ; k > 0; k-- {
		h.IntSlice[0]++
		heap.Fix(h, 0)
	}

	res := 1
	for _, x := range nums {
		res *= x
		res = res % mod
	}
	return res % mod
}

// https://leetcode.cn/problems/the-number-of-the-smallest-unoccupied-chair/
/*
有 n 个朋友在举办一个派对，这些朋友从 0 到 n - 1 编号。
派对里有 无数 张椅子，编号为 0 到 infinity 。当一个朋友到达派对时，他会占据 编号最小 且未被占据的椅子。
比方说，当一个朋友到达时，如果椅子 0 ，1 和 5 被占据了，那么他会占据 2 号椅子。
当一个朋友离开派对时，他的椅子会立刻变成未占据状态。如果同一时刻有另一个朋友到达，可以立即占据这张椅子。
给你一个下标从 0 开始的二维整数数组 times ，其中 times[i] = [arrivali, leavingi] 表示第 i 个朋友到达和离开的时刻，同时给你一个整数 targetFriend 。所有到达时间 互不相同 。
请你返回编号为 targetFriend 的朋友占据的 椅子编号 。
*/

func smallestChair(times [][]int, targetFriend int) int {
	// 按时间顺序，记录每个到达事件和离开事件相对应的朋友编号

	type timeRange struct {
		Arrive []int
		Leave  []int
	}
	events := make([]timeRange, 1e5+1)
	for i, t := range times {
		arrive, leave := t[0], t[1]
		events[arrive].Arrive = append(events[arrive].Arrive, i) // 朋友到达
		events[leave].Leave = append(events[leave].Leave, i)     // 朋友离开
	}

	// 初始化未被占据的椅子
	n := len(times)
	unoccupied := hp{make([]int, n)}
	for i := range unoccupied.IntSlice {
		unoccupied.IntSlice[i] = i
	}

	// 按时间顺序扫描每个事件
	belong := make([]int, n)
	for _, e := range events {
		for _, id := range e.Leave { // 朋友离开
			heap.Push(&unoccupied, belong[id]) // 返还椅子
		}
		for _, id := range e.Arrive { // 朋友到达
			belong[id] = heap.Pop(&unoccupied).(int) // 记录占据该椅子的朋友编号
			if id == targetFriend {
				return belong[id]
			}
		}
	}
	return 0
}

// 1801. 积压订单中的订单总数

type Order struct {
	price  int
	amount int
}

// 卖单堆：小顶堆（价格最低的在堆顶）
type SellHeap []Order

func (h SellHeap) Len() int           { return len(h) }
func (h SellHeap) Less(i, j int) bool { return h[i].price < h[j].price }
func (h SellHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *SellHeap) Push(x any) {
	*h = append(*h, x.(Order))
}

func (h *SellHeap) Pop() any {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}

// 买单堆：大顶堆（价格最高的在堆顶）
type BuyHeap []Order

func (h BuyHeap) Len() int           { return len(h) }
func (h BuyHeap) Less(i, j int) bool { return h[i].price > h[j].price }
func (h BuyHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *BuyHeap) Push(x any) {
	*h = append(*h, x.(Order))
}

func (h *BuyHeap) Pop() any {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}

/*
给你一个二维整数数组 orders ，其中每个 orders[i] = [pricei, amounti, orderTypei] 表示有 amounti 笔类型为 orderTypei 、价格为 pricei 的订单。

订单类型 orderTypei 可以分为两种：

0 表示这是一批采购订单 BuyHeap
1 表示这是一批销售订单 sell
注意，orders[i] 表示一批共计 amounti 笔的独立订单，这些订单的价格和类型相同。对于所有有效的 i ，由 orders[i] 表示的所有订单提交时间均早于 orders[i+1] 表示的所有订单。

存在由未执行订单组成的 积压订单 。积压订单最初是空的。提交订单时，会发生以下情况：

如果该订单是一笔采购订单 BuyHeap ，则可以查看积压订单中价格 最低 的销售订单 sell 。如果该销售订单 sell 的价格 低于或等于 当前采购订单 BuyHeap 的价格，则匹配并执行这两笔订单，并将销售订单 sell 从积压订单中删除。否则，采购订单 BuyHeap 将会添加到积压订单中。
反之亦然，如果该订单是一笔销售订单 sell ，则可以查看积压订单中价格 最高 的采购订单 BuyHeap 。如果该采购订单 BuyHeap 的价格 高于或等于 当前销售订单 sell 的价格，则匹配并执行这两笔订单，并将采购订单 BuyHeap 从积压订单中删除。否则，销售订单 sell 将会添加到积压订单中。
输入所有订单后，返回积压订单中的 订单总数 。由于数字可能很大，所以需要返回对 109 + 7 取余的结果。
*/
func getNumberOfBacklogOrders(orders [][]int) int {
	sells := &SellHeap{}
	buys := &BuyHeap{}
	heap.Init(sells)
	heap.Init(buys)

	for _, order := range orders {
		price := order[0]
		amount := order[1]
		orderType := order[2]

		if orderType == 0 {
			// 买单：匹配价格最低的卖单
			for amount > 0 && sells.Len() > 0 {
				top := (*sells)[0]
				if top.price > price {
					break // 卖单价格太高，无法匹配
				}

				if top.amount <= amount {
					// 完全消耗卖单
					amount -= top.amount
					heap.Pop(sells)
				} else {
					// 部分消耗卖单
					(*sells)[0].amount -= amount
					amount = 0
				}
			}
			// 剩余买单加入积压
			if amount > 0 {
				heap.Push(buys, Order{price: price, amount: amount})
			}
		} else {
			// 卖单：匹配价格最高的买单
			for amount > 0 && buys.Len() > 0 {
				top := (*buys)[0]
				if top.price < price {
					break // 买单价格太低，无法匹配
				}

				if top.amount <= amount {
					// 完全消耗买单
					amount -= top.amount
					heap.Pop(buys)
				} else {
					// 部分消耗买单
					(*buys)[0].amount -= amount
					amount = 0
				}
			}
			// 剩余卖单加入积压
			if amount > 0 {
				heap.Push(sells, Order{price: price, amount: amount})
			}
		}
	}

	// 统计积压订单总数
	total := 0
	for _, order := range *sells {
		total = (total + order.amount) % mod
	}
	for _, order := range *buys {
		total = (total + order.amount) % mod
	}
	return total
}

/*
给你一个二维整数数组 intervals ，其中 intervals[i] = [lefti, righti] 表示 闭 区间 [lefti, righti] 。

你需要将 intervals 划分为一个或者多个区间 组 ，每个区间 只 属于一个组，且同一个组中任意两个区间 不相交 。

请你返回 最少 需要划分成多少个组。

如果两个区间覆盖的范围有重叠（即至少有一个公共数字），那么我们称这两个区间是 相交 的。比方说区间 [1, 5] 和 [5, 8] 相交。
*/
func minGroups(intervals [][]int) int {
	h := &hp{}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	for _, x := range intervals {
		// 当前区间的左端点小于最小的右端点，只能创建一个新的组
		if h.Len() == 0 || x[0] < h.IntSlice[0] {
			heap.Push(h, x[1])
		} else {
			// 可以接到当前组的后面，更新右端点，并修复堆，使得堆顶仍然是最小的右端点
			h.IntSlice[0] = x[1]
			heap.Fix(h, 0)
		}
	}
	return h.Len()
}

func minGroupsByMap(intervals [][]int) int {
	diff := make(map[int]int)

	for _, p := range intervals {
		diff[p[0]]++   // 起点 +1
		diff[p[1]+1]-- // 终点下一位 -1
	}

	// 收集key并排序
	keys := make([]int, 0, len(diff))
	for k := range diff {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	// 前缀和找最大覆盖次数
	cur, ans := 0, 0
	for _, k := range keys {
		cur += diff[k]
		if cur > ans {
			ans = cur
		}
	}
	return ans
}

//https://leetcode.cn/problems/choose-k-elements-with-maximum-sum/description/
// 选出和最大的k个元素

func findMaxSum(nums1 []int, nums2 []int, k int) []int64 {
	ans := make([]int64, len(nums1))
	type tuple struct {
		Index int
		Var   int
		Var2  int
	}
	tuples := make([]tuple, 0, len(nums1))
	for i, x := range nums1 {
		tuples = append(tuples, tuple{
			Index: i,
			Var:   x,
			Var2:  nums2[i],
		})
	}
	sort.Slice(tuples, func(i, j int) bool {
		return tuples[i].Var > tuples[j].Var
	})
	n := len(nums1)
	h := &hp{}
	s := 0
	for i := 0; i < n; {
		start := i
		// 找到所有相同的 nums1[i]，这些数的答案都是一样的
		x := tuples[start].Var
		for ; i < n && tuples[i].Var == x; i++ {
			ans[tuples[i].Index] = int64(s)
		}
		// 把这些相同的 nums1[i] 对应的 nums2[i] 入堆
		// 同时用一个最小堆维护 nums[i] 的前 k 大元素：
		for ; start < i; start++ {
			y := tuples[start].Var2
			s += y
			heap.Push(h, y)
			if h.Len() > k {
				s -= heap.Pop(h).(int)
			}
		}
	}
	return ans

}

/*
分组循环

给你一个下标从 0 开始的整数数组 nums 和一个整数 threshold 。
请你从 nums 的子数组中找出以下标 l 开头、下标 r 结尾 (0 <= l <= r < nums.length) 且满足以下条件的 最长子数组 ：
nums[l] % 2 == 0
对于范围 [l, r - 1] 内的所有下标 i ，nums[i] % 2 != nums[i + 1] % 2
对于范围 [l, r] 内的所有下标 i ，nums[i] <= threshold
以整数形式返回满足题目要求的最长子数组的长度。
*/

func longestAlternatingSubarray(nums []int, threshold int) int {
	res := 0
	n := len(nums)
	i := 0
	for i < n {
		if nums[i] > threshold || nums[i]%2 != 0 {
			i++
			continue
		}
		start := i // 记录这一组开始的位置
		i++
		for i < n && nums[i] <= threshold && nums[i]%2 != nums[i-1]%2 {
			i++
		}
		res = max(res, i-start)
	}
	return res
}
