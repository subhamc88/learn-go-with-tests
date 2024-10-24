package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("subhamc88", "")
		want := "Hello, subhamc88"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'world'", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("In spanish", func(t *testing.T) {
		got := Hello("subhamc88", "Spanish")
		want := "Hola, subhamc88"
		assertCorrectMessage(t, got, want)
	})

	t.Run("In french", func(t *testing.T) {
		got := Hello("subhamc88", "French")
		want := "Bonjour, subhamc88"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
