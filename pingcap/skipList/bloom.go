package skipList

import (
	"io/ioutil"
	"os"

	"git.code.oa.com/geeker/awesome-work/pingcap/common"
	"github.com/willf/bloom"
	"k8s.io/klog"
)

type Slot struct {
	Filter *bloom.BloomFilter
	// 属于哪一个id
	Id int
	// 对应的index
	File string
	// 存储的file
	location string
}

func NewSlot(m, k uint, id int, location string, file string) *Slot {
	return &Slot{
		Filter:   bloom.New(m, k),
		Id:       id,
		location: location,
		File:     file,
	}
}
func FromFile(location string) *Slot {
	s := &Slot{
		Filter:   bloom.New(1, 1),
		location: location,
	}
	s.Load()
	return s
}

func (s *Slot) Add(key []byte) {
	s.Filter.Add(key)
}
func (s *Slot) Test(key []byte) bool {
	return s.Filter.Test(key)
}
func (s *Slot) Save() error {
	f, err := os.Create(s.location)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(s.Encode())
	return err
}
func (s *Slot) Load() error {
	f, err := os.OpenFile(s.location, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()
	body, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	s.Decode(body)
	return nil
}
func (s *Slot) Encode() []byte {
	d, err := s.Filter.MarshalJSON()
	if err != nil {
		klog.Errorf("bloom filter marshal failed,err:%v", err)
	}
	body := make([]byte, 4+4+len(s.File)+len(d))
	copy(body[:4], common.IntToBytes(s.Id))
	copy(body[4:8], common.IntToBytes(len(s.File)))
	copy(body[8:8+len(s.File)], s.File)
	copy(body[8+len(s.File):], d)
	return body
}

func (s *Slot) Decode(bytes []byte) {
	s.Id = common.BytesToInt(bytes[:4])
	length := common.BytesToInt(bytes[4:8])
	s.File = string(bytes[8 : 8+length])
	f := bloom.New(1, 1)
	err := f.UnmarshalJSON(bytes[8+length:])
	if err != nil {
		klog.Errorf("bloom filter unmarshal failed,err:%v", err)
	}
	s.Filter = f
}
