package ggnet

import "github.com/pkg/errors"

var (
	errIncompletePacket = errors.New("incomplete packet")

	//session 相关
	SessionClosedError  = errors.New("session has closed")
	SessionBlockedError = errors.New("session blocked")

	//frame 相关
	ErrTooLargeFrame = errors.New("too large frame")
)
