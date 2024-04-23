package _2_data_structure

type MinPQ struct {
	pq   []*ListNode
	size int
}

func NewMinPQ() *MinPQ {
	return &MinPQ{
		pq:   make([]*ListNode, 1),
		size: 0,
	}
}

/* 返回当前队列中最小元素 */
func (mpq *MinPQ) min() *ListNode {
	return mpq.pq[1]
}

/* 插入元素 e */
func (mpq *MinPQ) insert(e *ListNode) {
	mpq.size++
	// 先把元素加到最后
	mpq.pq = append(mpq.pq, e)
	//上浮到正确的位置
	mpq.swim(mpq.size)
}

/* 删除并返回当前队列中最小元素 */
func (mpq *MinPQ) pop() *ListNode {
	// 堆顶就是最小元素
	min := mpq.pq[1]
	// 将其换到最后并删除之
	mpq.swap(1, mpq.size)
	mpq.pq = mpq.pq[:mpq.size]
	mpq.size--
	// 将堆顶元素下沉到正确的位置
	mpq.sink(1)
	return min
}

/* 上浮第 x 个元素，以维护最小堆性质 */
func (mpq *MinPQ) swim(x int) {
	// 查看是否比自己的父节点小
	// 如果比自己的父节点小就与父节点交换位置
	for x > 1 && mpq.more(parent(x), x) {
		mpq.swap(parent(x), x)
		x = parent(x)
	}
}

/* 下沉第 x 个元素，以维护最小堆性质 */
func (mpq *MinPQ) sink(x int) {
	// 这里先比较左孩子节点，因为左孩子节点的index比较小
	for left(x) <= mpq.size {
		tempMin := left(x)

		// 如果tempMin比右孩子更大，则重置一下min的值
		if right(x) <= mpq.size && mpq.more(tempMin, right(x)) {
			tempMin = right(x)
		}

		// 节点x比两个孩子都小，不必下沉了
		if mpq.more(tempMin, x) {
			break
		}

		mpq.swap(x, tempMin)
		x = tempMin
	}
}

/* 交换数组的两个元素 */
func (mpq *MinPQ) swap(i, j int) {
	temp := mpq.pq[i]
	mpq.pq[i] = mpq.pq[j]
	mpq.pq[j] = temp
}

/* pq[i] 是否比 pq[j] 大？ */
func (mpq *MinPQ) more(i, j int) bool {
	return mpq.pq[i].Val > mpq.pq[j].Val
}

// 父节点的索引
func parent(root int) int {
	return root / 2
}

// 左孩子的索引
func left(root int) int {
	return root * 2
}

// 右孩子的索引
func right(root int) int {
	return root*2 + 1
}
