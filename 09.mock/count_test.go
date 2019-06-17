package main

import (
	"testing"
	"bytes"
	"reflect"
)

const (
	sleep = "sleep"
	write = "write"
)

type SpySleeperWithOrder struct {
	calls []string
}

func (s *SpySleeperWithOrder) Sleep() {
	s.calls = append(s.calls, sleep)
}

func (s *SpySleeperWithOrder) Write(p []byte) (n int, err error) {
	s.calls = append(s.calls, write)
	return
}

func TestCount(t *testing.T) {
	t.Run("count 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		sleeper := &SpySleeperWithOrder{}
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
		sleeper := &SpySleeperWithOrder{}

		count(sleeper, sleeper)

		got := sleeper.calls
		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got:[%v], expected:[%v]", got, want)
		}
	})
}
