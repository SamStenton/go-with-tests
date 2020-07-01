package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
        t.Helper()
		if got != want {
				t.Errorf("got %q want %q", got, want)
			}
		}
	
	t.Run("saying hello to someone", func(t *testing.T) {
		got := Hello("Sam", "")
		want := "Hello, Sam"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("...in spanish", func(t *testing.T) {
		got := Hello("Sam", "Spanish")
		want := "Hola, Sam"

		assertCorrectMessage(t, got, want)
	})

	t.Run("...in French", func(t *testing.T) {
		got := Hello("Sam", "French")
		want := "Bonjour, Sam"

		assertCorrectMessage(t, got, want)
	})
}
