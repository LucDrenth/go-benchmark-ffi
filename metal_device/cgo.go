package main

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Metal -framework CoreGraphics
#include <Metal/Metal.h>
#include <CoreGraphics/CoreGraphics.h>

uintptr_t getMTLDeviceCGO() {
    id<MTLDevice> device = MTLCreateSystemDefaultDevice();
    return (uintptr_t)device;
}
*/
import "C"

func getMTLDeviceCgo() uintptr {
	return uintptr(C.getMTLDeviceCGO())
}
