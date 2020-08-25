package tree

import "testing"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	bit := 0
	ptr := l1
	otherPtr := l2
	result := &ListNode{}
	guard := result
	for ptr != nil || otherPtr != nil {
		a := 0
		b := 0
		if ptr != nil {
			a = ptr.Val
		}
		if otherPtr != nil {
			b = otherPtr.Val
		}
		sum := a + b + bit
		if sum >= 10 {
			sum = sum - 10
			bit = 1
		} else {
			bit = 0
		}
		guard.Next = newListNode(sum)
		guard = guard.Next
		if ptr != nil {
			ptr = ptr.Next
		}
		if otherPtr != nil {
			otherPtr = otherPtr.Next
		}
	}
	if bit != 0 {
		guard.Next = newListNode(bit)
	}
	return result.Next
}
func newListNode(val int) *ListNode {
	return &ListNode{Val: val}
}

func Test_addTwoNumbers(t *testing.T) {
	l1 := newListNode(2)
	l2 := newListNode(5)
	addTwoNumbers(l1, l2)
}
