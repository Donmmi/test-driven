package main

import (
	"io"
	"encoding/json"
)

type FileSystemPlayerStore struct {
	database io.ReadWriteSeeker
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

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	league := f.GetLeague()
	for _, player := range league {
		if player.Name == name {
			return player.Wins
		}
	}
	return 0
}

func (f *FileSystemPlayerStore) Record(name string) {
	league := f.GetLeague()
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
