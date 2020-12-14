package _1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnit_check(t *testing.T) {
	str := "(]"
	assert.True(t, check(str))
}

//（（））（）（（）（）（））（）
func TestUnit_check_complex(t *testing.T) {
	str := "(())()(()()())()"
	assert.True(t, check(str))
}

func check(target string) bool {
	filter := make(map[node]interface{})
	bn := NewBigNode()
	mn := NewMiddleNode()
	filter[bn] = nil
	filter[mn] = nil
	for i := range target {
		for k := range filter {
			if k.filter(target[i]) {
				if k.count() < 0 {
					return false
				}
				break
			}
		}
	}

	return true
}

type node interface {
	filter(c byte) bool
	count() int
}
type bigNode struct {
	dict map[byte]int
	size int
}

func NewBigNode() *bigNode {
	m := make(map[byte]int)
	m['('] = 1
	m[')'] = -1
	return &bigNode{dict: m}
}
func (b *bigNode) filter(c byte) bool {
	if v, ok := b.dict[c]; ok {
		b.size = b.size + v
		return true
	}
	return false
}
func (b *bigNode) count() int {
	return b.size
}

type middleNode struct {
	dict map[byte]int
	size int
}

func NewMiddleNode() *middleNode {
	m := make(map[byte]int)
	m['['] = 1
	m[']'] = -1
	return &middleNode{dict: m}
}

func (m *middleNode) filter(c byte) bool {
	if v, ok := m.dict[c]; ok {
		m.size = m.size + v
		return true
	}
	return false
}
func (m *middleNode) count() int {
	return m.size
}
