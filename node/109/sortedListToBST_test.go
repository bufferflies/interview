package _09

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func TestUnit(t *testing.T) {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}
	r := sortedListToBST(head)
	assert.NotNil(t, r)
}
func sortedListToBST(head *ListNode) *TreeNode {
	s, middle := size(head)
	if s == 1 {
		return &TreeNode{Val: head.Val}
	}
	if s == 0 {
		return nil
	}
	root := &TreeNode{Val: middle.Next.Val}
	right := middle.Next.Next
	middle.Next = nil
	root.Left = sortedListToBST(head)
	root.Right = sortedListToBST(right)
	return root
}

func size(header *ListNode) (int, *ListNode) {
	slow, fast := &ListNode{}, header
	slow.Next = header
	size := 1
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		size = size + 2
	}
	if fast == nil {
		size--
	}
	return size, slow
}
