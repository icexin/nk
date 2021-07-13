#!/bin/bash

function compile {
    target=cnk/nk_${TARGET_GOARCH}.go
	./ccgo -o $target nuklear/nuklear.c \
		-export-externs "X" \
		-export-fields "" \
		-export-defines "" \
		-export-enums "" \
		-export-typedefs "" \
		-export-structs "" \
		-pkgname "cnk"

	sed -i -e 's/var win_bounds/var win_bounds Nk_rect/' $target
	sed -i -e 's/var iter_bounds/var iter_bounds Nk_rect/' $target
	sed -i -e 's|modernc.org/libc|github.com/icexin/nk/libc|' $target
	# sed -i -e 's/tls \*TLS,//g' $target
	# sed -i -e 's/tls \*TLS//g' $target
	# sed -i -e 's/\*TLS,//g' $target
	# sed -i -e 's/\*TLS//g' $target
	# sed -i -e 's/tls\.Alloc/tlsAlloc/g' $target
	# sed -i -e 's/tls\.Free/tlsFree/g' $target
	# sed -i -e 's/tls,//g' $target
	# sed -i -e 's/(tls)/()/g' $target

	go fmt $target
    mv cnk/capi_${TARGET_GOOS}_${TARGET_GOARCH}.go cnk/capi_${TARGET_GOARCH}.go
}

TARGET_GOOS=linux TARGET_GOARCH=386 compile
TARGET_GOOS=linux TARGET_GOARCH=amd64 compile
