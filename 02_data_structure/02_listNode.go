package _2_data_structure

// 把存在的重复元素全部去除
func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{Val: -1}
	p := dummy
	for head != nil {
		if head.Next != nil && head.Val == head.Next.Val {
			for head.Next != nil && head.Val == head.Next.Val {
				head = head.Next
			}
			// 把自己也去除掉
			head = head.Next
			if head == nil {
				p.Next = nil
			}

		} else {
			p.Next = head
			p = p.Next
			head = head.Next
		}
	}
	return dummy.Next
}

// 从未排序的链表中移除重复元素
func deleteDuplicatesUnsorted(head *ListNode) *ListNode {
	valCount := make(map[int]int)
	p := head
	for p != nil {
		if _, ok := valCount[p.Val]; !ok {
			valCount[p.Val] = 1
		} else {
			valCount[p.Val]++
		}
		p = p.Next
	}
	dummy := &ListNode{Val: -1}
	p = dummy
	for head != nil {
		for valCount[head.Val] > 1 {
			head = head.Next
		}
		p.Next = head
		p = p.Next
		head = head.Next
	}
	return dummy.Next
}

// https://leetcode.cn/problems/ugly-number-ii/description/
func nthUglyNumber(n int) int {
	p2, p3, p5 := 1, 1, 1
	product2, product3, product5 := 1, 1, 1
	// 可以理解为最终合并的有序链表（结果链表）
	ugly := make([]int, n+1)
	// 可以理解为结果链表上的指针
	for i := 1; i <= n; i++ {
		res := Mymin(product2, product3, product5)
		ugly[i] = res
		if res == product2 {
			product2 = ugly[p2] * 2
			p2++
		}
		if res == product3 {
			product3 = ugly[p3] * 3
			p3++
		}
		if res == product5 {
			product5 = ugly[p5] * 5
			p5++
		}
	}
	return ugly[n]
}

// 取三个数的最小值
func Mymin(i, j, k int) int {
	if i < j {
		if i < k {
			return i
		} else {
			return k
		}
	} else {
		if j < k {
			return j
		} else {
			return k
		}
	}
}

type kthSmallestPQ struct {
	ValList []matrixVal
	Length  int
}

type matrixVal struct {
	val int
	row int //数所在的行
	col int //数所在列
}

func newKthSmallestPQ() *kthSmallestPQ {
	return &kthSmallestPQ{
		ValList: make([]matrixVal, 1), // 一定要从1开始，用数组表示二叉树的需求
		Length:  0,
	}
}

func (k *kthSmallestPQ) insert(val matrixVal) {
	k.ValList = append(k.ValList, val)
	k.Length++

	// 上浮到正确的位置
	k.swim(k.Length)
}

func (k *kthSmallestPQ) pop() matrixVal {
	// 堆顶的元素就是最小的值
	minVal := k.ValList[1] // 注意堆顶的元素的的index为1，因为0被舍弃了

	//将其换到最后,并删除
	k.swap(1, k.Length)
	k.ValList = k.ValList[:k.Length]
	k.Length--
	// 将队首的元素下沉到正确位置
	k.sink(1)
	return minVal
}

// x表示在队列中的位置
// 下沉，将元素和自己的左右子节点比较，若比自己的左右节点小就swap
func (k *kthSmallestPQ) sink(x int) {

	// 优先比较左节点，因为左节点的索引比较小
	for left(x) <= k.Length {
		minVal := left(x)
		// 如果右节点比左节点小，更新最小值
		if right(x) <= k.Length && k.more(minVal, right(x)) {
			minVal = right(x)
		}
		// 左右节点比自己都大，退出循环
		if k.more(minVal, x) {
			break
		}
		k.swap(minVal, x)
		x = minVal
	}
}

// x表示在队列中的位置
func (k *kthSmallestPQ) swim(x int) {
	// x的父节点比自己大，将自己与父节点调换
	// 需要判断x>1,因为1就是已经是堆顶，它没有父节点了
	for x > 1 && k.more(parent(x), x) {
		// swap
		k.swap(x, parent(x))
		x = parent(x)
	}
}

