package main

import (
	"testing"
	"bytes"
)

func TestGreeting(t *testing.T) {
	buffer := &bytes.Buffer{}
	greeting(buffer, "Tom")
	got := buffer.String()
	want := "Hello, Tom"

	if got != want {
		t.Errorf("got:[%s], expected:[%s]", got, want)
	}
}
