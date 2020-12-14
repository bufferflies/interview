package _16

import (
	"testing"

	"git.code.oa.com/geeker/awesome-work/util"

	"github.com/stretchr/testify/assert"
)

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func TestUnit(t *testing.T) {
	root := &Node{Val: 1}
	root.Left = &Node{Val: 2}
	root.Right = &Node{Val: 3}
	r := connect(root)
	assert.NotNil(t, r)
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	head := root
	queue := util.NewMyQueue()
	queue.Enque(head)
	for queue.Size() > 0 {
		count := queue.Size()
		pre := queue.Peek().(*Node)
		for i := 0; i < count; i++ {
			n := queue.Deque().(*Node)
			if n.Left != nil {
				queue.Enque(n.Left)
			}
			if n.Right != nil {
				queue.Enque(n.Right)
			}
			if i != 0 {
				pre.Next = n
				pre = n
			}
		}
		pre.Next = nil
	}
	return head
}
