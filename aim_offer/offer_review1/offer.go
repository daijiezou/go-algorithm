package offer_review1

import (
	"fmt"
	"strconv"
	"strings"
)

func duplicate(numbers []int) int {
	for i := 0; i < len(numbers); i++ {
		for numbers[i] != i {
			if numbers[i] == numbers[numbers[i]] {
				return numbers[i]
			}
			numbers[i], numbers[numbers[i]] = numbers[numbers[i]], numbers[i]
		}
	}
	return -1
}

// 02
/*
在一个二维数组array中（每个一维数组的长度相同）
每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。请
[
[1,2,8,9],
[2,4,9,12],
[4,7,10,13],
[6,8,11,15]
]
*/
func Find(target int, array [][]int) bool {
	m := len(array)
	n := len(array[0])
	i := 0
	j := n - 1
	for i < m && j >= 0 {
		x := array[i][j]
		if x == target {
			return true
		} else if x > target {
			j--
		} else {
			i++
		}
	}
	return false
}

// 5. 替换空格
func replaceSpace(s string) string {
	res := strings.Builder{}
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			res.WriteByte(s[i])
		} else {
			res.WriteString("%20")
		}
	}
	return res.String()
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 6.从尾到头打印
func printListFromTailToHead(head *ListNode) []int {
	res := []int{}
	var dfs func(head *ListNode)
	dfs = func(head *ListNode) {
		if head == nil {
			return
		}
		dfs(head.Next)
		res = append(res, head.Val)
	}
	dfs(head)
	return res
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 7.重建二叉树
func reConstructBinaryTree(preOrder []int, vinOrder []int) *TreeNode {
	if len(preOrder) == 0 {
		return nil
	}
	root := &TreeNode{
		Val: preOrder[0],
	}
	leftLength := 0
	for i := 0; i < len(vinOrder); i++ {
		if vinOrder[i] == root.Val {
			leftLength = i
			break
		}
	}
	root.Left = reConstructBinaryTree(preOrder[1:leftLength+1], vinOrder[:leftLength])
	root.Right = reConstructBinaryTree(preOrder[1+leftLength:], vinOrder[leftLength+1:])
	return root
}

type TreeLinkNode struct {
	Val   int
	Left  *TreeLinkNode
	Right *TreeLinkNode
	Next  *TreeLinkNode // 指向父节点
}

// 8.二叉树的下一个节点
/*
给定一个二叉树其中的一个结点，
请找出中序遍历顺序的下一个结点并且返回。注意，树中的结点不仅包含左右子结点，同时包含指向父结点的next指针。
*/
func GetNext(pNode *TreeLinkNode) *TreeLinkNode {
	if pNode == nil {
		return nil
	}
	if pNode.Right != nil {
		cur := pNode.Right
		for cur.Left != nil {
			cur = cur.Left
		}
		return cur
	}
	var next *TreeLinkNode
	if pNode.Next != nil {
		cur := pNode
		parent := pNode.Next
		// 如果当前节点是其父节点的左子树，则退出循环par则为当前节点的下一个节点
		// 否则继续找它的父节点
		for parent != nil && cur == parent.Right {
			cur = parent
			parent = cur.Next
		}
		next = parent
	}
	return next
}

// 10.斐波那契数列
func Fibonacci(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 1
	}
	fib1 := 1
	fib2 := 1
	fibN := 0
	for i := 3; i <= n; i++ {
		fibN = fib1 + fib2
		fib1 = fib2
		fib2 = fibN

	}
	return fibN
}

// 11.旋转排序数组
func minNumberInRotateArray(nums []int) int {
	n := len(nums)
	left := -1
	right := n
	for left+1 < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[n-1] {
			left = mid
		} else if nums[mid] < nums[n-1] {
			right = mid
		} else {
			right--
		}
	}
	return nums[right]
}

// 12.矩阵中的路径
func hasPath(matrix [][]byte, word string) bool {
	// write code here
	pos := [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}
	m := len(matrix)
	n := len(matrix[0])
	var bfs func(i, j int, index int, visited []bool) bool
	bfs = func(i, j int, index int, visited []bool) bool {
		if matrix[i][j] != word[index] {
			return false
		}
		if index == len(word)-1 {
			return true
		}
		for _, v := range pos {
			newi := i + v[0]
			newj := j + v[1]

			if newi >= 0 && newi < m && newj >= 0 && newj < n {
				if visited[newi*n+newj] {
					continue
				}
				visited[newi*n+newj] = true
				if bfs(newi, newj, index+1, visited) {
					return true
				}
				visited[newi*n+newj] = false
			}
		}
		return false
	}
	vis := make([]bool, m*n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if bfs(i, j, 0, vis) {
				return true
			}
		}
	}
	return false
}

