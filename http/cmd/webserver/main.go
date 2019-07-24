package main

import (
	"net/http"
	"os"
	"donmmi/test-driven/http"
)

const dbFileName  = "game.json"

func main() {
	// store := NewInMemoryPlayerStore()
	file, err := os.OpenFile(dbFileName, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	store, err := poker.NewFileSystemPlayerStore(file)
	if err != nil {
		panic(err)
	}

	server := poker.NewPlayerServer(store)

	http.ListenAndServe(":5555", server)
}