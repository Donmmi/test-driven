package main

import "fmt"

const (
	defaultName = "World"
	france  = "france"
	spanish = "spanish"
	prefix = "Hello, "
	francePrefix = "FHello, "
	spanishPrefix = "SHello, "
)

func hello(language, name string) string {
	prefix := prefix

	if name == "" {
		name = defaultName
	}

	switch language {
	case france:
		prefix = francePrefix
	case spanish:
		prefix = spanishPrefix
	}
	return prefix + name
}

func main() {
	fmt.Println(hello("", "World"))
}
