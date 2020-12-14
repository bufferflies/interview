package _09

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnit(t *testing.T) {
	nums := []int{3, 4, 5}
	r := minSubArrayLen(11, nums)
	assert.Equal(t, 3, r)
}
func TestUnit_One(t *testing.T) {
	nums := []int{5}
	r := minSubArrayLen(3, nums)
	assert.Equal(t, 1, r)
}

func minSubArrayLen(s int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	result, tmp := 0, nums[0]
	for head, tail := 0, 0; tail < len(nums); {
		if tmp >= s {
			if result == 0 {
				result = tail - head + 1
			} else {
				result = min(result, tail-head+1)
			}
			tmp = tmp - nums[head]
			head++

		} else {
			tail++
			if tail < len(nums) {
				tmp = tmp + nums[tail]
			}
		}
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
