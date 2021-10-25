package store

import (
	"encoding/json"
	"fmt"
	"io"
)

// Player is a structure that represents a player with its Name and Score fields
type Player struct {
	Name  string
	Score int
}

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

/*
  NewLeague creates a league by decoding the given reader and returns it. If there is an error at decoding, it returns nil with the error
*/
func NewLeague(r io.Reader) (League, error) {
	var league []Player

	err := json.NewDecoder(r).Decode(&league)
	if err != nil {
		return nil, fmt.Errorf("could not parse league, %v", err)
	}

	return league, nil
}
