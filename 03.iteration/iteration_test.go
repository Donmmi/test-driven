package main

import (
	"testing"
	"fmt"
)

func TestIteration(t *testing.T) {
	t.Run("repeat", func(t *testing.T) {
		got := Repeat("a", 6)
		want := "aaaaaa"

		assertString(t, got, want)
	})

	t.Run("repeat with times", func(t *testing.T) {
		got := Repeat("a", 10)
		want := "aaaaaaaaaa"

		assertString(t, got, want)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}

func BenchmarkRepeatStandLib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RepeatStandLib("a", 10)
	}
}

func ExampleRepeat() {
	fmt.Println(Repeat("a", 6))
	// Output: aaaaaa
}

func assertString(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got:[%s], expected:[%s]", got, want)
	}
}
