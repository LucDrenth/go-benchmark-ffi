package main

import "testing"

func BenchmarkCGO(b *testing.B) {
	b.Run("CGO addition", func(b *testing.B) {
		for b.Loop() {
			addCgo(15, 20)
		}
	})

	b.Run("Go addition", func(b *testing.B) {
		for b.Loop() {
			addGo(15, 20)
		}
	})

	b.Run("Purego addition", func(b *testing.B) {
		err := loadPurego()
		if err != nil {
			b.Fatalf("failed to load purego: %v", err)
		}

		for b.Loop() {
			addPurego(15, 20)
		}
	})

	b.Run("Purego load", func(b *testing.B) {
		for b.Loop() {
			loadPurego()
		}
	})
}
