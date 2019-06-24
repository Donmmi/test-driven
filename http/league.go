package main

import (
	"io"
	"encoding/json"
)

func getLeague(rdr io.Reader) ([]Player, error) {
	var league []Player
	err := json.NewDecoder(rdr).Decode(&league)
	if err != nil {
		return nil, err
	}
	return league, nil
}
