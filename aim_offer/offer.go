package aim_offer

import (
	"container/heap"
	"fmt"
	"math/bits"
	"sort"
	"strconv"
	"strings"
	"sync"
)

var (
	x     string
	xOnce sync.Once
)

func singleton() string {
	if x == "" {
		xOnce.Do(func() {
			x = "1"
		})
	}
	return x
}

// 2.
/*
在一个长度为n的数组里的所有数字都在0~n-1的范围内。数组中某
些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了
几次。请找出数组中任意一个重复的数字。例如，如果输入长度为7的数
组｛2,3，1,0.2,5.3｝，那么对应的输出是重复的数字2或者3。
*/

func dup(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for nums[i] != i {
			if nums[nums[i]] == nums[i] {
				return nums[i]
			} else {
				nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
			}
		}
	}
	return -1
}

/*
在一个长度为n的数组里的所有数字都在0~n-1的范围内。数组中某
些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了
几次。请找出数组中任意一个重复的数字。例如，如果输入长度为7的数
组｛2,3，1,0.2,5.3｝，那么对应的输出是重复的数字2或者3。

不可以修改原数组
*/
func dup2(nums []int) int {
	start := 0
	end := len(nums) - 1
	for start <= end {
		mid := start + (end-start)/2
		cnt := getCnt(nums, start, mid)
		if start == end {
			if cnt > 1 {
				return start
			} else {
				break
			}
		}
		if cnt > mid-start+1 {
			end = mid
		} else {
			start = mid + 1
		}
	}
	return -1
}

func getCnt(nums []int, start, end int) int {
	cnt := 0
	for i := 0; i < len(nums); i++ {
		if start <= nums[i] && nums[i] <= end {
			cnt++
		}
	}
	return cnt
}

/*
题目：在一个二维数组中，每一行都按照从左到右递增的顺序排序，
每一列都按照从上到下递增的顺序排序。请完成一个函数，输入这样的一
个二维数组和一个整数，判断数组中是否含有该整数。
*/

func findTargetIn2DPlants(plants [][]int, target int) bool {
	row := len(plants)
	if row == 0 {
		return false
	}
	col := len(plants[0])
	if col == 0 {
		return false
	}
	x := 0
	y := col - 1
	for x < row && y >= 0 {
		if plants[x][y] == target {
			return true
		} else if plants[x][y] > target {
			// 如果大于目标数，则这一列都会大于该目标数
			y--
		} else if plants[x][y] < target {
			// 如果小于目标数，则这一行都会小于该目标数
			x++
		}
	}
	return false
}

func longestSubsequence(s string, k int) int {
	n := len(s)
	cnt := 0
	sum := 0

	bitlen := bits.Len(uint(k))
	for i := n - 1; i >= 0; i-- {
		index := n - i - 1
		if s[i] == '0' {
			cnt++
		} else {
			if index < bitlen && sum+(1<<index) <= k {
				cnt++
				sum += 1 << index
			}
		}
	}
	return cnt
}

func getSum(bytes []byte) int {
	base := 1
	sum := 0
	for i := 0; i < len(bytes); i++ {
		sum += int(bytes[i]) * base
		base *= 2
	}
	return sum
}

/*
请实现一个函数，将一个字符串s中的每个空格替换成“%20“。
例如，当字符串为“We Are Happy“.则经过替换之后的字符串为“We%20Are%20Happy“。
*/
func replaceSpace(s string) string {
	spaceCnt := strings.Count(s, " ")
	bytes := make([]byte, len(s)+spaceCnt*2)
	resI := len(s) + spaceCnt*2 - 1
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			bytes[resI] = '0'
			bytes[resI-1] = '2'
			bytes[resI-2] = '%'
			resI -= 3
		} else {
			bytes[resI] = s[i]
			resI--
		}
	}
	return string(bytes)
}

