package main

import "testing"

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
	err := loadPurego()
	if err != nil {
		t.Fatal(err)
	}

	result := addPurego(15, 20)
	if result != 35 {
		t.Fatal("incorrect result: ", result)
	}
}
