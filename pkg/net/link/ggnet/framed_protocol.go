package ggnet

import (
	"encoding/binary"
	"io"
	"math"
)

type FramedProtocol struct {
	SizeofLen   int
	maxRecv     int
	maxSend     int
	headDecoder func([]byte) int
	headEncoder func([]byte, int)
}

func (f *FramedProtocol) NewTransporter(rw io.ReadWriter) (t Transporter, err error) {
	r := &FramedTransporter{
		rw:       rw,
		protocol: f,
		//pool: &sync.Pool{
		//	New: func() interface{} {
		//		ret := make([]byte, 4)
		//		return ret
		//	},
		//},
	}
	r.readHeadBuf = r.readHead[:f.SizeofLen]
	r.sendHeadBuf = r.sendHead[:f.SizeofLen]

	return r, nil
}

func NewFramedProtocol(sizeofLen, maxRecv, maxSend int, byteOrder binary.ByteOrder) *FramedProtocol {
	protoObj := &FramedProtocol{
		SizeofLen: sizeofLen,
	}
	switch sizeofLen {
	case 1:
		if maxRecv > math.MaxUint8 {
			maxRecv = math.MaxUint8
		}
		if maxSend > math.MaxUint8 {
			maxSend = math.MaxUint8
		}
		protoObj.headDecoder = func(b []byte) int {
			return int(b[0])
		}
		protoObj.headEncoder = func(b []byte, size int) {
			b[0] = byte(size)
		}
	case 2:
		if maxRecv > math.MaxUint16 {
			maxRecv = math.MaxUint16
		}
		if maxSend > math.MaxUint16 {
			maxSend = math.MaxUint16
		}
		protoObj.headDecoder = func(b []byte) int {
			return int(byteOrder.Uint16(b))
		}
		protoObj.headEncoder = func(b []byte, size int) {
			byteOrder.PutUint16(b, uint16(size))
		}
	case 4:
		if maxRecv > math.MaxUint32 {
			maxRecv = math.MaxUint32
		}
		if maxSend > math.MaxUint32 {
			maxSend = math.MaxUint32
		}
		protoObj.headDecoder = func(b []byte) int {
			return int(byteOrder.Uint32(b))
		}
		protoObj.headEncoder = func(b []byte, size int) {
			byteOrder.PutUint32(b, uint32(size))
		}
	case 8:
		protoObj.headDecoder = func(b []byte) int {
			return int(byteOrder.Uint64(b))
		}
		protoObj.headEncoder = func(b []byte, size int) {
			byteOrder.PutUint64(b, uint64(size))
		}
	default:
		panic("FramedProtocol: unsupported head size")
	}
	protoObj.maxRecv = maxRecv
	protoObj.maxSend = maxSend

	return protoObj
}
