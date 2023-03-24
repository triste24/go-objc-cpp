package main

import "C"
import (
	"fmt"
)

//export Logo
func Logo(cs *C.char) {
	if cs != nil {
		fmt.Println(C.GoString(cs))
	}
}
