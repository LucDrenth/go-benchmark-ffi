package main

import (
	"github.com/ebitengine/purego"
)

var puregoAdd func(int64, int64) int64

func loadPurego() error {
	lib, err := purego.Dlopen("./c/libadd.dylib", purego.RTLD_LAZY)
	if err != nil {
		return err
	}

	sym, err := purego.Dlsym(lib, "do_add")
	if err != nil {
		return err
	}

	purego.RegisterFunc(&puregoAdd, sym)
	return nil
}

func addPurego(a, b int64) int64 {
	return puregoAdd(a, b)
}
