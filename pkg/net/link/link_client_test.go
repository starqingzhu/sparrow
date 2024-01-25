package link

import (
	"sparrow/pkg/net/link/ggnet"
	"sync"
	"testing"
	"time"
)

func TestTcpClient(t *testing.T) {
	var Addr = "127.0.0.1:8080"
	var wg sync.WaitGroup

	for i := 0; i < 1; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			client, err := ggnet.NewTcpClient()
			if err != nil {
				t.Errorf("NewTcpClient, id:%d, err:%s", j, err.Error())
				return
			}
			err = client.Connect(Addr)
			if err != nil {
				t.Errorf("Dial failed, id:%d, err:%s", j, err.Error())
				return
			}

			var sendBuf = "hello, world"
			t.Logf("send:%s", sendBuf)
			var c ggnet.Codec
			buf, err1 := c.Encode(1, []byte(sendBuf))
			if err1 != nil {
				t.Errorf("encode failed")
				return
			}
			_ = client.Conn.AsyncWrite(buf, nil)
			//client.Conn.AsyncWrite([]byte(sendBuf), nil)
			time.Sleep(time.Second * 20)
		}(i)
		time.Sleep(time.Millisecond * 1)
	}
	wg.Wait()
}
