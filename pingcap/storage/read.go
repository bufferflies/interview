package storage

import (
	"git.code.oa.com/geeker/awesome-work/pingcap/block"
	"git.code.oa.com/geeker/awesome-work/pingcap/common"
	"golang.org/x/exp/mmap"
	"k8s.io/klog"
)

type Read struct {
	path string
	r    *mmap.ReaderAt
}

func NewRead(path string) *Read {
	r, err := mmap.Open(path)
	if err != nil {
		klog.Errorf("open file failed,err:%v ", err)
		return nil
	}
	return &Read{
		r:    r,
		path: path,
	}
}
func (read *Read) Close() {
	read.r.Close()
}

func (read *Read) ReadNext(offset int64) (b block.Block, err error) {
	index := offset
	keySize := make([]byte, block.KeySizeLength)
	_, err = read.r.ReadAt(keySize, index)
	if err != nil {
		return block.Block{}, err
	}
	size := int64(common.BytesToInt(keySize))
	index = index + block.KeySizeLength

	key := make([]byte, size)
	read.r.ReadAt(key, index)
	valueSize := make([]byte, block.KeySizeLength)

	index = index + size
	read.r.ReadAt(valueSize, index)
	size = int64(common.BytesToInt(valueSize))

	index = index + block.KeySizeLength
	value := make([]byte, size)
	read.r.ReadAt(value, index)
	index = index + size

	ret := block.Block{
		KeySize:   len(key),
		Key:       key,
		ValueSize: len(value),
		Value:     value,
	}

	ret.Offset = offset
	ret.Length = index - offset
	return ret, nil
}
