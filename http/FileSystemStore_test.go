package main

import (
	"testing"
	"io"
	"io/ioutil"
	"os"
)

func createTmpFile(t *testing.T, initialData string) (io.ReadWriteSeeker, func()) {
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
	t.Run("test get league", func(t *testing.T) {
		database, clean := createTmpFile(t, `[{"Name":"Pepper","Wins":20},{"Name":"Floyd", "Wins":30}]`)
		defer clean()
		store := NewFileSystemPlayerStore(database)

		got := store.getLeague()
		want := []Player{{"Pepper",20},{"Floyd",30}}

		assertLeague(t, got, want)

		// read again
		got = store.getLeague()
		assertLeague(t, got, want)
	})

	t.Run("test get player score", func(t *testing.T) {
		database, clean := createTmpFile(t, `[{"Name":"Pepper","Wins":20},{"Name":"Floyd", "Wins":30}]`)
		defer clean()
		store := NewFileSystemPlayerStore(database)

		got := store.getPlayerScore("Pepper")
		want := 20
		assertPlayerScore(t, got, want)
	})

	t.Run("get non existing player score", func(t *testing.T) {
		database, clean := createTmpFile(t, `[{"Name":"Pepper","Wins":20},{"Name":"Floyd", "Wins":30}]`)
		defer clean()
		store := NewFileSystemPlayerStore(database)

		got := store.getPlayerScore("Apollo")
		want := 0
		assertPlayerScore(t, got, want)
	})

	t.Run("record player score", func(t *testing.T) {
		database, clean := createTmpFile(t, `[{"Name":"Pepper","Wins":20},{"Name":"Floyd", "Wins":30}]`)
		defer clean()
		store := NewFileSystemPlayerStore(database)

		store.record("Pepper")
		got := store.getPlayerScore("Pepper")
		want := 21
		assertPlayerScore(t, got, want)
	})
}

func assertPlayerScore(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got:[%d], expected:[%d]", got, want)
	}
}
