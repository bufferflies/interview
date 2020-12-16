package _206

import (
	"math"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnit(t *testing.T) {
	te := assert.New(t)
	sk := Constructor()
	sk.Add(1)
	sk.Add(2)
	sk.Add(3)
	te.False(sk.Search(0))
	sk.Add(4)
	te.True(sk.Search(1))
	te.False(sk.Erase(0))
	te.True(sk.Erase(1))
	te.False(sk.Search(1))
}

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
		left[i] = &Node{val: -1}
		right[i] = &Node{val: 10000}
	}
	// 从上到下 上面是第一层
	for i := maxLevel - 2; i >= 0; i-- {
		left[i].right = right[i]
		left[i].down = left[i+1]
		right[i].down = right[i+1]
	}
	left[maxLevel-1].right = right[maxLevel-1]
	return Skiplist{left[0]}
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
	lv := -1
	for ; head != nil; head = head.down {
		for head.right != nil && head.right.val < num {
			head = head.right
		}
		lv++
		pre[lv] = head
	}

	n := randLevel()
	arr := make([]*Node, n)
	tmp := &Node{val: -1}
	// 从上到下
	for i, a := range arr {
		p := pre[maxLevel-n+i]
		a = &Node{val: num, right: p.right}
		p.right = a
		tmp.down = a
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
