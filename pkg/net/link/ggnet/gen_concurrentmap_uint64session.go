package ggnet

import "sync"

var (
	SHARED_COUNT_SESSION_MAP = 32
)

type ConcurrentMapUint64Session []*sharedConcurrentMapUint64Session

type sharedConcurrentMapUint64Session struct {
	items map[uint64]*_session
	sync.RWMutex
}

// TupleConcurrentMapUint64Session Used by the Iter & IterBuffered functions to wrap two variables together over a channel,
type TupleConcurrentMapUint64Session struct {
	Key uint64
	Val *_session
}

func NewConcurrentMapUint64Session() ConcurrentMapUint64Session {
	this := make(ConcurrentMapUint64Session, SHARED_COUNT_SESSION_MAP)
	for i := 0; i < SHARED_COUNT_SESSION_MAP; i++ {
		this[i] = &sharedConcurrentMapUint64Session{items: make(map[uint64]*_session)}
	}
	return this
}

func (m ConcurrentMapUint64Session) GetShared(key uint64) *sharedConcurrentMapUint64Session {
	return m[func(key uint64) uint64 {
		return key % uint64(SHARED_COUNT_SESSION_MAP)
	}(key)]
}

func (m *ConcurrentMapUint64Session) Set(key uint64, value *_session) {
	shared := m.GetShared(key)
	shared.Lock()
	shared.items[key] = value
	shared.Unlock()
}

/*
SetNx
// like redis SETNX
// return true if the key was set
// return false if the key was not set
*/
func (m *ConcurrentMapUint64Session) SetNx(key uint64, value *_session) bool {
	shared := m.GetShared(key)
	shared.Lock()
	if _, ok := shared.items[key]; ok {
		shared.Unlock()
		return false
	}
	shared.items[key] = value
	shared.Unlock()
	return true
}

func (m ConcurrentMapUint64Session) Get(key uint64) (*_session, bool) {
	shared := m.GetShared(key)
	shared.Lock()
	val, ok := shared.items[key]
	shared.Unlock()

	return val, ok
}

func (m ConcurrentMapUint64Session) Count() int {
	count := 0
	for i := 0; i < SHARED_COUNT_SESSION_MAP; i++ {
		shared := m[i]
		shared.RLock()
		count += len(shared.items)
		shared.RUnlock()
	}
	return count
}

func (m ConcurrentMapUint64Session) Has(key uint64) bool {
	shared := m.GetShared(key)
	shared.RLock()
	_, ok := shared.items[key]
	shared.RUnlock()
	return ok
}

func (m *ConcurrentMapUint64Session) Remove(key uint64) {
	shared := m.GetShared(key)
	shared.Lock()
	delete(shared.items, key)
	shared.Unlock()
}

func (m *ConcurrentMapUint64Session) GetAndRemove(key uint64) (*_session, bool) {
	shared := m.GetShared(key)
	shared.Lock()
	val, ok := shared.items[key]
	delete(shared.items, key)
	shared.Unlock()
	return val, ok
}

// Iter Returns an iterator which could be used in a for range loop.
func (m ConcurrentMapUint64Session) Iter() <-chan TupleConcurrentMapUint64Session {
	ch := make(chan TupleConcurrentMapUint64Session)
	go func() {
		for _, shard := range m {
			shard.RLock()
			for key, val := range shard.items {
				ch <- TupleConcurrentMapUint64Session{key, val}
			}
			shard.RUnlock()
		}
		close(ch)
	}()
	return ch
}

// IterBuffered Returns a buffered iterator which could be used in a for range loop.
func (m ConcurrentMapUint64Session) IterBuffered() <-chan TupleConcurrentMapUint64Session {
	ch := make(chan TupleConcurrentMapUint64Session, m.Count())
	go func() {
		// Foreach shard.
		for _, shard := range m {
			// Foreach key, value pair.
			shard.RLock()
			for key, val := range shard.items {
				ch <- TupleConcurrentMapUint64Session{key, val}
			}
			shard.RUnlock()
		}
		close(ch)
	}()
	return ch
}
