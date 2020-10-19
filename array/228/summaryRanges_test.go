package _28

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestUnit_summaryRanges(t *testing.T) {
	nums := []int{0, 2, 3, 4, 6, 8, 9}
	result := summaryRanges(nums)
	expect := []string{"0", "2->4", "6", "8->9"}
	assert.Equal(t, expect, result)
}
func TestUnit_summaryRanges_single(t *testing.T) {
	nums := []int{1, 2}
	result := summaryRanges(nums)
	expect := []string{"1->2"}
	assert.Equal(t, expect, result)
}
func summaryRanges(nums []int) []string {
	// 处理异常情况
	if len(nums) == 1 {
		it := strconv.Itoa(nums[0])
		return []string{it}
	}
	if len(nums) == 0 {
		return []string{}
	}
	result := make([]string, 0)
	for i, pre := 0, 0; i < len(nums); i = i + 1 {
		if i != len(nums)-1 && nums[i+1]-nums[i] == 1 {
			continue
		} else {
			var it string
			if pre != i {
				it = fmt.Sprintf("%d->%d", nums[pre], nums[i])
			} else {
				it = fmt.Sprintf("%d", nums[pre])
			}
			result = append(result, it)
			pre = i + 1
		}
	}

	return result
}
