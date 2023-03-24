package objc

/*
#cgo CFLAGS: -mmacosx-version-min=11.0 -x objective-c
#cgo LDFLAGS: -mmacosx-version-min=11.0 -framework Cocoa -framework Foundation

#include "objc.h"
*/
import "C"

func SayObjcHello() {
	C.SayObjcHello()
}
