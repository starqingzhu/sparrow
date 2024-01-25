package link

import (
	"encoding/binary"
	"github.com/stretchr/testify/assert"
	"sparrow/pkg/net/link/ggnet"
	"testing"
)

func TestTcpServer(t *testing.T) {

	var protocol = ggnet.NewFramedProtocol(4, 64*1024*1024, 64*1024*1024, binary.LittleEndian)
	opts := ggnet.NewServerSpecWithOpts(ggnet.WithAddr("127.0.0.1:8080"), ggnet.WithNetwork(ggnet.Tcp), ggnet.WithSendChanSize(10))
	var server = ggnet.NewTcpServer(opts, protocol)

	err := server.Run()

	assert.NoError(t, err)
}
