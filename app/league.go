package main

import (
	"encoding/json"
	"fmt"
	"io"
)

// League stores a collection of players
type League []Player

// Find returns the player with the specified name, if it exists. If not it returns nil
func (l League) Find(name string) *Player {
	for i, p := range l {
		if p.Name == name {
			return &l[i]
		}
	}

	return nil
}

// NewLeague creates a JSON league and returns it
func NewLeague(r io.Reader) (League, error) {
	var league []Player

	err := json.NewDecoder(r).Decode(&league)
	if err != nil {
		return nil, fmt.Errorf("could not parse league, %v", err)
	}

	return league, nil
}
