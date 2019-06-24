package main

import (
	"io"
	"encoding/json"
)

type FileSystemPlayerStore struct {
	database io.ReadSeeker
}

func (f *FileSystemPlayerStore) GetLeague() []Player {
	_, err := f.database.Seek(0, 0)
	if err != nil {
		panic(err)
	}
	var league []Player
	err = json.NewDecoder(f.database).Decode(&league)
	if err != nil {
		panic(err)
	}
	return league
}