func replaceSpace2(s string) string {
	spaceCnt := strings.Count(s, " ")
	bytes := make([]byte, 0, len(s)+spaceCnt*2)
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			bytes = append(bytes, []byte{'%', '2', '0'}...)
		} else {
			bytes = append(bytes, s[i])
		}
	}
	return string(bytes)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func printListFromTailToHead(head *ListNode) []int {
	res := make([]int, 0)
	var reverse func(h *ListNode)
	reverse = func(h *ListNode) {
		if h == nil {
			return
		}
		reverse(h.Next)
		res = append(res, h.Val)
	}
	reverse(head)
	return res
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func reConstructBinaryTree(preOrder []int, vinOrder []int) *TreeNode {
	if len(preOrder) == 0 {
		return nil
	}
	root := &TreeNode{
		Val: preOrder[0],
	}
	index := 0
	for ; index < len(preOrder); index++ {
		if vinOrder[index] == preOrder[0] {
			break
		}
	}
	root.Left = reConstructBinaryTree(preOrder[1:index+1], vinOrder[:index])
	root.Right = reConstructBinaryTree(preOrder[index+1:], vinOrder[index+1:])
	return root
}

type TreeLinkNode struct {
	Val   int
	Left  *TreeLinkNode
	Right *TreeLinkNode
	Next  *TreeLinkNode // 指向父节点
}

// 获取中序遍历的下一个节点
// 左根右
func GetNext(pNode *TreeLinkNode) *TreeLinkNode {
	if pNode == nil {
		return nil
	}
	var next *TreeLinkNode
	// 查看节点的右子树是否存在，如存在则为节点右子树的左边节点
	if pNode.Right != nil {
		pRight := pNode.Right
		for pRight.Left != nil {
			pRight = pRight.Left
		}
		next = pRight
	} else if pNode.Next != nil {
		pCurrent := pNode
		pParent := pNode.Next
		// 直到找到当前节点是其父节点的左子树，则下一个节点就是其父节点
		for pParent != nil && pCurrent == pParent.Right {
			pCurrent = pParent
			pParent = pParent.Next
		}
		next = pParent
	}
	return next
}

var stack1 []int
var stack2 []int

func Push(node int) {
	stack1 = append(stack1, node)
}

func Pop() int {
	if len(stack2) != 0 {
		x := stack2[len(stack2)-1]
		stack2 = stack2[:len(stack2)-1]
		return x
	}
	if len(stack1) == 0 {
		return 0
	}
	for len(stack1) > 0 {
		leave := stack1[len(stack1)-1]
		stack1 = stack1[:len(stack1)-1]
		stack2 = append(stack2, leave)
	}
	x := stack2[len(stack2)-1]
	stack2 = stack2[:len(stack2)-1]
	return x
}

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

func minNumberInRotateArray(nums []int) int {
	// [0,n-2]
	//（-1，n-1)
	left, right := -1, len(nums)-1
	for left+1 < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[right] {
			left = mid
		} else if nums[mid] < nums[right] {
			right = mid
		} else {
			right--
		}
	}
	return nums[right]
}

func hasPath(matrix [][]byte, word string) bool {
	m := len(matrix)
	n := len(matrix[0])
	pos := [][]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}
	visited := make([]bool, m*n)
	var backtrack func(i, j int, start int) bool
	backtrack = func(i, j int, start int) bool {
		if matrix[i][j] != word[start] {
			return false
		}
		if start == len(word)-1 {
			return true
		}
		visited[i*m+j] = true
		for _, xy := range pos {
			nx := i + xy[0]
			ny := j + xy[1]
			if nx < 0 || nx >= m || ny < 0 || ny >= n {
				// 这里需要continue,去尝试其他方向
				continue
			}
			if visited[nx*m+j] {
				continue
			}
			if backtrack(nx, ny, start+1) {
				return true
			}
		}
		visited[i*m+j] = false
		return false
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if backtrack(i, j, 0) {
				return true // 搜到了！
			}
		}
	}
	return false // 没搜到

}

func movingCount(threshold int, rows int, cols int) int {
	visited := make([]bool, rows*cols)
	pos := [][]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if visited[i*cols+j] {
			return 0
		}
		if !check(threshold, i, j) {
			return 0
		}
		cnt := 1
		visited[i*cols+j] = true
		for _, xy := range pos {
			nx := i + xy[0]
			ny := j + xy[1]
			if nx < 0 || nx >= rows || ny < 0 || ny >= cols {
				// 这里需要continue,去尝试其他方向
				continue
			}
			cnt += dfs(nx, ny)
		}
		return cnt
	}
	return dfs(0, 0)
}

func check(threshold int, row int, col int) bool {
	sum := 0
	for row > 0 {
		sum += row % 10
		row /= 10
	}
	for col > 0 {
		sum += col % 10
		col /= 10
	}
	return sum <= threshold
}

