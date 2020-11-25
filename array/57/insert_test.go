package _7

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//给出一个无重叠的 ，按照区间起始端点排序的区间列表。
//
//在列表中插入一个新的区间，你需要确保列表中的区间仍然有序且不重叠（如果有必要的话，可以合并区间）。
//
//
//
//示例 11：
//
//输入：intervals = [[11,3],[6,9]], newInterval = [2,5]
//输出：[[11,5],[6,9]]
//示例 2：
//
//输入：intervals = [[11,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
//输出：[[11,2],[3,10],[12,16]]
//解释：这是因为新的区间 [4,8] 与 [3,5],[6,7],[8,10] 重叠。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/insert-interval
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func Test_insert(t *testing.T) {
	intervals := [][]int{{1, 3}, {6, 9}}
	newInterval := []int{4, 8}
	r := insert(intervals, newInterval)
	assert.Equal(t, []int{4, 8}, r)
}

type indexHelper struct {
	index int
	less  bool
}

func insert(intervals [][]int, newInterval []int) [][]int {
	index := findIndex(intervals, newInterval)
	result := intervals[:index[0].index-1]
	// 在区间内
	if index[0].less {
		result[index[0].index][0] = newInterval[0]
	}
	tmp := intervals[index[1].index:]
	if index[0].less {
		result[index[1].index-1][1] = newInterval[1]
	}
	result = append(result, tmp...)
	return result
}
func findIndex(intervals [][]int, newInterval []int) []*indexHelper {
	result := make([]*indexHelper, 2)
	lower := newInterval[0]
	high := newInterval[1]
	if lower < intervals[0][0] {
		result[0] = &indexHelper{index: 0, less: true}
	} else {
		for i := 0; i < len(intervals); i++ {
			if lower > intervals[i][0] && lower < intervals[i][1] {
				result[0] = &indexHelper{index: i, less: false}
				break
			}
			//不在区间内
			if i > 0 && lower > intervals[i-1][1] && lower < intervals[i][0] {
				result[0] = &indexHelper{index: i, less: true}
				break
			}
		}
	}

	for i := 1; i < len(intervals); i++ {
		if high > intervals[i][0] && high < intervals[i][1] {
			result[1] = &indexHelper{index: i, less: false}
			break
		}
		if high > intervals[i][1] && high < intervals[i+1][0] {
			result[1] = &indexHelper{index: i, less: true}
			break
		}
	}

	return result
}
