package _86

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnit_intervalIntersection(t *testing.T) {
	A := [][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}}
	B := [][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}}
	r := intervalIntersection(A, B)
	expect := [][]int{{1, 2}, {5, 5}, {8, 10}, {15, 23}, {24, 24}, {25, 25}}
	assert.Equal(t, expect, r)
}

func intervalIntersection(A [][]int, B [][]int) [][]int {
	result := make([][]int, 0)
	for ptrA, ptrB := 0, 0; ptrA < len(A) && ptrB < len(B); {
		i := max(A[ptrA][0], B[ptrB][0])
		j := min(A[ptrA][1], B[ptrB][1])
		if i <= j {
			result = append(result, []int{i, j})
		}
		if A[ptrA][1] < B[ptrB][1] {
			ptrA++
		} else {
			ptrB++
		}
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
