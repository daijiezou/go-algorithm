package _2_designStruct

// 双向链表节点
type Node struct {
	key, val   int
	next, prev *Node
}

// 双向链表
type DoubleList struct {
	// 头尾虚节点
	head, tail *Node
	// 链表元素数
	size int
}

func NewNode(k, v int) *Node {
	return &Node{key: k, val: v}
}

func NewDoubleList() *DoubleList {
	// 初始化双向链表的数据
	head := NewNode(0, 0)
	tail := NewNode(0, 0)
	head.next = tail
	tail.prev = head
	return &DoubleList{head: head, tail: tail, size: 0}
}

func (this *DoubleList) AddLast(x *Node) {
	x.next = this.tail
	x.prev = this.tail.prev
	this.tail.prev.next = x
	this.tail.prev = x
	this.size++
}

// 删除链表中的 x 节点（x 一定存在）
// 由于是双链表且给的是目标 Node 节点，时间 O(1)
func (this *DoubleList) Remove(x *Node) {
	x.prev.next = x.next
	x.next.prev = x.prev
	this.size--
}

// 删除链表中第一个节点，并返回该节点，时间 O(1)
func (this *DoubleList) RemoveFirst() *Node {
	if this.head.next == this.tail {
		return nil
	}
	first := this.head.next
	this.Remove(first)
	return first
}

type LRUCache struct {
	// key -> Node(key, val)
	_map map[int]*Node
	// Node(k1, v1) <-> Node(k2, v2)...
	cache *DoubleList
	// 最大容量
	cap int
}

// 为了节约篇幅，省略上文给出的代码部分...

// 将某个 key 提升为最近使用的
func (this *LRUCache) makeRecently(key int) {
	x := this._map[key]
	// 先从链表中删除这个节点
	this.cache.Remove(x)
	// 重新插到队尾
	this.cache.AddLast(x)
}

// 添加最近使用的元素
func (this *LRUCache) addRecently(key, val int) {
	x := NewNode(key, val)
	// 链表尾部就是最近使用的元素
	this.cache.AddLast(x)
	// 别忘了在 map 中添加 key 的映射
	this._map[key] = x
}

// 删除某一个 key
func (this *LRUCache) deleteKey(key int) {
	x := this._map[key]
	// 从链表中删除
	this.cache.Remove(x)
	// 从 map 中删除
	delete(this._map, key)
}

// 删除最久未使用的元素
func (this *LRUCache) removeLeastRecently() {
	// 链表头部的第一个元素就是最久未使用的
	deletedNode := this.cache.RemoveFirst()
	// 同时别忘了从 map 中删除它的 key
	deletedKey := deletedNode.key
	delete(this._map, deletedKey)
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		_map:  make(map[int]*Node, capacity),
		cache: NewDoubleList(),
		cap:   capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	node := this._map[key]
	if node == nil {
		return -1
	}
	// 将该数据提升为最近使用的
	this.addRecently(key, node.val)
	return node.val
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this._map[key]; ok {
		this.deleteKey(key)
		this.addRecently(key, value)
		return
	}
	if this.cache.size == this.cap {
		this.removeLeastRecently()
	}
	this.addRecently(key, value)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
