package _1_core_platform

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.cn/problems/minimum-depth-of-binary-tree/description/
// 二叉树的最小深度
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	nodeList := []*TreeNode{root}
	depth := 1
	for len(nodeList) > 0 {
		length := len(nodeList)
		for i := 0; i < length; i++ {
			cur := nodeList[0]
			nodeList = nodeList[1:]
			if cur.Left == nil && cur.Right == nil {
				return depth
			}
			if cur.Left != nil {
				nodeList = append(nodeList, cur.Left)
			}
			if cur.Right != nil {
				nodeList = append(nodeList, cur.Right)
			}
		}
		depth++
	}
	return depth
}

// https://leetcode.cn/problems/open-the-lock/submissions/521528158/
/*
	锁的初始数字为 '0000' ，一个代表四个拨轮的数字的字符串。
	列表 deadends 包含了一组死亡数字，一旦拨轮的数字和列表里的任何一个元素相同，这个锁将会被永久锁定，无法再被旋转。
	字符串 target 代表可以解锁的数字，你需要给出解锁需要的最小旋转次数，如果无论如何不能解锁，返回 -1 。
*/
func openLock(deadends []string, target string) int {
	visited := map[string]bool{}
	for _, deadend := range deadends {
		visited[deadend] = true
		if deadend == "0000" {
			return -1
		}
	}
	visited["0000"] = true
	var q []string
	q = append(q, "0000")

	step := 0
	for len(q) > 0 {
		sz := len(q)
		/* 将当前队列中的所有节点向周围扩散 */
		for i := 0; i < sz; i++ {
			cur := q[0]
			q = q[1:]

			/* 判断是否到达终点 */
			if cur == target {
				return step
			}

			/* 将一个节点的相邻节点加入队列 */
			for j := 0; j < 4; j++ {
				up := plusOne(cur, j)
				if _, ok := visited[up]; !ok {
					q = append(q, up)
					visited[up] = true
				}

				down := minusOne(cur, j)
				if _, ok := visited[down]; !ok {
					q = append(q, down)
					visited[down] = true
				}
			}
		}
		/* 在这里增加步数 */
		step++
	}
	return -1
}

func plusOne(s string, j int) string {
	ch := []byte(s)
	if ch[j] == '9' {
		ch[j] = '0'
	} else {
		ch[j] += 1
	}
	return string(ch)
}

// 将 s[i] 向下拨动一次
func minusOne(s string, j int) string {
	ch := []byte(s)
	if ch[j] == '0' {
		ch[j] = '9'
	} else {
		ch[j] -= 1
	}
	return string(ch)
}

/*
传统的 BFS 框架就是从起点开始向四周扩散，遇到终点时停止；
双向 BFS 则是从起点和终点同时开始扩散，当两边有交集的时候停止。
*/

func openLock2(deadends []string, target string) int {
	visited := map[string]bool{}
	for _, deadend := range deadends {
		visited[deadend] = true
		if deadend == "0000" {
			return -1
		}
	}
	visited["0000"] = true
	q := map[string]bool{
		"0000": true,
	}
	p := map[string]bool{
		target: true,
	}
	step := 0
	for len(q) > 0 && len(p) > 0 {
		temp := make(map[string]bool)
		/* 将当前队列中的所有节点向周围扩散 */
		for cur := range q {
			if _, ok := p[cur]; ok {
				return step
			}
			visited[cur] = true
			/* 将一个节点的相邻节点加入队列 */
			for j := 0; j < 4; j++ {
				up := plusOne(cur, j)
				if _, ok := visited[up]; !ok {
					temp[up] = true
				}

				down := minusOne(cur, j)
				if _, ok := visited[down]; !ok {
					temp[down] = true
				}
			}
		}
		q = p
		p = temp
		/* 在这里增加步数 */
		step++
	}
	return -1
}
