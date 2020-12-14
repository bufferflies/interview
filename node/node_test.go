package node

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type List interface {
	Get(int int) int
	AddAtHead(val int)
	AddAtTail(val int)
	AddAtIndex(index int, val int)
	DeleteAtIndex(index int)
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

type MyLinkedList struct {
	head *ListNode
	tail *ListNode
	size int
}

func Constructor() MyLinkedList {
	return MyLinkedList{size: 0}
}

func (l MyLinkedList) Get(index int) int {
	if index >= l.size || index < 0 {
		return -1
	}
	if index == 0 {
		return l.head.Val
	}
	tmp := l.head
	for tmp.Next != nil {
		index--
		if tmp.Next != nil {
			tmp = tmp.Next
		}
		if index == 0 {
			return tmp.Val
		}
	}
	return -1

}
func (l *MyLinkedList) AddAtHead(val int) {
	node := &ListNode{
		Val:  val,
		Next: l.head,
	}
	if l.size == 0 {
		l.tail = node
	}
	l.head = node
	l.size = l.size + 1
}
func (l *MyLinkedList) AddAtTail(val int) {
	node := &ListNode{
		Val:  val,
		Next: nil,
	}
	if l.size == 0 {
		l.head = node
	} else {
		l.tail.Next = node
	}
	l.tail = node
	l.size = l.size + 1
}
func (l *MyLinkedList) AddAtIndex(index int, val int) {
	if index <= 0 {
		l.AddAtHead(val)
		return
	} else if index > l.size {
		return
	} else if index == l.size {
		l.AddAtTail(val)
		return
	}
	tmp := l.head
	for tmp.Next != nil {
		index--
		if index == 0 {
			newNode := &ListNode{
				Val:  val,
				Next: tmp.Next,
			}
			tmp.Next = newNode
			l.size++
			return

		}
		tmp = tmp.Next
	}
}
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.size {
		return
	}
	if index == 0 {
		tmp := this.head
		this.head = this.head.Next
		tmp.Next = nil
		this.size = this.size - 1
		return
	}
	node := this.head
	for node.Next != nil {
		index--

		if index == 0 {
			// 如果node.Next.Next == nil 说明到最后一个节点了.相当于删除最后一个节点
			if node.Next.Next == nil {
				node.Next = nil
				this.tail = node
				this.size--
				return
			}
			// 其他情况就是删除中间一个节点(A->B->C),操作就是  A 直接指向 C 就行 (A->C)
			node2 := node.Next.Next
			node.Next = node2
			this.size--
			return
		}
		node = node.Next
	}
}

func TestUnit(t *testing.T) {
	linkedList := Constructor()
	linkedList.AddAtHead(1)
	linkedList.AddAtTail(3)
	linkedList.AddAtIndex(1, 2) //链表变为1-> 2-> 3
	val := linkedList.Get(1)    //返回2
	assert.Equal(t, 2, val)
	linkedList.DeleteAtIndex(1) //现在链表是1-> 3
	val = linkedList.Get(1)     //返回3
	assert.Equal(t, 3, val)

}
func TestUnit1(t *testing.T) {
	linkedList := Constructor()
	linkedList.AddAtHead(1)
	linkedList.DeleteAtIndex(0) //现在链表是1-> 3
}
func TestUnit2(t *testing.T) {
	linkedList := Constructor()
	linkedList.AddAtHead(1)
	linkedList.AddAtTail(3)
	linkedList.AddAtIndex(1, 2) //链表变为1-> 2-> 3
	val := linkedList.Get(1)    //返回2
	assert.Equal(t, 2, val)
	linkedList.DeleteAtIndex(0) //现在链表是2-> 3
	val = linkedList.Get(0)     //返回3
	assert.Equal(t, 2, val)
}
func TestUnit3(t *testing.T) {
	linkedList := Constructor()
	linkedList.AddAtIndex(0, 10) //链表变为
	linkedList.AddAtIndex(0, 20) //链表变为20->10
	linkedList.AddAtIndex(1, 30) //链表变为20->30->10
	val := linkedList.Get(0)     //返回2
	assert.Equal(t, 20, val)

}
