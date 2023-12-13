package ggnet

import (
	"context"
	"sparrow/pkg/log/zaplog"
	"sync/atomic"
)

type sessionKeyType struct{}

func (*sessionKeyType) String() string {
	return "session-key-bl9hi867"
}

var SessionKeyForContext = sessionKeyType(struct{}{})

func NewContextWithSession(ctx context.Context, s Session) context.Context {
	return context.WithValue(ctx, SessionKeyForContext, s)
}
func SessionFromContext(ctx context.Context) (s Session, ok bool) {
	s, ok = ctx.Value(SessionKeyForContext).(Session)
	return
}

type Session interface {
	ID() uint64
	Close()
	IsClose() bool

	Transporter() Transporter

	ReadFrame() ([]byte, error)
	WriteFrame(raw interface{}) error
}

type _session struct {
	transporter Transporter
	Id          uint64
	CloseFlag   atomic.Bool
	//send
	sendChan  chan interface{}
	closeChan chan struct{}
}

func NewSession(transporter Transporter, spec *Spec) *_session {
	session := &_session{
		Id:          globalSessionId.Add(1),
		transporter: transporter,
		closeChan:   make(chan struct{}),
	}
	session.CloseFlag.Store(false)

	if spec.SendChanSize > 0 {
		session.sendChan = make(chan interface{}, spec.SendChanSize)
		go session.sendLoop()
	}

	return session
}

func (s *_session) ID() uint64 {
	return s.Id
}

func (s *_session) Close() {
	s.transporter.Close()
	if s.CloseFlag.CompareAndSwap(false, true) {
		close(s.closeChan)
	}
	close(s.sendChan)
}

func (s *_session) IsClose() bool {
	return s.CloseFlag.Load()
}

func (s *_session) Transporter() Transporter {
	return s.transporter
}

func (s *_session) ReadFrame() ([]byte, error) {
	ret, err := s.transporter.Receive()
	if err != nil {
		_ = s.Close
	}

	remoteIp := invalidIpString
	if ra, ok := s.transporter.(RemoteAddr); ok {
		remoteIp = ra.RemoteAddr().String()
	}

	zaplog.LoggerSugar.Debugf("[_session::OnRecv] recv nLen:%d, buf:[%s], sessionId:[%d], remoteAddr:[%s]", len(ret), string(ret), s.ID(), remoteIp)

	return ret, err
}

func (s *_session) WriteFrame(raw interface{}) error {
	if s.IsClose() {
		return SessionClosedError
	}
	//无缓存
	if s.sendChan == nil {
		err := s.transporter.Send(raw)
		if err != nil {
			s.Close()
		}
		return err
	}
	//有缓存
	select {
	case s.sendChan <- raw:
		return nil
	default:
		s.Close()
		return SessionBlockedError
	}

}

func (s *_session) sendLoop() {
	defer func() {
		_ = s.Close
	}()

	for {
		select {
		case b, ok := <-s.sendChan:
			if !ok || b == nil {
				return
			}
			if err := s.transporter.Send(b); err != nil {
				return
			}
		case <-s.closeChan:
			return
		}
	}
}
