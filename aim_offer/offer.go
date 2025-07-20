package aim_offer

import (
	"fmt"
	"math/bits"
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
