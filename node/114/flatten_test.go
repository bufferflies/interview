package _14

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnit(t *testing.T) {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	//root.Left.Left = &TreeNode{Val: 3}
	//root.Left.Right = &TreeNode{Val: 4}
	root.Right = &TreeNode{Val: 5}
	//root.Right.Right = &TreeNode{Val: 6}
	flatten(root)
	assert.NotNil(t, root)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//var pre *TreeNode
type H struct {
	pre *TreeNode
}

func flatten(root *TreeNode) {
	h := H{}
	h.helper(root)
}
func (h *H) helper(root *TreeNode) {
	if root == nil {
		return
	}
	h.helper(root.Right)
	h.helper(root.Left)
	root.Right = h.pre
	root.Left = nil
	h.pre = root
}
