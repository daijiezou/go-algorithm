package main

func main() {

}

type DListNode struct {
	Key, Val int
	Pre      *DListNode
	Next     *DListNode
}

type LRUCache struct {
	cap   int
	Cache map[int]*DListNode
	Head  *DListNode
	Tail  *DListNode
}

func New(cap int) *LRUCache {
	if cap <= 0 {
		panic("cap value is invalid")
	}
	head := &DListNode{}
	tail := &DListNode{}
	head.Next = tail
	tail.Pre = head
	return &LRUCache{
		cap:   cap,
		Cache: make(map[int]*DListNode),
		Head:  head,
		Tail:  tail,
	}
}

func (l *LRUCache) DeleteNode(node *DListNode) {
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre
}

func (l *LRUCache) AddToHead(node *DListNode) {
	node.Next = l.Head.Next
	node.Pre = l.Head
	l.Head.Next.Pre = node
	l.Head.Next = node
}

func (l *LRUCache) RemoveTail() {
	node := l.Tail.Pre
	delete(l.Cache, node.Key)
	l.Tail.Pre.Pre.Next = l.Tail
	l.Tail.Pre = l.Tail.Pre.Pre
}

func (l *LRUCache) MovetoHead(node *DListNode) {
	l.DeleteNode(node)
	l.AddToHead(node)
}

func (l *LRUCache) Get(key int) int {
	if node, ok := l.Cache[key]; ok {
		l.MovetoHead(node)
		return node.Val

	}
	return -1
}

func (l *LRUCache) Put(key int, val int) {
	if node, ok := l.Cache[key]; ok {
		node.Val = val
		l.MovetoHead(node)
		return
	}
	node := &DListNode{
		Key: key,
		Val: val,
	}
	l.Cache[key] = node
	l.AddToHead(node)
	if len(l.Cache) > l.cap {
		l.RemoveTail()
	}
}
