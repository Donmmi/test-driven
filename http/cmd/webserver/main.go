package main

import (
	"net/http"
	"donmmi/test-driven/http"
	"log"
)

const dbFileName  = "game.json"

func main() {
	store, err := poker.NewFileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal("new store from file err:", err.Error())
	}
	server := poker.NewPlayerServer(store)

	http.ListenAndServe(":5555", server)
}