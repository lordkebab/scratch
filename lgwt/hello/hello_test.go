package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func (t *testing.T, got, want string) {
		t.Helper()

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	
	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Saying hello with an empty string", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Saying hello in spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Saying hello in french", func(t *testing.T) {
		got := Hello("Pierre", "French")
		want := "Bonjour, Pierre"

		assertCorrectMessage(t, got, want)
	})
}