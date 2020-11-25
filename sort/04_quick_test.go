package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Quick struct {
}

func TestUnit_Quick(t *testing.T) {
	b := &Quick{}
	src := []int{5, 3, 2, 1, 4}
	expect := []int{1, 2, 3, 4, 5}
	result := b.sort(src)
	assert.Equal(t, expect, result)
}
func (b Quick) sort(src []int) []int {
	spilt(src, 0, len(src)-1)
	return src
}
func spilt(src []int, start int, end int) {
	if end-start <= 1 {
		return
	}
	p := partition(src, start, end)

	spilt(src, start, p-1)
	spilt(src, p+1, end)

}
func partition(src []int, start int, end int) int {
	sentinel := start
	for i := start; i < end; i++ {
		if src[i] < src[end] {
			src[i], src[sentinel] = src[sentinel], src[i]
			sentinel++
		}
	}
	src[end], src[sentinel] = src[sentinel], src[end]
	return sentinel
}
