package libc

import (
	"sync"
	"unsafe"

	"github.com/icexin/nk/libc/sys/types"
)

var (
	memAllocator sync.Map
)

type TLS struct {
	stack []byte
	bp    int
}

func NewTLS() *TLS {
	return &TLS{
		stack: make([]byte, 40960),
	}
}

func (t *TLS) Alloc(n int) uintptr {
	if t.bp+n > len(t.stack) {
		panic("stack overflow")
	}
	old := t.bp
	t.bp += n
	return uintptr(unsafe.Pointer(&t.stack[old]))
}

func (t *TLS) Free(n int) {
	t.bp -= n
	if t.bp < 0 {
		panic("stack neg")
	}
}

func Xmalloc(t *TLS, n types.Size_t) uintptr {
	slice := make([]byte, n)
	ptr := uintptr(unsafe.Pointer(&slice[0]))
	memAllocator.Store(ptr, slice)
	return ptr
}

func Xfree(t *TLS, ptr uintptr) {
	memAllocator.Delete(ptr)
}
