package _2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/***
接雨水
*/

func TestUnit(t *testing.T) {
	height := []int{1, 0, 2}
	result := trap(height)
	assert.Equal(t, 1, result)
}
func trap(height []int) int {
	left := LeftView(height)
	right := RightView(height)
	result := 0
	for i := 0; i < len(height); i++ {
		result = result + min(left[i], right[i]) - height[i]
	}
	return result
}

func LeftView(height []int) []int {
	result := make([]int, len(height))
	result[0] = height[0]
	for i := 1; i < len(height); i++ {
		result[i] = max(result[i-1], height[i])
	}
	return result
}

func RightView(height []int) []int {
	result := make([]int, len(height))
	result[len(height)-1] = height[len(height)-1]
	for i := len(height) - 2; i >= 0; i-- {
		result[i] = max(result[i+1], height[i])
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
	if a > b {
		return b
	}
	return a
}
