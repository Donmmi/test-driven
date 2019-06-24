package main

import (
	"net/http"
	"os"
)

const dbFileName  = "game.json"

func main() {
	// store := NewInMemoryPlayerStore()
	file, err := os.OpenFile(dbFileName, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	store, err := NewFileSystemPlayerStore(file)
	if err != nil {
		panic(err)
	}

	server := NewPlayerServer(store)

	http.ListenAndServe(":5555", server)
}