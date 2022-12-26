package main

import (
	"C"
	"fmt"
)

//export ffiCheck
func ffiCheck(mes *C.char) *C.char {
	var i interface{}
	i = 123
	describe(i)

	tex := C.GoString(mes)
	return C.CString(tex)
}

func main() {
}

func describe(i interface{}) {
	fmt.Println(i)
}
