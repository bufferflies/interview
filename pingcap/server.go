package main

import (
	"sync"

	"git.code.oa.com/geeker/awesome-work/pingcap/block"
	"git.code.oa.com/geeker/awesome-work/pingcap/domain"
	"git.code.oa.com/geeker/awesome-work/pingcap/hash"
	"git.code.oa.com/geeker/awesome-work/pingcap/storage"
	"k8s.io/klog"
)

// (size int, id int, level int, stop chan struct{}, path string)
type Config struct {
	// chan size
	Size int
	// buffer size:2^Level default is 64M(16)
	Level int
	// segment location
	Path string
	// worker
	Worker int
	// Src
	Src string
}

type Server struct {
	config   Config
	register RegisterHandler
	read     *storage.Read
	Hash     *hash.Hash
	wg       *sync.WaitGroup
}
type RegisterHandler struct {
	handlerMap map[int]*WriteLoop
}

func NewServer(config Config) Server {
	r := storage.NewRead(config.Src)
	if r == nil {
		klog.Fatal("open src file  failed, please check ")
	}
	s := Server{
		config: config,
		register: RegisterHandler{
			handlerMap: make(map[int]*WriteLoop),
		},
		read: r,
		Hash: hash.NewHash(config.Worker),
		wg:   &sync.WaitGroup{},
	}
	s.wg.Add(config.Worker)
	for i := 0; i < config.Worker; i++ {
		w := NewWriteLoop(config.Size, i, config.Level, config.Path, s.wg)
		s.register.handlerMap[i] = w
	}
	return s
}
func (s *Server) Start() {
	s.PreHandler()
}
func (s *Server) PreHandler() {
	for _, v := range s.register.handlerMap {
		go v.start()
	}
	b, err := s.read.ReadNext(0)
	for offset := int64(0); err == nil; {
		hash, id := s.Hash.Hash(b.Key)
		entry := convert(b, hash)
		s.send(id, entry)
		offset = b.Offset + b.Length
		b, err = s.read.ReadNext(offset)
	}
	s.close()
	s.wg.Wait()
	klog.Infof("pre handle finish ")
}

//func (s *Server) Read(key []byte) []byte {
//	hash, id := s.Hash.Hash(b.Key)
//}
func (s *Server) send(id int, entry domain.Entry) {
	s.register.handlerMap[id].Send(entry)
}
func convert(b block.Block, hash uint64) domain.Entry {
	return domain.Entry{
		Offset: b.Offset,
		Key:    b.Key,
		Hash:   hash,
	}
}

func (s *Server) close() {
	for _, v := range s.register.handlerMap {
		v.Close()
	}
	s.read.Close()
}
