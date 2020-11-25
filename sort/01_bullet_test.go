package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Sort interface {
	sort(src []int) []int
}
type Bullet struct {
}

func TestUnit_Bullet(t *testing.T) {
	b := &Bullet{}
	src := []int{5, 3, 2, 1, 4}
	expect := []int{1, 2, 3, 4, 5}
	result := b.sort(src)
	assert.Equal(t, expect, result)
}
func (b Bullet) sort(src []int) []int {
	for i := 0; i < len(src); i++ {
		for j := 0; j < len(src)-i-1; j++ {
			if src[j] > src[j+1] {
				src[j], src[j+1] = src[j+1], src[j]
			}
		}
	}
	return src
}
