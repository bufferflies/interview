package _9

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	direct = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
)

func TestUnit_generateMatrix(t *testing.T) {
	matrix := generateMatrix(2)
	assert.Equal(t, [][]int{{1, 2}, {4, 3}}, matrix)
}
func TestUnit_generateMatrix_3(t *testing.T) {
	matrix := generateMatrix(3)
	assert.Equal(t, [][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}}, matrix)
}
func TestUnit_generateMatrix_1(t *testing.T) {
	matrix := generateMatrix(1)
	assert.Equal(t, [][]int{{1}}, matrix)
}
func generateMatrix(n int) [][]int {
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
	}
	start := 1
	for i, j := n-1, 0; i > 0; {
		round(result, i, j, start)
		start = start + i*4
		j++
		i = i - 2

	}
	if n&0x01 == 1 {
		j := n >> 1
		result[j][j] = start
	}
	return result
}

func round(nums [][]int, n, j, start int) {
	position := []int{j, j}
	for j := 0; j < len(direct); j++ {
		for i := 0; i < n; i++ {
			nums[position[0]][position[1]] = start
			position[0], position[1] = position[0]+direct[j][0], position[1]+direct[j][1]
			start++
		}
	}
}
