package main

import (
	"net/http"
	"fmt"
	"encoding/json"
)

type PlayerStore interface {
	getPlayerScore(name string) int
	record(name string)
	getLeague() []Player
}

type PlayerServer struct {
	store PlayerStore
	http.Handler
}

type Player struct {
	Name string
	Wins int
}

func NewPlayerServer(store PlayerStore) *PlayerServer {
	p := &PlayerServer{
		store:store,
	}
	router := http.NewServeMux()
	router.HandleFunc("/league", http.HandlerFunc(p.leagueHandler))
	router.HandleFunc("/players/", http.HandlerFunc(p.playersHandler))

	p.Handler = router

	return p
}

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(p.store.getLeague())
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
}

func (p *PlayerServer) playersHandler(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]
	switch r.Method {
	case http.MethodPost:
		p.processPlayerScore(w, player)
	case http.MethodGet:
		p.showPlayerScore(w, player)
	}
}

func (p *PlayerServer) showPlayerScore(w http.ResponseWriter, player string) {
	score := p.store.getPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	fmt.Fprint(w, score)
}

func (p *PlayerServer) processPlayerScore(w http.ResponseWriter, player string) {
	p.store.record(player)
	w.WriteHeader(http.StatusAccepted)
}
