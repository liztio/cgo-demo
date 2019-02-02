package main

// #cgo LDFLAGS: -laddlib
// #include "addlib.h"
import "C"

import (
	"fmt"
)

func main() {
	num := C.add_one(10)
	fmt.Printf("Go says: %d + one is %d\n", num, C.add_one(num))
}
