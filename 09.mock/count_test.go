package main

import (
	"testing"
	"bytes"
)

func TestCount(t *testing.T) {
	t.Run("count 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		count(buffer)

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
}