func (k *kthSmallestPQ) swap(i, j int) {
	// swap
	temp := k.ValList[i]
	k.ValList[i] = k.ValList[j]
	k.ValList[j] = temp
}

// 判断i是否大于j
// i>j:true
// i<j false
func (k *kthSmallestPQ) more(i, j int) bool {
	if k.ValList[i].val > k.ValList[j].val {
		return true
	} else {
		return false
	}
}

// https://leetcode.cn/problems/kth-smallest-element-in-a-sorted-matrix/description/
// 有序矩阵中的第k小元素
func kthSmallest(matrix [][]int, k int) int {
	pq := newKthSmallestPQ()
	for i := 0; i < len(matrix); i++ {
		pq.insert(matrixVal{row: i, col: 0, val: matrix[i][0]})
	}
	var res int
	for k > 0 && pq.Length > 0 {
		cur := pq.pop()
		res = cur.val
		k--
		row, col := cur.row, cur.col+1
		// 将数组中下一个元素加入堆中
		if col < len(matrix[row]) {
			pq.insert(matrixVal{
				val: matrix[row][col],
				row: row,
				col: col,
			})
		}
	}
	return res
}

// 查找和最小的 K 对数字
// https://leetcode.cn/problems/find-k-pairs-with-smallest-sums/description/
func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	pq := newkSmallestPairsPQ()
	for i := 0; i < len(nums1); i++ {
		val := []int{nums1[i], nums2[0], 0}
		pq.insert(val)
	}
	res := make([][]int, 0, k)
	for k > 0 && pq.Length > 0 {
		cur := pq.pop()
		res = append(res, []int{cur[0], cur[1]})
		index := cur[2] + 1
		if index < len(nums2) {
			pq.insert([]int{cur[0], nums2[index], index})
		}
		k--
	}
	return res
}

type kSmallestPairsPQ struct {
	ValList [][]int
	Length  int
}

func newkSmallestPairsPQ() *kSmallestPairsPQ {
	return &kSmallestPairsPQ{
		ValList: make([][]int, 1), // 一定要从1开始，用数组表示二叉树的需求
		Length:  0,
	}
}

func (k *kSmallestPairsPQ) insert(val []int) {
	k.ValList = append(k.ValList, val)
	k.Length++

	// 上浮到正确的位置
	k.swim(k.Length)
}

func (k *kSmallestPairsPQ) pop() []int {
	// 堆顶的元素就是最小的值
	minVal := k.ValList[1] // 注意堆顶的元素的的index为1，因为0被舍弃了

	//将其换到最后,并删除
	k.swap(1, k.Length)
	k.ValList = k.ValList[:k.Length]
	k.Length--
	// 将队首的元素下沉到正确位置
	k.sink(1)
	return minVal
}

// x表示在队列中的位置
// 下沉，将元素和自己的左右子节点比较，若比自己的左右节点小就swap
func (k *kSmallestPairsPQ) sink(x int) {

	// 优先比较左节点，因为左节点的索引比较小
	for left(x) <= k.Length {
		minVal := left(x)
		// 如果右节点比左节点小，更新最小值
		if right(x) <= k.Length && k.more(minVal, right(x)) {
			minVal = right(x)
		}
		// 左右节点比自己都大，退出循环
		if k.more(minVal, x) {
			break
		}
		k.swap(minVal, x)
		x = minVal
	}
}

// x表示在队列中的位置
func (k *kSmallestPairsPQ) swim(x int) {
	// x的父节点比自己大，将自己与父节点调换
	// 需要判断x>1,因为1就是已经是堆顶，它没有父节点了
	for x > 1 && k.more(parent(x), x) {
		// swap
		k.swap(x, parent(x))
		x = parent(x)
	}
}

func (k *kSmallestPairsPQ) swap(i, j int) {
	// swap
	temp := k.ValList[i]
	k.ValList[i] = k.ValList[j]
	k.ValList[j] = temp
}

func (k *kSmallestPairsPQ) more(i, j int) bool {
	if k.ValList[i][0]+k.ValList[i][1] > k.ValList[j][0]+k.ValList[j][1] {
		return true
	} else {
		return false
	}
}
