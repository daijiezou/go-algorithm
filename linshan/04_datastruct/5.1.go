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
    x := m.IntSlice[n-1]  // ✅ 取最后一个元素
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
