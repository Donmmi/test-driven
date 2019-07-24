package poker

import (
	"io"
	"encoding/json"
)

type league []Player

func (l league) find(name string) *Player {
	for i, p := range l {
		if p.Name == name {
			return &l[i]
		}
	}
	return nil
}

func getLeague(rdr io.Reader) ([]Player, error) {
	var league []Player
	err := json.NewDecoder(rdr).Decode(&league)
	if err != nil {
		return nil, err
	}
	return league, nil
}
