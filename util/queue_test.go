package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnit_Queue(t *testing.T) {
	te := assert.New(t)
	queue := NewMyQueue()
	te.Nil(queue.Peek())
	te.Nil(queue.Deque())
	queue.Enque(1)
	queue.Enque(2)
	queue.Enque(3)
	te.Equal(1, queue.Peek())
	te.Equal(1, queue.Deque())
	queue.Enque(3)
	te.Equal(2, queue.Peek())
	te.Equal(2, queue.Deque())
	te.Equal(2, queue.Size())

}
