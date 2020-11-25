package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnit_builder(t *testing.T) {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	result := buildTree(preorder, inorder)
	expect := buildExpect()
	assert.Equal(t, expect, result)
}
func buildExpect() *TreeNode {
	expect := &TreeNode{Val: 3}
	expect.Left = &TreeNode{Val: 9}
	expect.Right = &TreeNode{Val: 20}
	expect.Right.Left = &TreeNode{Val: 15}
	expect.Right.Right = &TreeNode{Val: 7}
	return expect
}
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	t := &TreeNode{Val: preorder[0]}
	length := len(preorder)
	helper(t, preorder, inorder, 0, length-1, 0, length-1)
	return t
}
func helper(t *TreeNode, preorder []int, inorder []int, preStart int, preEnd int, inStart int, inEnd int) {
	index := findIndex(inorder, t.Val)
	// 左边
	if index != inStart {
		// 计算左边的长度
		leftLength := index - inStart
		t.Left = &TreeNode{Val: preorder[preStart+1]}
		helper(t.Left, preorder, inorder, preStart+1, preStart+leftLength, inStart, index-1)
	}
	// 右边
	if index != inEnd {
		//
		rightLength := inEnd - index
		t.Right = &TreeNode{Val: preorder[preEnd-rightLength+1]}
		helper(t.Right, preorder, inorder, preEnd-rightLength+1, preEnd, index+1, inEnd)
	}
}
