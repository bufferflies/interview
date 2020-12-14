package _6

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

//给出一个区间的集合，请合并所有重叠的区间

func TestUnit(t *testing.T) {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	expect := [][]int{{1, 6}, {8, 10}, {15, 18}}
	r := merge(intervals)
	assert.Equal(t, expect, r)

}
func TestUnit_sort(t *testing.T) {
	intervals := [][]int{{1, 4}, {0, 4}}
	expect := [][]int{{0, 4}}
	r := merge(intervals)
	assert.Equal(t, expect, r)
	intervals = [][]int{{1, 4}, {2, 3}}
	expect = [][]int{{1, 4}}
	r = merge(intervals)
	assert.Equal(t, expect, r)

}
func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	sort.Sort(IntArraySlice(intervals))
	result := make([][]int, 0)
	tmp := intervals[0]
	for i := 1; i < len(intervals); i++ {
		sub, ok := isSub(tmp, intervals[i])
		if ok {
			tmp = sub
		} else {
			result = append(result, tmp)
			tmp = intervals[i]
		}
	}
	result = append(result, tmp)
	return result

}

type IntArraySlice [][]int

func (this IntArraySlice) Len() int {
	return len(this)
}

func (this IntArraySlice) Less(i, j int) bool {
	if this[i][0] < this[j][0] {
		return true
	} else if this[i][0] == this[j][0] && this[i][1] <= this[j][1] {
		return true
	}
	return false
}

func (this IntArraySlice) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

// 判断是否存在交集
func isSub(a, b []int) ([]int, bool) {
	left := max(a[0], b[0])
	right := min(a[1], b[1])
	if left < right {
		return []int{a[0], max(b[1], a[1])}, true
	} else {
		return b, false
	}
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
