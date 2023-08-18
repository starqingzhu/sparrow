package websocket

import (
	"fmt"
	"sparrow/internal/web"
	"sync"
	"testing"
)

func TestLogin(t *testing.T) {
	var wg sync.WaitGroup
	var clientCount = 500

	for i := 0; i < clientCount; i++ {
		wg.Add(1)

		var clientUser = fmt.Sprintf("user%d", i)
		client := web.WebLoginClient{
			Addr:     "192.168.59.184:9102",
			UserInfo: clientUser,
			WG:       &wg,
		}
		go client.Run()
	}

	wg.Wait()
}
