package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to Luke Skywalker", func(t *testing.T) {
		got := Hello("Luke Skywalker", "English")
		want := "Hello, Luke Skywalker"
		assertCorrectMessage(t, got, want)

	})

	t.Run("empty string defaults to 'World'", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in spanish", func(t *testing.T) {
		got := Hello("Jose", "Spanish")
		want := "Hola, Jose"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in french", func(t *testing.T) {
		got := Hello("Pierre", "French")
		want := "Bonjour, Pierre"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in japanese", func(t *testing.T) {
		got := Hello("重雄", "Japanese")
		want := "こんにちは, 重雄"
		assertCorrectMessage(t, got, want)
	})
}
