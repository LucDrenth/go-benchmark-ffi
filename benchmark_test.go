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
	err := loadPurego()
	if err != nil {
		b.Fatalf("failed to load purego: %v", err)
	}

	for b.Loop() {
		addPurego(15, 20)
	}
}

func BenchmarkPuregoLoad(b *testing.B) {
	for b.Loop() {
		loadPurego()
	}
}
