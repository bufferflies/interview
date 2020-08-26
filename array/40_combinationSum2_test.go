package array

import (
	"fmt"
	"sort"
	"testing"
)

//给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
//
//candidates 中的每个数字在每个组合中只能使用一次。
//
//说明：
//
//所有数字（包括目标数）都是正整数。
//解集不能包含重复的组合。
//示例 1:
//
//输入: candidates = [10,1,2,7,6,1,5], target = 8,
//所求解集为:
//[
//  [1, 7],
//  [1, 2, 5],
//  [2, 6],
//  [1, 1, 6]
//]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/combination-sum-ii
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func Test_combinationSum2(t *testing.T) {
	candidates := []int{2, 5, 2, 1, 2}
	target := 5
	result := combinationSum2(candidates, target)
	fmt.Printf("result:%v", result)
}

type helper2 struct {
	result [][]int
}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	result := make([][]int, 0)

	h := helper2{result: result}
	before := -1
	for i, v := range candidates {
		if v == before {
			continue
		}
		before = v
		array := make([]int, 0)
		array = append(array, v)
		h.travel(candidates, i, target, v, array)
	}
	return h.result
}

func (h *helper2) travel(candidates []int, ptr int, target int, sum int, array []int) {
	if sum == target {
		c := make([]int, len(array))
		copy(c, array)
		h.result = append(h.result, c)
		return
	}
	if sum > target {
		return
	}
	before := -1
	for i := ptr + 1; i < len(candidates); i = i + 1 {
		if candidates[i] == before {
			continue
		}
		before = candidates[i]
		sum = sum + candidates[i]
		array = append(array, candidates[i])
		h.travel(candidates, i, target, sum, array)
		sum = sum - array[len(array)-1]
		array = array[:len(array)-1]
	}
}