// JZ13 机器人的运动范围
func movingCount(threshold int, rows int, cols int) int {
	pos := [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}
	vis := make([]bool, rows*cols)
	var dfs func(i, j int) int
	vis[0] = true
	dfs = func(i, j int) int {
		if getScore(i, j) > threshold {
			return 0
		}
		res := 0
		for _, v := range pos {
			newi := i + v[0]
			newj := j + v[1]

			if newi >= 0 && newi < rows && newj >= 0 && newj < cols {
				if vis[newi*cols+newj] {
					continue
				}
				vis[newi*cols+newj] = true
				res += dfs(newi, newj)
			}
		}
		return res + 1
	}

	return dfs(0, 0)
}

func getScore(i, j int) int {
	sum := 0
	for i > 0 {
		sum += i % 10
		i = i / 10
	}
	for j > 0 {
		sum += j % 10
		j = j / 10
	}
	return sum
}

// 14.减绳子
func cutRope(n int) int {
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 2
	dp[3] = 3 // 选择不切分
	for i := 4; i <= n; i++ {
		for j := 1; j <= i/2; j++ {
			dp[i] = max(dp[i], dp[j]*dp[i-j])
		}
	}
	return dp[n]
}

// 16 快速幂算法
func Power(base float64, exponent int) float64 {
	if exponent > 0 {
		return pow(base, exponent)
	} else {
		return 1 / pow(base, -exponent)
	}
}

func pow(x float64, n int) float64 {
	res := float64(1)
	for n > 0 {
		if n&1 == 1 {
			res = x * res
		}
		x *= x
		n >>= 1
	}
	return res
}

func printNumbers(n int) []int {
	end := 1
	for n > 0 {
		end *= 10
		n--
	}
	res := make([]int, 0, end-1)
	for i := 1; i < end; i++ {
		res = append(res, i)
	}
	return res
}

func deleteNode(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Next: head}
	cur := dummy
	for cur != nil {
		if cur.Next != nil && cur.Next.Val == val {
			cur.Next = cur.Next.Next
			break
		}
		cur = cur.Next
	}
	return dummy.Next
}

func match(str string, pattern string) bool {
	// write code here
	var dfs func(si, pi int) bool
	dfs = func(si, pi int) bool {
		// 全部匹配完成
		if si == len(str) && pi == len(pattern) {
			return true
		}
		// 没有匹配结束
		if len(pattern) == pi {
			return false
		}

		if pi+1 < len(pattern) && pattern[pi+1] == '*' {
			if si < len(str) && (str[si] == pattern[pi] || pattern[pi] == '.') {
				// 匹配多个或者直接跳过
				return dfs(si+1, pi) || dfs(si, pi+2)
			} else {
				// 不满足，匹配0个
				return dfs(si, pi+2)
			}
		}

		if si < len(str) && (str[si] == pattern[pi] || pattern[pi] == '.') {
			// 匹配一个字符
			return dfs(si+1, pi+1)
		}

		return false
	}
	return dfs(0, 0)
}

// todo
func isNumeric() {
	return
}

// 21.调整数组顺序使得奇数位于偶数前面
func reOrderArray(array []int) []int {
	// write code here
	n := len(array)
	for i := 0; i < n; i++ {
		if array[i]%2 == 1 {
			j := i
			// 向前冒泡，直到前面不是偶数为止
			for j > 0 && array[j-1]%2 == 0 {
				array[j], array[j-1] = array[j-1], array[j]
				j--
			}
		}
	}
	return array
}

// 22.链表中倒数第K个节点
func FindKthToTail(pHead *ListNode, k int) *ListNode {
	// write code here
	fast := pHead
	slow := pHead
	for k > 0 {
		if fast == nil {
			return nil
		}
		fast = fast.Next
		k--
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}

func EntryNodeOfLoop(pHead *ListNode) *ListNode {
	fast := pHead
	slow := pHead
	hasLoop := false
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			hasLoop = true
			break
		}
	}
	if !hasLoop {
		return nil
	}
	meetingNode := fast
	nodesInLoop := 1
	for fast.Next != meetingNode {
		fast = fast.Next
		nodesInLoop++
	}
	fast = pHead
	slow = pHead
	for nodesInLoop > 0 {
		fast = fast.Next
	}
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}

