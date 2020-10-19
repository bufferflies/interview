package lru

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Node struct {
	key  int
	data int
	pre  *Node
	next *Node
}

func NewNode(key, val int) *Node {
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
	m        map[int]*Node
}

func Constructor(capacity int) LRUCache {
	cache := make(map[int]*Node, capacity)
	return LRUCache{
		capacity: capacity,
		m:        cache,
		size:     0,
		list:     NewNodeList(),
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.m[key]; ok {
		this.list.remove(node)
		this.list.addLastNode(node)
		return node.data
	} else {
		return -1
	}
}
func (this *LRUCache) addLastNode(key, value int) {
	node := NewNode(key, value)
	this.list.addLastNode(node)
	this.m[key] = node
	this.size++
}
func (this *LRUCache) RemoveFirst() {
	delete(this.m, this.list.head.key)
	this.list.remove(this.list.head)
	this.size--
}
func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.m[key]; ok {
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

func TestUnit(t *testing.T) {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	r := cache.Get(1)
	assert.Equal(t, 1, r)
	cache.Put(3, 3)
	r = cache.Get(2)
	assert.Equal(t, -1, r)
	cache.Put(4, 4)
	r = cache.Get(1)
	assert.Equal(t, -1, r)
	r = cache.Get(3)
	assert.Equal(t, 3, r)
	r = cache.Get(4)
	assert.Equal(t, 4, r)
}
func TestUnit_2(t *testing.T) {
	cache := Constructor(1)
	cache.Put(2, 1)
	r := cache.Get(2)
	assert.Equal(t, 1, r)
}
