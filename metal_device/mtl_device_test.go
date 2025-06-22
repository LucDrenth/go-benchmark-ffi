package main

import "testing"

func TestCgo(t *testing.T) {
	result := getMTLDeviceCgo()
	if result == 0 {
		t.Fatal("failed to get MTLDevice")
	}
}

func TestPurego(t *testing.T) {
	err := loadPuregoMTLCreateSystemDefaultDeviceSymbol()
	if err != nil {
		t.Fatal("failed to load purego")
	}

	t.Run("registered func", func(t *testing.T) {
		result := getMTLDevicePuregoRegisteredFunc()
		if result == 0 {
			t.Fatal("failed to get MTLDevice")
		}
	})

	t.Run("sys call", func(t *testing.T) {
		result := getMTLDevicePuregoSysCall()
		if result == 0 {
			t.Fatal("failed to get MTLDevice")
		}
	})
}
