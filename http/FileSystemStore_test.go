package main

import (
	"testing"
	"io/ioutil"
	"os"
)

func createTmpFile(t *testing.T, initialData string) (*os.File, func()) {
	t.Helper()
	tmpFile, err := ioutil.TempFile("", "db")
	if err != nil {
		panic(err)
	}
	_, err = tmpFile.Write([]byte(initialData))
	if err != nil {
		panic(err)
	}

	_, err = tmpFile.Seek(0 ,0)
	if err != nil {
		panic(err)
	}

	clean := func() {
		err = tmpFile.Close()
		if err != nil {
			panic(err)
		}
		err = os.Remove(tmpFile.Name())
		if err != nil {
			panic(err)
		}
	}
	return tmpFile, clean
}

func TestFileSystemPlayerStore(t *testing.T) {
	t.Run("get sorted league", func(t *testing.T) {
		database, clean := createTmpFile(t, `[{"Name":"Pepper","Wins":20},{"Name":"Floyd", "Wins":30}]`)
		defer clean()
		store, err := NewFileSystemPlayerStore(database)
		assertNoError(t, err)

		got := store.getLeague()
		want := []Player{{"Floyd",30}, {"Pepper",20}}

		assertLeague(t, got, want)

		// read again
		got = store.getLeague()
		assertLeague(t, got, want)
	})

	t.Run("test get player score", func(t *testing.T) {
		database, clean := createTmpFile(t, `[{"Name":"Pepper","Wins":20},{"Name":"Floyd", "Wins":30}]`)
		defer clean()
		store, err := NewFileSystemPlayerStore(database)
		assertNoError(t, err)

		got := store.getPlayerScore("Pepper")
		want := 20
		assertPlayerScore(t, got, want)
	})

	t.Run("get non existing player score", func(t *testing.T) {
		database, clean := createTmpFile(t, `[{"Name":"Pepper","Wins":20},{"Name":"Floyd", "Wins":30}]`)
		defer clean()
		store, err := NewFileSystemPlayerStore(database)
		assertNoError(t, err)

		got := store.getPlayerScore("Apollo")
		want := 0
		assertPlayerScore(t, got, want)
	})

	t.Run("record player score", func(t *testing.T) {
		database, clean := createTmpFile(t, `[{"Name":"Pepper","Wins":20},{"Name":"Floyd", "Wins":30}]`)
		defer clean()
		store, err := NewFileSystemPlayerStore(database)
		assertNoError(t, err)

		store.record("Pepper")
		got := store.getPlayerScore("Pepper")
		want := 21
		assertPlayerScore(t, got, want)
	})

	t.Run("record non existing score", func(t *testing.T) {
		database, clean := createTmpFile(t, `[{"Name":"Pepper","Wins":20},{"Name":"Floyd", "Wins":30}]`)
		defer clean()
		store, err := NewFileSystemPlayerStore(database)
		assertNoError(t, err)

		store.record("Apollo")
		got := store.getPlayerScore("Apollo")
		want := 1
		assertPlayerScore(t, got, want)
	})
}

func assertPlayerScore(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got:[%d], expected:[%d]", got, want)
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("unexpected err:", err)
	}
}
