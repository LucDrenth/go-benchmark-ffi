package main

import "github.com/ebitengine/purego"

var mtlCreateSystemDefaultDevice func() uintptr
var mtlCreateSystemDefaultDeviceSymbol uintptr

func loadPuregoMTLCreateSystemDefaultDeviceSymbol() error {
	_, err := purego.Dlopen("/System/Library/Frameworks/CoreGraphics.framework/CoreGraphics", purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		return err
	}

	metalFramework, err := purego.Dlopen("/System/Library/Frameworks/Metal.framework/Metal", purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		return err
	}

	sym, err := purego.Dlsym(metalFramework, "MTLCreateSystemDefaultDevice")
	if err != nil {
		return err
	}
	mtlCreateSystemDefaultDeviceSymbol = sym

	purego.RegisterFunc(&mtlCreateSystemDefaultDevice, sym)

	return nil
}

func getMTLDevicePuregoRegisteredFunc() uintptr {
	return mtlCreateSystemDefaultDevice()
}

func getMTLDevicePuregoSysCall() uintptr {
	mtlDeviceId, _, _ := purego.SyscallN(mtlCreateSystemDefaultDeviceSymbol)
	return mtlDeviceId
}
