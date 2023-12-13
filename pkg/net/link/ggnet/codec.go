package ggnet

import (
	"encoding/binary"
	"github.com/panjf2000/gnet/v2"
)

type Codec struct {
}

func (cd Codec) Encode(protoId uint16, buf []byte) ([]byte, error) {
	bodyOffset := TRANS_HEAD_LEN
	msgLen := bodyOffset + len(buf)

	data := make([]byte, msgLen)
	//binary.LittleEndian.PutUint16(data, protoId)
	binary.LittleEndian.PutUint32(data[:bodyOffset], uint32(len(buf)))
	copy(data[bodyOffset:msgLen], buf)
	return data, nil
}

func (cd *Codec) Decode(c gnet.Conn) (uint16, []byte, error) {
	bodyOffset := TRANS_HEAD_LEN
	buf, _ := c.Peek(bodyOffset)
	if len(buf) < bodyOffset {
		return 0, nil, errIncompletePacket
	}
	var protoId uint16 = 0
	//protoId := binary.LittleEndian.Uint16(buf[:HEAD_PROTOID_LEN])

	bodyLen := binary.LittleEndian.Uint32(buf[:bodyOffset])
	msgLen := bodyOffset + int(bodyLen)
	if c.InboundBuffered() < msgLen {
		return protoId, nil, errIncompletePacket
	}
	buf, _ = c.Peek(msgLen)
	_, _ = c.Discard(msgLen)

	return protoId, buf[bodyOffset:msgLen], nil
}

func (cd Codec) Unpack(buf []byte) ([]byte, error) {
	bodyOffset := TRANS_HEAD_LEN
	if len(buf) < bodyOffset {
		return nil, errIncompletePacket
	}

	bodyLen := binary.LittleEndian.Uint32(buf[:bodyOffset])
	msgLen := bodyOffset + int(bodyLen)
	if len(buf) < msgLen {
		return nil, errIncompletePacket
	}

	return buf[bodyOffset:msgLen], nil
}
