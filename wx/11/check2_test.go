package _1

import (
	"math/rand"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

//输入：4，[0,1] , 返回 2|3
func TestUnit_pick(t *testing.T) {
	black := []int{0, 1}
	r := pick(4, black)
	assert.True(t, equal([]int{2, 3}, r))
}
func equal(src, target []int) bool {
	sort.Ints(target)
	for i := range src {
		if src[i] != target[i] {
			return false
		}
	}
	return true
}

func pick(n int, black []int) []int {
	dick := make(map[int]interface{})
	for _, v := range black {
		dick[v] = nil
	}
	result := make([]int, n-len(black))
	for i := 0; i < n-len(black); i++ {
		j := randM(dick)
		result[i] = j
		dick[j] = nil
	}
	return result
}

// todo effective
// O(1) O(10)-->O(8)
// O(10)-->O(7) (1->7,2->8,3->9)

// 1. O(10) 2.(10) 10X10%8=4
// 1. 4X10 ==>
func randM(m map[int]interface{}) int {
	r := randN()
	_, flag := m[r]
	for flag {
		r = randN()
		_, flag = m[r]
	}
	return r
}

// randN
func randN() int {
	return rand.Intn(4)
}
