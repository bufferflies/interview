package _7

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnit(t *testing.T) {
	nums := []int{1, 2, 3, 1}
	nextPermutation(nums)
	assert.Equal(t, []int{1, 3, 1, 2}, nums)
	nums = []int{3, 2, 1}
	nextPermutation(nums)
	assert.Equal(t, []int{1, 2, 3}, nums)

}
func nextPermutation(nums []int) {
	reorder := -1
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			reorder = i
			break
		}
	}
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			reorder = i
			break
		}
	}
	if reorder >= 0 {
		for j := len(nums) - 1; j > reorder; j-- {
			if nums[j] > nums[reorder] {
				nums[j], nums[reorder] = nums[reorder], nums[j]
				break
			}
		}
	}

	reverse(nums[reorder+1:])
}

func reverse(nums []int) {
	length := len(nums)
	for i := 0; i < length/2; i++ {
		nums[i], nums[length-i-1] = nums[length-i-1], nums[i]
	}
}
