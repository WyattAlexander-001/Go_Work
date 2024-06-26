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

// USE THIS AS A TEMPLATE FOR TESTING
func TestHelloClean(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := HelloName("Chris")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'world'", func(t *testing.T) {
		got := HelloName("")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})

}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
