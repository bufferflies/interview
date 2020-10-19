package _23

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Max_Profit(t *testing.T) {
	arr := []int{3, 3, 5, 0, 0, 3, 1, 4}
	result := maxProfit(arr)
	assert.Equal(t, 6, result)
}
func maxProfit(prices []int) int {
	forward := maxForwardProfit(prices)
	back := maxBackProfit(prices)
	result := 0
	for i := range prices {
		profit := forward[i] + back[i]
		if profit > result {
			result = profit
		}
	}
	return result
}
func maxBackProfit(prices []int) []int {
	result := make([]int, len(prices))
	for max, i := prices[len(prices)-1], len(prices)-2; i >= 0; i-- {
		result[i] = max - prices[i]
		if result[i] < result[i+1] {
			result[i] = result[i+1]
		}
		if max <= prices[i] {
			max = prices[i]
		}
	}
	return result
}
func maxForwardProfit(prices []int) []int {
	result := make([]int, len(prices)+1)
	for i, min := 1, prices[0]; i < len(prices); i++ {
		result[i] = prices[i] - min
		if result[i] < result[i-1] {
			result[i] = result[i-1]
		}
		if min > prices[i] {
			min = prices[i]
		}
	}
	return result
}
