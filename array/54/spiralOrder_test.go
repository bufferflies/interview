package _4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//给定一个包含 m x n 个元素的矩阵（m 行, n 列），请按照顺时针螺旋顺序，返回矩阵中的所有元素。
//
//示例 11:
//
//输入:
//[
// [ 11, 2, 3 ],
// [ 4, 5, 6 ],
// [ 7, 8, 9 ]
//]
//输出: [11,2,3,6,9,8,7,4,5]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/spiral-matrix
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func Test_spiralOrderspiralOrder(t *testing.T) {
	matrix := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	result := spiralOrder(matrix)
	assert.Equal(t, []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}, result)
}
func Test_spiralOrderspiralOrder_column(t *testing.T) {
	matrix := [][]int{{6}, {9}, {7}}
	result := spiralOrder(matrix)
	assert.Equal(t, []int{6, 9, 7}, result)
}
func Test_spiralOrderspiralOrder_row(t *testing.T) {
	matrix := [][]int{{6, 9, 7}}
	result := spiralOrder(matrix)
	assert.Equal(t, []int{6, 9, 7}, result)
}

type Point struct {
	x         int
	y         int
	direction *Point
}

var (
	left    = &Point{x: -1, y: 0}
	right   = &Point{x: 1, y: 0}
	down    = &Point{x: 0, y: 1}
	up      = &Point{x: 0, y: -1}
	forward = &Point{x: 1, y: 1}
	back    = &Point{x: -1, y: -1}
)

func spiralOrder(matrix [][]int) []int {
	if len(matrix) < 1 {
		return nil
	}
	head := &Point{0, 0, forward}
	tail := &Point{len(matrix[0]) - 1, len(matrix) - 1, back}
	result := make([]int, 0)
	current := &Point{0, 0, right}
	for stop(head, tail) {
		r := travel(matrix, current, head, tail)
		result = append(result, r...)
		head.add()
		tail.add()
		current.direction = forward
		current.add()
		current.direction = right
	}
	if head.x == tail.x && head.y == tail.y {
		result = append(result, matrix[current.y][current.x])
	}

	return result
}

func travel(matrix [][]int, current *Point, head *Point, tail *Point) (result []int) {
	result = make([]int, 0)
	for current.x < tail.x {
		result = append(result, matrix[current.y][current.x])
		current.add()
	}
	result = append(result, matrix[current.y][current.x])
	// 只有一行
	if head.y == tail.y {
		return result
	}
	current.direction = down
	current.add()
	for current.y < tail.y {
		result = append(result, matrix[current.y][current.x])
		current.add()
	}
	// 只有一列
	result = append(result, matrix[current.y][current.x])
	if head.x == tail.x {
		return result
	}
	current.direction = left
	current.add()
	for current.x > head.x {
		result = append(result, matrix[current.y][current.x])
		current.add()
	}
	result = append(result, matrix[current.y][current.x])
	current.direction = up
	current.add()
	for current.y > head.y {
		result = append(result, matrix[current.y][current.x])
		current.add()
	}
	return result
}
func (p *Point) add() {
	p.x = p.x + p.direction.x
	p.y = p.y + p.direction.y
}

func stop(p1, p2 *Point) bool {
	if p1.x <= p2.x && p1.y <= p2.y {
		return true
	}
	return false
}
