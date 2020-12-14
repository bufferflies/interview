package bytedance

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnit(t *testing.T) {
	i, j := 1, 3
	expect := "0.(3)"
	r := mod(i, j)
	assert.Equal(t, expect, r)
}
func TestUnit_7(t *testing.T) {
	i, j := 1, 7
	expect := "0.(142857)"
	r := mod(i, j)
	assert.Equal(t, expect, r)
}
func TestUnit_big(t *testing.T) {
	i, j := 1, 13
	expect := "0.(076923)"
	r := mod(i, j)
	assert.Equal(t, expect, r)
}
func TestUnit_0(t *testing.T) {
	i, j := 1, 2
	expect := "0.5"
	r := mod(i, j)
	assert.Equal(t, expect, r)
}

// 1/7
// 0,1
// 10/7 ==> 1,3
// 30/7==> 4,2
func mod(i, j int) string {
	dick := make(map[int]int)
	result := make([]byte, 0)
	k, m := compute(i, j)
	result = append(result, []byte(strconv.Itoa(k))...)
	if m == 0 {
		return string(result)
	}
	result = append(result, '.')
	index := len(result)
	for {
		max := 0
		for m = m * 10; m <= j; {
			max++
			m = m * 10
		}
		_, ok := dick[m]
		if ok {
			break
		}
		for l := 0; l < max; l++ {
			result = append(result, '0')
		}
		dick[m] = index - 1
		k, m = compute(m, j)
		result = append(result, []byte(strconv.Itoa(k))...)
		if m == 0 {
			break
		}
		index = len(result)
	}
	if m == 0 {
		return string(result)
	}
	result = append(result, ')')
	return insert(result, dick[m])
}

// charu
func insert(str []byte, index int) string {
	return string(str[:index+1]) + "(" + string(str[index+1:])
} // compute
func compute(i, j int) (k, m int) {
	k = i / j
	m = i % j
	return k, m
}
