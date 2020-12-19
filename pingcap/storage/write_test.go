package storage

import (
	"testing"

	"git.code.oa.com/geeker/awesome-work/pingcap/domain"

	"github.com/stretchr/testify/assert"
)

/**
segment_1_0.index:
0 0
0 -1
1 1
2 2
3 3
4 4
5 5
6 6
7 7
*/
func TestUnit_Should_Search_Key(t *testing.T) {
	te := assert.New(t)
	file := "./segment_1_0.index"
	r, err := Search(file, uint64(0))
	te.Nil(err)
	te.True(len(r)%16 == 0)
	te.True(len(r)>>4 == 2)
	index := &domain.Index{}
	index.Decode(r)

	te.Equal(2, len(index.Entries))
	te.Equal(uint64(0), index.Entries[1].Hash)
	te.Equal(int64(-1), index.Entries[1].Offset)
	te.Equal(uint64(0), index.Entries[0].Hash)
	te.Equal(int64(0), index.Entries[0].Offset)

	r, err = Search(file, uint64(10))
	te.NotNil(err)

	r, err = Search(file, uint64(4))
	te.Nil(err)
	index = &domain.Index{}
	index.Decode(r)
	te.Equal(1, len(index.Entries))
	te.Equal(uint64(4), index.Entries[0].Hash)
	te.Equal(int64(4), index.Entries[0].Offset)
}
