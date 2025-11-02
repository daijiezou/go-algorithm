package offer_review1

import (
	"container/heap"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

// 01.数组中重复的数字
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
			// 说明这一列都大于目标数字
			j--
		} else {
			// 说明这一行都小于目标数字
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

// 11.搜索旋转排序数组中最小的数字
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
				// 做选择
				visited[newi*n+newj] = true
				if bfs(newi, newj, index+1, visited) {
					return true
				}
				// 回溯,撤销选择
				visited[newi*n+newj] = false
			}
		}
		return false
	}
	// 用一维数组来存放二维数组的访问状态
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

// 删除指定的节点
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
				// 匹配1个或者直接跳过
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
// 解法一：类似插入排序，时间 O(n^2)，空间 O(1)，稳定
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

// 解法二：使用辅助数组，时间 O(n)，空间 O(n)，稳定
func reOrderArray2(array []int) []int {
	if len(array) == 0 {
		return []int{}
	}
	res := make([]int, 0, len(array))
	// 第一次遍历，放入所有奇数
	for _, num := range array {
		if num%2 == 1 {
			res = append(res, num)
		}
	}
	// 第二次遍历，放入所有偶数
	for _, num := range array {
		if num%2 == 0 {
			res = append(res, num)
		}
	}
	return res
}

// 解法三：双指针，时间 O(n)，空间 O(1)，不稳定
func reOrderArray3(array []int) []int {
	left, right := 0, len(array)-1
	for left < right {
		// 从左向右找偶数
		for left < right && array[left]%2 == 1 {
			left++
		}
		// 从右向左找奇数
		for left < right && array[right]%2 == 0 {
			right--
		}
		// 交换
		if left < right {
			array[left], array[right] = array[right], array[left]
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

// 25.合并两个排序的链表
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

		// 避免重复打印
		if upper <= lower {
			for i := right; i >= left; i-- {
				res = append(res, matrix[lower][i])
			}
		}
		lower--

		// 避免重复打印
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
	// 判断右子树是不是都小于根节点
	right = append(right, sequence[index:n-1]...)
	for i := 0; i < len(right); i++ {
		if right[i] < root {
			return false
		}
	}
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

// 38.字符串的排列
func Permutation(str string) []string {
	res := []string{}
	n := len(str)
	// 将字符串转换为字符数组并排序
	chars := []byte(str)
	sort.Slice(chars, func(i, j int) bool {
		return chars[i] < chars[j]
	})
	vis := make([]bool, n)
	var backtrack func(cur []byte)
	backtrack = func(cur []byte) {
		if len(cur) == n {
			res = append(res, string(cur))
			return
		}
		for i := 0; i < n; i++ {
			if vis[i] {
				continue
			}
			// 在使用第二个同样的字符时，必须要第一个也使用了，保证顺序
			if i > 0 && chars[i] == chars[i-1] && !vis[i-1] {
				continue
			}
			vis[i] = true
			cur = append(cur, chars[i])
			backtrack(cur)
			vis[i] = false
			cur = cur[:len(cur)-1]
		}
	}
	backtrack([]byte{})

	return res
}

// 39.数组中出现次数超过一半的数组
func MoreThanHalfNum_Solution(numbers []int) int {
	if len(numbers) <= 1 {
		return -1
	}
	// write code here
	candidate := numbers[0]
	cnt := 1
	for i := 1; i < len(numbers); i++ {
		if numbers[i] != candidate {
			cnt--
			// 换一个候选人
			if cnt <= 0 {
				cnt = 1
				candidate = numbers[i]
			}
		} else {
			cnt++
		}
	}
	return candidate
}

// 40.最小的K个数
func GetLeastNumbers_Solution(input []int, k int) []int {
	if k <= 0 || k > len(input) {
		return []int{}
	}
	left := 0
	right := len(input) - 1
	index := part(input, left, right)
	for index != k-1 {
		if index > k-1 {
			right = index - 1
			index = part(input, left, right)
		} else {
			left = index + 1
			index = part(input, left, right)
		}
	}
	return input[:k]
}

func part(nums []int, left, right int) int {
	flag := nums[right]
	index := left
	for i := left; i < right; i++ {
		if nums[i] < flag {
			nums[index], nums[i] = nums[i], nums[index]
			index++
		}
	}
	nums[index], nums[right] = nums[right], nums[index]
	return index
}

// MaxHeap 大顶堆实现
type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] } // 大顶堆：父节点大于子节点
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[0 : n-1]
	return x
}

