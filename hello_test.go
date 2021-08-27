package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to the people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying 'Hello, World!' when no string is given", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello in spanish", func(t *testing.T) {
		got := Hello("Manuel", "Spanish")
		want := "Hola, Manuel!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello in french", func(t *testing.T) {
		got := Hello("Judie", "French")
		want := "Bonjour, Judie!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello in polish", func(t *testing.T) {
		got := Hello("Maya", "Polish")
		want := "Cześć, Maya!"
		assertCorrectMessage(t, got, want)
	})
}
