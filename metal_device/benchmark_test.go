package main

import "testing"

func BenchmarkGetMTLDeviceCgo(b *testing.B) {
	for b.Loop() {
		getMTLDeviceCgo()
	}
}

func BenchmarkGetMTLDevicePurego(b *testing.B) {
	err := loadPuregoMTLCreateSystemDefaultDeviceSymbol()
	if err != nil {
		b.Fatalf("failed to load purego: %v", err)
	}

	b.Run("registered func", func(b *testing.B) {
		for b.Loop() {
			getMTLDevicePuregoRegisteredFunc()
		}
	})

	b.Run("sys call", func(b *testing.B) {
		for b.Loop() {
			getMTLDevicePuregoSysCall()
		}
	})
}

func BenchmarkGetMTLDevicePuregoLoad(b *testing.B) {
	for b.Loop() {
		loadPuregoMTLCreateSystemDefaultDeviceSymbol()
	}
}
