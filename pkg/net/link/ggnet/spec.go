package ggnet

import "time"

type Spec struct {
	Addr            string
	Network         NetworkType
	TcpNoDelay      bool
	TcpWriteBuffer  int
	TcpReadBuffer   int
	KeepAlivePeriod time.Duration
	TCPLingerSecond int
	SendChanSize    int64 //发送缓存size
}

func NewServerSpecWithOpts(opts ...ServerConfOption) *Spec {
	s := defaultServerSpec
	for _, opt := range opts {
		opt(&s)
	}
	return &s
}

type ServerConfOption func(*Spec)

func WithAddr(addr string) ServerConfOption {
	return func(o *Spec) {
		o.Addr = addr
	}
}

func WithNetwork(t NetworkType) ServerConfOption {
	return func(o *Spec) {
		o.Network = t
	}
}

func WithTcpNoDelay(noDelay bool) ServerConfOption {
	return func(o *Spec) {
		o.TcpNoDelay = noDelay
	}
}

func WithTcpWriteBuffer(writeBuff int) ServerConfOption {
	return func(o *Spec) {
		o.TcpWriteBuffer = writeBuff
	}
}

func WithTcpReadBuffer(readBuff int) ServerConfOption {
	return func(o *Spec) {
		o.TcpReadBuffer = readBuff
	}
}

func WithTCPLingerSecond(t int) ServerConfOption {
	return func(o *Spec) {
		o.TCPLingerSecond = t
	}
}

func WithKeepAlivePeriod(t time.Duration) ServerConfOption {
	return func(o *Spec) {
		o.KeepAlivePeriod = t
	}
}
func WithSendChanSize(n int64) ServerConfOption {
	return func(o *Spec) {
		o.SendChanSize = n
	}
}

const (
	AcceptTimeout = 500 * time.Millisecond
	//ReadTimeout zero for not set read deadline for Conn (better  performance)
	ReadTimeout = 0 * time.Millisecond
	//WriteTimeout zero for not set write deadline for Conn (better performance)
	WriteTimeout = 0 * time.Millisecond
	//IdleTimeout idle timeout
	IdleTimeout = 600000 * time.Millisecond
	//QueuePerInvoker queue gap
	QueuePerInvoker int = 10000000
	KeepAlivePeriod     = 0 * time.Millisecond

	//TCPReadBuffer link read buffer length
	TCPReadBuffer = 128 * 1024 * 1024
	//TCPWriteBuffer link write buffer length
	TCPWriteBuffer = 128 * 1024 * 1024
	// true for single user connection(better performance)
	TCPNoDelay          = true
	MaxInvokePerSession = -1
	TCPLingerSecond     = 0
	FlushTimeout        = time.Second * 5
	SendChanSize        = 16
)

var defaultServerSpec = Spec{
	Addr:            "127.0.0.1:8080",
	Network:         Tcp,
	TcpNoDelay:      TCPNoDelay,
	TcpWriteBuffer:  TCPWriteBuffer,
	TcpReadBuffer:   TCPReadBuffer,
	TCPLingerSecond: TCPLingerSecond,
	KeepAlivePeriod: KeepAlivePeriod,
	SendChanSize:    SendChanSize,
}
