package ggnet

import "io"

type Protocol interface {
	NewTransporter(rw io.ReadWriter) (Transporter, error)
}
type ProtocolFunc func(rw io.ReadWriter) (Transporter, error)

func (pf ProtocolFunc) NewTransporter(rw io.ReadWriter) (Transporter, error) {
	return pf(rw)
}
