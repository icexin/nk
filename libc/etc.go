// Copyright 2020 The Libc Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package libc

import (
	"fmt"
	"unsafe"

	"github.com/icexin/nk/libc/sys/types"
)

func Bool32(b bool) int32 {
	if b {
		return 1
	}

	return 0
}

func CString(s string) uintptr {
	n := len(s)
	p := Xmalloc(nil, types.Size_t(n)+1)
	if p == 0 {
		return 0
	}

	copy((*RawMem)(unsafe.Pointer(p))[:n:n], s)
	*(*byte)(unsafe.Pointer(p + uintptr(n))) = 0
	return p
}

func GoString(s uintptr) string {
	if s == 0 {
		return ""
	}

	var buf []byte
	for {
		b := *(*byte)(unsafe.Pointer(s))
		if b == 0 {
			return string(buf)
		}

		buf = append(buf, b)
		s++
	}
}

// GoBytes returns a byte slice from a C char* having length len bytes.
func GoBytes(s uintptr, len int) []byte {
	if len == 0 {
		return nil
	}

	return (*RawMem)(unsafe.Pointer(s))[:len:len]
}

func X__assert_fail(tls *TLS, assertion, file uintptr, line uint32, function uintptr) {
	panic(fmt.Sprintf("Assert failed:%s %s:%d.%s", GoString(assertion), GoString(file), line, GoString(function)))
}
