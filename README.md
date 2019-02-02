# Cgo Demos

`make all`

## [Addlib](blob/master/src/addlib)

_A .so for adding two numbers together_

```shell
make -C src/addlib
```

## [Addbin](blob/master/cmd/addbin)
A go library that calls [Addlib](#Addlib)

```
make addbin
LD_LIBRARY_PATH=src/addlib addbin
```

## [Imagemagick] (blob/master/cmd/imagemagick)
_A small example Go application using libimagemagick_

May require `libmagickwand-6-headers` to compile and `libmagickwand-6.q16`

```shell
make imagemagick
./imagemagick
```

## [C_from_go](blob/master/cmd/c_from_go)
_An example of calling C from Go_

```shell
make c_from_go
./c_from_go
```

## [gofromc](blob/master/src/gofromc)
An example of calling go from C

```shell
make -C src/gofromc
./gofromc
```
