package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {		
		got := Hello("world")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is suppplied", func(t *testing.T) {
		want := "Hello, World"
		got := Hello("")
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper() // With this we tell this method is a helper
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}