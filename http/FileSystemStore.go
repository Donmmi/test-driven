package main

import (
	"encoding/json"
	"os"
)

type FileSystemPlayerStore struct {
	database *json.Encoder
	league league
}

func NewFileSystemPlayerStore(database *os.File) *FileSystemPlayerStore {
	_, err := database.Seek(0, 0)
	if err != nil {
		panic(err)
	}
	league, err := getLeague(database)
	if err != nil {
		panic(err)
	}
	f := &FileSystemPlayerStore{
		database:json.NewEncoder(&tape{database}),
		league:league,
	}
	return f
}

func (f *FileSystemPlayerStore) getLeague() []Player {
	return f.league
}

func (f *FileSystemPlayerStore) getPlayerScore(name string) int {
	player := f.league.find(name)
	if player != nil {
		return player.Wins
	}

	return 0
}

func (f *FileSystemPlayerStore) record(name string) {
	player := f.league.find(name)
	if player != nil {
		player.Wins++
	} else {
		f.league = append(f.league, Player{name, 1})
	}

	err := f.database.Encode(f.league)
	if err != nil {
		panic(err)
	}
}
