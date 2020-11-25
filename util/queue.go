package util

import "container/list"

type Queue interface {
	Enque(element interface{})
	Deque() interface{}
	Peek() interface{}
	Size() int
}
type MyQueue struct {
	list *list.List
}

func NewMyQueue() *MyQueue {
	return &MyQueue{list: list.New()}
}

func (m MyQueue) Enque(element interface{}) {
	m.list.PushBack(element)
}
func (m MyQueue) Peek() interface{} {
	e := m.list.Front()
	if e != nil {
		return e.Value
	}
	return nil
}
func (m MyQueue) Deque() interface{} {
	e := m.list.Front()
	if e != nil {
		m.list.Remove(e)
		return e.Value
	}
	return nil
}

func (m MyQueue) Size() int {
	return m.list.Len()
}
