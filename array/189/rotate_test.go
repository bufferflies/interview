package _89

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnit_Rotate(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	rotate(nums, 3)
	assert.Equal(t, []int{2, 3, 4, 1}, nums)
}
func TestUnit_N_Less_K(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 3)
	assert.Equal(t, []int{5, 6, 7, 1, 2, 3, 4}, nums)
}
func rotate(nums []int, k int) {
	length := len(nums)
	// 断裂的地址
	step := length - k%length
	tmp := make([]int, length-step)
	copy(tmp, nums[step:])
	for i := step - 1; i >= 0; i-- {
		nums[i+k] = nums[i]
	}
	for i := 0; i < len(tmp); i++ {
		nums[i] = tmp[i]
	}
}
