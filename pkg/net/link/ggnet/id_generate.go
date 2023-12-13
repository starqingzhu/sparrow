package ggnet

import "sync/atomic"

type IDUint64 uint64

func (id *IDUint64) Get() uint64 {
	return atomic.LoadUint64((*uint64)(id))
}

func (id *IDUint64) Set(val uint64) {
	atomic.StoreUint64((*uint64)(id), val)
}

func (id *IDUint64) Add(n uint64) uint64 {
	return atomic.AddUint64((*uint64)(id), n)
}

func (id *IDUint64) CAS(old, new uint64) bool {
	return atomic.CompareAndSwapUint64((*uint64)(id), old, new)
}
