package gants

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestGants(t *testing.T) {
	defer ants.Release()

	runtimes := 10
	var sum int32

	var wg sync.WaitGroup
	syncCalculateSum := func() {
		func() {
			time.Sleep(10 * time.Millisecond)
			fmt.Println("hello world")
		}()
		wg.Done()
	}

	for i := 0; i < runtimes; i++ {
		wg.Add(1)
		_ = ants.Submit(syncCalculateSum)
	}

	wg.Wait()

	fmt.Printf("running goroutines: %d\n", ants.Running())
	fmt.Printf("finish all tasks.\n")

	p, _ := ants.NewPoolWithFunc(10, func(i interface{}) {
		func(i interface{}) {
			n := i.(int32)
			atomic.AddInt32(&sum, n)
			fmt.Printf("run with %d\n", n)
		}(i)
		wg.Done()
	})
	defer p.Reboot()

	for i := 0; i < runtimes; i++ {
		wg.Add(1)
		_ = p.Invoke(int32(i))
	}
	wg.Wait()
	fmt.Printf("running goroutines: %d\n", p.Running())
	fmt.Printf("finish all tasks, result is %d\n", sum)

}
