package main

import (
	"testing"
	"fmt"
)

func TestAdd(t *testing.T) {
	got := Add(4, 2)
	want := 6

	if got != want {
		t.Errorf("got:[%d], expected:[%d]", got, want)
	}
}

func ExampleAdd() {
	fmt.Println(Add(7, 9))
	// Output: 16
}
