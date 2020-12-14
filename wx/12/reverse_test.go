package _2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Node struct {
	val  int
	Next *Node
}

func TestUnit_Normal(t *testing.T) {
	//1->2->3->4->5
	root := &Node{val: 1}
	root.Next = &Node{val: 2}
	root.Next.Next = &Node{val: 3}
	root.Next.Next.Next = &Node{val: 4}
	root.Next.Next.Next.Next = &Node{val: 5}
	expect := &Node{val: 5}
	expect.Next = &Node{val: 4}
	expect.Next.Next = &Node{val: 3}
	expect.Next.Next.Next = &Node{val: 2}
	expect.Next.Next.Next.Next = &Node{val: 1}

	//r := reverse(root)
	r := reverse(root)
	assert.Equal(t, expect, r)
}
func TestUnit_3(t *testing.T) {
	//1->2->3->4->5
	root := &Node{val: 1}
	root.Next = &Node{val: 2}
	root.Next.Next = &Node{val: 3}
	root.Next.Next.Next = &Node{val: 4}
	expect := &Node{val: 4}
	expect.Next = &Node{val: 3}
	expect.Next.Next = &Node{val: 2}
	expect.Next.Next.Next = &Node{val: 1}

	//r := reverse(root)
	r := reverse(root)
	assert.Equal(t, expect, r)
}

// reverse
func reverse(root *Node) *Node {
	if root == nil {
		return nil
	}
	tail := root
	head := &Node{Next: root}
	for tail.Next != nil {
		tmp := tail.Next
		tail.Next = tail.Next.Next
		tmp.Next = head.Next
		head.Next = tmp
		//head.Next.Next = tail
	}
	return head.Next
}
