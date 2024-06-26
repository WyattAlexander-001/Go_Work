// hello_test.go
package main

import "testing"

func TestAdd(t *testing.T) {
	got := add(2, 3)
	want := 5

	if got != want {
		t.Errorf("add(2, 3) = %d; want %d", got, want)
	}
}

func TestSubtract(t *testing.T) {
	got := subtract(2, 3)
	want := -1

	if got != want {
		t.Errorf("subtract(2, 3) = %d; want %d", got, want)
	}
}

func TestMultiply(t *testing.T) {
	got := multiply(2, 3)
	want := 6

	if got != want {
		t.Errorf("multiply(2, 3) = %d; want %d", got, want)
	}
}
