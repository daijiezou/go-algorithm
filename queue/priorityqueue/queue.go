package priorityqueue

type MaxPQ struct {
	pq   []int
	size int
}

// NewMinPQ 初始化一个小顶堆
func NewMinPQ(cap int) *MaxPQ {
	// 索引 0 不用，所以多分配一个空间
	pq := make([]int, cap+1)
	return &MaxPQ{pq: pq}
}

/* 插入元素 e */
func (mpq *MaxPQ) insert(e int) {
	mpq.size++
	// 先把元素加到最后
	mpq.pq = append(mpq.pq, e)
	//上浮到正确的位置
	mpq.swim(mpq.size)
}

/* 删除并返回当前队列中最小元素 */
func (mpq *MaxPQ) pop() int {
	// 堆顶就是最小元素
	minVal := mpq.pq[1]

	// 将其换到最后并删除之
	mpq.swap(1, mpq.size)
	mpq.pq = mpq.pq[:mpq.size]
	mpq.size--

	// 将堆顶元素下沉到正确的位置
	mpq.sink(1)
	return minVal
}

/* 上浮第 x 个元素，以维护最小堆性质 */
func (mpq *MaxPQ) swim(x int) {
	// 查看是否比自己的父节点小
	// 如果比自己的父节点小就与父节点交换位置
	for x > 1 && mpq.more(mpq.parent(x), x) {
		mpq.swap(mpq.parent(x), x)
		x = mpq.parent(x)
	}
}

/* 下沉第 x 个元素，以维护最小堆性质 */
func (mpq *MaxPQ) sink(x int) {
	for mpq.left(x) <= mpq.size {
		min := mpq.left(x)
		// 如果tempMin比左孩子更大，则重置一下min的值
		if mpq.right(x) <= mpq.size && mpq.more(min, mpq.right(x)) {
			min = mpq.right(x)
		}
		// 节点x比两个孩子都小，不必下沉了
		if mpq.more(min, x) {
			break
		}
		mpq.swap(x, min)
		x = min
	}
}

/* 交换数组的两个元素 */
func (mpq *MaxPQ) swap(i, j int) {
	temp := mpq.pq[i]
	mpq.pq[i] = mpq.pq[j]
	mpq.pq[j] = temp
}

/* pq[i] 是否比 pq[j] 大？ */
func (mpq *MaxPQ) more(i, j int) bool {
	return mpq.pq[i] > mpq.pq[j]
}

// 父节点的索引
func (mpq *MaxPQ) parent(root int) int {
	return root / 2
}

// 左孩子的索引
func (mpq *MaxPQ) left(root int) int {
	return root * 2
}

// 右孩子的索引
func (mpq *MaxPQ) right(root int) int {
	return root*2 + 1
}
