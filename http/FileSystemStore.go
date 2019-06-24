package main

import (
	"io"
	"encoding/json"
)

type FileSystemPlayerStore struct {
	database io.Writer
	league league
}

func NewFileSystemPlayerStore(database io.ReadWriteSeeker) *FileSystemPlayerStore {
	_, err := database.Seek(0, 0)
	if err != nil {
		panic(err)
	}
	league, err := getLeague(database)
	if err != nil {
		panic(err)
	}
	f := &FileSystemPlayerStore{
		database:&tape{database},
		league:league,
	}
	return f
}

func (f *FileSystemPlayerStore) getLeague() league {
	return f.league
}

func (f *FileSystemPlayerStore) getPlayerScore(name string) int {
	league := f.getLeague()

	player := league.find(name)
	if player != nil {
		return player.Wins
	}

	return 0
}

func (f *FileSystemPlayerStore) record(name string) {
	league := f.getLeague()

	player := league.find(name)
	if player != nil {
		player.Wins++
	}

	err := json.NewEncoder(f.database).Encode(&league)
	if err != nil {
		panic(err)
	}
}
