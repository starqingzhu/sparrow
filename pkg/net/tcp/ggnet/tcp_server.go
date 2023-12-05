package ggnet

import (
	"bytes"
	"github.com/panjf2000/gnet/v2"
	"sparrow/pkg/log/zaplog"
	"time"
)

type TcpServer struct {
	Addr string `json:"addr,default=127.0.0.1:8080"`
	gnet.BuiltinEventEngine
}

// OnBoot fires when the engine is ready for accepting connections.
// The parameter engine has information and various utilities.
func (s *TcpServer) OnBoot(eng gnet.Engine) (action gnet.Action) {
	fd, err := eng.Dup()
	zaplog.LoggerSugar.Infof("server is listening, fd:%d, err:%v", fd, err)
	return
}

// OnShutdown fires when the engine is being shut down, it is called right after
// all event-loops and connections are closed.
func (s *TcpServer) OnShutdown(eng gnet.Engine) {
	fd, err := eng.Dup()
	zaplog.LoggerSugar.Infof("server is shutdowning, fd:%d, err:%v", fd, err)
	return
}

// OnOpen fires when a new connection has been opened.
//
// The Conn c has information about the connection such as its local and remote addresses.
// The parameter out is the return value which is going to be sent back to the peer.
// Sending large amounts of data back to the peer in OnOpen is usually not recommended.
func (s *TcpServer) OnOpen(c gnet.Conn) (out []byte, action gnet.Action) {
	zaplog.LoggerSugar.Infof("ip:%s connect server", c.RemoteAddr())
	return
}

// OnClose fires when a connection has been closed.
// The parameter err is the last known connection error.
func (s *TcpServer) OnClose(c gnet.Conn, err error) (action gnet.Action) {
	zaplog.LoggerSugar.Infof("ip:%s disconnect server, err:%v", c.RemoteAddr(), err)
	return
}

// OnTraffic fires when a socket receives data from the peer.
//
// Note that the []byte returned from Conn.Peek(int)/Conn.Next(int) is not allowed to be passed to a new goroutine,
// as this []byte will be reused within event-loop after OnTraffic() returns.
// If you have to use this []byte in a new goroutine, you should either make a copy of it or call Conn.Read([]byte)
// to read data into your own []byte, then pass the new []byte to the new goroutine.
func (s *TcpServer) OnTraffic(c gnet.Conn) (action gnet.Action) {

	var buffer *bytes.Buffer
	var buff []byte

	buf, err := c.Next(-1)
	buff = make([]byte, len(buf))
	copy(buff, buf)

	buffer = bytes.NewBuffer(buff)
	if err != nil {
		return gnet.Close
	}

	zaplog.LoggerSugar.Infof("ip:%s, buff:%s", c.RemoteAddr(), buffer)
	return
}

// OnTick fires immediately after the engine starts and will fire again
// following the duration specified by the delay return value.
func (s *TcpServer) OnTick() (delay time.Duration, action gnet.Action) {
	zaplog.LoggerSugar.Infof("ontick now:%d", time.Now().UnixMilli())
	return time.Second * 5, gnet.None
}
