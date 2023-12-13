package ggnet

import (
	"github.com/panjf2000/gnet/v2"
)

func TCPConnSetup(conn gnet.Conn, spec *Spec) gnet.Conn {
	_ = conn.SetNoDelay(spec.TcpNoDelay)
	_ = conn.SetWriteBuffer(spec.TcpWriteBuffer)
	_ = conn.SetReadBuffer(spec.TcpReadBuffer)
	_ = conn.SetLinger(spec.TCPLingerSecond)
	if spec.KeepAlivePeriod != 0 {
		_ = conn.SetKeepAlivePeriod(spec.KeepAlivePeriod)
	}
	return conn
}
