package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("hello world", func(t *testing.T) {
		got := hello("", "")
		want := "Hello, World"

		assertCorrectName(t, got, want)
	})

	t.Run("hello some one", func(t *testing.T) {
		got := hello("", "Chris")
		want := "Hello, Chris"

		assertCorrectName(t, got, want)
	})

	t.Run("hello from france", func(t *testing.T) {
		got := hello("france", "Ted")
		want := "FHello, Ted"

		assertCorrectName(t, got, want)
	})

	t.Run("hello from spanish", func(t *testing.T) {
		got := hello("spanish", "Tom")
		want := "SHello, Tom"

		assertCorrectName(t, got, want)
	})
}

func assertCorrectName(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got:[%s], expected:[%s]", got, want)
	}
}
