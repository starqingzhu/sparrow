package ggnet

import "sync"

type SessionManager struct {
	sms         ConcurrentMapUint64Session
	disposeOnce sync.Once
	disposeWait sync.WaitGroup
	disposed    bool
}

func NewSessionManager() *SessionManager {
	manager := &SessionManager{}
	manager.sms = NewConcurrentMapUint64Session()
	return manager
}

func (m *SessionManager) Dispose() {
	m.disposeOnce.Do(func() {
		m.disposed = true
		for tuple := range m.sms.Iter() {
			_ = tuple.Val.Close
		}
		m.disposeWait.Wait()
	})
}

func (m *SessionManager) SessionCount() int {
	if m.disposed {
		return 0
	}
	return m.sms.Count()
}

func (m *SessionManager) Range(f func(Session) bool) {
	for tuple := range m.sms.Iter() {
		if !f(tuple.Val) {
			break
		}
	}
}

func (m *SessionManager) NewSession(trans Transporter, spec *Spec) *_session {
	session := NewSession(trans, spec)
	m.putSession(session)
	return session
}

func (m *SessionManager) GetSession(sessionID uint64) *_session {
	session, _ := m.sms.Get(sessionID)
	return session
}

func (m *SessionManager) putSession(session *_session) {
	if m.disposed {
		session.Close()
		return
	}
	if m.sms.Has(session.ID()) {
		session.Close()
		m.delSession(session)
	}
	m.disposeWait.Add(1)
	m.sms.Set(session.ID(), session)
}

func (m *SessionManager) delSession(session *_session) {
	if m.sms.Has(session.ID()) {
		m.sms.Remove(session.ID())
		m.disposeWait.Done()
	}
}
