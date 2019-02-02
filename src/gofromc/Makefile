all: gofromc

_cgo_export.h: greet.go
	go tool cgo -exportheader $@ $^

greet.a: greet.go
	go build -buildmode=c-archive $^

gofromc: main.c _cgo_export.h greet.a
	$(CC) -o $@ $^ -lpthread

clean:
	rm -fr _cgo_export.h greet.a _obj gofromc
