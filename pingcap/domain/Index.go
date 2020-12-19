package domain

import "git.code.oa.com/geeker/awesome-work/pingcap/common"

const SlotLength = 16

type Index struct {
	Entries []Entry
}
type Entry struct {
	hash   uint64
	offset int64
}

func (i *Index) Encode() []byte {
	ret := make([]byte, SlotLength*len(i.Entries))
	for i, v := range i.Entries {
		copy(ret[i*SlotLength:(i+1)*SlotLength], v.Encode())
	}
	return ret
}

func (i *Index) Decode(bytes []byte) {
	size := len(bytes) / SlotLength
	entries := make([]Entry, size)
	for i := 0; i < size; i++ {
		t := &Entry{}
		t.Decode(bytes[i*SlotLength : (i+1)*SlotLength])
		entries[i] = *t
	}
	i.Entries = entries
}

func (e *Entry) Encode() []byte {
	ret := make([]byte, 16)
	copy(ret[:8], common.Uint64ToBytes(e.hash))
	copy(ret[8:], common.Int64ToBytes(e.offset))
	return ret
}

func (e *Entry) Decode(bytes []byte) {
	e.hash = common.BytesToUint64(bytes[:8])
	e.offset = common.BytesToInt64(bytes[8:])
}
