package main

import "C"
import "fmt"

//export Greet
func Greet(chars *C.char) {
	str := C.GoString(chars)
	fmt.Printf("Hello from Go, %s!\n", str)
}

func main() {} // required for main package

type Point struct {
	x, y float32
}

func getPoint() *Point {
	return &Point{
		x: 10,
		y: 12,
	}
}

func usePoint(pt *Point) {
	fmt.Printf("x: %f, y: %f\n", pt.x, pt.y)
	// Goes out of scope
}
