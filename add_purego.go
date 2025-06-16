package main

import (
	"github.com/ebitengine/purego"
)

var puregoAdd func(int64, int64) int64
var puregoAddSymbol uintptr

func loadPuregoAdd() error {
	lib, err := purego.Dlopen("./c/libadd.dylib", purego.RTLD_LAZY)
	if err != nil {
		return err
	}

	sym, err := purego.Dlsym(lib, "do_add")
	if err != nil {
		return err
	}
	puregoAddSymbol = sym

	purego.RegisterFunc(&puregoAdd, sym)
	return nil
}

func addPuregoRegisteredFunc(a, b int64) int64 {
	return puregoAdd(a, b)
}

func addPuregoSysCall(a, b int64) int64 {
	r1, _, _ := purego.SyscallN(puregoAddSymbol, uintptr(a), uintptr(b))
	return int64(r1)
}
