package websocket

import (
	"fmt"
	"sparrow/internal/web"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestBinLogin(t *testing.T) {
	var wg sync.WaitGroup
	var totalClient atomic.Int64

	totalClient.Store(0)
	var clientCount int64 = 1000

	var incr int64 = 0

	var continueFlag bool = true

	for continueFlag {
		for totalClient.Load() < clientCount {
			wg.Add(1)

			var clientUser = fmt.Sprintf("user%d", incr)
			totalClient.Add(1)
			client := web.WebLoginClient{
				//Addr:        "192.168.59.184:9102",
				Addr:        "10.110.55.211:9108",
				UserInfo:    clientUser,
				WG:          &wg,
				Mt:          &sync.Mutex{},
				TotalClient: &totalClient,
			}
			go client.RunBin()
			incr++
			//time.Sleep(10 * time.Millisecond)

		}
		time.Sleep(300 * time.Millisecond)
	}

	wg.Wait()
}
