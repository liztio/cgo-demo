.PHONY: all clean run

GO?="go"
addlib := $(shell pwd)/src/addlib

all:
	@make -C src/addlib
	@make -C src/addbin
	@make -C src/gofromc
	@make addbin
	@make imagemagick
	@make cfromgo

clean:
	rm -f addbin c_from_go imagemagick
	@make -C src/addlib clean
	@make -C src/gofromc clean
	@make -C src/addbin clean

src/addlib/addlib.a: src/addlib/addlib.c
	$(CC)

go_bin: addbin

addbin: cmd/addbin/main.go src/addlib/addlib.o
	CGO_CFLAGS=-I$(addlib) CGO_LDFLAGS=-L$(addlib) $(GO) build github.com/liztio/cgo-demo/cmd/addbin

imagemagick: cmd/imagemagick/main.go
	$(GO) build github.com/liztio/cgo-demo/cmd/imagemagick

cfromgo: cmd/c_from_go/addlib.c	cmd/c_from_go/addlib.h	cmd/c_from_go/main.go
	$(GO) build github.com/liztio/cgo-demo/cmd/c_from_go