// Top 返回堆顶元素（不删除）
func (h MaxHeap) Top() int {
	if len(h) == 0 {
		return 0
	}
	return h[0]
}

// MinHeap 小顶堆实现
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] } // 小顶堆：父节点小于子节点
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[0 : n-1]
	return x
}

func (h MinHeap) Top() int {
	if len(h) == 0 {
		return 0
	}
	return h[0]
}

var maxh = &MaxHeap{}
var minh = &MinHeap{}

func init() {
	heap.Init(maxh)
	heap.Init(minh)
}

// 41.数据流中的中位数
func Insert(num int) {
	// 默认往大顶堆塞
	if (minh.Len()+maxh.Len())%2 == 0 {
		if minh.Len() > 0 && num > minh.Top() {
			heap.Push(minh, num)
			x := heap.Pop(minh)
			heap.Push(maxh, x)
		} else {
			heap.Push(maxh, num)
		}
	} else {
		if maxh.Len() > 0 && num < maxh.Top() {
			x := heap.Pop(maxh)
			heap.Push(minh, x)
			heap.Push(maxh, num)
		} else {
			heap.Push(minh, num)
		}
	}
}

func GetMedian() float64 {
	if (minh.Len()+maxh.Len())%2 == 0 {
		return (float64(minh.Top()) + float64(maxh.Top())) / 2
	} else {
		// 因为默认往大顶堆去塞，所以如果最后的总数是单数的话，中位数就会出现在大顶堆的首位
		return float64(maxh.Top())
	}
}

// 42.连续子数组的最大和
func FindGreatestSumOfSubArray(array []int) int {
	if len(array) <= 0 {
		return -1
	}
	sum := 0
	res := math.MinInt
	for i := 0; i < len(array); i++ {
		sum += array[i]
		res = max(res, sum)
		if sum < 0 {
			sum = 0
		}
	}
	return res
}

// 43.整数中1出现的次数,比如 n=12,那么1出现了5次,11算两次
func NumberOf1Between1AndN_Solution(n int) int {
	// 边界：n<=0 时没有正数，直接返回 0
	if n <= 0 {
		return 0
	}
	count := 0
	// factor 表示正在统计的位权：
	// 1 表示个位，10 表示十位，100 表示百位，依此类推
	// 思路：对每一位单独计算“这一位上出现数字1”的次数，然后把各位的次数相加
	for factor := 1; factor <= n; factor *= 10 {
		// 把 n 按当前位拆成三段： high | cur | low
		// 以 factor = 10（十位）为例：
		//   n = high * (factor*10) + cur * factor + low
		high := n / (factor * 10) // 当前位左侧的高位数值
		cur := (n / factor) % 10  // 当前位的数字（0..9）
		low := n % factor         // 当前位右侧的低位数值

		// 针对“当前位”为 1 的出现次数，有三种情况：
		if cur == 0 {
			// 情况1：当前位小于1（即为 0）
			// 高位可完整循环 high 次，每次这一位会有 factor 个数取到“1”
			// 例如：统计十位为1，有 0..9 的低位变化共 factor 种
			// 故贡献：high * factor
			count += high * factor
		} else if cur == 1 {
			// 情况2：当前位等于1
			// 除了 high 次完整循环外，最后一次“未完整循环”也会落在当前位为1，
			// 这一次能贡献 (low + 1) 个（从低位 0 计到 low）
			// 故贡献：high*factor + (low + 1)
			count += high*factor + low + 1
		} else {
			// 情况3：当前位大于1（2..9）
			// 表示已经“跨过了1”这个段，所以等价于完整循环了 (high+1) 次
			// 故贡献：(high + 1) * factor
			count += (high + 1) * factor
		}
	}
	return count
}

// 44.数字序列中某一位的数字
// 44. 数字序列中某一位的数字
func findNthDigit(n int) int {
	if n < 10 { // 0..9 直接返回
		return n
	}

	// 跳过一位数（0..9 共 10 个字符）
	n -= 10

	// 依次尝试 2 位数段、3 位数段、...
	digitLen := 2
	countInBlock := 90 // 2 位数有 90 个：10..99
	for n > digitLen*countInBlock {
		n -= digitLen * countInBlock
		digitLen++
		countInBlock *= 10 // 下一段数量 ×10：90, 900, 9000...
	}

	// 该段起始数字，如 2位段=10，3位段=100
	firstOfBlock := pow10(digitLen - 1)

	// 在该段内第 indexInBlock 个字符（从 0 开始）
	indexInBlock := n
	// 落到的具体第几个数字（从 0 开始）
	numberIndex := indexInBlock / digitLen
	// 在该数字内的第几位（从 0 开始，左→右）
	offsetInNumber := indexInBlock % digitLen

	// 得到目标数字
	target := firstOfBlock + numberIndex

	// 取出第 offsetInNumber 位
	s := strconv.Itoa(target)
	return int(s[offsetInNumber] - '0')
}

