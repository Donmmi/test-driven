package main

import (
	"io"
	"os"
	"fmt"
)

func count(w io.Writer) {
	for i := 3; i > 0; i-- {
		fmt.Fprintln(w, i)
	}

	fmt.Fprintln(w, "Go!")
}

func main() {
	count(os.Stdout)
}
