package _2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//样例输入：4 10 15 13
//样例输出：12
func TestUnit_compute(t *testing.T) {
	arr := []int{10, 15, 13}
	expect := 12
	r := compute(4, arr)
	assert.Equal(t, expect, r)
}

func compute(num int, array []int) int {
	left := 1
	right := power(num) - 1
	result := getMiddle(left, right)
	for tmp := findLeftOrRight(result, array); tmp != 0; {
		if tmp > 0 {
			left = result + 1
		} else if tmp < 0 {
			right = result - 1
		} else {
			break
		}
		tmp = findLeftOrRight(result, array)
		if tmp == 0 {
			break
		}
		result = getMiddle(left, right)
	}
	return result
}
func power(nums int) int {
	return 1 << nums
}

func getMiddle(left int, right int) int {
	return (left + right) / 2
}

// findLeftOrRight -1-->left 1:right  0:mix not same side
func findLeftOrRight(target int, array []int) int {
	count := 0
	for _, v := range array {
		if v > target {
			count++
		} else {
			count--
		}
	}
	if count == 3 {
		return 1
	} else if count == -3 {
		return -1
	} else {
		return 0
	}
}
