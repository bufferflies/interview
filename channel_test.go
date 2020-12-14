package awesome_work

import (
	"fmt"
	"sync"
	"testing"
)

func TestUnit_Channel(t *testing.T) {
	ch := make(chan int, 2)
	wg := sync.WaitGroup{}
	wg.Add(2)
	ch <- 1
	ch <- 2
	close(ch)
	go func() {
		m := <-ch
		fmt.Printf("receive-A message:%d \n\r", m)
		wg.Done()
	}()
	go func() {
		m := <-ch
		fmt.Printf("receive-B message:%d \n\r", m)
		wg.Done()
	}()

	wg.Wait()
	//time.Sleep(time.Second * 2)
}
