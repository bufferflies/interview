package _99

import "git.code.oa.com/geeker/awesome-work/util"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	head := root
	queue := util.NewMyQueue()
	queue.Enque(head)
	for queue.Size() > 0 {
		count := queue.Size()
		for i := 0; i < count; i++ {
			n := queue.Deque().(*TreeNode)
			if n.Left != nil {
				queue.Enque(n.Left)
			}
			if n.Right != nil {
				queue.Enque(n.Right)
			}
			if i == count-1 {
				result = append(result, n.Val)
			}
		}

	}
	return result
}
