package mypq

type Pq struct {
	nums []int
	size int
}

func NewPQ() *Pq {
	return &Pq{
		// 索引0不使用
		nums: make([]int, 1),
		size: 0,
	}
}

func (q *Pq) insert(x int) {
	q.nums = append(q.nums, x)
	q.size++
	q.swim(q.size)
}

func (q *Pq) Pop() int {
	minVal := q.nums[1]

	// 将最后一个元素放到头部
	q.nums[1], q.nums[q.size] = q.nums[q.size], q.nums[1]
	// 移除最后一个元素
	q.nums = q.nums[:q.size]
	q.size--
	q.Sink(1)
	return minVal
}

func (q *Pq) Sink(x int) {
	for left(x) <= q.size {
		minIndex := left(x)
		if right(x) <= q.size && q.nums[minIndex] > q.nums[right(x)] {
			minIndex = right(x)
		}
		if q.nums[minIndex] > q.nums[x] {
			break
		}
		q.nums[x], q.nums[minIndex] = q.nums[minIndex], q.nums[x]
		x = minIndex
	}
}

func (q *Pq) swim(x int) {
	for x > 1 && q.nums[parent(x)] > q.nums[x] {
		q.nums[x], q.nums[parent(x)] = q.nums[parent(x)], q.nums[x]
		x = parent(x)
	}
}

func parent(x int) int {
	return x / 2
}

func left(x int) int {
	return x * 2
}

func right(x int) int {
	return x*2 + 1
}
