package ggnet

import (
	"github.com/panjf2000/gnet/v2"
	"sparrow/pkg/log/zaplog"
	"time"
)

type TcpServer struct {
	*gnet.BuiltinEventEngine
	eng            gnet.Engine
	sessionManager *SessionManager
	handler        Handler
	protocol       Protocol
	spec           *Spec
	//workerPool     *goPool.Pool `json:"workerPool"`
}

func NewTcpServer(s *Spec, protocol Protocol, handler Handler) *TcpServer {
	return &TcpServer{
		sessionManager: NewSessionManager(),
		spec:           s,
		protocol:       protocol,
		handler:        handler,
		//workerPool:     goPool.Default(),
	}
}

func (s *TcpServer) Run() error {
	return gnet.Run(s,
		s.spec.Addr,
		gnet.WithLockOSThread(true),
		gnet.WithMulticore(true),
		gnet.WithReusePort(false),
		gnet.WithReuseAddr(false),
		gnet.WithTicker(true),
		gnet.WithTCPKeepAlive(time.Minute*1),
		gnet.WithTCPNoDelay(gnet.TCPDelay),
		gnet.WithLoadBalancing(gnet.LeastConnections))
}

// OnBoot fires when the engine is ready for accepting connections.
// The parameter engine has information and various utilities.
func (s *TcpServer) OnBoot(eng gnet.Engine) (action gnet.Action) {
	s.eng = eng
	fd, err := eng.Dup()
	zaplog.LoggerSugar.Infof("[TcpServer::OnBoot] server start success, %s, fd:%d, err:%v", s.spec.Addr, fd, err)
	return
}

// OnShutdown fires when the engine is being shut down, it is called right after
// all event-loops and connections are closed.
func (s *TcpServer) OnShutdown(eng gnet.Engine) {
	fd, err := eng.Dup()
	zaplog.LoggerSugar.Infof("[TcpServer::OnShutdown] server is shutdowning, fd:%d, err:%v", fd, err)
	s.sessionManager.Dispose()
	return
}

// OnOpen fires when a new connection has been opened.
//
// The Conn c has information about the connection such as its local and remote addresses.
// The parameter out is the return value which is going to be sent back to the peer.
// Sending large amounts of data back to the peer in OnOpen is usually not recommended.
func (s *TcpServer) OnOpen(c gnet.Conn) (out []byte, action gnet.Action) {

	//构建session
	trans, err := s.protocol.NewTransporter(TCPConnSetup(c, s.spec))
	if err != nil {
		zaplog.LoggerSugar.Errorf("[TcpServer::OnOpen] ip:%s  NewTransporter failed", c.RemoteAddr())
		_ = c.Close()
		return
	}
	session := s.sessionManager.NewSession(trans, s.spec)
	c.SetContext(session.ID())
	zaplog.LoggerSugar.Infof("[TcpServer::OnOpen] ip:%s connect server, total:%d", c.RemoteAddr(), s.sessionManager.SessionCount())

	//设置业务层代理回调
	s.handler.HandleSession(session)

	////todo 测试代码
	//_ = session.WriteFrame([]byte("server back"))
	return
}

// OnClose fires when a connection has been closed.
// The parameter err is the last known connection error.
func (s *TcpServer) OnClose(c gnet.Conn, err error) (action gnet.Action) {
	id, okCtx := c.Context().(uint64)
	if !okCtx {
		zaplog.LoggerSugar.Panicf("[TcpServer::OnClose] context error, net fd:%d", c.Fd())
		_ = c.Close()
		return
	}

	oldSession := s.sessionManager.GetSession(id)
	if oldSession == nil {
		zaplog.LoggerSugar.Panicf("[TcpServer::OnClose] not find Session, net id:%d", id)
		_ = c.Close()
		return
	}
	oldSession.Close()
	s.sessionManager.delSession(oldSession)

	zaplog.LoggerSugar.Infof("[TcpServer::OnClose] ip:%s disconnect server, total:%d, err:%v", c.RemoteAddr(), s.sessionManager.SessionCount(), err)

	if s.sessionManager.SessionCount() == 0 {
		zaplog.LoggerSugar.Infof("[TcpServer::OnClose] onclose clients is empty")
	}

	return
}

// OnTraffic fires when a socket receives data from the peer.
//
// Note that the []byte returned from Conn.Peek(int)/Conn.Next(int) is not allowed to be passed to a new goroutine,
// as this []byte will be reused within event-loop after OnTraffic() returns.
// If you have to use this []byte in a new goroutine, you should either make a copy of it or call Conn.Read([]byte)
// to read data into your own []byte, then pass the new []byte to the new goroutine.
func (s *TcpServer) OnTraffic(c gnet.Conn) (action gnet.Action) {

	//_ = c.InboundBuffered()
	//_ = c.OutboundBuffered()
	//_, _ = c.Discard(1)

	id, okCtx := c.Context().(uint64)
	if !okCtx {
		zaplog.LoggerSugar.Panicf("[TcpServer::OnTraffic] context error, net fd:%d", c.Fd())
	}
	session := s.sessionManager.GetSession(id)
	if session == nil {
		zaplog.LoggerSugar.Errorf("[TcpServer::OnTraffic] not find Session, net id:%d", id)
		_ = c.Close()
		return
	}
	_, err := session.ReadFrame()
	if err != nil {
		zaplog.LoggerSugar.Errorf("[TcpServer::OnTraffic] OnRecv error, net id:%d", id)
		return
	}
	c.RemoteAddr()
	/*
		buf := bbPool.Get()
		nLen, err := c.WriteTo(buf)
		//nLen, err := c.Read(buf.Bytes())
		if err != nil {
			zaplog.LoggerSugar.Errorf("[TcpServer::OnTraffic] WriteTo failed, err:[%s], addr:[%s]", err.Error(), c.RemoteAddr())
			c.Close()
			bbPool.Put(buf)
			return
		}
		zaplog.LoggerSugar.Debugf("[TcpServer::OnTraffic] recv nLen:%d, buf:[%s], client:[%s]", nLen, buf.Bytes(), c.RemoteAddr())
		_ = s.workerPool.Submit(
			func() {
				err = c.AsyncWrite(buf.Bytes(), func(c gnet.Conn, err error) error {
					if err != nil {
						zaplog.LoggerSugar.Errorf("[TcpServer::OnTraffic] conn=[%s] done writev:[%v]", c.RemoteAddr().String(), err)
						bbPool.Put(buf)
					}
					return nil
				})
				zaplog.LoggerSugar.Debugf("[TcpServer::OnTraffic] send nLen:%d, buf:[%s], client:[%s]", buf.Len(), buf.Bytes(), c.RemoteAddr())
				time.Sleep(time.Millisecond * 100)
				bbPool.Put(buf)
			},
		)
	*/

	//var buffer *bytes.Buffer
	//var buff []byte
	//
	//buf, err := c.Next(-1)
	//buff = make([]byte, len(buf))
	//copy(buff, buf)
	//
	//buffer = bytes.NewBuffer(buff)
	//if err != nil {
	//	return gnet.Close
	//}
	//
	//zaplog.LoggerSugar.Infof("ip:%s, buff:%s", c.RemoteAddr(), buffer)
	return
}

// OnTick fires immediately after the engine starts and will fire again
// following the duration specified by the delay return value.
func (s *TcpServer) OnTick() (delay time.Duration, action gnet.Action) {
	//zaplog.LoggerSugar.Infof("[TcpServer::OnTick] ontick now:%d", time.Now().UnixMilli())
	return time.Second * 5, gnet.None
}
