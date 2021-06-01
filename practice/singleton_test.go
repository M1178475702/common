package practice

import (
	"sync"
	"sync/atomic"
	"testing"
	"unsafe"
)

type Si struct {

}

var si *Si

func GetSi_1() *Si {
	if si != nil {
		return nil
	}
	ptr := unsafe.Pointer(si)
	//每次都会创建一个对象，不好
	if atomic.CompareAndSwapPointer(&ptr, unsafe.Pointer(nil), unsafe.Pointer(&Si{})) {
		return si
	}

	si = &Si{}
	return si
}
var once sync.Once
func GetSi() *Si {
	if si != nil {
		return si
	}

	once.Do(func() {
		si = &Si{}
	})
	return si
}



func TestSingleton(t *testing.T) {

}
