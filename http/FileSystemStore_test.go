package main

import (
	"testing"
	"strings"
)

func TestFileSystemPlayerStore(t *testing.T) {
	t.Run("test get league", func(t *testing.T) {
		database := strings.NewReader(`[{"Name":"Pepper","Wins":20},{"Name":"Floyd", "Wins":30}]`)
		store := FileSystemPlayerStore{database}

		got := store.GetLeague()
		want := []Player{{"Pepper",20},{"Floyd",30}}

		assertLeague(t, got, want)

		// read again
		got = store.GetLeague()
		assertLeague(t, got, want)
	})

	t.Run("test get player score", func(t *testing.T) {
		database := strings.NewReader(`[{"Name":"Pepper","Wins":20},{"Name":"Floyd", "Wins":30}]`)
		store := FileSystemPlayerStore{database}

		got := store.GetPlayerScore("Pepper")
		want := 20
		assertPlayerScore(t, got, want)
	})
}

func assertPlayerScore(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got:[%d], expected:[%d]", got, want)
	}
}
