package cache

type Lfu struct {
}

func (l Lfu) Get(key []byte) (value []byte, err error) {
	panic("implement me")
}
