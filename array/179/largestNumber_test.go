package _79

import (
	"sort"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnit_largestNumber(t *testing.T) {
	nums := []int{10, 2}
	r := largestNumber(nums)
	assert.Equal(t, "210", r)
}

type IntSlice []int

func (c IntSlice) Len() int {
	return len(c)
}
func (c IntSlice) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
func (c IntSlice) Less(i, j int) bool {
	a := strconv.Itoa(c[i])
	b := strconv.Itoa(c[j])
	return a > b
}
func largestNumber(nums []int) string {
	if isZero(nums) {
		return "0"
	}
	//arr := newIntSlice()
	sort.Sort(IntSlice(nums))
	var b strings.Builder
	for _, v := range nums {
		b.WriteString(strconv.Itoa(v))
	}
	return b.String()
}
func isZero(nums []int) bool {
	for _, v := range nums {
		if v != 0 {
			return false
		}
	}
	return true
}
