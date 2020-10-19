package _52

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnit_Max_Product(t *testing.T) {
	nums := []int{-1, -2, -9, -6}
	result := maxProduct(nums)
	assert.Equal(t, 108, result)
}

func maxProduct(nums []int) int {
	maxProduct := nums[0]
	minProduct := nums[0]
	result := nums[0]
	for i := 1; i < len(nums); i++ {
		tmp := max(nums[i], maxProduct*nums[i], minProduct*nums[i])
		minProduct = min(nums[i], maxProduct*nums[i], minProduct*nums[i])
		maxProduct = tmp
		result = max(result, maxProduct)
	}
	return result
}
func max(nums ...int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if max < nums[i] {
			max = nums[i]
		}
	}
	return max
}
func min(nums ...int) int {
	min := nums[0]
	for i := 1; i < len(nums); i++ {
		if min > nums[i] {
			min = nums[i]
		}
	}
	return min
}