/*
设：

mu = 从头节点到环入口的距离
lambda = 环的长度
t = 快慢指针第一次相遇时慢指针走的步数
第一次相遇时有两件事实：

慢指针走了 t 步
快指针走了 2t 步
由于快指针比慢指针多走了整圈数，故 2t − t = k·lambda，即 t ≡ 0 (mod lambda) 相对于慢指针在环内的偏移量
把路径分成“进环前”和“环内”两段：

慢指针在相遇时的位置可以写成：t = mu + x，其中 x 是它进入环后又走的步数（0 ≤ x < lambda）
由 2t − t = k·lambda 得：

t = k·lambda
因此 mu + x = k·lambda
推出 mu = k·lambda − x
这条等式的含义：从“相遇点”再往前走 (lambda − x) 步就到达环入口；而从链表头走 mu 步也到达环入口。并且因为 mu = k·lambda − x，与环长同余，意味着：

如果一个指针从头结点出发走 mu 步，另一个指针从相遇点出发每次走 1 步，它们的相遇点就是环的入口
*/

func EntryNodeOfLoop2(pHead *ListNode) *ListNode {
	fast := pHead
	slow := pHead
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			fast = pHead
			for fast != slow {
				fast = fast.Next
				slow = slow.Next
			}
			return fast
		}
	}
	return nil
}

// 24 反转链表
func ReverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func Merge(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	if pHead1 == nil {
		return pHead2
	}
	if pHead2 == nil {
		return pHead1
	}

	if pHead1.Val < pHead2.Val {
		head := &ListNode{
			Val: pHead1.Val,
		}
		head.Next = Merge(pHead1.Next, pHead2)
		return head
	} else {
		head := &ListNode{
			Val: pHead2.Val,
		}
		head.Next = Merge(pHead2.Next, pHead1)
		return head
	}

}

// 26.树的子结构
func HasSubtree(pRoot1 *TreeNode, pRoot2 *TreeNode) bool {
	res := false
	if pRoot1 != nil && pRoot2 != nil {
		if pRoot1.Val == pRoot2.Val {
			res = Tree1HasTree2(pRoot1.Left, pRoot2.Left) && Tree1HasTree2(pRoot1.Right, pRoot2.Right)
		}
		if !res {
			res = HasSubtree(pRoot1.Left, pRoot2) || HasSubtree(pRoot1.Right, pRoot2)
		}
	}
	return res
}

func Tree1HasTree2(pRoot1 *TreeNode, pRoot2 *TreeNode) bool {
	if pRoot2 == nil {
		return true
	}
	if pRoot1 == nil {
		return false
	}
	if pRoot1.Val != pRoot2.Val {
		return false
	}
	return Tree1HasTree2(pRoot1.Left, pRoot2.Left) && Tree1HasTree2(pRoot1.Right, pRoot2.Right)
}

// 27.二叉树镜像
func Mirror(pRoot *TreeNode) *TreeNode {
	// write code here
	if pRoot == nil {
		return nil
	}
	left := pRoot.Left
	right := pRoot.Right

	pRoot.Right = Mirror(left)
	pRoot.Left = Mirror(right)

	return pRoot
}

// 28.判断一棵树是否是对称的
func isSymmetrical(pRoot *TreeNode) bool {
	if pRoot == nil {
		return true
	}
	return IsSame(pRoot.Left, pRoot.Right)
}

func IsSame(p1, p2 *TreeNode) bool {
	if p1 == nil && p2 == nil {
		return true
	}
	if p1 == nil {
		return false
	}
	if p2 == nil {
		return false
	}
	if p1.Val != p2.Val {
		return false
	}
	return IsSame(p1.Left, p2.Right) && IsSame(p1.Right, p2.Left)
}

// 29.顺时针打印矩阵
func printMatrix(matrix [][]int) []int {
	m := len(matrix)
	n := len(matrix[0])
	res := make([]int, 0)
	upper := 0
	lower := m - 1
	left := 0
	right := n - 1
	hang := true
	lie := false
	hangOrder := true
	lieOrder := true
	for upper <= lower && left <= right {
		if hang {
			hang = !hang
			lie = true
			if hangOrder {
				for i := left; i <= right; i++ {
					res = append(res, matrix[upper][i])
				}
				hangOrder = !hangOrder
				upper++
			} else {
				for i := right; i >= left; i-- {
					res = append(res, matrix[lower][i])
				}
				hangOrder = !hangOrder
				lower--
			}
			continue
		} else {
			lie = !lie
			hang = true
			if lieOrder {
				for i := upper; i <= lower; i++ {
					res = append(res, matrix[i][right])
				}
				lieOrder = !lieOrder
				right--
			} else {
				for i := lower; i >= upper; i-- {
					res = append(res, matrix[i][left])
				}
				lieOrder = !lieOrder
				left++
			}
		}
	}
	return res
}

