package _21

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Max_Profit(t *testing.T) {
	arr := []int{7, 6, 4, 3, 1}
	result := maxProfit(arr)
	assert.Equal(t, 0, result)
}

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	result := 0
	for i, min := 1, 0; i < len(prices); i++ {
		profit := prices[i] - prices[min]
		if result < profit {
			result = profit
		}
		if prices[min] > prices[i] {
			min = i
		}
	}
	return result
}
