package priorityqueue

type MaxPQ struct {
	pq   []int
	size int
}

// NewMinPQ 初始化一个小顶堆
// 堆顶是最小的元素
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
	// 查看父节点是否大于自己，如果大于则交换
	for x > 1 && (mpq.pq[mpq.parent(x)] > mpq.pq[x]) {
		mpq.swap(mpq.parent(x), x)
		x = mpq.parent(x)
	}
}

/* 下沉第 x 个元素，以维护最小堆性质 */
func (mpq *MaxPQ) sink(x int) {
	for mpq.left(x) <= mpq.size {
		minIndex := mpq.left(x)
		if mpq.right(x) <= mpq.size && mpq.pq[mpq.right(x)] < mpq.pq[minIndex] {
			minIndex = mpq.right(x)
		}
		if mpq.pq[x] <= mpq.pq[minIndex] {
			break
		}
		mpq.swap(x, minIndex)
		x = minIndex
	}
}

/* 交换数组的两个元素 */
func (mpq *MaxPQ) swap(i, j int) {
	temp := mpq.pq[i]
	mpq.pq[i] = mpq.pq[j]
	mpq.pq[j] = temp
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
