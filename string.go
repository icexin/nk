package nk

import "unsafe"

// RawMem represents the biggest byte array the runtime can handle
type RawMem [1<<30 - 1]byte

// RawMem64 represents the biggest uint64 array the runtime can handle.
type RawMem64 [unsafe.Sizeof(RawMem{}) / unsafe.Sizeof(uint64(0))]uint64

// size_t strlen(const char *s)
func Xstrlen(s uintptr) (r int) {
	if s == 0 {
		return 0
	}

	for ; *(*int8)(unsafe.Pointer(s)) != 0; s++ {
		r++
	}
	return r
}

// void *memset(void *s, int c, size_t n)
func Xmemset(s uintptr, c int32, n size_t) uintptr {
	if n != 0 {
		c := byte(c & 0xff)

		//this will make sure that on platforms where they are not equally alligned
		//we clear out the first few bytes until allignment
		bytesBeforeAllignment := s % unsafe.Alignof(uint64(0))
		if bytesBeforeAllignment > uintptr(n) {
			bytesBeforeAllignment = uintptr(n)
		}
		b := (*RawMem)(unsafe.Pointer(s))[:bytesBeforeAllignment:bytesBeforeAllignment]
		n -= size_t(bytesBeforeAllignment)
		for i := range b {
			b[i] = c
		}
		if n >= 8 {
			i64 := uint64(c) + uint64(c)<<8 + uint64(c)<<16 + uint64(c)<<24 + uint64(c)<<32 + uint64(c)<<40 + uint64(c)<<48 + uint64(c)<<56
			b8 := (*RawMem64)(unsafe.Pointer(s + bytesBeforeAllignment))[: n/8 : n/8]
			for i := range b8 {
				b8[i] = i64
			}
		}
		if n%8 != 0 {
			b = (*RawMem)(unsafe.Pointer(s + bytesBeforeAllignment + uintptr(n-n%8)))[: n%8 : n%8]
			for i := range b {
				b[i] = c
			}
		}
	}
	return s
}

// void *memcpy(void *dest, const void *src, size_t n);
func Xmemcpy(dest, src uintptr, n size_t) (r uintptr) {
	if n == 0 {
		return dest
	}

	s := (*RawMem)(unsafe.Pointer(src))[:n:n]
	d := (*RawMem)(unsafe.Pointer(dest))[:n:n]
	copy(d, s)
	return dest
}