func cutRope(n int) int {
	// write code here
	if n < 2 {
		return 0
	}
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
	dp[3] = 3 // 表示可以切，也可以直接返回,3的
	for i := 4; i <= n; i++ {
		for j := 2; j <= i/2; j++ {
			dp[i] = max(dp[i], dp[j]*dp[i-j])
		}
	}
	return dp[n]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

/*
把一个整数减去 1 之后再和原来的整数做位与运算，得到的结果相当
于把整数的二进制表示中最右边的1变成0。很多二进制的问题都可以用这
种思路解决。
*/

func NumberOf1(n int) int {
	cnt := 0
	for n != 0 {
		cnt++
		// 相当于把最右边的1变成了0
		// 1100 -1 = 1001
		// 1001 & 1100  = 1000
		// 计算能做多少次这样的操作，就等于这个数字中有多少个1
		n = (n - 1) & n
	}

	return cnt
}

func toTwosComplement(n int, bits int) string {
	if n >= 0 {
		// 正数直接转换为二进制
		return fmt.Sprintf("%0*b", bits, n)
	}

	// 负数转换为补码
	// 1. 取绝对值
	abs := -n
	// 2. 转换为二进制
	binary := fmt.Sprintf("%0*b", bits, abs)
	// 3. 按位取反
	complement := ""
	for _, bit := range binary {
		if bit == '0' {
			complement += "1"
		} else {
			complement += "0"
		}
	}
	// 4. 加1
	result := ""
	carry := 1
	for i := len(complement) - 1; i >= 0; i-- {
		sum := int(complement[i]-'0') + carry
		result = string(rune('0'+sum%2)) + result
		carry = sum / 2
	}

	return result
}

func Power(base float64, exponent int) float64 {
	if exponent == 1 {
		return base
	}
	if exponent == 0 {
		return 1
	}
	if exponent > 0 {
		return pow(base, exponent)
	} else {
		return 1 / pow(base, -exponent)
	}
}

func pow(x float64, y int) float64 {
	res := float64(1)
	for y > 0 {
		if y&1 == 1 {
			res *= x
		}
		x = x * x
		y >>= 1
	}
	return res
}

func printNumbers(n int) []string {
	res := make([]string, 0)
	if n <= 0 {
		return res
	}
	var dfs func(index int, preStr []byte)
	dfs = func(index int, preStr []byte) {
		if index == n {
			s := string(preStr)
			i := 0
			for i < len(s) && s[i] == '0' {
				i++
			}
			if i < len(s) { // 至少有一位不是0
				res = append(res, s[i:])
			}
			return
		}
		for i := 0; i < 10; i++ {
			preStr[index] = byte('0' + i)
			dfs(index+1, preStr)
		}
	}
	pre := make([]byte, n)
	dfs(0, pre)
	return res
}

func deleteNode(head *ListNode, val int) *ListNode {
	dummy := &ListNode{
		Val:  -1,
		Next: head,
	}
	p := dummy
	for p != nil {
		if p.Next.Val == val {
			p.Next = p.Next.Next
			break
		}
		p = p.Next
	}
	return dummy.Next
}

// 删除重复链表
func deleteDuplication(pHead *ListNode) *ListNode {
	dummy := &ListNode{Next: pHead}
	prev := dummy
	curr := pHead
	for curr != nil {
		duplicate := false
		// 检查是否有重复
		for curr.Next != nil && curr.Val == curr.Next.Val {
			// 跳过重复节点
			curr = curr.Next
			duplicate = true
		}
		if duplicate {
			// 把自身也跳过
			prev.Next = curr.Next
		} else {
			// 相当于接上不重复的
			prev = prev.Next
		}
		curr = curr.Next
	}
	return dummy.Next
}

func deleteDupKeepFirstSorted(head *ListNode) *ListNode {
	curr := head
	for curr != nil && curr.Next != nil {
		if curr.Val == curr.Next.Val {
			// 跳过重复节点
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}
	return head
}

func match(str string, pattern string) bool {
	// write code here
	return matchCore(str, 0, pattern, 0)
}
func matchCore(str string, sIndex int, pattern string, pIndex int) bool {
	// 匹配完成
	if sIndex == len(str) && pIndex == len(pattern) {
		return true
	}
	if pIndex == len(pattern) {
		return false
	}
	if pIndex+1 < len(pattern) && pattern[pIndex+1] == '*' {
		if sIndex < len(str) && (pattern[pIndex] == str[sIndex] || pattern[pIndex] == '.') {
			// 匹配0个或者多个
			return matchCore(str, sIndex+1, pattern, pIndex) ||
				//略过当前的*
				matchCore(str, sIndex, pattern, pIndex+2)
		} else {
			//和当前字符不匹配，略过当前的*
			return matchCore(str, sIndex, pattern, pIndex+2)
		}

		//if sIndex < len(str) && (pattern[pIndex] == str[sIndex] || pattern[pIndex] == '.') {
		//    // 匹配1次及以上 或 匹配0次
		//    return matchCore(str, sIndex+1, pattern, pIndex) || matchCore(str, sIndex, pattern, pIndex+2)
		//} else {
		//    // 匹配0次
		//    return matchCore(str, sIndex, pattern, pIndex+2)
		//}
	}
	if sIndex < len(str) && (pattern[pIndex] == str[sIndex] || pattern[pIndex] == '.') {
		return matchCore(str, sIndex+1, pattern, pIndex+1)
	}
	return false
}

func reOrderArray(array []int) []int {
	res := make([]int, 0, len(array))
	even := make([]int, 0)
	for i := 0; i < len(array); i++ {
		if array[i]%2 != 0 {
			res = append(res, array[i])
		} else {
			even = append(even, array[i])
		}
	}
	return append(res, even...)
}

func reOrderArray1(array []int) []int {
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

// 不稳定，会打乱顺序
func reOrderArray2(array []int) []int {
	// write code here
	left := 0
	right := len(array) - 1
	for left < right {
		for left < right && array[left]%2 != 0 {
			left++
		}
		for left < right && array[right]%2 == 0 {
			right--
		}
		if left < right {
			array[left], array[right] = array[right], array[left]
		}
		left++
		right--
	}
	return array
}

func FindKthToTail(pHead *ListNode, k int) *ListNode {
	if k == 0 || pHead == nil {
		return nil
	}
	fast := pHead
	slow := pHead

	for i := 0; i < k; i++ {
		// k 大于链表长度
		if fast == nil {
			return fast
		}
		fast = fast.Next
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
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			// 2. 找入口
			entry := pHead
			for entry != slow {
				entry = entry.Next
				slow = slow.Next
			}
			return entry
		}
	}

	return nil
}

func EntryNodeOfLoop2(pHead *ListNode) *ListNode {
	fast := pHead
	slow := pHead
	var meetingNode *ListNode
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			meetingNode = slow
			break
		}
	}
	// 没有环
	if meetingNode == nil {
		return nil
	}
	nodesInLoop := 1
	node1 := meetingNode
	for node1.Next != meetingNode {
		node1 = node1.Next
		nodesInLoop++
	}
	fast = pHead
	for i := 0; i < nodesInLoop; i++ {
		fast = fast.Next
	}
	slow = pHead
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}

