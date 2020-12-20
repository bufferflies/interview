package skipList

import (
	"testing"

	"git.code.oa.com/geeker/awesome-work/pingcap/domain"

	"github.com/stretchr/testify/assert"
)

func TestUnit(t *testing.T) {
	te := assert.New(t)
	sk := Constructor(4)
	sk.Add(1, &domain.Entry{Hash: 1})
	sk.Add(2, &domain.Entry{Hash: 2})
	sk.Add(3, &domain.Entry{Hash: 3})
	te.False(sk.Search(0))
	sk.Add(4, &domain.Entry{Hash: 4})
	sk.Add(5, &domain.Entry{Hash: 5})
	te.True(sk.Search(1))
	te.False(sk.Erase(0))
	te.True(sk.Erase(1))
	te.False(sk.Search(1))
	arr := sk.ToValues()
	expect := []int{2, 3, 4, 5}
	for i, v := range arr {
		e := v.(*domain.Entry)
		te.Equal(expect[i], int(e.Hash))
	}

}
