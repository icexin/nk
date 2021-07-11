package nk

import (
	"sort"
	"unsafe"
)

type sorter struct {
	len  int
	base uintptr
	sz   uintptr
	f    func(uintptr, uintptr) int32
}

func (s *sorter) Len() int { return s.len }

func (s *sorter) Less(i, j int) bool {
	return s.f(s.base+uintptr(i)*s.sz, s.base+uintptr(j)*s.sz) < 0
}

func (s *sorter) Swap(i, j int) {
	p := uintptr(s.base + uintptr(i)*s.sz)
	q := uintptr(s.base + uintptr(j)*s.sz)
	for i := 0; i < int(s.sz); i++ {
		*(*byte)(unsafe.Pointer(p)), *(*byte)(unsafe.Pointer(q)) = *(*byte)(unsafe.Pointer(q)), *(*byte)(unsafe.Pointer(p))
		p++
		q++
	}
}

// void qsort(void *base, size_t nmemb, size_t size, int (*compar)(const void *, const void *));
func Xqsort(base uintptr, nmemb, size size_t, compar uintptr) {
	sort.Sort(&sorter{
		len:  int(nmemb),
		base: base,
		sz:   uintptr(size),
		f: (*struct {
			f func(uintptr, uintptr) int32
		})(unsafe.Pointer(&struct{ uintptr }{compar})).f,
	})
}
