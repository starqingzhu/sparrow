package pool

import (
	"encoding/json"
	"sync"
	"testing"
)

/*
保存和复用临时对象，减少内存分配，降低 GC 压力。
*/

type Student struct {
	Name   string
	Age    int32
	Remark [1024]byte
}

func BenchmarkUnmarshal(b *testing.B) {
	var buf = make([]byte, 0)
	for n := 0; n < b.N; n++ {
		stu := &Student{}
		json.Unmarshal(buf, stu)
	}
}

func BenchmarkUnmarshalWithPool(b *testing.B) {
	var studentPool = sync.Pool{
		New: func() interface{} {
			return new(Student)
		},
	}
	var buf = make([]byte, 0)
	for n := 0; n < b.N; n++ {
		stu := studentPool.Get().(*Student)
		json.Unmarshal(buf, stu)
		studentPool.Put(stu)
	}
}
