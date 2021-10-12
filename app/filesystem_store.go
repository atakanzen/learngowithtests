package main

import (
	"io"
)

type FileSystemPlayerStore struct {
	DB io.ReadSeeker
}

func (f *FileSystemPlayerStore) GetLeague() []Player {
	f.DB.Seek(0, 0)
	league, _ := NewLeague(f.DB)
	return league
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	var score int

	for _, player := range f.GetLeague() {
		if player.Name == name {
			score = player.Score
			break
		}
	}

	return score
}
