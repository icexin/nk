#!/bin/bash

function compile {
    target=nk_${TARGET_GOOS}_${TARGET_GOARCH}.go
	ccgo -o $target nuklear/nuklear.c \
		-export-externs "X" \
		-export-fields "" \
		-export-defines "" \
		-export-enums "" \
		-export-typedefs "" \
		-export-structs "" \
		-nostdlib \
		-pkgname "nk"

	sed -i -e 's/var win_bounds/var win_bounds Nk_rect/' $target
	sed -i -e 's/var iter_bounds/var iter_bounds Nk_rect/' $target
	sed -i -e 's/tls \*TLS,//g' $target
	sed -i -e 's/tls \*TLS//g' $target
	sed -i -e 's/\*TLS,//g' $target
	sed -i -e 's/\*TLS//g' $target
	sed -i -e 's/tls\.Alloc/tlsAlloc/g' $target
	sed -i -e 's/tls\.Free/tlsFree/g' $target
	sed -i -e 's/tls,//g' $target
	sed -i -e 's/(tls)/()/g' $target
	sed -i -e 's/Nk__panic/Nk__panic_empty/g' $target

	go fmt $target
}

TARGET_GOOS=linux TARGET_GOARCH=386 compile
TARGET_GOOS=linux TARGET_GOARCH=amd64 compile
