package main

import "net/http"

type InMemoryPlayerStore struct {
	store map[string]int
}

func (i *InMemoryPlayerStore) getPlayerScore(name string) int {
	return i.store[name]
}

func (i *InMemoryPlayerStore) record(name string) {
	i.store[name]++
}

func (i *InMemoryPlayerStore) getLeague() []Player {
	var league []Player
	for name, wins := range i.store {
		league = append(league, Player{name, wins})
	}
	return league
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

func main() {
	store := NewInMemoryPlayerStore()
	server := &PlayerServer{store:store}

	http.ListenAndServe(":5555", server)
}