package cpp

/*
#cgo CFLAGS: -mmacosx-version-min=11.0
#cgo LDFLAGS: -mmacosx-version-min=11.0

#include "cpp.h"
*/
import "C"

func SayCppHello() {
	C.SayCppHello()
}
