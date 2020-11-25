package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Merge struct {
}

func TestUnit_Merge(t *testing.T) {
	b := &Merge{}
	src := []int{5, 3, 2, 1, 4}
	expect := []int{1, 2, 3, 4, 5}
	result := b.sort(src)
	assert.Equal(t, expect, result)
}
func (b Merge) sort(src []int) []int {
	if len(src) <= 1 {
		return src
	}
	middle := len(src) / 2
	left := b.sort(src[:middle])
	right := b.sort(src[middle:])
	return b.merge(left, right)
}

func (b Merge) merge(left []int, right []int) []int {
	result := make([]int, 0)
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	if i != len(left) {
		result = append(result, left[i:]...)
	}
	if j != len(right) {
		result = append(result, right[j:]...)
	}
	return result
}
