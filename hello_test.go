package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("say hello with a person's name", func(t *testing.T) {

		got := Hello("Chris", "")
		want := "Hello, Chris"

		if got != want {
			printError(t, got, want)
		}

	})

	t.Run("Don't supply a name default to hello world in English", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		if got != want {
			printError(t, got, want)
		}
	})


	t.Run("Don't supply a name default to hello world in Spanish", func(t *testing.T) {
		got := Hello("", "Spanish")
		want := "Hola, World"

		if got != want {
			printError(t, got, want)
		}
	})

	t.Run("in Spanish", func(t *testing.T){
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		if got != want {
			printError(t, got, want)
		}
	})

	t.Run("in French", func(t *testing.T){
		got := Hello("Elodie", "French")
		want := "Bonjour, Elodie"
		if got != want {
			printError(t, got, want)
		}
	})

	t.Run("in Japanese", func(t *testing.T){
		got := Hello("Elodie", "Japanese")
		want := "Kon'nichiwa, Elodie"
		if got != want {
			printError(t, got, want)
		}
	})
	
}




func printError(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
