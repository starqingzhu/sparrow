package ggnet

import (
	"encoding/binary"
	"io"
	"net"
	"sparrow/pkg/log/zaplog"
	"time"
)

type FramedTransporter struct {
	protocol *FramedProtocol
	rw       io.ReadWriter

	readHeadBuf []byte
	sendHeadBuf []byte
	readHead    [TRANS_HEAD_LEN]byte
	sendHead    [TRANS_HEAD_LEN]byte
	//pool        *sync.Pool
}

func (f *FramedTransporter) Receive() ([]byte, error) {
	var err error

	if _, err = io.ReadFull(f.rw, f.readHeadBuf); err != nil {
		return nil, err
	}

	size := f.protocol.headDecoder(f.readHeadBuf)
	if size > f.protocol.maxRecv {
		return nil, ErrTooLargeFrame
	}

	frame := make([]byte, size)
	if _, err = io.ReadFull(f.rw, frame); err != nil {
		return nil, err
	}

	zaplog.LoggerSugar.Debugf("[FramedTransporter::Receive] err:%v, frame:%s", err, frame)

	return frame, nil
}

func (f *FramedTransporter) Send(msg interface{}) (err error) {
	allLen := 0
	switch msg.(type) {
	case [][]byte:
		buffers := msg.([][]byte)
		for i := 0; i < len(buffers); i++ {
			allLen += len(buffers[i])
		}

		binary.LittleEndian.PutUint32(f.sendHeadBuf, uint32(allLen))
		_, err = f.rw.Write(f.sendHeadBuf)
		if err != nil {
			return nil
		}

		netBuf := net.Buffers(buffers)
		_, err = netBuf.WriteTo(f.rw)
	case []byte:
		body := msg.([]byte)
		buffers := [][]byte{f.sendHeadBuf, body}
		allLen = len(body)
		binary.LittleEndian.PutUint32(buffers[0], uint32(allLen))
		netBuf := net.Buffers(buffers)

		_, err = netBuf.WriteTo(f.rw)
	}
	return nil
}

func (f *FramedTransporter) Close() error {
	if closer, ok := f.rw.(io.Closer); ok {
		return closer.Close()
	}

	return nil
}

func (f *FramedTransporter) SetReadDeadline(t time.Time) error {
	if rd, ok := f.rw.(SetReadDeadline); ok {
		return rd.SetReadDeadline(t)
	}
	return nil
}

func (f *FramedTransporter) RemoteAddr() net.Addr {
	if ra, ok := f.rw.(RemoteAddr); ok {
		return ra.RemoteAddr()
	}
	return &dummyAddr{}
}