func printMatrix2(matrix [][]int) []int {
	m := len(matrix)
	n := len(matrix[0])
	res := make([]int, 0)
	upper := 0
	lower := m - 1
	left := 0
	right := n - 1
	for upper <= lower && left <= right {
		for i := left; i <= right; i++ {
			res = append(res, matrix[upper][i])
		}
		upper++
		for i := upper; i <= lower; i++ {
			res = append(res, matrix[i][right])
		}
		right--
		if upper <= lower {
			for i := right; i >= left; i-- {
				res = append(res, matrix[lower][i])
			}
		}
		lower--

		if left <= right {
			for i := lower; i >= upper; i-- {
				res = append(res, matrix[i][left])
			}
		}
		left++
	}
	return res
}

// 31.栈的压入、弹出序列
func IsPopOrder(pushV []int, popV []int) bool {
	cur := make([]int, 0)
	curIndex := 0
	for i := 0; i < len(pushV); i++ {
		x := pushV[i]
		if x == popV[curIndex] {
			curIndex++
		} else {
			cur = append(cur, x)
		}
		for len(cur) > 0 && cur[len(cur)-1] == popV[curIndex] {
			curIndex++
			cur = cur[:len(cur)-1]
		}
	}
	return curIndex == len(popV)
}

// 32.从上往下打印二叉树
func PrintFromTopToBottom(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		res = append(res, cur.Val)
		if cur.Left != nil {
			q = append(q, cur.Left)
		}
		if cur.Right != nil {
			q = append(q, cur.Right)
		}
	}
	return res
}

// 33.判断数组是不是二叉搜索树的后续遍历
func VerifySquenceOfBST(sequence []int) bool {
	if len(sequence) == 0 {
		return false
	}
	return checkSubSquenceOfBST(sequence)
}

func checkSubSquenceOfBST(sequence []int) bool {
	if len(sequence) <= 1 {
		return true
	}
	n := len(sequence)
	root := sequence[n-1]
	left := make([]int, 0)
	right := make([]int, 0)
	index := 0
	for i := 0; i < n-1; i++ {
		if sequence[i] < root {
			left = append(left, sequence[i])
			index = i + 1
		} else {
			index = i
			break
		}
	}
	right = append(right, sequence[index:n-1]...)
	for i := 0; i < len(right); i++ {
		if right[i] < root {
			return false
		}
	}
	fmt.Println(left, right)
	return checkSubSquenceOfBST(left) && checkSubSquenceOfBST(right)
}

// 34.二叉树中和为某一值的路径
func FindPath(root *TreeNode, target int) [][]int {
	res := make([][]int, 0)
	var dfs = func(root *TreeNode) {}
	cur := 0
	curPath := make([]int, 0)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		cur += root.Val
		curPath = append(curPath, root.Val)
		defer func() {
			cur -= root.Val
			curPath = curPath[:len(curPath)-1]
		}()
		if root.Left == nil && root.Right == nil && cur == target {
			temp := make([]int, len(curPath))
			copy(temp, curPath)
			res = append(res, temp)
			return
		}

		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)
	return res
}

type RandomListNode struct {
	Label  int
	Next   *RandomListNode
	Random *RandomListNode
}

// 35.复杂链表的复制
func Clone(head *RandomListNode) *RandomListNode {
	//write your code here
	m := make(map[*RandomListNode]*RandomListNode)
	cur := head
	for cur != nil {
		m[cur] = &RandomListNode{Label: cur.Label}
		cur = cur.Next
	}
	cur = head
	for cur != nil {
		if cur.Next != nil {
			m[cur].Next = m[cur.Next]
		}
		if cur.Random != nil {
			m[cur].Random = m[cur.Random]
		}
		cur = cur.Next
	}
	return m[head]
}

// 36.将二叉搜索树转为双向链表
func Convert(pRootOfTree *TreeNode) *TreeNode {
	// write code here
	var head *TreeNode
	var dfs func(cur *TreeNode)
	dfs = func(cur *TreeNode) {
		if cur == nil {
			return
		}
		dfs(cur.Right)
		cur.Right = head
		if head != nil {
			head.Left = cur
		}
		head = cur
		dfs(cur.Left)
	}
	dfs(pRootOfTree)
	return head
}

// 37.二叉树的序列化与反序列化
func Serialize(root *TreeNode) string {
	list := []string{}
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			list = append(list, "#")
			return
		}
		val := strconv.Itoa(root.Val)
		list = append(list, val)
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)
	return strings.Join(list, ",")
}

func Deserialize(s string) *TreeNode {
	var dfs func() *TreeNode
	i := 0
	list := strings.Split(s, ",")
	dfs = func() *TreeNode {
		x := list[i]
		i++
		if x == "#" {
			return nil
		}
		val, _ := strconv.Atoi(x)
		root := &TreeNode{Val: val}
		root.Left = dfs()
		root.Right = dfs()
		return root
	}
	return dfs()
}
