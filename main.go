package main

import (
	"cgoDemo/cpp"
	"cgoDemo/objc"
)

func main() {
	objc.SayObjcHello()
	cpp.SayCppHello()
}
