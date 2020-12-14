package my

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var dins map[string]chan int
var lock sync.Mutex

func register(din string, ch chan int) {
	lock.Lock()
	defer lock.Unlock()
	dins[din] = ch
}
func deregister(din string) {
	lock.Lock()
	defer lock.Unlock()
	ch, ok := dins[din]
	if ok {
		close(ch)
		delete(dins, din)
	}

}
func getChan(din string) (r chan int, flag bool) {
	lock.Lock()
	defer lock.Unlock()
	ch, ok := dins[din]
	return ch, ok
}
func wait(din string) {
	ch := make(chan int)
	register(din, ch)

	select {
	case r := <-ch:
		fmt.Printf("notify di,r:%d \n\r", r)
		break
	case <-time.After(time.Minute):
		fmt.Print("time over")
		break
	}

}
func notify(din string) {
	ch, ok := getChan(din)
	if ok {
		ch <- 1
		deregister(din)
	}
}
func TestUnit_Consumer(t *testing.T) {
	dins = make(map[string]chan int)
	din := "test"
	go wait(din)
	time.Sleep(time.Second * 10)
	notify(din)
	notify(din)
}
