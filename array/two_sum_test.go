package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func twoSum(nums []int, target int) []int {
	ptr := 0
	expect := make(map[int]int, len(nums))
	for ptr < len(nums) {
		_, ok := expect[nums[ptr]]
		if ok {
			break
		} else {
			expect[target-nums[ptr]] = ptr
		}
		ptr++
	}
	return []int{expect[nums[ptr]], ptr}
}
func Test_twoSum(t *testing.T) {
	c := []int{3, 2, 4}
	result := twoSum(c, 6)
	assert.Equal(t, result, []int{1, 2})
}
