package _1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnit(t *testing.T) {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}
	r := rotateRight(head, 2)
	assert.NotNil(t, r)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	length := 1
	slow, fast := head, head
	for fast.Next != nil {
		fast = fast.Next
		length++
	}
	fast.Next = head
	step := length - k%length - 1
	for ; step > 0; step-- {
		slow = slow.Next
	}
	result := slow.Next
	slow.Next = nil
	return result

}