func ReverseList(head *ListNode) *ListNode {
	// write code here
	var pre, next *ListNode
	cur := head
	for cur != nil {
		// 先保存下一个节点
		next = cur.Next

		// 翻转，将下一个链表指向上一个
		cur.Next = pre

		pre = cur
		cur = next
	}
	// 最后结束的时候cur为nil，pre为cur的上一个节点，返回pre
	return pre
}

func Merge(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	if pHead1 == nil {
		return pHead2
	}
	if pHead2 == nil {
		return pHead1
	}
	head := &ListNode{}
	if pHead1.Val < pHead2.Val {
		head = pHead1
		head.Next = Merge(pHead1.Next, pHead2)
	} else {
		head = pHead2
		head.Next = Merge(pHead1, pHead2.Next)
	}
	return head
}

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

func Mirror(pRoot *TreeNode) *TreeNode {
	if pRoot == nil {
		return pRoot
	}
	pRoot.Right, pRoot.Left = pRoot.Left, pRoot.Right
	Mirror(pRoot.Right)
	Mirror(pRoot.Left)
	return pRoot
}

func isSymmetrical(pRoot *TreeNode) bool {
	if pRoot == nil {
		return true
	}
	return checkIsSymmetrical(pRoot.Left, pRoot.Right)
}

