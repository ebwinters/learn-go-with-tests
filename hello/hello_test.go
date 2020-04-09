package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMsg := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("say hello to someone", func(t *testing.T) {
		got := Hello("Ethan")
		want := "Hello, Ethan"
		assertCorrectMsg(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"
		assertCorrectMsg(t, got, want)
	})

}
