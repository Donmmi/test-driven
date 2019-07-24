package poker_test

import (
	"testing"
	"donmmi/test-driven/http"
)

func TestFileSystemPlayerStore(t *testing.T) {
	t.Run("get sorted League", func(t *testing.T) {
		database, clean := poker.CreateTmpFile(t, `[{"Name":"Pepper","Wins":20},{"Name":"Floyd", "Wins":30}]`)
		defer clean()
		store, err := poker.NewFileSystemPlayerStore(database)
		poker.AssertNoError(t, err)

		got := store.GetLeague()
		want := []poker.Player{{"Floyd",30}, {"Pepper",20}}

		poker.AssertLeague(t, got, want)

		// read again
		got = store.GetLeague()
		poker.AssertLeague(t, got, want)
	})

	t.Run("test get player Score", func(t *testing.T) {
		database, clean := poker.CreateTmpFile(t, `[{"Name":"Pepper","Wins":20},{"Name":"Floyd", "Wins":30}]`)
		defer clean()
		store, err := poker.NewFileSystemPlayerStore(database)
		poker.AssertNoError(t, err)

		got := store.GetPlayerScore("Pepper")
		want := 20
		poker.AssertPlayerScore(t, got, want)
	})

	t.Run("get non existing player Score", func(t *testing.T) {
		database, clean := poker.CreateTmpFile(t, `[{"Name":"Pepper","Wins":20},{"Name":"Floyd", "Wins":30}]`)
		defer clean()
		store, err := poker.NewFileSystemPlayerStore(database)
		poker.AssertNoError(t, err)

		got := store.GetPlayerScore("Apollo")
		want := 0
		poker.AssertPlayerScore(t, got, want)
	})

	t.Run("Record player Score", func(t *testing.T) {
		database, clean := poker.CreateTmpFile(t, `[{"Name":"Pepper","Wins":20},{"Name":"Floyd", "Wins":30}]`)
		defer clean()
		store, err := poker.NewFileSystemPlayerStore(database)
		poker.AssertNoError(t, err)

		store.Record("Pepper")
		got := store.GetPlayerScore("Pepper")
		want := 21
		poker.AssertPlayerScore(t, got, want)
	})

	t.Run("Record non existing Score", func(t *testing.T) {
		database, clean := poker.CreateTmpFile(t, `[{"Name":"Pepper","Wins":20},{"Name":"Floyd", "Wins":30}]`)
		defer clean()
		store, err := poker.NewFileSystemPlayerStore(database)
		poker.AssertNoError(t, err)

		store.Record("Apollo")
		got := store.GetPlayerScore("Apollo")
		want := 1
		poker.AssertPlayerScore(t, got, want)
	})
}