func checkIsSymmetrical(t1, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil {
		return false
	}
	if t1.Val != t2.Val {
		return false
	}
	return checkIsSymmetrical(t1.Left, t2.Right) && checkIsSymmetrical(t1.Right, t2.Left)
}

// 顺时针打印矩阵
func printMatrix(matrix [][]int) []int {
	// write code here
	res := make([]int, 0)
	m := len(matrix)
	if m <= 0 {
		return res
	}
	n := len(matrix[0])
	top, bottom := 0, m-1
	left, right := 0, n-1
	for top <= bottom && left <= right {
		for i := left; i <= right; i++ {
			res = append(res, matrix[top][i])
		}
		top++
		for i := top; i <= bottom; i++ {
			res = append(res, matrix[i][right])
		}
		right--

		// 这里需要额外进行判断。避免重复打印
		if top <= bottom {
			for i := right; i >= left; i-- {
				res = append(res, matrix[bottom][i])
			}
			bottom--
		}

		// 这里需要外进行判断，避免重复打印
		if right >= left {
			for i := bottom; i >= top; i-- {
				res = append(res, matrix[i][left])
			}
			left++
		}

	}
	return res
}

var minStack, stack []int

func Push1(node int) {
	stack = append(stack, node)
	if len(minStack) == 0 || node <= minStack[len(minStack)-1] {
		minStack = append(minStack, node)
	}
}
func Pop1() {
	x := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	if len(minStack) > 0 && x == minStack[len(minStack)-1] {
		minStack = minStack[:len(minStack)-1]
	}
}
func Top() int {
	if len(stack) > 0 {
		return stack[len(stack)-1]
	}
	return 0
}
func Min() int {
	if len(minStack) > 0 {
		return minStack[len(minStack)-1]
	}
	return 0
}

func IsPopOrder(pushV []int, popV []int) bool {
	var stack3 []int
	j := 0
	for _, x := range pushV {
		if x == popV[j] {
			j++
		} else {
			stack3 = append(stack3, x)
		}
		for len(stack3) > 0 && stack3[len(stack3)-1] == popV[j] {
			j++
			stack3 = stack3[:len(stack3)-1]
		}
	}
	return j == len(popV)
}

func PrintFromTopToBottom(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		res = append(res, cur.Val)
		if cur.Left != nil {
			queue = append(queue, cur.Left)
		}
		if cur.Right != nil {
			queue = append(queue, cur.Right)
		}
	}
	return res
}

func Print(pRoot *TreeNode) [][]int {
	res := make([][]int, 0)
	if pRoot == nil {
		return res
	}
	queue := []*TreeNode{pRoot}
	for len(queue) > 0 {
		size := len(queue)
		temp := make([]int, 0)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			temp = append(temp, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, temp)
	}
	return res
}

func Print2(pRoot *TreeNode) [][]int {
	res := make([][]int, 0)
	if pRoot == nil {
		return res
	}
	queue := []*TreeNode{pRoot}
	flag := false
	for len(queue) > 0 {
		size := len(queue)
		temp := make([]int, 0)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			temp = append(temp, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}

		}
		if flag {
			reverse(temp)
		}
		flag = !flag

		res = append(res, temp)
	}
	return res
}

func reverse(nums []int) {
	n := len(nums)
	for i := 0; i < n/2; i++ {
		nums[i], nums[n-i-1] = nums[n-i-1], nums[i]
	}
}

func VerifySquenceOfBST(sequence []int) bool {
	// write code here
	if len(sequence) == 0 {
		return false
	}
	return checkSquenceOfBST(sequence)
}

func checkSquenceOfBST(sequence []int) bool {
	if len(sequence) == 0 {
		return true
	}
	n := len(sequence)
	root := sequence[n-1]
	i := 0
	left := make([]int, 0)
	for ; i < n-1; i++ {
		if sequence[i] > root {
			break
		} else {
			left = append(left, sequence[i])
		}
	}
	right := make([]int, 0)
	for j := i; j < n-1; j++ {
		if sequence[j] < root {
			return false
		}
		right = append(right, sequence[j])
	}
	return checkSquenceOfBST(left) && checkSquenceOfBST(right)
}

