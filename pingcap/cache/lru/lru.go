package lru

type Node struct {
	key  []byte
	data []byte
	pre  *Node
	next *Node
}

func NewNode(key, val []byte) *Node {
	return &Node{data: val, key: key}
}

type NodeList struct {
	size int
	head *Node
	last *Node
}

func NewNodeList() *NodeList {
	return &NodeList{}
}
func (list *NodeList) addLastNode(node *Node) {
	if list.size == 0 {
		list.head = node
	} else {
		list.last.next = node
		node.pre = list.last
	}
	list.last = node
	list.size++
}
func (list *NodeList) remove(src *Node) {
	if list.size == 1 {
		list.head = nil
		list.last = nil
	} else if src == list.last {
		last := src.pre
		list.last = last
		last.next = nil
		src.pre = nil
	} else if src == list.head {
		next := src.next
		list.head = next
		next.pre = nil
		src.next = nil
	} else {
		pre := src.pre
		next := src.next
		src.pre = nil
		src.next = nil
		pre.next = next
		next.pre = pre
	}
	list.size--
}

type LRUCache struct {
	capacity int
	size     int
	list     *NodeList
	m        map[string]*Node
}

func Constructor(capacity int) LRUCache {
	cache := make(map[string]*Node, capacity)
	return LRUCache{
		capacity: capacity,
		m:        cache,
		size:     0,
		list:     NewNodeList(),
	}
}

func (this *LRUCache) Get(key []byte) []byte {
	if node, ok := this.m[string(key)]; ok {
		this.list.remove(node)
		this.list.addLastNode(node)
		return node.data
	} else {
		return nil
	}
}
func (this *LRUCache) addLastNode(key, value []byte) {
	node := NewNode(key, value)
	this.list.addLastNode(node)
	this.m[string(key)] = node
	this.size++
}
func (this *LRUCache) RemoveFirst() {
	delete(this.m, string(this.list.head.key))
	this.list.remove(this.list.head)
	this.size--
}
func (this *LRUCache) Add(key, value []byte) {
	if node, ok := this.m[string(key)]; ok {
		this.list.remove(node)
		node.data = value
		this.list.addLastNode(node)
		return
	}
	if this.size >= this.capacity {
		this.RemoveFirst()
	}
	this.addLastNode(key, value)
}
