package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Bucket struct {
}

func TestUnit_operate(t *testing.T) {
	assert.Equal(t, 0, getParent(1))
	assert.Equal(t, 0, getParent(2))
	assert.Equal(t, 1, getParent(3))
	assert.Equal(t, 1, getParent(4))
}

func TestUnit_MaxBucket(t *testing.T) {
	b := &Bucket{}
	src := []int{5, 3, 2, 1, 4}
	b.buildMaxBucket(src)
	expect := []int{5, 4, 2, 1, 3}
	assert.Equal(t, expect, src)
}
func TestUnit_MinBucket(t *testing.T) {
	b := &Bucket{}
	src := []int{5, 3, 2, 1, 4}
	b.buildMinBucket(src)
	expect := []int{1, 2, 3, 5, 4}
	assert.Equal(t, expect, src)
}

func TestUnit_Bucket_Sort_AES(t *testing.T) {
	b := &Bucket{}
	src := []int{5, 3, 2, 1, 4}
	expect := []int{1, 2, 3, 4, 5}
	result := b.sort(src)
	assert.Equal(t, expect, result)
}
func TestUnit_Bucket_Sort_DESC(t *testing.T) {
	b := &Bucket{}
	src := []int{5, 3, 2, 1, 4}
	expect := []int{5, 4, 3, 2, 1}
	result := b.sortDesc(src)
	assert.Equal(t, expect, result)
}
func (b Bucket) buildMaxBucket(src []int) {
	for i := range src {
		upMax(src, i)
	}
}
func (b Bucket) buildMinBucket(src []int) {
	for i := range src {
		upMin(src, i)
	}
}

func (b Bucket) sort(src []int) []int {
	b.buildMaxBucket(src)
	for i, length := 0, len(src); i < length; i++ {
		src[0], src[length-1-i] = src[length-1-i], src[0]
		sinkMax(src, 0, length-i-1)
	}
	return src
}
func (b Bucket) sortDesc(src []int) []int {
	b.buildMinBucket(src)
	for i, length := 0, len(src); i < length; i++ {
		src[0], src[length-1-i] = src[length-1-i], src[0]
		sinkMin(src, 0, length-i-1)
	}
	return src
}
func sinkMax(src []int, index int, end int) {
	if index > end {
		return
	}
	for left, right := getLeftChild(index), getRightChild(index); left < end && right < end; left, right = getLeftChild(index), getRightChild(index) {
		if src[left] < src[index] && src[right] < src[index] {
			break
		} else if src[left] > src[right] {
			src[left], src[index] = src[index], src[left]
			index = left
		} else {
			src[right], src[index] = src[index], src[right]
			index = right
		}
	}
}
func sinkMin(src []int, index int, end int) {
	if index > end {
		return
	}
	for left, right := getLeftChild(index), getRightChild(index); left < end && right < end; left, right = getLeftChild(index), getRightChild(index) {
		if src[left] > src[index] && src[right] > src[index] {
			break
		} else if src[left] < src[right] {
			src[left], src[index] = src[index], src[left]
			index = left
		} else {
			src[right], src[index] = src[index], src[right]
			index = right
		}
	}
}
func upMin(src []int, index int) {
	if index <= 0 {
		return
	}
	for parent := getParent(index); parent >= 0 && src[parent] > src[index]; parent = getParent(index) {
		src[parent], src[index] = src[index], src[parent]
		index = parent
	}
}
func upMax(src []int, index int) {
	if index <= 0 {
		return
	}
	for parent := getParent(index); parent >= 0 && src[parent] < src[index]; parent = getParent(index) {
		src[parent], src[index] = src[index], src[parent]
		index = parent
	}
}

func getParent(index int) int {
	return (index - 1) >> 1
}
func getLeftChild(index int) int {
	return index<<1 + 1
}
func getRightChild(index int) int {
	return index<<1 + 2
}
