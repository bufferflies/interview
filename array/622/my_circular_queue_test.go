package _22

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type MyCircularQueue struct {
	size  int
	count int
	//未消费的索引
	index  int
	buffer []int
}

/** Initialize your data structure here. Set the size of the queue to be k. */
func Constructor(k int) MyCircularQueue {
	buffer := make([]int, k)
	return MyCircularQueue{size: k, buffer: buffer}
}

/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	index := (this.index + this.count) % this.size
	this.buffer[index] = value
	this.count++
	return true
}

/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	this.count--
	this.index++
	return true
}

/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	index := (this.index - 1) % this.size
	return this.buffer[index]
}

/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	index := (this.index + this.count - 1) % this.size
	return this.buffer[index]
}

/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
	return this.count < 1
}

/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueue) IsFull() bool {
	return this.count >= this.size
}

func TestUnit_1(t *testing.T) {
	circularQueue := Constructor(3)
	assert.True(t, circularQueue.EnQueue(1))

	assert.True(t, circularQueue.EnQueue(2))
	assert.True(t, circularQueue.EnQueue(3))
	assert.False(t, circularQueue.EnQueue(4))
	assert.Equal(t, 3, circularQueue.Rear())
	assert.True(t, circularQueue.IsFull())
	assert.True(t, circularQueue.DeQueue())
	assert.True(t, circularQueue.EnQueue(4))
	assert.Equal(t, 4, circularQueue.Rear())

}
