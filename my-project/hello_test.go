package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		assertCorrectMessage(t, Hello("Chris", ""), "Hello, Chris")
	})

	t.Run("empty string defaults to 'world'", func(t *testing.T) {
		assertCorrectMessage(t, Hello("", ""), "Hello, World")
	})

	t.Run("saying hello to people in Spanish", func(t *testing.T) {
		got := Hello("Chris", "Spanish")
		want := "Hola, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to people in French", func(t *testing.T) {
		got := Hello("Chris", "French")
		want := "Bonjour, Chris"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