// 简单的 10 的幂（非负）
func pow10(k int) int {
	res := 1
	for k > 0 {
		res *= 10
		k--
	}
	return res
}

// 45.把数组排成最小的数
func PrintMinNumber(numbers []int) string {
	// write code here
	strs := make([]string, len(numbers))
	for i, num := range numbers {
		strs[i] = strconv.Itoa(num)
	}
	sort.Slice(strs, func(i, j int) bool {
		return strs[i]+strs[j] < strs[j]+strs[i]
	})
	return strings.Join(strs, "")
}

// 46.把数字翻译成字符串
func solve(nums string) int {
	// write code here
	n := len(nums)
	dp := make([]int, n)
	if nums[n-1] != '0' {
		dp[n-1] = 1
	}
	for i := n - 2; i >= 0; i-- {
		if nums[i] != '0' {
			dp[i] = dp[i+1]
		}
		// 处理两个数字
		if nums[i] == '1' || (nums[i] == '2' && nums[i+1] <= '6') {
			if i == n-2 {
				dp[i] += 1
			} else {
				dp[i] += dp[i+2]
			}
		}
	}
	return dp[0]
}

// 47. 最大礼物价值
func maxValue(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	dp[1][1] = grid[0][0]
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			dp[i][j] = max(dp[i-1][j], dp[i][j-1]) + grid[i-1][j-1]
		}
	}
	return dp[m][n]
}

// 48.最长不含重复字符的子字符串
func lengthOfLongestSubstring(s string) int {

	cnts := make(map[uint8]int)
	left := 0
	res := 1
	for i := 0; i < len(s); i++ {
		x := s[i]
		cnts[x]++
		for cnts[x] > 1 {
			leave := s[left]
			left++
			cnts[leave]--
		}
		res = max(res, i-left+1)
	}
	return res
}

// 49.丑数
func GetUglyNumber_Solution(index int) int {
	// write code here
	ugly2 := 1
	p2 := 0
	ugly3 := 1
	p3 := 0
	ugly5 := 1
	p5 := 0
	uglys := make([]int, index)
	for i := 0; i < index; i++ {
		curUgly := min(ugly2, ugly3, ugly5)
		uglys[i] = curUgly

		if curUgly >= ugly2 {
			ugly2 = uglys[p2] * 2
			p2++
		}

		if curUgly >= ugly3 {
			ugly3 = uglys[p3] * 3
			p3++
		}

		if curUgly >= ugly5 {
			ugly5 = uglys[p5] * 5
			p5++
		}
	}
	return uglys[index-1]
}

// 50.数组中第一个出现一次的字符
func FirstNotRepeatingChar(str string) int {
	// write code here
	maps := make(map[int32][]int)
	for i, x := range str {
		maps[x] = append(maps[x], i)
	}
	minIndex := len(str)
	for _, x := range maps {
		if len(x) == 1 {
			minIndex = min(minIndex, x[0])
		}
	}
	if minIndex == len(str) {
		return -1
	}
	return minIndex
}

// 数组中的逆序对
func InversePairs(nums []int) int {
	// 归并排序计数，时间 O(n log n)，空间 O(n)
	const mod = 1000000007
	n := len(nums)
	if n < 2 {
		return 0
	}
	tmp := make([]int, n)

	var mergeSort func(l, r int) int
	mergeSort = func(l, r int) int {
		if l >= r {
			return 0
		}
		m := l + (r-l)/2
		left := mergeSort(l, m) % mod
		right := mergeSort(m+1, r) % mod

		// 合并并统计跨区间逆序对
		i, j, k := l, m+1, l
		cnt := 0
		for i <= m && j <= r {
			if nums[i] <= nums[j] {
				tmp[k] = nums[i]
				// 当左边元素放入时，右边已放入的元素个数为 (j - (m+1))，
				// 它们都小于等于当前左元素，不构成新的逆序对
				k++
				i++
			} else {
				// nums[i] > nums[j]，形成 (m - i + 1) 个逆序对
				tmp[k] = nums[j]
				cnt += (m - i + 1)
				if cnt >= mod { // 减少大数累积
					cnt %= mod
				}
				k++
				j++
			}
		}
		for i <= m {
			tmp[k] = nums[i]
			k++
			i++
		}
		for j <= r {
			tmp[k] = nums[j]
			k++
			j++
		}
		// 拷回原数组
		for p := l; p <= r; p++ {
			nums[p] = tmp[p]
		}
		total := left + right
		total %= mod
		total += cnt
		total %= mod
		return total
	}

	return mergeSort(0, n-1) % mod
}

