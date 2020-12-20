package main

import (
	"bytes"
	"os"
	"strconv"
	"sync"

	"git.code.oa.com/geeker/awesome-work/pingcap/constant"

	"git.code.oa.com/geeker/awesome-work/pingcap/domain"
	"git.code.oa.com/geeker/awesome-work/pingcap/skipList"
	"k8s.io/klog"
)

type WriteLoop struct {
	id         int
	blockChan  chan domain.Entry
	stopCh     chan struct{}
	bufferSize int

	list  skipList.SkipList
	count int
	//path  string
	dir  string
	slot *skipList.Slot

	wg *sync.WaitGroup
}

func NewWriteLoop(size int, id int, level int, dir string, wg *sync.WaitGroup) *WriteLoop {

	w := &WriteLoop{
		blockChan:  make(chan domain.Entry, size),
		stopCh:     make(chan struct{}),
		id:         id,
		list:       skipList.Constructor(level),
		bufferSize: 1 << level,
		dir:        dir,
		wg:         wg,
	}
	err := os.MkdirAll(w.dir+"/"+strconv.Itoa(id), os.ModePerm)
	if err != nil {
		klog.Fatalf("create dir failed ,err%v", err)
	}
	w.slot = skipList.NewSlot(constant.BloomM, constant.BloomK, w.id,
		w.getBloom(), w.getFile())
	return w
}
func (w *WriteLoop) start() {
	klog.Infof("write loop %d start  ", w.id)
	defer func() {
		klog.Infof("write loop %d end  ", w.id)
	}()
	for {
		select {
		case m, ok := <-w.blockChan:
			if !ok {
				w.write()
				w.wg.Done()
				return
			}
			if w.list.Size() >= w.bufferSize {
				w.write()
			}
			w.slot.Test(m.Key)
			w.list.Add(m.Hash, &m)
		case <-w.stopCh:
			klog.Infof("write %d stop", w.id)
			w.write()
			w.wg.Done()
			return
		}
	}
}
func (w *WriteLoop) Send(entry domain.Entry) {
	w.blockChan <- entry
}
func (w *WriteLoop) Close() {
	close(w.blockChan)
}
func (w *WriteLoop) Stop() {
	w.stopCh <- struct{}{}
}
func (w *WriteLoop) write() {
	f, err := os.Create(w.getFile())
	if err != nil {
		klog.Errorf("create file constant:%v", err)
	}
	defer f.Close()
	b := bytes.Buffer{}
	for _, v := range w.list.ToValues() {
		_, err := b.Write(v.Encode())
		if err != nil {
			klog.Errorf("write entry failed err:%v", err)
		}
	}
	f.Write(b.Bytes())
	w.count++
	w.slot.Save()
	w.list = skipList.Constructor(w.bufferSize)
	w.slot = skipList.NewSlot(constant.BloomM, constant.BloomK, w.id,
		w.getBloom(), w.getFile())
}
func (w *WriteLoop) getFile() string {
	return w.dir + "/" + strconv.Itoa(w.id) + "/segment_" + strconv.Itoa(w.count) + ".index"
}
func (w *WriteLoop) getBloom() string {
	return w.dir + "/" + strconv.Itoa(w.id) + "/segment_" + strconv.Itoa(w.count) + ".bloom"
}
