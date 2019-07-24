package poker

import (
	"net/http"
	"fmt"
	"encoding/json"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
	Record(name string)
	GetLeague() []Player
}

const ContentTypeJson  = "application/json"

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
	err := json.NewEncoder(w).Encode(p.store.GetLeague())
	if err != nil {
		panic(err)
	}
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
	score := p.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	fmt.Fprint(w, score)
}

func (p *PlayerServer) processPlayerScore(w http.ResponseWriter, player string) {
	p.store.Record(player)
	w.Header().Set("content-type", ContentTypeJson)
	w.WriteHeader(http.StatusAccepted)
}
