package main

// #include <stdint.h>
// double *global_val;
// void setNum(double *val) {
//     global_val = val;
//}
// double *getNum() {
//     return global_val;
// }
import "C"
import "fmt"

func setNum() {
	var val = C.double(10)
	C.setNum(&val)
}

func main() {
	setNum()
	fmt.Printf("%d", int(*C.getNum()))
}