func FindPath(root *TreeNode, target int) [][]int {
	cur := make([]int, 0)
	sum := 0
	res := make([][]int, 0)
	var dfs = func(root *TreeNode) {}
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		sum += root.Val
		cur = append(cur, root.Val)
		if root.Right == nil && root.Left == nil {
			if sum == target {
				fmt.Println(cur)
				temp := make([]int, len(cur))
				copy(temp, cur)
				res = append(res, temp)
			}
		}
		dfs(root.Left)
		dfs(root.Right)
		sum -= root.Val
		cur = cur[:len(cur)-1]
	}
	dfs(root)
	return res
}

type RandomListNode struct {
	Label  int
	Next   *RandomListNode
	Random *RandomListNode
}

/**
 *
 * @param pHead RandomListNode类
 * @return RandomListNode类
 */
func Clone(head *RandomListNode) *RandomListNode {
	if head == nil {
		return nil
	}

	// 第一步：创建所有节点的映射，只复制Label值
	m := make(map[*RandomListNode]*RandomListNode)
	cur := head
	for cur != nil {
		m[cur] = &RandomListNode{
			Label: cur.Label,
		}
		cur = cur.Next
	}

	// 第二步：设置Next和Random指针，指向新创建的节点
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

func Convert(pRootOfTree *TreeNode) *TreeNode {
	if pRootOfTree == nil {
		return nil
	}

	var pLastNodeInList *TreeNode

	// 内联的ConvertNode函数逻辑
	var convertNode func(pCurrent *TreeNode)
	convertNode = func(pCurrent *TreeNode) {
		if pCurrent == nil {
			return
		}

		// 递归处理左子树
		convertNode(pCurrent.Left)

		// 处理当前节点：建立双向链接
		pCurrent.Left = pLastNodeInList
		if pLastNodeInList != nil {
			pLastNodeInList.Right = pCurrent
		}

		// 更新pLastNodeInList为当前节点
		pLastNodeInList = pCurrent

		// 递归处理右子树
		convertNode(pCurrent.Right)
	}

	// 开始转换
	convertNode(pRootOfTree)

	// pLastNodeInList指向双向链表的尾节点，
	// 我们需要返回头节点
	pHeadOfList := pLastNodeInList
	for pHeadOfList != nil && pHeadOfList.Left != nil {
		pHeadOfList = pHeadOfList.Left
	}

	return pHeadOfList
}

func Convert2(pRootOfTree *TreeNode) *TreeNode {
	if pRootOfTree == nil {
		return nil
	}
	var dfs func(root *TreeNode)
	var pHeadOfList *TreeNode
	dfs = func(cur *TreeNode) {
		if cur == nil {
			return
		}
		dfs(cur.Right)
		cur.Right = pHeadOfList
		if pHeadOfList != nil {
			pHeadOfList.Left = cur
		}
		pHeadOfList = cur
		dfs(cur.Left)
	}
	dfs(pRootOfTree)
	return pHeadOfList
}

// Serialize encodes a tree to a single string.
func Serialize(root *TreeNode) string {
	var result []string
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			result = append(result, "#")
			return
		}
		result = append(result, strconv.Itoa(node.Val))
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	// 不能直接将数字转为字符串，否则无法区分 12,3 和1，23
	return strings.Join(result, ",")
}

// Deserialize decodes your encoded data to tree.
func Deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}

	values := strings.Split(data, ",")
	index := 0

	var dfs func() *TreeNode
	dfs = func() *TreeNode {
		if index >= len(values) || values[index] == "#" {
			index++
			return nil
		}

		val, _ := strconv.Atoi(values[index])
		index++

		node := &TreeNode{Val: val}
		node.Left = dfs()
		node.Right = dfs()

		return node
	}

	return dfs()
}

