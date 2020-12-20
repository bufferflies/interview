package skipList

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnit_Slot_Should_Codec(t *testing.T) {
	te := assert.New(t)
	key := []byte("test")
	s := NewSlot(4, 1, 1, "../testdata/slot/1/1.slot")
	s.Add(key)
	te.True(s.Test(key))
	te.True(s.Test([]byte("fake")))

	s.Save()
	c := NewSlot(4, 1, 1, "../testdata/slot/1/1.slot")
	c.Load()
	te.Equal(c, s)
	//os.RemoveAll(c.File)
}
