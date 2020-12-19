package skipList

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnit(t *testing.T) {
	te := assert.New(t)
	sk := Constructor(4)
	sk.Add(1, 1)
	sk.Add(2, 2)
	sk.Add(3, 3)
	te.False(sk.Search(0))
	sk.Add(4, 4)
	sk.Add(5, 5)
	te.True(sk.Search(1))
	te.False(sk.Erase(0))
	te.True(sk.Erase(1))
	te.False(sk.Search(1))
	arr := sk.ToValues()
	expect := []interface{}{2, 3, 4, 5}
	te.Equal(expect, arr)
}
