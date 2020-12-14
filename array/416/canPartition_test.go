package _16

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnit(t *testing.T) {
	arr := []int{1, 5, 11, 5}
	r := canPartition(arr)
	assert.True(t, r)
}

func TestUnit_11(t *testing.T) {
	arr := []int{14, 2, 8, 2, 4}
	r := canPartition(arr)
	assert.False(t, r)
}
func TestUnit_1234(t *testing.T) {
	arr := []int{1, 2, 3, 5}
	r := canPartition(arr)
	assert.False(t, r)
}

func canPartition(nums []int) bool {
	sum, ok := validate(nums)
	if !ok {
		return ok
	}
	target := sum>>1 + 1
	// dp[i][j]: 表示总重量等于j情况下，i是否满足
	dp := make([][]bool, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]bool, target)
	}
	// 1.初始化 一个都不装 j=0
	// i=0  一个都不选都是false
	for i := 0; i < len(nums); i++ {
		dp[i][0] = true
	}
	dp[0][nums[0]] = true
	for i := 1; i < len(nums); i++ {
		v := nums[i]
		for j := 1; j < target; j++ {
			if j >= v {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-v]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[len(nums)-1][target-1]
}

// validate
func validate(nums []int) (int, bool) {
	if len(nums) <= 1 {
		return 0, false
	}
	sum, max := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		sum = sum + nums[i]
		if max < nums[i] {
			max = nums[i]
		}
	}

	if sum&0x01 == 1 {
		return 0, false
	}
	ave := sum >> 1
	if max > ave {
		return 0, false
	}
	return sum, true
}