// 52.两个链表的第一个公共节点
func FindFirstCommonNode(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	if pHead1 == nil || pHead2 == nil {
		return nil
	}
	// write code here
	a := pHead1
	b := pHead2
	for pHead1 != pHead2 {
		if pHead1 == nil {
			pHead1 = b
		} else {
			pHead1 = pHead1.Next
		}

		if pHead2 == nil {
			pHead2 = a
		} else {
			pHead2 = pHead2.Next
		}
	}
	return pHead1
}

// 数字在升序数组中出现的次数
func GetNumberOfK(nums []int, k int) int {
	first := lowerBound(nums, k)
	if first == len(nums) {
		return 0
	}
	last := lowerBound(nums, k+1) - 1
	return last - first + 1
}

func lowerBound(nums []int, k int) int {
	left := -1
	right := len(nums)
	for left+1 < right {
		mid := left + (right-left)/2
		if nums[mid] >= k {
			right = mid
		} else {
			left = mid
		}
	}
	return right
}

// 54. 二叉搜索树的第K小的值
func KthNode(proot *TreeNode, k int) int {
	// write code here
	var dfs func(root *TreeNode)
	res := -1
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		dfs(root.Left)
		k--
		if k == 0 {
			res = root.Val
			return
		}
		dfs(root.Right)
	}
	dfs(proot)
	return res
}

// 55.二叉树的深度
func TreeDepth(pRoot *TreeNode) int {
	if pRoot == nil {
		return 0
	}
	leftDepth := TreeDepth(pRoot.Left) + 1
	rightDepth := TreeDepth(pRoot.Right) + 1
	return max(leftDepth, rightDepth)
}

// 56.数组中只出现一次的两个数字
func FindNumsAppearOnce(nums []int) []int {
	res := 0
	for _, x := range nums {
		res ^= x
	}
	mask := res & -res

	a, b := 0, 0
	for _, x := range nums {
		if x&mask == 0 {
			a ^= x
		} else {
			b ^= x
		}
	}

	if a < b {
		return []int{a, b}
	} else {
		return []int{b, a}
	}

}

// 57.和为s的两个数字
func FindNumbersWithSum(array []int, sum int) []int {
	// write code here
	target := map[int]struct{}{}
	for _, x := range array {
		if _, ok := target[sum-x]; ok {
			return []int{x, sum - x}
		} else {
			target[x] = struct{}{}
		}
	}
	return []int{}
}

// 58.左旋字符串
func LeftRotateString(str string, n int) string {
	if str == "" {
		return str
	}
	lenth := len(str)
	n = lenth - (n % lenth)
	strs := []byte(str)
	reverse(strs)
	reverse(strs[:n])
	reverse(strs[n:])
	return string(strs)
}

