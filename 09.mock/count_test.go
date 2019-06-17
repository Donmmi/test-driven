package main

import (
	"testing"
	"bytes"
)

type SpySleeper struct {
	calls int
}

func (s *SpySleeper) Sleep() {
	s.calls++
}

func TestCount(t *testing.T) {
	t.Run("count 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		sleeper := &SpySleeper{}
		count(buffer, sleeper)

		got := buffer.String()
		want := `3
2
1
Go!
`

		if got != want {
			t.Errorf("got:[%s], expected:[%s]", got, want)
		}
	})

	t.Run("inspect sleep", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		sleeper := &SpySleeper{}

		count(buffer, sleeper)

		got := sleeper.calls
		want := 3

		if got != want {
			t.Errorf("got:[%d], expected:[%d]", got, want)
		}
	})
}
