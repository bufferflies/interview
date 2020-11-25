package base

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func Test_Context(t *testing.T) {
	ctx, f := context.WithCancel(context.Background())
	go wait(ctx)
	time.Sleep(time.Second * 3)
	f()

}
func wait(ctx context.Context) {
	select {
	// parent thread done
	case <-ctx.Done():
		fmt.Errorf("done")
	// wait less than 2s
	case <-time.After(time.Second * 2):
		fmt.Errorf("time out ")
	}
}