func reverse(s []byte) {

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// 59.滑动窗口的最大值
// 给定一个长度为 n 的数组 num 和滑动窗口的大小 size ，找出所有滑动窗口里数值的最大值。
func maxInWindows(num []int, size int) []int {
	n := len(num)
	if size == 0 || size > n {
		return []int{}
	}
	win := make([]int, 0)
	res := make([]int, 0)
	for i := 0; i < n; i++ {
		for len(win) > 0 && num[i] > num[win[len(win)-1]] {
			win = win[:len(win)-1]
		}
		win = append(win, i)
		for win[0] < i-size+1 {
			win = win[1:]
		}
		if i >= size-1 {
			res = append(res, num[win[0]])
		}

	}
	// write code here
	return res
}

// 61.扑克牌顺子
func IsContinuous(numbers []int) bool {
	sort.Ints(numbers)
	zeroCnt := 0
	need := 0
	for i := 0; i < 5; i++ {
		if numbers[i] == 0 {
			zeroCnt++
		}
	}
	for i := zeroCnt + 1; i < 5; i++ {
		if numbers[i] == numbers[i-1] {
			return false
		}
		need += numbers[i] - numbers[i-1] - 1
	}
	return need <= zeroCnt
}

// 62 孩子们的游戏(圆圈中最后剩下的数)
func LastRemaining_Solution(n int, m int) int {
	if n == 1 {
		return 0
	}
	// 无论怎么样，存活的人就只有一个，我们需要做的就是把这个人的编号映射到原始的编号
	// 在5人圈 0,1,2,3,4 中，第一个淘汰的是 2。
	// 剩下 3, 4, 0, 1。为了套用 f(4,3) 的解，我们把 3,4,0,1 重新编号为 0', 1', 2', 3'。
	// 我们刚算出 f(4,3) 的解是 0。这意味着在新编号 0', 1', 2', 3' 中，幸存者是 0'。
	// 现在，我们看 0' 对应回5人圈的原始编号是谁？
	// 0' -> 3 = (0 + m) % n
	// 1' -> 4 = (1 + m) % n
	// 2' -> 0 = (2 + m) % n
	// 3' -> 1 = (3 + m) % n
	// 幸存者 0' 对应的原始编号是 3。
	// 下面按此递推实现：
	// f(n, m) = (f(n-1, m) + m) % n
	res := 0
	for i := 2; i <= n; i++ {
		res = (res + m) % i
	}
	return res
}

// LastRemaining_Debug: 迭代版，边循环边打印“编码偏移”的映射过程
// 等价递推：ans_i 表示 f(i, m)，有 ans_1 = 0，ans_i = (ans_{i-1} + m) % i。
// 打印内容：
//   - i：当前规模（总人数）
//   - kill = (m-1)%i：本轮第一个被淘汰的原始位置
//   - start = (kill+1)%i = m%i：下一轮重编号的起点
//   - ans_{i-1}：规模 i-1 的“新编号体系”下的答案
//   - 映回原编号：ans_i = (ans_{i-1} + start) % i = (ans_{i-1} + m) % i
func LastRemaining_Debug(n, m int) int {
	if n <= 0 {
		fmt.Println("n 必须为正")
		return -1
	}
	ans := 0 // f(1,m) = 0
	fmt.Printf("i=%d, ans=%d (初始)\n", 1, ans)
	for i := 2; i <= n; i++ {
		kill := (m - 1) % i
		start := (kill + 1) % i // 等价于 m % i
		prev := ans
		ans = (ans + m) % i
		fmt.Printf("i=%d | kill=%d start=%d | ans_{i-1}=%d -> ans_i=(%d+%d)%%%d=%d\n",
			i, kill, start, prev, prev, m, i, ans)
	}
	fmt.Printf("最终答案 f(%d,%d)=%d\n", n, m, ans)
	return ans
}

// 63.买卖股票的最好时机
func maxProfit(prices []int) int {
	// write code here
	preMin := prices[0]
	res := 0
	for i := 1; i < len(prices); i++ {
		res = max(res, prices[i]-preMin)
		preMin = min(preMin, prices[i])

	}
	return res
}

// 64.求1+2+3+...+n
func Sum_Solution(n int) int {
	// write code here
	if n <= 1 {
		return 1
	}
	return Sum_Solution(n-1) + n
}

// 65.不用加减乘除做加法
func Add(num1 int, num2 int) int {
	// write code here
	for num2 != 0 {
		sum := num1 ^ num2
		// 计算进位
		carry := (num1 & num2) << 1
		num1 = sum
		num2 = carry
	}
	return num1
}

// 67.把字符串转换成整数
func StrToInt(str string) int {
	// 1. 去前导空格
	i := 0
	for i < len(str) && str[i] == ' ' {
		i++
	}

	// 如果全是空格
	if i == len(str) {
		return 0
	}

	// 2. 处理符号
	sign := 1
	if str[i] == '-' {
		sign = -1
		i++
	} else if str[i] == '+' {
		i++
	}

	// 3. 解析数字
	var res int
	for i < len(str) && str[i] >= '0' && str[i] <= '9' {
		digit := int(str[i] - '0')

		// 4. 溢出判断
		// math.MaxInt32 = 2147483647
		if res > 214748364 || (res == 214748364 && digit > 7) {
			if sign == 1 {
				return 2147483647
			} else {
				// 如果是负数，且溢出，应返回 math.MinInt32
				// 此处 res*sign 会是 -214748364 * 10 - digit，必然小于 MinInt32
				return -2147483648
			}
		}

		res = res*10 + digit
		i++
	}

	return res * sign
}

// 68.二叉搜索树的最近公共祖先
func lowestCommonAncestor(root *TreeNode, p int, q int) int {
	for root != nil {
		if root.Val > p && root.Val < q {
			return root.Val
		} else if root.Val < p && root.Val > q {
			return root.Val
		} else if root.Val == p || root.Val == q {
			return root.Val
		} else if root.Val > p && root.Val > q {
			// p,q 都在左子树
			root = root.Left
		} else {
			// p,q 都在右子树
			root = root.Right
		}
	}
	return -1
}

func lowestCommonAncestor_2(root *TreeNode, p int, q int) int {
	if root == nil {
		return -1
	}
	if root.Val > p && root.Val < q {
		return root.Val
	} else if root.Val > q && root.Val < p {
		return root.Val
	} else if root.Val == p || root.Val == q {
		return root.Val
	} else if root.Val > p && root.Val > q {
		return lowestCommonAncestor(root.Left, p, q)
	} else {
		// p,q 都在右子树
		return lowestCommonAncestor(root.Right, p, q)
	}
}

// 69.跳楼梯
func jumpFloor(number int) int {
	if number == 1 {
		return 1
	}
	if number == 2 {
		return 2
	}
	dp := make([]int, number+1)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 2
	for i := 2; i <= number; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[number]
}

// 70.矩形覆盖
func rectCover(number int) int {
	if number == 0 {
		return 0
	}
	if number == 1 {
		return 1
	}
	if number == 2 {
		return 2
	}
	// write code here
	dp := make([]int, number+1)
	dp[1] = 1
	dp[2] = 2

	for i := 3; i <= number; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[number]
}

// 71.跳台阶拓展问题
func jumpFloorII(number int) int {
	if number == 1 {
		return 1
	}
	if number == 2 {
		return 2
	}
	dp := make([]int, number+1)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 2
	sum := 3
	for i := 2; i <= number; i++ {
		dp[i] = sum + 1 // 可以跳上n级
		sum += dp[i]    // 统计之前的总数
	}
	// 位运算
	// f(n) = f(n-1) + f(n-2) + ... + f(1) + f(0)
	// f(n-1) = f(n-2) + ... + f(1) + f(0)
	// f(1) = 1
	// f(2) = 2
	// f(3) = 4
	// f(4) = 8
	// f(n) = 2^(n-1)
	// return 1 << (number - 1)
	return dp[number]
}

// 73.翻转单词序列
func ReverseSentence(str string) string {
	// write code here
	bytes := []byte(str)
	reverse(bytes)
	start := 0
	for i := 0; i < len(str); i++ {
		if bytes[i] == ' ' {
			reverse(bytes[start:i])
			start = i + 1
		}
	}
	reverse(bytes[start:])
	return string(bytes)
}

func FindContinuousSequence(sum int) [][]int {
	window := []int{}
	curSum := 0
	res := make([][]int, 0)
	for i := 1; i <= sum/2+1; i++ {
		window = append(window, i)
		curSum += i
		for curSum > sum {
			x := window[0]
			window = window[1:]
			curSum -= x
		}
		if curSum == sum && len(window) > 1 {
			temp := make([]int, len(window))
			copy(temp, window)
			res = append(res, temp)
		}
	}
	return res
}

type DoubleListNode struct {
	Val  int
	Next *ListNode
	Pre  *ListNode
}

// 使用 map 统计频率
var counts map[byte]int

// 使用队列保持“只出现一次”的字符的顺序
var queue []byte

func init() {
	// 在程序启动时或需要重置状态时调用
	counts = make(map[byte]int)
	queue = make([]byte, 0)
}

// 75.字符流中第一个不重复的字符
func Insert1(ch byte) {
	// 计数值加一
	counts[ch]++
	// 如果是第一次出现，则加入队列
	if counts[ch] == 1 {
		queue = append(queue, ch)
	}
}

func FirstAppearingOnce() byte {
	// 循环检查队首元素，如果它已经不是“只出现一次”，则将其从队首移除
	for len(queue) > 0 && counts[queue[0]] > 1 {
		queue = queue[1:] // 出队
	}

	// 如果队列不为空，队首元素就是答案
	if len(queue) > 0 {
		return queue[0]
	}

	// 如果队列为空，说明没有只出现一次的字符
	return '#'
}

// 删除链表中重复的节点
func deleteDuplication(pHead *ListNode) *ListNode {
	// 创建一个哨兵节点，以防头节点就是重复节点
	dummy := &ListNode{Next: pHead}
	cur := dummy

	// 确保 cur.Next 存在，我们总是判断 cur 后面的节点
	for cur.Next != nil && cur.Next.Next != nil {
		// 发现重复
		if cur.Next.Val == cur.Next.Next.Val {
			x := cur.Next.Val // 记录下重复值
			// 循环删除所有值为 x 的节点
			for cur.Next != nil && cur.Next.Val == x {
				cur.Next = cur.Next.Next
			}
		} else {
			// 没有发现重复，cur 安全前进
			cur = cur.Next
		}
	}
	return dummy.Next
}

// 删除排序链表中的重复元素（保留一个）
func deleteDuplicates_keepOne(pHead *ListNode) *ListNode {
	if pHead == nil {
		return nil
	}
	cur := pHead
	for cur != nil && cur.Next != nil {
		// 如果当前节点和下一个节点的值相同
		if cur.Val == cur.Next.Val {
			// 跳过下一个节点
			cur.Next = cur.Next.Next
		} else {
			// 否则，正常前进
			cur = cur.Next
		}
	}
	return pHead
}

// 77.按之字型打印二叉树
func Print(pRoot *TreeNode) [][]int {
	res := make([][]int, 0)
	if pRoot == nil {
		return res
	}
	q := []*TreeNode{pRoot}
	flag := false
	for len(q) > 0 {
		size := len(q)
		temp := make([]int, size)
		flag = !flag
		for i := 0; i < size; i++ {
			x := q[0]
			q = q[1:]
			if flag {
				temp[i] = x.Val
			} else {
				temp[size-i-1] = x.Val
			}
			if x.Left != nil {
				q = append(q, x.Left)
			}
			if x.Right != nil {
				q = append(q, x.Right)
			}
		}
		res = append(res, temp)
	}
	return res
}

func IsBalanced_Solution(pRoot *TreeNode) bool {
	// 调用辅助函数，如果返回 -1 则不平衡，否则平衡
	return getHeightAndCheck(pRoot) != -1
}

// 辅助函数：自底向上计算高度，并同时检查平衡性
// 返回值：-1 表示不平衡，否则返回树的高度
func getHeightAndCheck(node *TreeNode) int {
	// 基准情况：空树是平衡的，高度为 0
	if node == nil {
		return 0
	}

	// 递归获取左子树的高度
	leftHeight := getHeightAndCheck(node.Left)
	// 如果左子树已经不平衡，直接返回 -1，进行“剪枝”
	if leftHeight == -1 {
		return -1
	}

	// 递归获取右子树的高度
	rightHeight := getHeightAndCheck(node.Right)
	// 如果右子树已经不平衡，直接返回 -1
	if rightHeight == -1 {
		return -1
	}

	// 检查当前节点是否平衡
	if abs(leftHeight-rightHeight) > 1 {
		return -1
	}

	// 返回当前树的高度
	return max(leftHeight, rightHeight) + 1
}

// abs 返回整数的绝对值
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// 调整数组顺序使得奇数位于偶数之前
func reOrderArrayTwo(array []int) []int {
	// write code here
	left := 0
	right := len(array)
	for left < right {
		for left < right && array[left]%2 == 0 {
			array[left], array[right] = array[right], array[left]
			right--
		}
		left++
	}
	return array
}

// 82.路径和（从根节点开始）
func hasPathSum(root *TreeNode, sum int) bool {
	// write code here
	if root == nil {
		return false
	}
	sum -= root.Val
	if root.Left == nil && root.Right == nil && sum == 0 {
		return true
	}
	if hasPathSum(root.Left, sum) {
		return true
	}
	if hasPathSum(root.Right, sum) {
		return true
	}
	return false
}

// 二叉树路径和，从任意节点开始
func FindPath3(root *TreeNode, sum int) int {
	// write code here
	cnts := make(map[int]int)
	cnts[0] = 1
	res := 0
	var dfs func(root *TreeNode, cur int, cnts map[int]int)
	dfs = func(root *TreeNode, cur int, cnts map[int]int) {
		if root == nil {
			return
		}
		cur += root.Val
		res += cnts[cur-sum]
		// 记录之前路径的值
		cnts[cur]++
		dfs(root.Left, cur, cnts)
		dfs(root.Right, cur, cnts)
	}
	dfs(root, 0, cnts)
	return res
}

// 85连续子数组的最大和
func FindGreatestSumOfSubArray2(array []int) []int {
	if len(array) == 0 {
		return []int{}
	}

	// 初始化全局最大和与当前最大和
	globalMax := array[0]
	currentMax := array[0]

	// 初始化最终返回的子数组的边界
	startIndex := 0
	endIndex := 0

	// 用于追踪当前子数组的起始位置
	currentStartIndex := 0

	for i := 1; i < len(array); i++ {
		// 判断是延续当前子数组，还是从当前元素开始一个新子数组
		if array[i] > currentMax+array[i] {
			currentMax = array[i]
			currentStartIndex = i
		} else {
			currentMax = currentMax + array[i]
		}

		// 如果当前子数组的和大于已知的全局最大和，或者和相等但长度更长，
		// 则更新全局最大和及其边界。
		if currentMax > globalMax || (currentMax == globalMax && (i-currentStartIndex+1) > (endIndex-startIndex+1)) {
			globalMax = currentMax
			startIndex = currentStartIndex
			endIndex = i
		}
	}

	return array[startIndex : endIndex+1]
}

// 86.二叉树的最近公共祖先
func lowestCommonAncestor2(root *TreeNode, o1 int, o2 int) int {
	// 基准情况：如果树为空，或者找到了 o1 或 o2，则返回当前节点的值（或-1）
	if root == nil {
		return -1
	}
	if root.Val == o1 || root.Val == o2 {
		return root.Val
	}

	// 递归搜索左、右子树
	left := lowestCommonAncestor2(root.Left, o1, o2)
	right := lowestCommonAncestor2(root.Right, o1, o2)

	// 情况1：o1 和 o2 分别在左右子树中，那么当前 root 就是 LCA
	if left != -1 && right != -1 {
		return root.Val
	}

	// 情况2：o1 和 o2 都在左子树中，返回左子树的结果
	if left != -1 {
		return left
	}

	// 情况3：o1 和 o2 都在右子树中，返回右子树的结果
	if right != -1 {
		return right
	}

	// 情况4：左右子树都没有找到 o1 或 o2
	return -1
}

// 剪绳子
func cutRope2(number int64) int64 {
	// 贪心策略：尽可能多地剪出长度为 3 的段
	const MOD = 998244353

	if number <= 3 {
		return number - 1
	}

	a := number / 3
	b := number % 3

	if b == 0 {
		// 完美分解为 a 个 3
		return quickPow(3, a, MOD)
	} else if b == 1 {
		// 把这个1，当成是3+1=4,所以3要少一个
		// 分解为 a-1 个 3 和一个 4 (2*2)
		return (quickPow(3, a-1, MOD) * 4) % MOD
	} else { // b == 2
		// 分解为 a 个 3 和一个 2
		return (quickPow(3, a, MOD) * 2) % MOD
	}
}

// 快速幂算法（用于计算 x^n % mod）
func quickPow(x, n, mod int64) int64 {
	res := int64(1)
	for n > 0 {
		if n%2 == 1 {
			res = (res * x) % mod
		}
		x = (x * x) % mod
		n >>= 1
	}
	return res
}

// LRU 缓存实现
// 双向链表节点
type DListNode struct {
	key, value int
	prev, next *DListNode
}

// LRU 缓存结构
type LRUCache struct {
	capacity int
	cache    map[int]*DListNode
	head     *DListNode // 虚拟头节点
	tail     *DListNode // 虚拟尾节点
}

// 构造函数，初始化 LRU 缓存
func Constructor(capacity int) LRUCache {
	head := &DListNode{}
	tail := &DListNode{}
	head.next = tail
	tail.prev = head

	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*DListNode),
		head:     head,
		tail:     tail,
	}
}

