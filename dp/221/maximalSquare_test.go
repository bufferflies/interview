package _21

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnit_max(t *testing.T) {
	matrix := [][]byte{{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'}}
	r := maximalSquare(matrix)
	assert.Equal(t, 4, r)
}
func maximalSquare(matrix [][]byte) int {
	dp := make([][]int, len(matrix))
	result := 0
	for i := range dp {
		dp[i] = make([]int, len(matrix[0]))

	}
	for i := 0; i < len(matrix); i++ {
		dp[i][0] = int(matrix[i][0]) - '0'
		if result == 0 && dp[i][0] == 1 {
			result = 1
		}
	}
	for j := 0; j < len(matrix[0]); j++ {
		dp[0][j] = int(matrix[0][j]) - '0'
		if result == 0 && dp[0][j] == 1 {
			result = 1
		}
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][j] == '0' {
				dp[i][j] = 0
			} else {
				dp[i][j] = Min(dp[i-1][j], Min(dp[i][j-1], dp[i-1][j-1])) + 1
				result = Max(dp[i][j], result)
			}

		}
	}
	return result * result
}
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
