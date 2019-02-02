package main

// #include "addlib.h"
import "C"
import "fmt"

func main() {
	num := 10

	added := C.add_one(C.uint32_t(num))
	fmt.Printf("%d + 1 is %d!\n", num, added)
}