func Permutation(str string) []string {
	if len(str) == 0 {
		return []string{}
	}

	// 将字符串转换为字符数组并排序
	chars := []byte(str)
	sort.Slice(chars, func(i, j int) bool {
		return chars[i] < chars[j]
	})

	n := len(chars)
	res := make([]string, 0)
	used := make([]bool, n)

	var backtrack func(track []byte)
	backtrack = func(track []byte) {
		if len(track) == n {
			res = append(res, string(track))
			return
		}

		for i := 0; i < n; i++ {
			// 跳过已使用的字符
			if used[i] {
				continue
			}

			// 关键：跳过重复字符
			// 如果当前字符与前一个字符相同，且前一个字符还没有被使用，则跳过
			/*
				举个具体例子
				假设我们有字符串"aab"，排序后是['a', 'a', 'b']，索引分别是[0, 1, 2]。

				情况分析：
				当我们处理索引1的'a'时：

				chars[1] == chars[0]（都是'a'）
				此时有两种情况：
				如果used[0] == false（第一个'a'还没用）：
				我们跳过索引1的'a'
				强制要求必须先用索引0的'a'
				如果used[0] == true（第一个'a'已经用了）：
				我们可以使用索引1的'a'
				因为已经按顺序了
			*/
			if i > 0 && chars[i] == chars[i-1] && !used[i-1] {
				continue // "如果你想用第二个'a'，但第一个'a'还没用，那就等等！"
			}

			used[i] = true
			track = append(track, chars[i])
			backtrack(track)
			track = track[:len(track)-1]
			used[i] = false
		}
	}

	backtrack([]byte{})
	return res
}

/*
*摩尔投票算法（Boyer-Moore Majority Vote Algorithm）**来实现O(n)时间复杂度和O(1)空间复杂度的解法。
这个算法的核心思想是：如果一个数字出现次数超过数组长度的一半，那么它在"投票对抗"中一定会胜出。
空间效率：相比哈希表O(n)空间，只需要O(1)空间
时间效率：仍然保持O(n)时间复杂度
简洁性：代码简单，逻辑清晰
实用性：在大数据场景下特别有用，内存占用极小
*/

func MoreThanHalfNum_Solution(numbers []int) int {
	n := len(numbers)
	if n < 1 {
		return -1
	}
	cnt := 1
	candidate := numbers[0]
	for i := 1; i < n; i++ {
		if cnt == 0 {
			candidate = numbers[i]
			cnt = 1
		} else if candidate == numbers[i] {
			cnt++
		} else {
			cnt--
		}
	}
	return candidate
}

// GetLeastNumbers_Solution 使用快速选择算法实现
func GetLeastNumbers_Solution(input []int, k int) []int {
	n := len(input)
	if n <= 0 || k > n || k <= 0 {
		return []int{}
	}
	start := 0
	end := n - 1
	index := Part(input, start, end)
	for index != k-1 {
		if index > k-1 {
			end = index - 1
			index = Part(input, start, end)
		} else {
			start = index + 1
			index = Part(input, start, end)
		}
	}
	return input[:k]
}

// GetLeastNumbers_Heap 使用大顶堆实现（推荐）
// 时间复杂度: O(n log k), 空间复杂度: O(k)
func GetLeastNumbers_Heap(input []int, k int) []int {
	if len(input) <= 0 || k > len(input) || k <= 0 {
		return []int{}
	}

	// 使用大顶堆维护k个最小元素
	maxHeap := &MaxHeap{}

	for _, num := range input {
		if maxHeap.Len() < k {
			// 堆未满，直接加入
			heap.Push(maxHeap, num)
		} else if num < maxHeap.Top() {
			// 当前元素比堆顶小，替换
			heap.Pop(maxHeap)
			heap.Push(maxHeap, num)
		}
	}

	// 提取结果
	result := make([]int, 0, k)
	for maxHeap.Len() > 0 {
		result = append(result, heap.Pop(maxHeap).(int))
	}

	//// 由于是大顶堆，出来的顺序是从大到小，需要反转
	//for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
	//	result[i], result[j] = result[j], result[i]
	//}

	return result
}

// GetLeastNumbers_MinHeap 使用小顶堆实现（空间效率低）
// 时间复杂度: O(n log n), 空间复杂度: O(n)
func GetLeastNumbers_MinHeap(input []int, k int) []int {
	if len(input) <= 0 || k > len(input) || k <= 0 {
		return []int{}
	}

	// 将所有元素放入小顶堆
	minHeap := &MinHeap{}
	for _, num := range input {
		heap.Push(minHeap, num)
	}

	// 取出k个最小元素
	result := make([]int, 0, k)
	for i := 0; i < k && minHeap.Len() > 0; i++ {
		result = append(result, heap.Pop(minHeap).(int))
	}

	return result
}

func Part(arr []int, left, right int) int {
	// 选择最左边的元素作为基准值
	pivot := arr[right]

	i := left
	// 遍历数组，将小于基准值的元素移到左侧
	for j := left; j < right; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	// 将基准值放到正确的位置
	arr[right], arr[i] = arr[i], arr[right]
	return i
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
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
