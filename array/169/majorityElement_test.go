package _69

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnit(t *testing.T) {
	arr := []int{1, 1, 2, 2, 2}
	assert.NotNil(t, 2, majorityElement(arr))
}

// majorityElement
func majorityElement(nums []int) int {
	candidate := nums[0]
	count := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == candidate {
			count++
		} else {
			count--
		}
		if count == 0 {
			candidate = nums[i]
			count = 1
		}

	}
	return candidate
}
