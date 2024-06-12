package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		assertCorrectMessage(t, Hello("Chris"), "Hello, Chris")
	})

	t.Run("Hello defaults to world if no name is supplied", func(t *testing.T) {
		assertCorrectMessage(t, Hello("tg"), "Hello, World")
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
