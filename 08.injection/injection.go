package main

import (
	"fmt"
	"io"
	"os"
	"net/http"
)

func greeting(w io.Writer, name string) {
	fmt.Fprint(w, "Hello, ", name)
}

func greetingHandler(w http.ResponseWriter, r *http.Request) {
	greeting(w, "World")
}

func main() {
	greeting(os.Stdout, "donmmi")

	http.ListenAndServe(":5555", http.HandlerFunc(greetingHandler))
}
