package ggnet

import (
	"net"
	"time"
)

type Transporter interface {
	Receive() ([]byte, error)
	Send(interface{}) error
	Close() error
}

type SetReadDeadline interface {
	SetReadDeadline(t time.Time) error
}

type dummyAddr struct {
}

func (d *dummyAddr) Network() string {
	return "io.ReadWriter"
}
func (d *dummyAddr) String() string {
	return "io.ReadWriter"
}

type RemoteAddr interface {
	RemoteAddr() (addr net.Addr)
}
