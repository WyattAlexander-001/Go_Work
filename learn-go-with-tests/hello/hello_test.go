// For testing : go mod init hello
package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, world"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestHelloName(t *testing.T) {
	got := HelloName("Wyatt")
	want := "Hello, Wyatt"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestIdealHello(t *testing.T) {
	got := IdealHello("Wyatt", "Spanish")
	want := "Hola, Wyatt"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	got = IdealHello("", "Spanish")
	want = "Hola, world"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	got = IdealHello("", "")
	want = "Hi, world"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
