package poker

import (
	"encoding/json"
	"os"
	"sort"
)

type FileSystemPlayerStore struct {
	database *json.Encoder
	league league
}

func initFileSystemPlayerStore(file *os.File) error {
	_, err := file.Seek(0, 0)
	if err != nil {
		return err
	}

	info, err := os.Stat(file.Name())
	if err != nil {
		return err
	}
	if info.Size() == 0 {
		_, err := file.Write([]byte(`[]`))
		if err != nil {
			return err
		}
		_, err = file.Seek(0, 0)
		if err != nil {
			return err
		}
	}
	return nil
}

func NewFileSystemPlayerStore(file *os.File) (*FileSystemPlayerStore, error) {
	err := initFileSystemPlayerStore(file)
	if err != nil {
		return nil, err
	}

	league, err := getLeague(file)
	if err != nil {
		return nil, err
	}
	f := &FileSystemPlayerStore{
		database:json.NewEncoder(&tape{file}),
		league:league,
	}
	return f, nil
}

func (f *FileSystemPlayerStore) getLeague() []Player {
	sort.Slice(f.league, func(i, j int) bool {
		return f.league[i].Wins > f.league[j].Wins
	})
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
