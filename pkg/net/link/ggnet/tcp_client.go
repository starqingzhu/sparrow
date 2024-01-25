package ggnet

import (
	"bytes"
	"github.com/panjf2000/gnet/v2"
	"github.com/panjf2000/gnet/v2/pkg/logging"
	"sparrow/pkg/log/zaplog"
	"time"
)

type TcpClientEvent struct {
	gnet.Conn
}

type TcpClient struct {
	*TcpClientEvent
	*gnet.Client
}

func NewTcpClient() (*TcpClient, error) {
	event := &TcpClientEvent{}
	c := &TcpClient{TcpClientEvent: event}
	cli, err := gnet.NewClient(event, gnet.WithLogLevel(logging.DebugLevel),
		gnet.WithLockOSThread(true),
		gnet.WithTicker(true))
	if err != nil {
		zaplog.LoggerSugar.Errorf("NewClient failed, err:%s", err.Error())
		return nil, err
	}
	c.Client = cli
	c.Client.Start()

	return c, nil
}

func (c *TcpClient) Connect(addr string) error {
	conn, err := c.Client.Dial("link", addr)
	c.TcpClientEvent.Conn = conn
	return err
}

// OnBoot fires when the engine is ready for accepting connections.
// The parameter engine has information and various utilities.
func (c *TcpClientEvent) OnBoot(eng gnet.Engine) (action gnet.Action) {
	fd, err := eng.Dup()
	zaplog.LoggerSugar.Infof("client is Onboot, fd:%d, err:%v", fd, err)

	return
}

// OnShutdown fires when the engine is being shut down, it is called right after
// all event-loops and connections are closed.
func (c *TcpClientEvent) OnShutdown(eng gnet.Engine) {
	fd, err := eng.Dup()
	zaplog.LoggerSugar.Infof("client is shutdowning, fd:%d, err:%v", fd, err)
	return
}

// OnOpen fires when a new connection has been opened.
//
// The Conn c has information about the connection such as its local and remote addresses.
// The parameter out is the return value which is going to be sent back to the peer.
// Sending large amounts of data back to the peer in OnOpen is usually not recommended.
func (c *TcpClientEvent) OnOpen(cn gnet.Conn) (out []byte, action gnet.Action) {

	fd, err := cn.Dup()
	zaplog.LoggerSugar.Infof("fd:%d, err:%v, ip:%s connect server ip:%s", fd, err, cn.LocalAddr(), cn.RemoteAddr())
	return
}

// OnClose fires when a connection has been closed.
// The parameter err is the last known connection error.
func (c *TcpClientEvent) OnClose(cn gnet.Conn, err error) (action gnet.Action) {
	zaplog.LoggerSugar.Infof("ip:%s disconnect server:%s, err:%v", cn.LocalAddr(), cn.RemoteAddr(), err)
	cn.Close()
	return
}

// OnTraffic fires when a socket receives data from the peer.
//
// Note that the []byte returned from Conn.Peek(int)/Conn.Next(int) is not allowed to be passed to a new goroutine,
// as this []byte will be reused within event-loop after OnTraffic() returns.
// If you have to use this []byte in a new goroutine, you should either make a copy of it or call Conn.Read([]byte)
// to read data into your own []byte, then pass the new []byte to the new goroutine.
func (c *TcpClientEvent) OnTraffic(cn gnet.Conn) (action gnet.Action) {

	var buffer *bytes.Buffer
	var buff []byte

	buf, err := cn.Next(-1)
	buff = make([]byte, len(buf))
	copy(buff, buf)

	buffer = bytes.NewBuffer(buff)
	if err != nil {
		return gnet.Close
	}

	zaplog.LoggerSugar.Infof("recv from server ip:%s, buff:%s", cn.RemoteAddr(), buffer)
	return
}

// OnTick fires immediately after the engine starts and will fire again
// following the duration specified by the delay return value.
func (c *TcpClientEvent) OnTick() (delay time.Duration, action gnet.Action) {
	//zaplog.LoggerSugar.Infof("ontick now:%d", time.Now().UnixMilli())
	return time.Second * 5, gnet.None
}
