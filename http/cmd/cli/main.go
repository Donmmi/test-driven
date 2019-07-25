package main

import (
	"fmt"
	"os"
	"donmmi/test-driven/http"
	"log"
)

const gameFileName  = "game.json"

func main() {
	fmt.Println("Let's play poker.")

	store, err := poker.NewFileSystemPlayerStoreFromFile(gameFileName)
	if err != nil {
		log.Fatal("new store from file err:", err.Error())
	}

	cli := poker.NewCLI(store, os.Stdin)
	cli.PlayPoker()
}
