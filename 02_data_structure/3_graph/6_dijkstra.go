package __graph

import "container/heap"

type State struct {
	id            int
	distFromStart int
}

func adj(s int) []int {
	// 输入节点 s 返回 s 的相邻节点
}

func weight(from, to int) int {
	// 返回节点 from 到节点 to 之间的边的权重
}

type PriorityQueue []*State

func dijkstra(start int, graph [][]int) []int {
	// 图中节点的个数
	V := len(graph)
	// 记录最短路径的权重，你可以理解为 dp table
	// 定义：distTo[i] 的值就是节点 start 到达节点 i 的最短路径权重
	distTo := make([]int, V)
	// 求最小值，所以 dp table 初始化为正无穷
	for i := range distTo {
		distTo[i] = -1
	}
	// base case，start 到 start 的最短距离就是 0
	distTo[start] = 0

	// 优先级队列，distFromStart 较小的排在前面
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &State{start, 0})

	for pq.Len() > 0 {
		curState := heap.Pop(&pq).(*State)
		// extend down -200
		// <div class="img-content"><img src="/images/dijkstra/4.jpeg" class="myimage"/></div>
		curNodeID := curState.id
		curDistFromStart := curState.distFromStart

		if curDistFromStart > distTo[curNodeID] {
			// 已经有一条更短的路径到达 curNode 节点了
			continue
		}
		// 将 curNode 的相邻节点装入队列
		for _, nextNodeID := range adj(curNodeID) {
			// 看看从 curNode 达到 nextNode 的距离是否会更短
			distToNextNode := distTo[curNodeID] + weight(curNodeID, nextNodeID)
			if distTo[nextNodeID] < 0 || distTo[nextNodeID] > distToNextNode {
				// 更新 dp table
				distTo[nextNodeID] = distToNextNode
				// 将这个节点以及距离放入队列
				heap.Push(&pq, &State{nextNodeID, distToNextNode})
			}
		}
	}

	return distTo
}

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].distFromStart < pq[j].distFromStart
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*State)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	item := (*pq)[len(*pq)-1]
	*pq = (*pq)[0 : len(*pq)-1]
	return item
	//old := *pq
	//n := len(old)
	//item := old[n-1]
	//*pq = old[:n-1]
	//return item
}
