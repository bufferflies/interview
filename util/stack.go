package util

import "container/list"

type Stack interface {
	Push(element interface{})
	Pop() interface{}
	Peek() interface{}
	Size() int
}

type MyStack struct {
	list *list.List
}

func NewMyStack() *MyStack {
	return &MyStack{
		list: list.New(),
	}
}

func (s *MyStack) Push(element interface{}) {
	s.list.PushBack(element)
}
func (s *MyStack) Size() int {
	return s.list.Len()
}
func (s *MyStack) Pop() interface{} {
	e := s.list.Back()
	if e != nil {
		s.list.Remove(e)
		return e.Value
	}
	return e
}

func (s *MyStack) Peek() interface{} {
	e := s.list.Back()
	if e == nil {
		return e
	}
	return s.list.Back().Value
}
