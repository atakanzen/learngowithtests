package main

import (
	"encoding/json"
	"fmt"
	"io"
)

func NewLeague(r io.Reader) ([]Player, error) {
	var league []Player

	err := json.NewDecoder(r).Decode(&league)
	if err != nil {
		return nil, fmt.Errorf("could not parse league, %v", err)
	}

	return league, nil
}
