package skipList

import (
	"math"
	"math/rand"
)

const (
	maxLevel = 4
	maxRand  = 15
)

func randLevel() int {
	return maxLevel - int(math.Log2(1.0+maxRand*rand.Float64()))
}

// Skiplist
type Skiplist struct {
	head *Node
}

// Node
type Node struct {
	right *Node
	down  *Node
	val   int
}

func Constructor() Skiplist {
	left := make([]*Node, maxLevel)
	right := make([]*Node, maxLevel)
	for i := 0; i < maxLevel; i++ {
		left[i] = &Node{val: math.MinInt16}
		right[i] = &Node{val: math.MaxInt16}
	}
	// 第0层为最下层
	for i := 0; i < maxLevel; i++ {
		left[i].right = right[i]
		if i != 0 {
			left[i].down = left[i-1]
			right[i].down = right[i-1]
		}

	}
	return Skiplist{left[maxLevel-1]}
}

func (this *Skiplist) Search(target int) bool {
	head := this.head
	for head != nil {
		if head.right != nil {
			if head.right.val < target {
				head = head.right
			} else if head.right.val == target {
				return true
			} else {
				head = head.down
			}
		} else {
			head = head.down
		}
	}
	return false
}

func (this *Skiplist) Add(num int) {
	pre := make([]*Node, maxLevel)
	head := this.head
	// 找到每次最小于num的node
	for lv := maxLevel; head != nil; head = head.down {
		for head.right != nil && head.right.val < num {
			head = head.right
		}
		lv--
		pre[lv] = head
	}

	n := randLevel()
	arr := make([]*Node, n)
	tmp := &Node{val: math.MinInt16}
	// 从上面开始创建
	for i, a := range arr {
		p := pre[i]
		a = &Node{val: num, right: p.right}
		p.right = a
		// 0层不应该有向下指针
		//t := tmp
		if i != 0 {
			a.down = tmp
		}
		tmp = a
	}
}

func (this *Skiplist) Erase(num int) bool {
	ret := false
	head := this.head
	for head != nil {
		if head.right != nil {
			if head.right.val < num {
				head = head.right
			} else if head.right.val == num {
				head.right = head.right.right
				head = head.down
				ret = true
			} else {
				head = head.down
			}
		} else {
			head = head.down
		}
	}
	return ret
}
