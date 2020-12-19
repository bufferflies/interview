package skipList

import (
	"math"
	"math/rand"
)

// SkipList
type SkipList struct {
	head     *Node
	tail     *Node
	size     int
	maxLevel int
	maxRand  int
}

// Node
type Node struct {
	right *Node
	down  *Node
	key   uint64
	val   interface{}
}

func Constructor(maxLevel int) SkipList {
	left := make([]*Node, maxLevel)
	right := make([]*Node, maxLevel)
	for i := 0; i < maxLevel; i++ {
		left[i] = &Node{key: 0}
		right[i] = &Node{key: 1<<64 - 1}
	}
	// 第0层为最下层
	for i := 0; i < maxLevel; i++ {
		left[i].right = right[i]
		if i != 0 {
			left[i].down = left[i-1]
			right[i].down = right[i-1]
		}
	}
	return SkipList{
		head:     left[maxLevel-1],
		tail:     left[0],
		size:     0,
		maxLevel: maxLevel,
		maxRand:  1<<maxLevel - 1,
	}
}
func (this *SkipList) ToValues() []interface{} {
	tail := this.tail.right
	ret := make([]interface{}, this.size)
	for i := 0; tail != nil && tail.right != nil; i++ {
		ret[i] = tail.val
		tail = tail.right
	}
	return ret
}

func (this *SkipList) randLevel() int {
	return this.maxLevel - int(math.Log2(float64(1.0+float64(this.maxRand)*rand.Float64())))
}
func (this *SkipList) Search(target uint64) bool {
	head := this.head
	for head != nil {
		if head.right != nil {
			if head.right.key < target {
				head = head.right
			} else if head.right.key == target {
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
func (this *SkipList) Add(num uint64, val interface{}) {
	pre := make([]*Node, this.maxLevel)
	head := this.head
	// 找到每次最小于num的node
	for lv := this.maxLevel; head != nil; head = head.down {
		for head.right != nil && head.right.key < num {
			head = head.right
		}
		lv--
		pre[lv] = head
	}

	n := this.randLevel()
	arr := make([]*Node, n)
	tmp := &Node{key: 0}
	// 从上面开始创建
	for i, a := range arr {
		p := pre[i]
		a = &Node{key: num, val: val, right: p.right}
		p.right = a
		// 0层不应该有向下指针
		if i != 0 {
			a.down = tmp
		}
		tmp = a
	}
	this.size++
}

func (this *SkipList) Erase(num uint64) bool {
	ret := false
	head := this.head
	for head != nil {
		if head.right != nil {
			if head.right.key < num {
				head = head.right
			} else if head.right.key == num {
				head.right = head.right.right
				head = head.down
				ret = true
				this.size--
			} else {
				head = head.down
			}
		} else {
			head = head.down
		}
	}
	return ret
}
