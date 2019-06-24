package main

import "net/http"

func main() {
	store := NewInMemoryPlayerStore()
	server := &PlayerServer{store:store}

	http.ListenAndServe(":5555", server)
}