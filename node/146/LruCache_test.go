package _46

import (
	"container/list"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnit_lru(t *testing.T) {
	te := assert.New(t)
	l := Constructor(2)
	l.Put(1, 1)
	l.Put(2, 2)
	te.Equal(1, l.Get(1))
	l.Put(3, 3)
	te.Equal(-1, l.Get(2))
}
func TestUnit_lru_modify(t *testing.T) {
	te := assert.New(t)
	l := Constructor(2)
	l.Put(2, 1)
	l.Put(2, 2)
	te.Equal(2, l.Get(2))
	l.Put(3, 3)
	te.Equal(2, l.Get(2))
}

type LRUCache struct {
	list     *list.List
	dic      map[int]*list.Element
	size     int
	capacity int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{size: capacity,
		dic:  make(map[int]*list.Element, capacity),
		list: list.New(),
	}
}

type node struct {
	key   int
	value int
}

func newNode(key, value int) *node {
	return &node{key: key, value: value}
}

func (this *LRUCache) Get(key int) int {
	v, ok := this.dic[key]
	if !ok {
		return -1
	}
	this.list.MoveToBack(v)
	return v.Value.(*node).value
}
func (this *LRUCache) moveBack(e *list.Element) {
	this.list.Remove(e)
	this.list.PushBack(e)
}
func (this *LRUCache) Put(key int, value int) {
	v, ok := this.dic[key]
	if ok {
		n := v.Value.(*node)
		n.value = value
		this.list.MoveToBack(v)
	} else {
		n := newNode(key, value)
		e := this.list.PushBack(n)
		this.dic[key] = e
		this.capacity++
	}
	if this.capacity > this.size {
		del := this.list.Front()
		this.list.Remove(del)
		delete(this.dic, del.Value.(*node).key)
		this.capacity--
	}
}
