// Copyright 2020 The Libc Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package nk

import (
	"fmt"
	"unsafe"
)

func Bool32(b bool) int32 {
	if b {
		return 1
	}

	return 0
}

func gostr(s uintptr) string {
	var ret []byte
	for c := s; ; c++ {
		p := (*byte)(unsafe.Pointer(c))
		if *p == 0 {
			break
		}
		ret = append(ret, *p)
	}
	return string(ret)
}

func X__assert_fail(assertion, file uintptr, line uint32, function uintptr) {
	panic(fmt.Sprintf("Assert failed:%s %s:%d.%s", gostr(assertion), gostr(file), line, gostr(function)))
}
