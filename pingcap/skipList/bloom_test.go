package skipList

import (
	"testing"

	"git.code.oa.com/geeker/awesome-work/pingcap/constant"

	"github.com/stretchr/testify/assert"
)

func TestUnit_Slot_Should_Codec(t *testing.T) {
	te := assert.New(t)
	key := []byte("test")
	s := NewSlot(constant.BloomM, constant.BloomK, 1, "../testdata/db/0/segment_0.bloom", "/")
	s.Add(key)
	te.True(s.Test(key))
	te.False(s.Test([]byte("fake")))

	s.Save()
	c := FromFile("../testdata/db/0/segment_0.bloom")
	te.True(c.Test(key))
	te.False(c.Test([]byte("fake")))
}