// 将节点移动到链表头部（最近使用）
func (lru *LRUCache) moveToHead(node *DListNode) {
	// 先从当前位置删除
	lru.removeNode(node)
	// 添加到头部
	lru.addToHead(node)
}

// 从链表中删除节点
func (lru *LRUCache) removeNode(node *DListNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

// 在链表头部添加节点
func (lru *LRUCache) addToHead(node *DListNode) {
	node.prev = lru.head
	node.next = lru.head.next
	lru.head.next.prev = node
	lru.head.next = node
}

// 删除链表尾部节点（最久未使用）
func (lru *LRUCache) removeTail() *DListNode {
	node := lru.tail.prev
	lru.removeNode(node)
	return node
}

// Get 操作：获取键对应的值
func (lru *LRUCache) Get(key int) int {
	if node, exists := lru.cache[key]; exists {
		// 将访问的节点移到头部
		lru.moveToHead(node)
		return node.value
	}
	return -1 // 不存在返回 -1
}

// Put 操作：插入或更新键值对
func (lru *LRUCache) Put(key int, value int) {
	if node, exists := lru.cache[key]; exists {
		// 键已存在，更新值并移到头部
		node.value = value
		lru.moveToHead(node)
	} else {
		// 键不存在，创建新节点
		newNode := &DListNode{key: key, value: value}
		lru.cache[key] = newNode
		lru.addToHead(newNode)

		// 检查容量，超出则删除最久未使用的节点
		if len(lru.cache) > lru.capacity {
			tailNode := lru.removeTail()
			delete(lru.cache, tailNode.key)
		}
	}
}
