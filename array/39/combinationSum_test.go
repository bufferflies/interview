package _9

import (
	"fmt"
	"testing"
)

//给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
//
//candidates 中的数字可以无限制重复被选取。
//
//说明：
//
//所有数字（包括 target）都是正整数。
//解集不能包含重复的组合。
//示例 11：
//
//输入：candidates = [2,3,6,7], target = 7,
//所求解集为：
//[
//  [7],
//  [2,2,3]
//]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/combination-sum
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func Test_combinationSum(t *testing.T) {
	candidates := []int{2, 3, 6, 7}
	target := 7
	result := combinationSum(candidates, target)
	fmt.Printf("result:%v", result)
}

type helper struct {
	result [][]int
}

func combinationSum(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	h := helper{result: result}
	for i, v := range candidates {
		array := make([]int, 0)
		array = append(array, v)
		h.travel(candidates, i, target, v, array)
	}
	return h.result
}

func (h *helper) travel(candidates []int, ptr int, target int, sum int, array []int) {
	if sum == target {
		c := make([]int, len(array))
		copy(c, array)
		h.result = append(h.result, c)
		return
	}
	if sum > target {
		return
	}

	for i := ptr; i < len(candidates); i = i + 1 {
		sum = sum + candidates[i]
		array = append(array, candidates[i])
		h.travel(candidates, i, target, sum, array)
		sum = sum - array[len(array)-1]
		array = array[:len(array)-1]
	}
}
