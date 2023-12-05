package tcp

import (
	"github.com/panjf2000/gnet/v2"
	"github.com/panjf2000/gnet/v2/pkg/logging"
	"github.com/stretchr/testify/assert"
	"sparrow/pkg/net/tcp/ggnet"
	"testing"
	"time"
)

func TestTcpClient(t *testing.T) {

	client := &ggnet.TcpClient{
		Addr: "127.0.0.1:8080",
	}
	cli, err := gnet.NewClient(client, gnet.WithLogLevel(logging.DebugLevel),
		gnet.WithLockOSThread(true),
		gnet.WithTicker(true))

	assert.NoError(t, err)

	err = cli.Start()
	assert.NoError(t, err)
	defer cli.Stop() //nolint:errcheck

	client.Conn, err = cli.Dial("tcp", client.Addr)
	assert.NoError(t, err)

	client.Conn.AsyncWrite([]byte("hello, world"), nil)

	time.Sleep(time.Second * 10)
}
