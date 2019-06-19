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

func main() {
	store := &InMemoryPlayerStore{map[string]int{}}
	server := &PlayerServer{store:store}

	http.ListenAndServe(":5555", server)
}