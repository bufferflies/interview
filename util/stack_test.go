package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnit_Mystack(t *testing.T) {
	te := assert.New(t)
	stack := NewMyStack()
	te.Nil(stack.Peek())
	te.Nil(stack.Pop())
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	te.Equal(3, stack.Peek())
	te.Equal(3, stack.Pop())
	te.Equal(2, stack.Pop())
	stack.Push(4)
	te.Equal(4, stack.Peek())
	te.Equal(4, stack.Pop())
}
