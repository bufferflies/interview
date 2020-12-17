package skipList

import (
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
	sk.Add(5)
	te.True(sk.Search(1))
	te.False(sk.Erase(0))
	te.True(sk.Erase(1))
	te.False(sk.Search(1))
}
