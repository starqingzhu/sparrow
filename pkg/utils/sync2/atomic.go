package sync2

import "sync/atomic"

type AtomicUint64 uint64

func (i *AtomicUint64) Add(n uint64) uint64 {
	return atomic.AddUint64((*uint64)(i), n)
}

func (i *AtomicUint64) Set(n uint64) {
	atomic.StoreUint64((*uint64)(i), n)
}

func (i *AtomicUint64) Get() uint64 {
	return atomic.LoadUint64((*uint64)(i))
}

func (i *AtomicUint64) CompareAndSwap(oldval, newval uint64) (swapped bool) {
	return atomic.CompareAndSwapUint64((*uint64)(i), oldval, newval)
}
