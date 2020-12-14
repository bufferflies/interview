package _2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Tree struct {
	val   int
	left  *Tree
	right *Tree
}

//前序遍历 preorder = [3,9,20,15,7]
//left->mid->right
//中序遍历 inorder = [9,3,15,20,7]
func TestUnit(t *testing.T) {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	r := rebuild(preorder, inorder)
	assert.Equal(t, nil, r)
}

func rebuild(preorder []int, inorder []int) *Tree {
	return helper(preorder, inorder)
}

func helper(preorder []int, inorder []int) *Tree {
	if len(preorder) == 1 {
		return &Tree{val: preorder[0]}
	}
	index := findIndex(preorder[0], inorder)
	r := &Tree{val: preorder[0]}
	// bug legth
	r.left = helper(preorder[1:index+1], inorder[:index])
	r.right = helper(preorder[index+1:], inorder[index+1:])
	return r
}

// findIndex
func findIndex(target int, array []int) int {
	for i, v := range array {
		if v == target {
			return i
		}
	}
	return -1
}
