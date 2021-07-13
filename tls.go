package nk

import (
	"sync"
	"unsafe"
)

var memAllocator sync.Map
var (
	stack = make([]byte, 40960)
	bp    int
)

func tlsAlloc(n int) uintptr {
	if bp+n > len(stack) {
		panic("stack overflow")
	}
	old := bp
	bp += n
	return uintptr(unsafe.Pointer(&stack[old]))
}

func tlsFree(n int) {
	bp -= n
	if bp < 0 {
		panic("stack neg")
	}
}

func Xmalloc(n Nk_size) uintptr {
	slice := make([]byte, n)
	ptr := uintptr(unsafe.Pointer(&slice[0]))
	memAllocator.Store(ptr, slice)
	return ptr
}

func Xfree(n uintptr) {
	memAllocator.Delete(n)
}
