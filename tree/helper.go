package tree

// 树结构
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 寻找数组
func FindIndex(arr []int, target int) (index int) {
	for i, v := range arr {
		if v == target {
			return i
		}
	}
	return -1
}
