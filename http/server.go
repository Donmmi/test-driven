package main

import (
	"net/http"
	"fmt"
)

type PlayerStore interface {
	getPlayerScore(string) int
}

type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]

	score := p.store.getPlayerScore(player)

	fmt.Fprint(w, score)
}
