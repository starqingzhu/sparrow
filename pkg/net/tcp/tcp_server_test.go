package tcp

import (
	"github.com/panjf2000/gnet/v2"
	"github.com/stretchr/testify/assert"
	"sparrow/pkg/net/tcp/ggnet"
	"testing"
	"time"
)

func TestTcpServer(t *testing.T) {
	var server = &ggnet.TcpServer{
		Addr: "127.0.0.1:8080",
	}

	err := gnet.Run(server,
		server.Addr,
		gnet.WithLockOSThread(true),
		gnet.WithMulticore(true),
		gnet.WithReusePort(false),
		gnet.WithReuseAddr(false),
		gnet.WithTicker(true),
		gnet.WithTCPKeepAlive(time.Minute*1),
		gnet.WithTCPNoDelay(gnet.TCPDelay),
		gnet.WithLoadBalancing(gnet.LeastConnections))

	assert.NoError(t, err)
	//if err == nil {
	//	t.Logf("gnet run failed, err:%s\n", err.Error())
	//	return
	//}
}
