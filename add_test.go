package main

import (
	"testing"
)

func TestGo(t *testing.T) {
	result := addGo(15, 20)
	if result != 35 {
		t.Fatal("incorrect result: ", result)
	}
}

func TestCGO(t *testing.T) {
	result := addCgo(15, 20)
	if result != 35 {
		t.Fatal("incorrect result: ", result)
	}
}

func TestPurego(t *testing.T) {
	err := loadPuregoAdd()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("registered func", func(t *testing.T) {
		result := addPuregoRegisteredFunc(15, 20)
		if result != 35 {
			t.Fatal("incorrect result: ", result)
		}
	})

	t.Run("sys call", func(t *testing.T) {
		result := addPuregoSysCall(15, 20)
		if result != 35 {
			t.Fatal("incorrect result: ", result)
		}
	})
}
