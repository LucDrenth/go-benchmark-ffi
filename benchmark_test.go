package main

import "testing"

func BenchmarkGo(b *testing.B) {
	for b.Loop() {
		addGo(15, 20)
	}
}
func BenchmarkCGO(b *testing.B) {
	for b.Loop() {
		addCgo(15, 20)
	}
}

func BenchmarkPurego(b *testing.B) {
	err := loadPuregoAdd()
	if err != nil {
		b.Fatalf("failed to load purego: %v", err)
	}

	b.Run("registered func", func(b *testing.B) {
		for b.Loop() {
			addPuregoRegisteredFunc(15, 20)
		}
	})

	b.Run("sys call", func(b *testing.B) {
		for b.Loop() {
			addPuregoSysCall(15, 20)
		}
	})
}

func BenchmarkPuregoLoad(b *testing.B) {
	for b.Loop() {
		loadPuregoAdd()
	}
}
