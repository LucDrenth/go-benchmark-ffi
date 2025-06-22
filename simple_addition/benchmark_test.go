package main

import "testing"

func BenchmarkAdditionGo(b *testing.B) {
	for b.Loop() {
		addGo(15, 20)
	}
}

func BenchmarkAdditionCgo(b *testing.B) {
	for b.Loop() {
		addCgo(15, 20)
	}
}

func BenchmarkAdditionPurego(b *testing.B) {
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

func BenchmarkAdditionPuregoLoad(b *testing.B) {
	for b.Loop() {
		loadPuregoAdd()
	}
}
