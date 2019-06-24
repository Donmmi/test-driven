package main

import (
	"io"
	"encoding/json"
)

type FileSystemPlayerStore struct {
	database io.ReadWriteSeeker
	league []Player
}

func NewFileSystemPlayerStore(database io.ReadWriteSeeker) *FileSystemPlayerStore {
	league, err := getLeague(database)
	if err != nil {
		panic(err)
	}
	f := &FileSystemPlayerStore{
		database:database,
		league:league,
	}
	return f
}

func (f *FileSystemPlayerStore) getLeague() []Player {
	return f.league
}

func (f *FileSystemPlayerStore) getPlayerScore(name string) int {
	league := f.getLeague()
	for _, player := range league {
		if player.Name == name {
			return player.Wins
		}
	}
	return 0
}

func (f *FileSystemPlayerStore) record(name string) {
	league := f.getLeague()
	for i, player := range league {
		if player.Name == name {
			league[i].Wins++
			break
		}
	}

	// 不做seek会导致从文件尾继续写入，需要从文件头开始覆盖写
	_, err := f.database.Seek(0, 0)
	if err != nil {
		panic(err)
	}

	err = json.NewEncoder(f.database).Encode(&league)

	if err != nil {
		panic(err)
	}
}
