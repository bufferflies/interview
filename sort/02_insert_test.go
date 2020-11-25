package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Insert struct {
}

func TestUnit_Insert(t *testing.T) {
	b := &Insert{}
	src := []int{5, 3, 2, 1, 4}
	expect := []int{1, 2, 3, 4, 5}
	result := b.sort(src)
	assert.Equal(t, expect, result)
}
func (b Insert) sort(src []int) []int {
	for i := 0; i < len(src); i++ {
		for j := i; j > 0; j-- {
			if src[j] < src[j-1] {
				src[j], src[j-1] = src[j-1], src[j]
			}
		}
	}
	return src
}
