package nk

import (
	"sync"
	"unsafe"
)

var tlsAllocator sync.Map

func tlsAlloc(n int) uintptr {
	slice := make([]byte, n)
	ptr := uintptr(unsafe.Pointer(&slice[0]))
	tlsAllocator.Store(ptr, slice)
	return ptr
}

func tlsFree(n uintptr) {
	tlsAllocator.Delete(n)
}

func Xmalloc(n Nk_size) uintptr {
	return tlsAlloc(int(n))
}

func Xfree(n uintptr) {
	tlsFree(n)
}
