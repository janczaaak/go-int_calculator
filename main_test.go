package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	x := add(2, 2)
	if x != 4 {
		t.Error("Expected", 4, "Got\n", x)
	}

}
func TestSubtract(t *testing.T) {
	x := subtract(5, 2)
	if x != 3 {
		t.Error("Expected", 3, "Got", x)
	}
}
func TestMultiply(t *testing.T) {
	x := multiply(2, 3)
	if x != 6 {
		t.Error("Expected", 6, "Got", x)
	}
}
func TestDivide(t *testing.T) {
	x := divide(6, 3)
	if x != 2 {
		t.Error("Expected", 2, "Got", x)
	}
}
